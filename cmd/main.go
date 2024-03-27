package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/Kimoto-Norihiro/access-control-system/controller"
	"github.com/Kimoto-Norihiro/access-control-system/database"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	controller := controller.NewController(db)

	r := gin.Default()
	// cors
	r.Use(cors.Default())

	// ユーザー登録
	r.POST("/user", controller.CreateUser)
	// 入室
	r.POST("/entry", controller.Entry)
	// 退室
	r.PUT("/exit", controller.Exit)

	r.Run(":4040")
}
