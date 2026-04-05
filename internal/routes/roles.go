package routes

import (
	"net/http"

	"github.com/Ahmed-Armaan/finance-dashboard-api.git/internal/cache"
	"github.com/Ahmed-Armaan/finance-dashboard-api.git/internal/database"
	"github.com/Ahmed-Armaan/finance-dashboard-api.git/internal/database/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RoleChangeRequest struct {
	RequestedRole models.UserRole `json:"requested_role" binding:"required"`
}

type updateRoleRequest struct {
	UserID uuid.UUID       `json:"user_id" binding:"required"`
	Role   models.UserRole `json:"role" binding:"required"`
}

func RequestRoleChange(db database.DatabaseStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RoleChangeRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validRoles := map[models.UserRole]bool{
			models.Viewer:  true,
			models.Analyst: true,
			models.Admin:   true,
		}
		if !validRoles[req.RequestedRole] {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid role requested"})
			return
		}

		userID := c.MustGet("userID").(uuid.UUID)

		if err := db.CreateRoleRequest(userID, req.RequestedRole); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "could not submit role request"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "role change request submitted"})
	}
}

func UpdateUserRole(db database.DatabaseStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.MustGet("role").(string)
		if models.UserRole(role) != models.SuperAdmin {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			return
		}

		var req updateRoleRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validRoles := map[models.UserRole]bool{
			models.Viewer:  true,
			models.Analyst: true,
			models.Admin:   true,
		}
		if !validRoles[req.Role] {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid role"})
			return
		}

		if err := db.UpdateUserRole(req.UserID, req.Role); err != nil {
			if err == database.ErrUserInvalid {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "user not found"})
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "could not update role"})
			return
		}

		cache.Set(req.UserID, req.Role)

		c.JSON(http.StatusOK, gin.H{"message": "role updated"})
	}
}
