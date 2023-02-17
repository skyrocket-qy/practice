const SQL = require('sequelize');
const translation = require('./translation');
const settings = require('../settings');

class StoreObjects {
  #db
  static #lockInit = true
  constructor(db) {
    if (StoreObjects.#lockInit) {
      throw "can not use 'new' directly";
    }
    this.#db = db;
    StoreObjects.#lockInit = true;
  }

  get connection() {
    return this.#db;
  }

  buildupModel(tbName, definition, fieldsLangExtendsion) {
    if (!fieldsLangExtendsion) {
      //reset the value evaluated as false to be undefined
      // since fieldExtend function only skips 'undefined' value
      fieldsLangExtendsion = undefined;
    }
    else if (Object.getPrototypeOf(fieldsLangExtendsion) !== Array.prototype) {
      throw `Stop bulding up model because ${fieldsLangExtendsion} is not array`;
    }

    // in fieldsLangExtendsion: decides fields that need language extensions -> nameEnUs, nameZhHk, nameZhTw
    return this.#db.define(tbName, translation.fieldExtend(tbName, definition, fieldsLangExtendsion));
  }

  static create(db) {
    StoreObjects.#lockInit = false;
    return new this(db);
  }
}

class UMSStoreObjects extends StoreObjects {
  // objects in databases
  //cms_*
  #tags;
  #categories;
  #photos;
  #memberLevels;
  #gameGroups;
  #ratings;
  #pegiRatings;
  #locales;
  #projects;
  #games;
  #templates;
  #gameLists;
  #gamelistGames;
  #banners;
  #rankings;
  #rankingGames;
  #bannerLists;
  #bannerlistBanners;
  #pageBlocks;
  #pages;
  #pagePageblocks;
  #configs;
  //auth
  #authUsers;
  //ams_*
  #users;
  #avatars;
  #favorites;
  #recentPlays;
  #leaderboard;
  #familyMembers;
  #families;

  constructor(db) {
    super(db);
    //build up models through getter
    //cms_*
    this.tags;
    this.categories;
    this.photos;
    this.memberLevels;
    this.gameGroups;
    this.ratings;
    this.pegiRatings;
    this.locales;
    this.projects;
    this.games;
    this.templates;
    this.gameLists;
    this.gamelistGames;
    this.banners;
    this.rankings;
    this.rankingGames;
    this.bannerLists;
    this.bannerlistBanners;
    this.pageBlocks;
    this.pages;
    this.pagePageblocks;
    this.configs;
    //auth
    this.authUsers;
    //ams_*
    this.users;
    this.avatars;
    this.favorites;
    this.recentPlays;
    this.leaderboard;
    this.familyMembers;
    this.families;
  }

