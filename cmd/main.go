package main

import (
	"todo/internal/config"
	"todo/internal/database"
	"todo/internal/router"
)

func main() {
	//init config
	config := config.LoadConfig("./config/config.json")
	dbConf := config.DataBase

	//init DataBase
	dataBase := database.New(dbConf)

	//init routers
	router := router.SetupRoutes(dataBase)

	//run server
	router.Run(":" + config.Port)

}
