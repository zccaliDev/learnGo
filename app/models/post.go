package models

import (
	"time"
)

type Post struct {
	ID    		int			`gorm:"AUTO_INCREMENT,primary_key"`
	Title 		string
	Body 		string
	UserID		int64
	Comments	[]Comment
	Likes		[]Likes
	User		User                	`json:"user"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time 			`sql:"index"`
}
