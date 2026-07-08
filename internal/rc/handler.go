package rc

import (
	"erro
	"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) ListActive(c *gin.Context) {
	resp, err := h.service.ListActive(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetProforma(c *gin.Context) {
	// Get the recruitment cycle ID from the URL parameter
	rcID := c.Param("id")
	resp, err := h.service.repo.GetProformasByRecruitmentCycleID(c.Request.Context(), rcID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetProformaByRole returns the proformas assigned to the authenticated user (apc/coco)
// Requires the AuthMiddleware to have set "userID" in the gin context.
func (h *Handler) GetProformaByRole(c *gin.Context) {
	rcID := c.Param("id")
	role := c.Param("role")

	// Validate role early for a clear error message
	if role != "apc" && role != "coco" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "role must be 'apc' or 'coco'"})
		return
	}

	// userID is set by AuthMiddleware
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing user identity"})
		return
	}

	uid, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user identity"})
		return
	}

	proformas, err := h.service.GetProformasByRole(c.Request.Context(), rcID, uid, role)
	if err != nil {
		if errors.Is(err, ErrInvalidRole) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, proformas)
}}		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, resp)
}		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, resp)
}