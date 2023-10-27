package main

import (
	"log"

	"github.com/abinay-ps/ginframeworkexample/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	handlers.CallHandlers(router)
	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}
