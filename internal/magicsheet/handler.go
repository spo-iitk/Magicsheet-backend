package magicsheet

import "github.com/gin-gonic/gin"

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetMagicSheet(c *gin.Context) {}

func (h *Handler) RegisterCandidate(c *gin.Context) {}

func (h *Handler) CheckIn(c *gin.Context) {}

func (h *Handler) CheckOut(c *gin.Context) {}

func (h *Handler) UpdateSessionResult(c *gin.Context) {}

func (h *Handler) CreateRound(c *gin.Context) {}
