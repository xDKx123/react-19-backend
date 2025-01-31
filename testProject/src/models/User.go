package models

import "testProject/src/middleware"

type User struct {
	BasicModel
	Name     string `gorm:"type:varchar(100);not null"`
	Surname  string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);not null"`
	Password string `gorm:"type:varchar(100);not null" json:"-"`
	Salt     string `gorm:"type:varchar(100);not null" json:"-"`
}

func (u User) CheckPassword(password string) bool {
	return middleware.CheckPasswordHash(password, u.Salt, u.Password)
}
