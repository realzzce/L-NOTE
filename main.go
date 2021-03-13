package main

import (
	"other/L-NOTE/models"
	"other/L-NOTE/router"
)

func main() {

	models.InitConnectDB()
	router.InitRouter()
}
