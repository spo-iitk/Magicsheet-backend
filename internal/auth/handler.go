package auth

import (
	"net/http"
	"os"
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


func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "invalid request body",
		})
		return
	}

	resp, err := h.service.Login(c.Request.Context(), req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	}

	maxAgeSeconds := 45 * 60
	if envMinutes := os.Getenv("ACCESS_TOKEN_EXPIRE_MINUTES"); envMinutes != "" {
		if minutes, parseErr := strconv.Atoi(envMinutes); parseErr == nil && minutes > 0 {
			maxAgeSeconds = minutes * 60
		}
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("access_token", resp.AccessToken, maxAgeSeconds, "/", "", false, true)

	c.JSON(http.StatusOK, resp)

}

func (h *Handler) Logout(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("access_token", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message" : "logout endpoint reached",
	})
}

func (h *Handler) Me(c *gin.Context) {
	userID := c.GetUint("userID")

	resp, err := h.service.Me(c.Request.Context(), userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) CreateUser(c *gin.Context) {
	var req CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	resp, err := h.service.CreateUser(c.Request.Context(), req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}
