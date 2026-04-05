package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Ahmed-Armaan/finance-dashboard-api.git/internal/database"
	"github.com/Ahmed-Armaan/finance-dashboard-api.git/internal/database/models"
	"github.com/gin-gonic/gin"
)

func GetRequests(db database.DatabaseStore) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, ok := ctx.MustGet("role").(string)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user role not found"})
			return
		}

		fmt.Println(role)
		if role != string(models.SuperAdmin) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			return
		}

		offsetStr := ctx.Query("offset")
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid offset"})
			return
		}

		requests, err := db.ListRequests(offset)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "fetch failed"})
			return
		} else {
			ctx.JSON(200, requests)
		}
	}
}
