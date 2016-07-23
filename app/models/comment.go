package models

import (
	"time"
)

type Comment struct {
	ID		int		`gorm:"AUTO_INCREMENT,primary_key"`
	Body 		string                `json:"body"`
	PostID		int
	UserID 		int64
	User		User                `json:"user"`
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}
