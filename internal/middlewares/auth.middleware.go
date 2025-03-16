package middlewares

import (
	"context"
	services "ecommerce/internal/services/user"
	"ecommerce/internal/utils/auth"

	"github.com/gin-gonic/gin"
)

type ContextKey string

const UserUUIDKey ContextKey = "userUUID"

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check headers authorization
		accessToken, err := c.Cookie("access_token")
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"code": 40001, "err": "Unauthorized", "description": ""})
			return
		}
		// validate jwt token by subject
		claims, err := auth.VerifyTokenJWT(accessToken)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"code": 40001, "err": "invalid token", "description": ""})
			return
		}
		// update claims to context
		ctx := context.WithValue(c.Request.Context(), UserUUIDKey, claims.Subject)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func AdminMiddleware(adminService services.IAdminService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userUUID := c.Request.Context().Value(UserUUIDKey).(string)
		user, err := adminService.CheckAdmin(userUUID)
		if err != nil || !user {
			c.AbortWithStatusJSON(401, gin.H{"code": 401, "error": "Unauthorized"})
			return
		}
		c.Next()
	}
}
