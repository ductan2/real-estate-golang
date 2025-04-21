package middlewares

import (
	"context"
	services "ecommerce/internal/services/user"
	"ecommerce/internal/utils/auth"
	"ecommerce/pkg/response"

	"github.com/gin-gonic/gin"
)

type ContextKey string

const UserUUIDKey ContextKey = "userUUID"

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check headers authorization
		accessToken, err := c.Cookie("access_token")
		if err != nil {
			response.ErrorResponse(c, response.Unauthorized, "Unauthorized")
			return
		}
		// validate jwt token by subject
		claims, err := auth.VerifyTokenJWT(accessToken)
		if err != nil {
			response.ErrorResponse(c, response.Unauthorized, "Invalid token")
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
		println("userUUID", userUUID)
		user, err := adminService.CheckAdmin(userUUID)
		if err != nil || !user {
			response.ErrorResponse(c, response.Unauthorized, "Unauthorized")
			return
		}
		c.Next()
	}
}

func OptionalAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        if accessToken, err := c.Cookie("access_token"); err == nil {
            if claims, err := auth.VerifyTokenJWT(accessToken); err == nil {
                ctx := context.WithValue(c.Request.Context(), UserUUIDKey, claims.Subject)
                c.Request = c.Request.WithContext(ctx)
            }
        }
        c.Next()
    }
}