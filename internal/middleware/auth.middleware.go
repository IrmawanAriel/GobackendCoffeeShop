package middleware

import (
	"IrmawanAriel/goBackendCoffeeShop/pkg"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthJwtMiddleware(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := pkg.NewResponse(ctx)

		// Ambil header Authorization
		header := ctx.GetHeader("Authorization")
		if header == "" {
			response.Unauthorized("Authorization header is missing", nil)
			ctx.Abort()
			return
		}

		// Pastikan header memiliki token Bearer
		if !strings.HasPrefix(header, "Bearer ") {
			response.Unauthorized("Invalid Bearer token format", nil)
			ctx.Abort()
			return
		}

		// Ambil token dari header
		token := strings.TrimPrefix(header, "Bearer ")

		// Verifikasi token
		check, err := pkg.VerifyToken(token)
		if err != nil {
			response.Unauthorized("Invalid Bearer token", err.Error())
			ctx.Abort()
			return
		}

		// Set userId dan email ke context
		ctx.Set("userId", check.Id)
		ctx.Set("email", check.Email)

		// Jika role spesifik diperlukan, verifikasi role
		if role != "" && check.Role != role {
			response.Unauthorized("Invalid role for token", nil)
			ctx.Abort()
			return
		}

		// Lanjutkan ke handler berikutnya jika tidak ada masalah
		ctx.Next()
	}
}
