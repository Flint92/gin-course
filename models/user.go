package models

import (
	"gin-course/dao"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (User) TableName() string {
	return "user"
}

func GetUserById(id int) (*User, error) {
	var user User
	err := dao.Db.First(&user, "id=?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func AddUser(name string) (*User, error) {
	user := User{
		Name: name,
	}
	err := dao.Db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
