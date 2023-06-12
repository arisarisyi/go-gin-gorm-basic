package middlewares

import (
	"fmt"
	"go-crud-restapi/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware(c *gin.Context) {
	// Get cookie
	tokenString, err := c.Cookie("token")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	claims := &config.JWTClaim{}
	// parsing token

	fmt.Println(claims,"AWAL")
	token,err := jwt.ParseWithClaims(tokenString,claims,func(t *jwt.Token)(interface{},error){
		return config.JWT_KEY,nil
	})

	if err != nil{
		v,_ := err.(*jwt.ValidationError)
		switch v.Errors{
		case jwt.ValidationErrorSignatureInvalid:
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		case jwt.ValidationErrorExpired:
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized token expired"})
			return
		default :
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}
	}

	if !token.Valid{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

	c.Next()
}