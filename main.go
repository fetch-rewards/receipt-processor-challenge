package main

import (
	"log"
	"receipt-processor-challenge/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	receiptsController := controller.ReceiptController{}
	receiptsController.AddRoutes(router)

	log.Fatal(
		router.Run(),
	)
}
