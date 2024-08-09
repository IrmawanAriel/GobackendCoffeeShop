package middleware

import (
	"IrmawanAriel/goBackendCoffeeShop/pkg"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthJwtMiddleware(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := pkg.NewResponse(ctx)
		var header string

		if header = ctx.GetHeader("Authorization"); header == "" {
			response.Unauthorized("Unauthorized", nil)
			return
		}

		if !strings.Contains(header, "Bearer") {
			response.Unauthorized("Inavlid Bearer Token", nil)
			return
		}

		token := strings.Replace(header, "Bearer ", "", -1) // hapus Bearer token hingga tokennya saja yang tersisa

		check, err := pkg.VerifyToken(token, role)
		if err != nil {
			response.Unauthorized("Inavlid Bearer Token", nil)
			return
		}

		ctx.Set("userId", check.Id)
		ctx.Set("email", check.Email)
		ctx.Next()
	}
}
