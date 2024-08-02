package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	route_v1 "telegram-bot/routes/v1"
)

func main() {

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	app := gin.Default()

	api := app.Group("/api")
	route_v1.Route(&gin.Context{}, api)

	app.Run(":9090")
}
