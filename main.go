package main

import (
	"belajar-jwt/database"
	"belajar-jwt/router"
)

func main() {
	database.StartDB()

	router.New().Run(":3000")
}
