package models

import (
	"time"
)

type Comment struct {
	ID		int		`gorm:"AUTO_INCREMENT,primary_key"`
	Body 		string
	PostID		int
	UserID 		int64
	User		User
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}
