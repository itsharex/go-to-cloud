package middlewares

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go-to-cloud/conf"
	repo "go-to-cloud/internal/repositories"
	"net/http"
	"time"
)

type login struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var jwtMiddleware *jwt.GinJWTMiddleware

func GinJwtMiddleware() *jwt.GinJWTMiddleware {
	return jwtMiddleware
}

func AuthHandler() gin.HandlerFunc {

	m, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       conf.GetJwtKey().Realm,
		Key:         []byte(conf.GetJwtKey().Security),
		Timeout:     time.Hour * 2,
		MaxRefresh:  time.Hour / 2,
		IdentityKey: "jti",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*repo.User); ok {
				return jwt.MapClaims{
					"jti": v.ID,
					"sub": v.Account,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVal login
			if err := c.ShouldBind(&loginVal); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			account := loginVal.UserName
			password := loginVal.Password

			user := repo.GetUser(&account, &password)

			if user != nil {
				return user, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*repo.User); ok && v.ID > 0 {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		LoginResponse: func(c *gin.Context, code int, message string, time time.Time) {
			c.AbortWithStatus(http.StatusTemporaryRedirect)
			c.Redirect(http.StatusTemporaryRedirect, "/login")
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer:",
	})

	jwtMiddleware = m

	return m.MiddlewareFunc()
}
