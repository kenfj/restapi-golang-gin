package main

import (
	"example/web-service-gin/app"
	"log"
)

func main() {
	router := app.SetRouter()

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
