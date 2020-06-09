package domain

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"nowim.user/internal/db"
)
import "github.com/jinzhu/gorm"

type PostgresUserRepo struct {
	db *gorm.DB
}

func NewPostgresUserRepo() *PostgresUserRepo {
	return &PostgresUserRepo{db: db.DB().AutoMigrate(&User{})}
}

func (p PostgresUserRepo) AddUser(user *User) error {
	return p.db.Create(user).Error
}

func (p PostgresUserRepo) UserByName(username string) (*User, error) {
	var user User
	err := p.db.Where("username = ?", username).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	return &user, err
}

func (p PostgresUserRepo) UserByID(id int64) (*User, error) {
	var user User
	err := p.db.Where("userID = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	return &user, nil
}

func (p PostgresUserRepo) AllUsers() (map[int64]*User, error) {
	users := make([]*User, 0)
	err := p.db.Find(&users).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	result := make(map[int64]*User)
	for _, u := range users {
		result[u.UserID] = u
	}

	return result, nil
}
