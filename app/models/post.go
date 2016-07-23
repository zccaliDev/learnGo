package models

import (
	"time"
)

type Post struct {
	ID    		int			`gorm:"AUTO_INCREMENT,primary_key" json:"id"`
	Title 		string                        `json:"title"`
	Body 		string                        `json:"body"`
	UserID		int64
	Comments	[]Comment
	Likes		[]Likes
	Like		int                        `gorm:"-" json:"like"`
	Comment		int                        `gorm:"-" json:"comment"`
	User		User                	`json:"user"`
	CreatedAt time.Time                        `json:"createdAt"`
	UpdatedAt time.Time
	DeletedAt *time.Time 			`sql:"index"`
}
