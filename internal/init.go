package internal

import (
	"log"

	"github.com/Ahmed-Armaan/finance-dashboard-api.git/internal/database"
	"github.com/Ahmed-Armaan/finance-dashboard-api.git/internal/middleware"
	"github.com/Ahmed-Armaan/finance-dashboard-api.git/internal/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(db database.DatabaseStore) {
	r := gin.Default()
	r.Use(cors.Default())

	public := r.Group("/")
	{
		public.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "pong"})
		})
		public.POST("/signup", routes.Signup(db))
		public.POST("/login", routes.Login(db))
	}

	protected := r.Group("/")
	protected.Use(middleware.VerifyJWTMiddleware(), middleware.CheckCacheMiddleware())
	{
		protected.GET("/data")
		protected.GET("/insights")
		protected.GET("/records")
		protected.PUT("/records")
		protected.GET("/requests", routes.GetRequests(db))
		protected.POST("/request")
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error: Failed to start server\n%v\n", err)
	}
}
