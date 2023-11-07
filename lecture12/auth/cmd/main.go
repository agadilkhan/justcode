package main

import (
	"lecture12/auth/internal/app"
	"log"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Printf("app run err: %v", err)
	}
}
