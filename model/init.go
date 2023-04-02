package model

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Database(connString string) {
   newLogger := logger.New(
   	  log.New(os.Stdout, "\r\n", log.LstdFlags),
   	  logger.Config{
		  SlowThreshold:             time.Second, // Slow SQL threshold
		  LogLevel:                  logger.Info, // Log level(这里记得根据需求改一下)
		  IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
		  Colorful:                  false,
	  },
   	)

   db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
   	Logger: newLogger,
   })

   if err != nil {
   	 panic(any("连接异常"))
   }

   sqlDB, err := db.DB()


   sqlDB.SetMaxIdleConns(10)
   sqlDB.SetMaxOpenConns(20)

   DB = db
   migration()
}


