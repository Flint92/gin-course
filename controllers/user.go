package controllers

import (
	"gin-course/models"
	"strconv"
)

type UserController struct{}

func (userController *UserController) UserInfo(id string) *JsonStruct {
	idInt, _ := strconv.Atoi(id)
	user, _ := models.GetUserById(idInt)
	return ReturnSuccess(user)
}

func (userController *UserController) AddUser(name string) *JsonStruct {
	user, _ := models.AddUser(name)
	return ReturnSuccess(user)
}