  buildRelation(isPreview = false) {
    if (isPreview) {
      this.families.belongsToMany(this.users, {
        as: 'members', through: {
          model: this.familyMembers,
          unique: false,
        }, foreignKey: 'familyId'
      });
      translation.updateTransMap('ams_family', [{'members': 'ams_userprofile'}]);
      this.users.belongsToMany(this.families, {
        as: 'members', through: {
          model: this.familyMembers,
          unique: false,
        }, foreignKey: 'userprofileId'
      });
      translation.updateTransMap('ams_userprofile', [{'members': 'ams_family'}]);

      this.users.belongsTo(this.authUsers, { as: 'user', foreignKey: 'userId' });
      this.users.belongsTo(this.memberLevels, { as: 'level', foreignKey: 'levelId' });
      this.users.belongsTo(this.avatars, { as: 'avatar', foreignKey: 'avatarId' });
      this.users.belongsTo(this.ratings, {as: 'rating', foreignKey: { name: 'ratingId', allowNull: true }});
      this.users.belongsTo(this.pegiRatings, {as: 'pegiRating', foreignKey: { name: 'pegiRatingId', allowNull: true }});
      translation.updateTransMap('ams_userprofile', [{'user': 'auth_user'}, {'level': 'cms_memberlevel'}, {'avatar': 'ams_avatar'}]);
      this.users.belongsToMany(this.games, {
        as: 'favorites', through: {
          model: this.favorites,
          unique: false,
        }, foreignKey: 'userprofileId'
      });
      translation.updateTransMap('ams_userprofile', [{'favorites': 'cms_game'}]);
      this.games.belongsToMany(this.users, {
        as: 'favorites', through: {
          model: this.favorites,
          unique: false,
        }, foreignKey: 'gameId'
      });
      translation.updateTransMap('cms_game', [{'favorites': 'ams_userprofile'}]);
      this.users.belongsToMany(this.games, {
        as: 'recentPlay', through: {
          model: this.recentPlays,
          unique: false,
        }, foreignKey: 'userId'
      });
      translation.updateTransMap('ams_userprofile', [{'recentPlay': 'cms_game'}]);
      this.games.belongsToMany(this.users, {
        as: 'recentPlay', through: {
          model: this.recentPlays,
          unique: false,
        }, foreignKey: 'gameId'
      });
      translation.updateTransMap('cms_game', [{'recentPlay': 'ams_userprofile'}]);
      this.users.belongsToMany(this.games, {
        as: 'leaderboard', through: {
          model: this.leaderboard,
          unique: false,
        }, foreignKey: 'userId'
      });
      translation.updateTransMap('ams_userprofile', [{'leaderboard': 'cms_game'}]);
      this.games.belongsToMany(this.users, {
        as: 'leaderboard', through: {
          model: this.leaderboard,
          unique: false,
        }, foreignKey: 'gameId'
      });
      translation.updateTransMap('cms_game', [{'leaderboard': 'ams_userprofile'}]);
    }
    
    this.memberLevels.belongsTo(this.photos, { as: 'icon', foreignKey: 'iconId' });
    translation.updateTransMap('cms_memberlevel', [{'icon': 'cms_photo'}]);

    this.gameGroups.belongsTo(this.photos, { as: 'icon', foreignKey: 'iconId' });
    // Table name will exceeded the limit, the following 2 lines won't work!!!
    //gameGroups.belongsToMany(gameGroups, {as: 'children', through: 'cms_gamegroup_children', foreignKey: 'fromGamegroupId', constraints: false});
    //gameGroups.belongsToMany(gameGroups, {as: 'gameGroupSet', through: 'cms_gamegroup_children', foreignKey: 'toGamegroupId'});
    this.gameGroups.belongsTo(this.memberLevels, { as: 'level', foreignKey: 'levelId' });
    translation.updateTransMap('cms_gamegroup', [{'icon': 'cms_photo'}, {'level': 'cms_memberlevel'}]);

    this.ratings.belongsTo(this.photos, { as: 'icon', foreignKey: 'iconId' });
    translation.updateTransMap('cms_rating', [{'icon': 'cms_photo'}]);
    this.pegiRatings.belongsTo(this.photos, { as: 'icon', foreignKey: 'iconId' });
    translation.updateTransMap('cms_pegirating', [{'icon': 'cms_photo'}]);

    this.games.belongsTo(this.photos, { as: 'logo', foreignKey: 'logoId' });
    this.games.belongsTo(this.photos, { as: 'header', foreignKey: 'headerId' });
    this.games.belongsTo(this.photos, { as: 'splash', foreignKey: 'splashId' });
    this.games.belongsTo(this.photos, { as: 'videoPreview', foreignKey: 'videoPreviewId' });
    this.games.belongsTo(this.ratings, { as: 'rating', foreignKey: 'ratingId' });
    this.games.belongsTo(this.pegiRatings, { as: 'ratingPegi', foreignKey: 'ratingPegiId' });
    translation.updateTransMap('cms_game', [{'logo': 'cms_photo'}, {'header': 'cms_photo'}, {'splash': 'cms_photo'}, {'videoPreview': 'cms_photo'}, {'rating': 'cms_rating'}, {'ratingPegi': 'cms_pegirating'}]);

    this.games.belongsToMany(this.photos, { as: 'imageGallery', through: 'cms_game_image_gallery', foreignKey: 'gameId' });
    translation.updateTransMap('cms_game', [{'imageGallery': 'cms_photo'}]);
    this.photos.belongsToMany(this.games, { as: 'gameSet', through: 'cms_game_image_gallery', foreignKey: 'photoId' });
    translation.updateTransMap('cms_photo', [{'gameSet': 'cms_game'}]);

    this.games.belongsToMany(this.categories, { as: 'categories', through: 'cms_game_categories', foreignKey: 'gameId' });
    translation.updateTransMap('cms_game', [{'categories': 'cms_category'}]);
    this.categories.belongsToMany(this.games, { as: 'gameSet', through: 'cms_game_categories', foreignKey: 'categoryId' });
    translation.updateTransMap('cms_category', [{'gameSet': 'cms_game'}]);

    this.games.belongsToMany(this.tags, { as: 'tags', through: 'cms_game_tags', foreignKey: 'gameId' });
    translation.updateTransMap('cms_game', [{'tags': 'cms_tag'}]);
    this.tags.belongsToMany(this.games, { as: 'gameSet', through: 'cms_game_tags', foreignKey: 'tagId' });
    translation.updateTransMap('cms_tag', [{'gameSet': 'cms_game'}]);

    this.games.belongsToMany(this.locales, { as: 'locales', through: 'cms_game_locales', foreignKey: 'gameId' });
    translation.updateTransMap('cms_game', [{'locales': 'cms_gamelocale'}]);
    this.locales.belongsToMany(this.games, { as: 'gameSet', through: 'cms_game_locales', foreignKey: 'gamelocaleId' });
    translation.updateTransMap('cms_gamelocale', [{'gameSet': 'cms_game'}]);

    this.games.belongsToMany(this.projects, { as: 'projects', through: 'cms_game_projects', foreignKey: 'gameId' });
    translation.updateTransMap('cms_game', [{'projects': 'cms_project'}]);
    this.projects.belongsToMany(this.games, { as: 'gameSet', through: 'cms_game_projects', foreignKey: 'projectId' });
    translation.updateTransMap('cms_project', [{'gameSet': 'cms_game'}]);

    this.games.belongsToMany(this.gameGroups, { as: 'group', through: 'cms_game_group', foreignKey: 'gameId' });
    translation.updateTransMap('cms_game', [{'group': 'cms_gamegroup'}]);
    this.gameGroups.belongsToMany(this.games, { as: 'gameSet', through: 'cms_game_group', foreignKey: 'gamegroupId' });
    translation.updateTransMap('cms_gamegroup', [{'gameSet': 'cms_game'}]);

    this.gameLists.belongsToMany(this.games, {
      as: 'games', through: {
        model: this.gamelistGames,
        unique: false,
      }, foreignKey: 'gamelistId'
    });
    translation.updateTransMap('cms_gamelist', [{'games': 'cms_game'}]);
    this.games.belongsToMany(this.gameLists, {
      as: 'gamelistSet', through: {
        model: this.gamelistGames,
        unique: false,
      }, foreignKey: 'game_id'
    });
    translation.updateTransMap('cms_game', [{'gamelistSet': 'cms_gamelist'}]);

    this.rankings.belongsToMany(this.games, {
      as: 'games', through: {
        model: this.rankingGames,
        unique: false,
      }, foreignKey: 'rankingId'
    });
    translation.updateTransMap('cms_ranking', [{'games': 'cms_game'}]);
    this.games.belongsToMany(this.rankings, {
      as: 'rankingSet', through: {
        model: this.rankingGames,
        unique: false,
      }, foreignKey: 'game_id'
    });
    translation.updateTransMap('cms_game', [{'rankingSet': 'cms_ranking'}]);

    this.banners.belongsTo(this.tags, { as: 'tag', foreignKey: 'tagId' });
    this.banners.belongsTo(this.gameLists, { as: 'gameList', foreignKey: 'gameListId' });
    this.banners.belongsTo(this.pages, { as: 'page', foreignKey: 'pageId' });
    translation.updateTransMap('cms_banner', [{'tag': 'cms_tag'}, {'gameList': 'cms_gamelist'}, {'page': 'cms_page'}]);

    this.bannerLists.belongsToMany(this.banners, {
      as: 'banners', through: {
        model: this.bannerlistBanners,
        unique: false,
      }, foreignKey: 'bannerlistId'
    });
    translation.updateTransMap('cms_bannerlist', [{'banners': 'cms_banner'}]);
    this.banners.belongsToMany(this.bannerLists, {
      as: 'bannerSet', through: {
        model: this.bannerlistBanners,
        unique: false,
      }, foreignKey: 'banner_id'
    });
    translation.updateTransMap('cms_banner', [{'bannerSet': 'cms_bannerlist'}]);

    this.pageBlocks.belongsTo(this.templates, { as: 'template', foreignKey: 'templateId' });
    this.pageBlocks.belongsTo(this.gameLists, { as: 'gameList', foreignKey: 'gameListId' });
    this.pageBlocks.belongsTo(this.banners, { as: 'banner', foreignKey: 'bannerId' });
    this.pageBlocks.belongsTo(this.bannerLists, { as: 'bannerList', foreignKey: 'bannerListId' });
    this.pageBlocks.belongsTo(this.pages, { as: 'target', foreignKey: 'targetId' });
    this.pageBlocks.belongsTo(this.gameGroups, { as: 'group', foreignKey: 'groupId' });
    translation.updateTransMap('cms_pageblock', [{'template': 'cms_template'}, {'gameList': 'cms_gamelist'}, {'banner': 'cms_banner'}, {'bannerList': 'cms_bannerlist'}, {'target': 'cms_page'}, {'group': 'cms_gamegroup'}]);

    this.pages.belongsToMany(this.pageBlocks, {
      as: 'blocks', through: {
        model: this.pagePageblocks,
        unique: false,
      }, foreignKey: 'pageId'
    });
    translation.updateTransMap('cms_page', [{'blocks': 'cms_pageblock'}]);
    this.pageBlocks.belongsToMany(this.pages, {
      as: 'pageSet', through: {
        model: this.pagePageblocks,
        unique: false,
      }, foreignKey: 'pageblock_id'
    });
    translation.updateTransMap('cms_pageblock', [{'pageSet': 'cms_page'}]);
  }

