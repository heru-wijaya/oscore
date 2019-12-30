package main

import (
	"ovo-score/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/v1/oscore/entity", controller.CreateEntityHandler)
	router.Run("localhost:8000")
}
