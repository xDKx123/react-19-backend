package models

import "testProject/src/middleware"

type User struct {
	BasicModel
	Name     string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);not null"`
	Password string `gorm:"type:varchar(100);not null" json:"-"`
	Salt     string `gorm:"type:varchar(100);not null" json:"-"`
}

func (u User) CheckPassword(password string) bool {
	return middleware.CheckPasswordHash(password, u.Salt, u.Password)
}

func (u User) CreateUser(password string) {
	salt, err := middleware.GenerateSalt()

	if err != nil {
		panic(err)
	}

	u.Salt = salt
	u.Password, _ = middleware.HashPassword(password, salt)
}
