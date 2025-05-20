package main

import (
	"karting-challenge/router"
	"log"
)

func main() {
	r := router.SetupRouter()
	log.Println("🚀 Server running on :8080")
	log.Fatal(r.Run(":8080"))
}
