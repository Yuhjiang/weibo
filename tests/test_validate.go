package main

import (
	"github.com/Yuhjiang/weibo/utils"
	"github.com/go-playground/validator/v10"
)

type Login struct {
	Username string `validate:"required,min=10,max=20"`
	Password string `validate:"required,min=10,max=20"`
}

func main() {
	login := Login{"hello", "password"}
	err := utils.Validate.Struct(login).(validator.ValidationErrors)
	if err != nil {
		utils.HandleValidateError(err)
	}
}
