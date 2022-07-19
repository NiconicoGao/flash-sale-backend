package router

import (
	"flash-sale-backend/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/ping", handler.Default)
		api.POST("/upload", handler.Upload)
		api.POST("/item", handler.AddItemHandler)
		api.GET("/spuinfo", handler.GetAllItemHandler)
		api.POST("/activity", handler.AddActivityHandler)
		api.GET("/activity", handler.GetActivityInfoHandler)
		api.GET("/order", handler.PlaceOrderHandler)
	}

	r.StaticFS("/api/image", http.Dir("./upload"))

}
