package rc

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/Magicsheet-backend/internal/auth"
	"github.com/spo-iitk/Magicsheet-backend/internal/middleware"
)

func RegisterRoutes(api *gin.RouterGroup, handler *Handler) {
	r := api.Group("/rc")

	// Public-ish: OPC proforma list and active cycles
	r.GET("/", handler.ListActive)
	r.GET("/:id/opc/proforma", handler.GetProforma)

	// Protected: APC and CoCo role-filtered proforma list
	protected := r.Group("/:id")
	protected.Use(auth.AuthMiddleware())
	protected.Use(middleware.RequireRoles("apc", "coco", "god"))
	protected.GET("/:role/proforma", handler.GetProformaByRole)

	// Protected: assign a proforma to a user (opc or apc only)
	assign := r.Group("/:id")
	assign.Use(auth.AuthMiddleware())
	assign.Use(middleware.RequireRoles("god", "opc", "apc"))
	assign.POST("/assign", handler.AssignMagicsheet)
	assign.GET("/unassigned/:role", handler.GetUnassignedUsers)
}
