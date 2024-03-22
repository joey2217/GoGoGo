package middleware

import (
	"net/http"
	"web/dao/model"
	"web/utils"

	"github.com/gin-gonic/gin"
)

const AUTH_COOKIE = "auth_token"
const CTX_AUTH = "user_id"

func Auth(role model.Role) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Request.Cookie(AUTH_COOKIE)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		claims, err2 := utils.ParseToken(cookie.Value)
		if err2 != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err2.Error()})
			return
		}
		if claims.Role == model.Admin || role == claims.Role {
			ctx.Set(CTX_AUTH, claims.UserID)
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		}
	}
}
