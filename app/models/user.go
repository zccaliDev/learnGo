package models

type User struct {
	ID 		int	`gorm:"AUTO_INCREMENT,primary_key"`
	Name 		string        `json:"name"`
	Email 		string
	Password	string
	Posts		[]Post
	Comment		[]Comment
}


