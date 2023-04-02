package conf

import (
	"fmt"
	"gostudy/model"
)

func Init() {
	s := "root:admin@/blog?charset=utf8&parseTime=True&loc=Local"
	fmt.Println(s)
	model.Database(s)
}
