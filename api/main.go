package api

import (
	"gostudy/model"
	"gostudy/serializer"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(200 , serializer.Response{
		Code:  0,
		Msg:   "pong",
	})
}

func CurrentUser(c *gin.Context) *model.User {
   if user, _ := c.Get("user"); user != nil {
   	  if u, ok := user.(*model.User); ok {
   	  	return u
	  }
   }
   return nil
}

