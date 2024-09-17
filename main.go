package main

import (
	"log"
	"receipt-processor-challenge/controller"
	"receipt-processor-challenge/logic"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	receiptsService := logic.NewReceiptsService()
	receiptsController := controller.NewReceiptsController(receiptsService)
	receiptsController.AddRoutes(router)

	log.Fatal(
		router.Run(),
	)
}
