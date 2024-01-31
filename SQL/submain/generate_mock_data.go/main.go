package main

import (
	"fmt"
	"math/rand"
	"sql-practice/model"
	"sql-practice/pkg/orm"

	generate_random_string "github.com/skyrocketOoO/go-utility/generate_random_string"

	"gorm.io/gorm"
)

func main() {
	db := orm.GetDB()

	numUsers := 10
	numPosts := numUsers * 7
	numComments := numUsers * 11

	createUser(db, numUsers)
	for i := 0; i < numPosts; i++ {
		createPosts(db, rand.Intn(numUsers)+1)
	}
	for i := 0; i < numComments; i++ {
		createComments(db, rand.Intn(numUsers)+1, rand.Intn(numPosts)+1)
	}

	// Close the database connection
	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to close the database connection")
	}
	sqlDB.Close()
}

func createUser(db *gorm.DB, numUsers int) {
	for i := 0; i < numUsers; i++ {
		user := model.User{
			Username: fmt.Sprintf("user%d", i),
		}

		db.Create(&user)
	}
}

func createPosts(db *gorm.DB, userID int) {
	db.Create(&model.Post{
		Title:   generate_random_string.RandStringBytesMaskImprSrcUnsafe(10),
		Content: generate_random_string.RandStringBytesMaskImprSrcUnsafe(20),
		UserID:  uint(userID),
	})

}

func createComments(db *gorm.DB, userID int, postID int) {
	db.Create(&model.Comment{
		Content: generate_random_string.RandStringBytesMaskImprSrcUnsafe(15),
		UserID:  uint(userID),
		PostID:  uint(postID),
	})
}
