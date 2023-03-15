package main

import (
	"fmt"

	"github.com/sanyewudezhuzi/tiktok/conf"
	"github.com/sanyewudezhuzi/tiktok/model"
	"github.com/sanyewudezhuzi/tiktok/router"
)

func init() {
	conf.LoadEnvironment()
	model.Mysqlini()
	// model.AutomigrateMySQL()
	fmt.Println("continue")
}

func main() {
	fmt.Println("hello tiktok")

	r := router.Router()
	r.Run(conf.Port)
}
