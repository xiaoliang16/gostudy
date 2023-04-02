package serializer

import "gostudy/model"

type User struct {
	ID uint `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
}

func BuildUser(user model.User) User {
	return User {
		ID: user.ID,
		UserName: user.UserName,
		NickName: user.UserName,
	}
}

func BuildUserResponse(user model.User) Response {
	return Response{
		Data:  BuildUser(user),
	}
}