  // ================ definition of cms_* ================
  get tags() {
    this.#tags = this.#tags || this.buildupModel('cms_tag', {
      id: {
        type: SQL.UUID,
        primaryKey: true,
      },
      name: SQL.STRING,
    }, ['name']);

    return this.#tags;
  }

  get categories() {
    this.#categories = this.#categories || this.buildupModel('cms_category', {
      id: {
        type: SQL.UUID,
        primaryKey: true,
      },
      name: SQL.STRING,
      referent: SQL.STRING,
      order: SQL.INTEGER,
    }, ['name']);

    return this.#categories;
  }

  get photos() {
    this.#photos = this.#photos || this.buildupModel('cms_photo', {
      id: {
        type: SQL.UUID,
        primaryKey: true,
      },
      name: SQL.STRING,
      image: SQL.STRING,
      url: SQL.STRING,
    }, ['image', 'url']);

    return this.#photos;
  }

  get memberLevels() {
    this.#memberLevels = this.#memberLevels || this.buildupModel('cms_memberlevel', {
      id: {
        type: SQL.UUID,
        primaryKey: true,
      },
      name: SQL.STRING,
    }, ['name']);

    return this.#memberLevels;
  }

  get gameGroups() {
    this.#gameGroups = this.#gameGroups || this.buildupModel('cms_gamegroup', {
      id: {
        type: SQL.UUID,
        primaryKey: true,
      },
      name: SQL.STRING,
      priority: SQL.INTEGER,
    }, ['name']);

    return this.#gameGroups;
  }

  get ratings() {
    this.#ratings = this.#ratings || this.buildupModel('cms_rating', {
      id: {
        type: SQL.UUID,
        primaryKey: true,
      },
      name: SQL.STRING,
      age: SQL.INTEGER,
    }, ['name']);

    return this.#ratings;
  }

  get pegiRatings() {
    this.#pegiRatings = this.#pegiRatings || this.buildupModel('cms_pegirating', {
      id: {
        type: SQL.UUID,
        primaryKey: true,
      },
      name: SQL.STRING,
      age: SQL.INTEGER,
      referent: SQL.STRING,
    });

    return this.#pegiRatings;
  }

  get locales() {
    this.#locales = this.#locales || this.buildupModel('cms_gamelocale', {
      id: {
        type: SQL.UUID,
        primaryKey: true,
      },
      name: SQL.STRING,
      label: SQL.STRING,
    });

    return this.#locales;
  }

  get projects() {
    this.#projects = this.#projects || this.buildupModel('cms_project', {
      id: {
        type: SQL.UUID,
        primaryKey: true,
      },
      name: SQL.STRING,
      label: SQL.STRING,
      target: SQL.STRING,
    });

    return this.#projects;
  }

  get games() {
    this.#games = this.#games || this.buildupModel('cms_game', {
      id: {
        type: SQL.UUID,
        primaryKey: true,
      },
      label: SQL.STRING,
      author: SQL.STRING,
      name: SQL.STRING,
      video: SQL.STRING,
      description: SQL.STRING,
      abstract: SQL.STRING,
      purchaseType: SQL.STRING,
      score: SQL.STRING,
      pubDate: SQL.DATE,
      endDate: SQL.DATE,
      multiPlay: SQL.INTEGER,
      launchBase: SQL.INTEGER,
      weekLaunchBase: SQL.INTEGER,
      monthLaunchBase: SQL.INTEGER,
      launch: SQL.INTEGER,
      weekLaunch: SQL.INTEGER,
      monthLaunch: SQL.INTEGER,
      favoriteCount: SQL.INTEGER,
      isKeyboardMouseControl: SQL.BOOLEAN,
      isJoystickControl: SQL.BOOLEAN,
      isTouchScreenControl: SQL.BOOLEAN,
      includeInGameBundle: SQL.BOOLEAN,
      isPlayable: SQL.BOOLEAN,
      isGryoControl: SQL.BOOLEAN,
      inputMethod: SQL.INTEGER,
      isPortrait: SQL.BOOLEAN,
    }, ['author', 'name', 'video', 'description', 'abstract']);

    return this.#games;
  }

  get templates() {
    this.#templates = this.#templates || this.buildupModel('cms_template', {
      id: {
        type: SQL.UUID,
        primaryKey: true,
      },
      name: SQL.STRING,
    });

    return this.#templates;
  }

  get gameLists() {
    this.#gameLists = this.#gameLists || this.buildupModel('cms_gamelist', {
      id: {
        type: SQL.UUID,
        primaryKey: true,
      },
      name: SQL.STRING,
    });

    return this.#gameLists;
  }

  get gamelistGames() {
    this.#gamelistGames = this.#gamelistGames || this.buildupModel('cms_gamelist_games', {
      id: {
        type: SQL.INTEGER,
        primaryKey: true,
        autoIncrement: true
      },
      sortValue: SQL.INTEGER,
      gamelistId: SQL.UUID,
      gameId: SQL.UUID,
    });

    return this.#gamelistGames;
  }

  get banners() {
    this.#banners = this.#banners || this.buildupModel('cms_banner', {
      id: {
        type: SQL.UUID,
        primaryKey: true,
      },
      title: SQL.STRING,
      image: SQL.STRING,
      video: SQL.STRING,
      targetType: {
        type: SQL.INTEGER,
        get() {
          return 'A_' + this.getDataValue('targetType');
        }
      },
      url: SQL.STRING,
    }, ['title', 'url']);

    return this.#banners;
  }

  get rankings() {
    this.#rankings = this.#rankings || this.buildupModel('cms_ranking', {
      id: {
        type: SQL.UUID,
        primaryKey: true,
      },
      name: SQL.STRING,
    });

    return this.#rankings;
  }

  get rankingGames() {
    this.#rankingGames = this.#rankingGames || this.buildupModel('cms_ranking_games', {
      id: {
        type: SQL.INTEGER,
        primaryKey: true,
        autoIncrement: true
      },
      sortValue: SQL.INTEGER,
      rankingId: SQL.UUID,
      gameId: SQL.UUID,
    });

    return this.#rankingGames;
  }

  get bannerLists() {
    this.#bannerLists = this.#bannerLists || this.buildupModel('cms_bannerlist', {
      id: {
        type: SQL.UUID,
        primaryKey: true,
      },
      name: SQL.STRING,
    });

    return this.#bannerLists;
  }

  get bannerlistBanners() {
    this.#bannerlistBanners = this.#bannerlistBanners || this.buildupModel('cms_bannerlist_banners', {
      id: {
        type: SQL.INTEGER,
        primaryKey: true,
        autoIncrement: true
      },
      sortValue: SQL.INTEGER,
      bannerlistId: SQL.UUID,
      bannerId: SQL.UUID,
    });

    return this.#bannerlistBanners;
  }

  get pageBlocks() {
    this.#pageBlocks = this.#pageBlocks || this.buildupModel('cms_pageblock', {
      id: {
        type: SQL.UUID,
        primaryKey: true,
      },
      title: SQL.STRING,
    }, ['title']);

    return this.#pageBlocks;
  }

  get pages() {
    this.#pages = this.#pages || this.buildupModel('cms_page', {
      id: {
        type: SQL.UUID,
        primaryKey: true,
      },
      title: SQL.STRING,
    }, ['title']);

    return this.#pages;
  }

  get pagePageblocks() {
    this.#pagePageblocks = this.#pagePageblocks || this.buildupModel('cms_page_blocks', {
      id: {
        type: SQL.INTEGER,
        primaryKey: true,
        autoIncrement: true
      },
      sortValue: SQL.INTEGER,
      pageId: SQL.UUID,
      pageblockId: SQL.UUID,
    });

    return this.#pagePageblocks;
  }

  get configs() {
    this.#configs = this.#configs || this.buildupModel('cms_config', {
      id: {
        type: SQL.UUID,
        primaryKey: true,
      },
      name: SQL.STRING,
      value: SQL.STRING,
    });

    return this.#configs;
  }

  // ================ definition of auth_* ================
  get authUsers() {
    this.#authUsers = this.#authUsers || this.buildupModel('auth_user', {
      id: {
        type: SQL.INTEGER,
        primaryKey: true,
        autoIncrement: true
      },
      password: SQL.STRING,
      lastLogin: SQL.DATE,
      isSuperuser: SQL.BOOLEAN,
      username: SQL.STRING,
      firstName: SQL.STRING,
      lastName: SQL.STRING,
      email: SQL.INTEGER,
      isStaff: SQL.BOOLEAN,
      isActive: SQL.BOOLEAN,
      dateJoined: SQL.DATE,
    });

    return this.#authUsers;
  }

  // ================ definition of cms_* ================
  get users() {
    this.#users = this.#users || this.buildupModel('ams_userprofile', {
      id: {
        type: SQL.UUID,
        primaryKey: true,
      },
      ubilive: SQL.STRING,
      identification: SQL.STRING,
      player_id: SQL.STRING,
      nickname: SQL.STRING,
      gender: {
        type: SQL.INTEGER,
        get() {
          const value = this.getDataValue('gender');
          if(value !== null) {
            return 'A_' + value;
          }
          return value;
        }
      },
      signature: SQL.STRING,
      region: SQL.STRING,
      mobile: SQL.STRING,
      email: SQL.STRING,
      laboite: SQL.STRING,
      pincode: SQL.STRING,
      birthday: SQL.DATE,
      age: SQL.INTEGER,
      rating_uuid: {
        type: SQL.VIRTUAL,
        get() {
          var rate;
          if (settings.RATING_SPEC === "pegi")
            rate = this.ratingPegi;
          else
            rate = this.rating;
          return rate ? rate.id : '';
        }
      },
    });

    return this.#users;
  }

  get avatars() {
    this.#avatars = this.#avatars || this.buildupModel('ams_avatar', {
      id: {
        type: SQL.UUID,
        primaryKey: true,
      },
      name: SQL.STRING,
      image: SQL.STRING,
    });

    return this.#avatars;
  }

  get favorites() {
    this.#favorites = this.#favorites || this.buildupModel('ams_userprofile_favorites', {
      id: {
        type: SQL.INTEGER,
        primaryKey: true,
        autoIncrement: true
      },
      sortValue: SQL.INTEGER,
      userprofileId: SQL.UUID,
      gameId: SQL.UUID,
    });

    return this.#favorites;
  }

  get recentPlays() {
    this.#recentPlays = this.#recentPlays || this.buildupModel('ams_userprofile_recent_play', {
      id: {
        type: SQL.INTEGER,
        primaryKey: true,
        autoIncrement: true
      },
      played: SQL.DATE,
      count: SQL.INTEGER,
      userId: SQL.UUID,
      gameId: SQL.UUID,
    });

    return this.#recentPlays;
  }

  get leaderboard() {
    this.#leaderboard = this.#leaderboard || this.buildupModel('ams_userprofile_recent_play', {
      id: {
        type: SQL.INTEGER,
        primaryKey: true,
        autoIncrement: true
      },
      played: SQL.DATE,
      count: SQL.INTEGER,
      userId: SQL.UUID,
      gameId: SQL.UUID,
    });

    return this.#leaderboard;
  }

  get familyMembers() {
    this.#familyMembers = this.#familyMembers || this.buildupModel('ams_family_members', {
      id: {
        type: SQL.INTEGER,
        primaryKey: true,
        autoIncrement: true
      },
      sortValue: SQL.INTEGER,
      familyId: SQL.UUID,
      userprofileId: SQL.UUID,
    });

    return this.#familyMembers;
  }

  get families() {
    this.#families = this.#families || this.buildupModel('ams_family', {
      id: {
        type: SQL.UUID,
        primaryKey: true,
      },
      ubilive: SQL.STRING,
    });

    return this.#families;
  }
}

