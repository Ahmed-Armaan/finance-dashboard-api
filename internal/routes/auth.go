package routes

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Ahmed-Armaan/finance-dashboard-api.git/internal/auth"
	"github.com/Ahmed-Armaan/finance-dashboard-api.git/internal/database"
	"github.com/gin-gonic/gin"
)

type loginResponseType struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(db database.DatabaseStore) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jsonData, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to read request body"})
			return
		}

		var request loginResponseType
		if err = json.Unmarshal(jsonData, &request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to read request body"})
			return
		}

		user, err := db.GetUser(request.Username)
		if err != nil {
			if err == database.ErrUserInvalid {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "user checkup failed"})
			}
			return
		}

		if auth.Compare([]byte(request.Password), user.Password) {
			jwtToken, err := auth.SignJwt(user.ID, string(user.Role))
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "login failed"})
			}

			http.SetCookie(ctx.Writer, &http.Cookie{
				Name:     "session",
				Value:    jwtToken,
				Path:     "/",
				MaxAge:   60 * 60 * 24 * 7,
				Secure:   true,
				HttpOnly: true,
				SameSite: http.SameSiteNoneMode,
			})
		}
	}
}
