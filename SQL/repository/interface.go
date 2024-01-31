package repository

import "sql-practice/model"

// RepositoryInterface
type RepositoryInterface interface {
	// CRUD operations for model.User
	CreateUser(user *model.User) error
	GetUserByID(userID uint) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(userID uint) error

	// Additional queries for model.User
	GetUserByUsername(username string) (*model.User, error)
	GetAllUsers() ([]model.User, error)

	// CRUD operations for model.Post
	CreatePost(post *model.Post) error
	GetPostByID(postID uint) (*model.Post, error)
	UpdatePost(post *model.Post) error
	DeletePost(postID uint) error

	// Additional queries for model.Post
	GetPostsByUserID(userID uint) ([]model.Post, error)

	// CRUD operations for model.Comment
	CreateComment(comment *model.Comment) error
	GetCommentByID(commentID uint) (*model.Comment, error)
	UpdateComment(comment *model.Comment) error
	DeleteComment(commentID uint) error

	// Additional queries for model.Comment
	GetCommentsByUserID(userID uint) ([]model.Comment, error)
	GetCommentsByPostID(postID uint) ([]model.Comment, error)
}
