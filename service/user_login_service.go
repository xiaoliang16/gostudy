package service

import (
	"gostudy/model"
	"gostudy/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required, min=5, max=30"`
	Password string `form:"password" json:"password" binding:"required, min=8, max=40"`
}

func (service *UserLoginService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	var user model.User

	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
         return serializer.ParamErr("", nil)
	}

	service.setSession(c, user)

	return serializer.BuildUserResponse(user)
}