package middleware

import (
	"encoding/json"
	"log"
	"other/L-NOTE/models"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var identityKey = "id"

// AuthMiddleware .
func GinJWTMiddlewareInit() (authMiddleware *jwt.GinJWTMiddleware) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Minute * 7,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.UserModel); ok {
				v.UserClaims = models.GetUserClaims(v.Username)
				jsonClaim, _ := json.Marshal(v.UserClaims)
				return jwt.MapClaims{
					"userName":   v.Username,
					"userClaims": string(jsonClaim),
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			jsonClaim := claims["userClaims"].(string)
			var userClaims []models.ClaimsModel
			json.Unmarshal([]byte(jsonClaim), &userClaims)
			return &models.UserModel{
				Username:   claims["userName"].(string),
				UserClaims: userClaims,
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals models.UserModel
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			username := loginVals.Username
			password := loginVals.Password

			// check data login
			if models.CheckUser(username, password) {
				return &models.UserModel{
					Username: username,
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		// Authorizator: jwtAuthorizator.HandleAuthorizator,
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*models.UserModel); ok {
				if v.RoleID <= 10 {
					return true
				}
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
	}

	// errInit := authMiddleware.MiddlewareInit()
	//
	// if errInit != nil {
	// 	log.Fatal("authMiddleware.MiddlewareInit() Error: " + errInit.Error())
	// }
	return
}
