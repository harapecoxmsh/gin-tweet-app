package main

import (
	"gin-tweet-app/infra"
	"gin-tweet-app/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	if err := db.AutoMigrate(&models.Post{}); err != nil {
		panic("Faild to migrate database")
	}

}
