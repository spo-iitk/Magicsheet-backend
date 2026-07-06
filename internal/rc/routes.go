package rc

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup, handler *Handler) {
	r := api.Group("/rc")
	r.GET("/", handler.ListActive)
	r.GET("/:id/opc/proforma",handler.GetProforma)
}
