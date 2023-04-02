package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Session(secret string) gin.HandlerFunc {
    store := cookie.NewStore([]byte(secret))
    store.Options(sessions.Options{HttpOnly:true, MaxAge: 7* 10000, Path:"/"})
    return sessions.Sessions("gin-session", store)
}
