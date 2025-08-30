package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-dilve/config"
	"github.com/go-dilve/routes"
)

func init() {
	err := godotenv.Load()
		if err != nil {
		log.Println("No .env file found, using system environment variables")
	}
	config.ConnectDB()
}
func main() {
	r := gin.Default()

	routes.Routes(r)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
