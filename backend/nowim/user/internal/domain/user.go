package domain

type User struct {
	UserID   int64  `json:"userID" gorm:"column:userID;AUTO_INCREMENT;PRIMARY_KEY"`
	Username string `json:"username" gorm:"column:username;UNIQUE;index:username_index"`
	Password string `json:"password" gorm:"column:password"`
}

func NewUser(userID int64, username string, password string) *User {
	return &User{UserID: userID, Username: username, Password: password}
}
