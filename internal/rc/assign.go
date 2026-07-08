package rc

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/Magicsheet-backend/internal/database"
)


// GetProformaByIDAndCycle fetches a proforma only if it belongs to the given RC.
func (r *Repository) GetProformaByIDAndCycle(ctx context.Context, proformaID uint, rcID uint) (*database.Proforma, error) {
	var p database.Proforma
	err := r.db.WithContext(ctx).
		Where("id = ? AND recruitment_cycle_id = ?", proformaID, rcID).
		First(&p).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetUserByEmailAndRole fetches a user only if their email and role both match.
func (r *Repository) GetUserByEmailAndRole(ctx context.Context, email string, role database.UserRole) (*database.User, error) {
	var u database.User
	err := r.db.WithContext(ctx).
		Where("email = ? AND role = ?", email, role).
		First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// GetUnassignedUsersByRole returns all active users with the given role that have
// no active CoordinatorAssignment for the specified proforma.
func (r *Repository) GetUnassignedUsersByRole(ctx context.Context, proformaID uint, role database.UserRole) ([]database.User, error) {
	var users []database.User
	err := r.db.WithContext(ctx).
		Where("role = ? AND is_active = ?", role, true).
		Where("id NOT IN (?)",
			r.db.Model(&database.CoordinatorAssignment{}).
				Select("user_id").
				Where("proforma_id = ? AND is_active = ?", proformaID, true),
		).
		Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetAssignment returns an existing CoordinatorAssignment for (proforma, user, role), or nil if not found.
func (r *Repository) GetAssignment(ctx context.Context, proformaID, userID uint, role database.AssignmentRole) (*database.CoordinatorAssignment, error) {
	var a database.CoordinatorAssignment
	err := r.db.WithContext(ctx).
		Where("proforma_id = ? AND user_id = ? AND role = ?", proformaID, userID, role).
		First(&a).Error
	if err != nil {
		return nil, err
	}
	return &a, nil
}

// CreateAssignment inserts a new CoordinatorAssignment row.
func (r *Repository) CreateAssignment(ctx context.Context, a *database.CoordinatorAssignment) error {
	return r.db.WithContext(ctx).Create(a).Error
}

// UpdateAssignment persists changes to an existing CoordinatorAssignment.
func (r *Repository) UpdateAssignment(ctx context.Context, a *database.CoordinatorAssignment) error {
	return r.db.WithContext(ctx).Save(a).Error
}


// GetUnassignedUsers handles GET /rc/:id/unassigned/:role
// Returns users with the given role not yet assigned to the proforma (passed as query param).
func (h *Handler) GetUnassignedUsers(c *gin.Context) {
	rcIDStr := c.Param("id")
	roleStr := c.Param("role")
	proformaIDStr := c.Query("proforma_id")

	var rcID, proformaID uint
	if _, err := fmt.Sscan(rcIDStr, &rcID); err != nil || rcID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid recruitment cycle id"})
		return
	}
	if _, err := fmt.Sscan(proformaIDStr, &proformaID); err != nil || proformaID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "proforma_id query param required"})
		return
	}

	var role database.UserRole
	switch roleStr {
	case "apc":
		role = database.RoleApc
	case "coco":
		role = database.RoleCoco
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "role must be 'apc' or 'coco'"})
		return
	}

	// Confirm proforma belongs to this RC
	if _, err := h.service.repo.GetProformaByIDAndCycle(c.Request.Context(), proformaID, rcID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "proforma does not belong to this recruitment cycle"})
		return
	}

	users, err := h.service.repo.GetUnassignedUsersByRole(c.Request.Context(), proformaID, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	opts := make([]UserOption, 0, len(users))
	for _, u := range users {
		opts = append(opts, UserOption{ID: u.ID, Name: u.Name, Email: u.Email})
	}
	c.JSON(http.StatusOK, opts)
}

// AssignMagicsheet handles POST /rc/:id/assign
// opc can assign apc or coco; apc can only assign coco.
func (h *Handler) AssignMagicsheet(c *gin.Context) {
	rcIDStr := c.Param("id")
	var rcID uint
	if _, err := fmt.Sscan(rcIDStr, &rcID); err != nil || rcID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid recruitment cycle id"})
		return
	}

	callerID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing user identity"})
		return
	}
	uid, ok := callerID.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user identity"})
		return
	}

	callerRole, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing role"})
		return
	}
	role, ok := callerRole.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid role"})
		return
	}

	var req AssignMagicsheetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.AssignMagicsheet(c.Request.Context(), rcID, uid, role, req)
	if err != nil {
		switch err {
		case ErrInvalidRole, ErrForbiddenAssign, ErrProformaNotInRC, ErrTargetUserInvalid:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, resp)
}


func toAssignResponse(a *database.CoordinatorAssignment) *AssignMagicsheetResponse {
	return &AssignMagicsheetResponse{
		AssignmentID: a.ID,
		ProformaID:   a.ProformaID,
		AssignedToID: a.UserID,
		Role:         string(a.Role),
		IsActive:     a.IsActive,
	}
}
