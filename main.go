package main

import "rest-api/routers"

func main() {
	routers.StartRouter().Run(":8080")
}
