package routes

import (
	"net/http"

	"github.com/Ahmed-Armaan/finance-dashboard-api.git/internal/auth"
	"github.com/Ahmed-Armaan/finance-dashboard-api.git/internal/database"
	"github.com/Ahmed-Armaan/finance-dashboard-api.git/internal/database/models"
	"github.com/gin-gonic/gin"
)

type signupRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func Signup(db database.DatabaseStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req signupRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hashedPassword, err := auth.HashAndSalt([]byte(req.Password))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "could not process password"})
			return
		}

		user, err := db.AddUser(req.UserName, hashedPassword, models.Viewer)
		if err != nil {
			if err == database.ErrUserExists {
				c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "username already taken"})
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"id":       user.ID,
			"username": user.UserName,
			"role":     user.Role,
		})
	}
}
