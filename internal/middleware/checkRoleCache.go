package middleware

import (
	"net/http"

	"github.com/Ahmed-Armaan/finance-dashboard-api.git/internal/cache"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CheckCacheMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userIdStr, exist := ctx.Get("userID")
		if !exist {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "user role unknown"})
			return
		}

		userId, _ := userIdStr.(uuid.UUID)

		role, exist := cache.Get(userId)
		if exist {
			ctx.Set("role", role)
		}

		ctx.Next()
	}
}
