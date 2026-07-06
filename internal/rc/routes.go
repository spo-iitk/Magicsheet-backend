package rc

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/Magicsheet-backend/internal/auth"
	"github.com/spo-iitk/Magicsheet-backend/internal/middleware"
)

func RegisterRoutes(api *gin.RouterGroup, handler *Handler) {
	r := api.Group("/rc")

	// Public-ish: OPC proforma list and active cycles (uses its own auth elsewhere)
	r.GET("/", handler.ListActive)
	r.GET("/:id/opc/proforma", handler.GetProforma)

	// Protected: APC and CoCo role-filtered proforma list
	// GET /api/rc/:id/:role/proforma  (role = "apc" | "coco")
	protected := r.Group("/:id")
	protected.Use(auth.AuthMiddleware())
	protected.Use(middleware.RequireRoles("apc", "coco", "god"))
	protected.GET("/:role/proforma", handler.GetProformaByRole)
}
