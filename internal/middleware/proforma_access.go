package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProformaAccessChecker interface {
	HasProformaAccess(userID uint, proformaID uint) (bool, error)
}

func RequireProformaAccess(checker ProformaAccessChecker) gin.HandlerFunc {
	return func(c *gin.Context) {

		userID := c.GetUint("userID")

		proformaIDStr := c.Param("proformaID")

		proformaID64, err := strconv.ParseUint(proformaIDStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid proforma id",
			})
			c.Abort()
			return
		}

		proformaID := uint(proformaID64)

		hasAccess, err := checker.HasProformaAccess(userID, proformaID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to verify access",
			})
			c.Abort()
			return
		}

		if !hasAccess {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "access denied",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
