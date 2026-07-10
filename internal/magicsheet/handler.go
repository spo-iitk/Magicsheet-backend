package magicsheet

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetMagicSheet(c *gin.Context) {
	ctx := c.Request.Context()

	proformaID, err := getUintParam(c, "proformaID")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid proforma id",
		})
		return
	}

	response, err := h.service.GetMagicSheet(ctx, proformaID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) RegisterCandidate(c *gin.Context) {
	ctx := c.Request.Context()

	proformaID, err := getUintParam(c, "proformaID")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid proforma id",
		})
		return
	}

	var req RegisterCandidateRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	userID := c.GetUint("userID")
	candidate, err := h.service.RegisterCandidate(ctx, userID, proformaID, req.RollNumber)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, candidate)
}

func (h *Handler) CheckIn(c *gin.Context) {
	ctx := c.Request.Context()

	proformaID, err := getUintParam(c, "proformaID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid proforma id",
		})
		return
	}

	sessionID, err := getUintParam(c, "sessionID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid session id",
		})
		return
	}

	response, err := h.service.CheckIn(ctx, proformaID, sessionID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) CheckOut(c *gin.Context) {
	ctx := c.Request.Context()

	proformaID, err := getUintParam(c, "proformaID")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid proforma id",
		})
		return
	}

	sessionID, err := getUintParam(c, "sessionID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid session id",
		})
		return
	}

	response, err := h.service.CheckOut(ctx, proformaID, sessionID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)

}

func (h *Handler) UpdateSessionResult(c *gin.Context) {
	ctx := c.Request.Context()

	proformaID, err := getUintParam(c, "proformaID")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid proforma id",
		})
		return
	}

	sessionID, err := getUintParam(c, "sessionID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid session id",
		})
		return
	}

	var req UpdateSessionResultRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	response, err := h.service.UpdateSessionResult(ctx, proformaID, sessionID, req.Status)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) CreateRound(c *gin.Context) {
	ctx := c.Request.Context()

	proformaID, err := getUintParam(c, "proformaID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid proforma id",
		})
		return
	}

	var req CreateRoundRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	response, err := h.service.CreateRound(
		ctx,
		proformaID,
		req.Name,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response)
}

func getUintParam(c *gin.Context, key string) (uint, error) {
	value, err := strconv.ParseUint(c.Param(key), 10, 32)

	if err != nil {
		return 0, err
	}

	return uint(value), nil
}
