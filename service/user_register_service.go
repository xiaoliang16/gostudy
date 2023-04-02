package service

import (
	"gostudy/model"
	"gostudy/serializer"
)

type UserRegisterService struct {
	Nickname string `form:"nickname" json:"nickname" binding:"required, min=2, max=30"`
    UserName string `form:"user_name" json:"user_name" binding:"required, min=5, max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

func (service *UserRegisterService) valid() *serializer.Response {
	if service.Password != service.PasswordConfirm {
		return &serializer.Response{
			Code:  40001,
			Msg:   "两次密码不同",
		}
	}

	count := int64(0)
	model.DB.Model(&model.User{}).Where("nickname = ?", service.Nickname).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code:  40002,
			Msg:   "",
		}
	}

	count = 0
	model.DB.Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code:  40003,
			Msg:   "",
		}
	}

	return nil
}

func (service *UserRegisterService) Register() serializer.Response {
	user := model.User{
		UserName: service.UserName,
		Status:   model.Active,
	}

	if err := service.valid(); err != nil {
         return *err
	}

	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.ParamErr("注册失败", err)
	}

	return serializer.BuildUserResponse(user)
}