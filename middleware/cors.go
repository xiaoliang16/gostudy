package middleware

import (
	"regexp"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
   config := cors.DefaultConfig()

   config.AllowMethods = []string{"GET", "POST"}
   config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}

   if gin.Mode() == gin.ReleaseMode {
   	  config.AllowOrigins = []string{"http://www.example.com"}
   } else {
   	  config.AllowOriginFunc = func(origin string) bool {
		  if regexp.MustCompile(`^http://127\.0\.0\.1:\d+$`).MatchString(origin) {
		  	 return true
		  }
		  if regexp.MustCompile(`^http://localhost:\d+$`).MatchString(origin) {
		  	return  true
		  }
		  return false
	  }
   }

   config.AllowCredentials = true
   return cors.New(config)
}
