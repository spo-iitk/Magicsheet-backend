package magicsheet

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/Magicsheet-backend/internal/auth"
	"github.com/spo-iitk/Magicsheet-backend/internal/middleware"
)

func RegisterRoutes(router *gin.RouterGroup, handler *Handler, accessChecker middleware.ProformaAccessChecker) {

	protected := router.Group("/magicsheet/:proformaID")
	protected.Use(auth.AuthMiddleware())

	protected.Use(middleware.RequireProformaAccess(accessChecker))

	protected.GET("", handler.GetMagicSheet)
	protected.POST("/register", handler.RegisterCandidate)
	protected.POST("/sessions/:sessionID/check-in", handler.CheckIn)
	protected.POST("/sessions/:sessionID/check-out", handler.CheckOut)
	protected.POST("/sessions/:sessionID/result", handler.UpdateSessionResult)

	protected.POST("/rounds", handler.CreateRound)
	protected.PATCH("/rounds/:roundID", handler.UpdateRound)
	protected.POST(
    "/candidates/:candidateID/rounds/:roundID/check-in",
    handler.CreateAndCheckIn,
)
}
