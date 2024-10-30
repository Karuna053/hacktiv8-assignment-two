package main

import (
	"rest-api/database"
	"rest-api/routers"
)

func main() {
	database.InitDB()
	routers.StartRouter().Run(":8080")
}
