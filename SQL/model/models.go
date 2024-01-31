package model

import "time"

type Post struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"not null"`
	Content   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UserID    uint
	Comments  []Comment
}

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"unique;not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Posts     []Post
	Comments  []Comment
}

type Comment struct {
	ID        uint      `gorm:"primaryKey"`
	Content   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UserID    uint
	PostID    uint
}
