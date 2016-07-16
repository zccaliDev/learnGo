package models

type Likes struct {
	ID		int		`gorm:"AUTO_INCREMENT,primary_key"`
	PostID		int
	UserID 		int64
}
