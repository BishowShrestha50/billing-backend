package middleware

import (
	"net/http"
	"strconv"

	"billing-backend/internal/app/models"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(r models.ProductDetailsInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.GetHeader("x-user-id")
		if id == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
			return
		}
		uid, _ := strconv.ParseUint(id, 10, 32)
		if uid <= 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
			return
		}
		ctx.Next()
	}
}