class StoreConnection {
  #username
  #password
  #host
  constructor({ username, password, host }) {
    this.#username = username;
    this.#password = password;
    this.#host = host;
  }

  get username() {
    return this.#username;
  }

  get password() {
    return this.#password;
  }

  get host() {
    return this.#host;
  }

  launch(database) {
    throw 'not implemented';
  }
}

class PostgresStoreConnection extends StoreConnection {
  #logging
  constructor({ username, password, host, logging }) {
    super({ username, password, host });
    this.#logging = logging;
  }

  launch(database) {
    return new SQL({
      database: database,
      username: this.username,
      password: this.password,
      dialect: 'postgres',
      host: this.host,
      logging: this.#logging,
      define: {
        charset: 'utf8',
        collate: 'utf8_general_ci',
        freezeTableName: true,
        timestamps: false,
        underscored: true,
      }
    });
  }
}

class StoreManager {
  #connections
  #dbconn
  #storeCls
  constructor(dbconn, storeCls) {
    this.#connections = new Map();
    this.#dbconn = dbconn;
    this.#storeCls = storeCls;
  }

  i18nName(name) {
    const Langs = translation.Langs; // [['en-us', 'EnUs'], ['zh-hk', 'ZhHk'], ['zh-tw', 'ZhTw']]
    const i18nMap = new Map(Langs);

    return i18nMap.get((name || '').toLowerCase());
  }

  get(database, isPreview = false) {
    if (!this.#connections.has(database)) {
      const db = this.#dbconn.launch(database);

      const store = this.#storeCls.create(db);
      store.buildRelation(isPreview);
      this.#connections.set(database, store);
    }

    return this.#connections.get(database);
  }
}

module.exports = { StoreManager, PostgresStoreConnection, StoreConnection, UMSStoreObjects, StoreObjects };
