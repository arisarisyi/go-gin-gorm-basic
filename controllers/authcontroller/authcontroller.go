package authcontroller

import (
	"fmt"
	"go-crud-restapi/config"
	"go-crud-restapi/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	var userInput models.User

	if err := c.ShouldBindJSON(&userInput); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"message":err.Error()})
		return
	}

	fmt.Println(userInput.Role)

	var user models.User
	if err := models.DB.Where("email = ?",userInput.Email).First(&user).Error;err != nil{
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound,gin.H{"message":"Email atau password salah"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"message":err.Error()})
			return
		}
	}

	// validate pssword
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(userInput.Password)); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"message":"Email atau password salah"})
		return
	}

	
	// generate token
	expTime := time.Now().Add(time.Minute *4)
	claims := &config.JWTClaim{
		Email: user.Email,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "go-crud-restapi",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// deklarasi logic signin
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	token,err := tokenAlgo.SignedString(config.JWT_KEY)
	if err !=nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"message":err.Error()})
			return
	}
	  c.SetCookie("token",token,3600,"/","localhost",false,true)
		c.JSON(http.StatusOK,gin.H{"message":"Successfully"})
}

func Register(c *gin.Context) {
	var userInput models.User

	if err := c.ShouldBindJSON(&userInput); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"message":err.Error()})
		return
	}

	if userInput.Role==""{
		userInput.Role = "User"
	}

	if err := models.DB.Where("email = ?",userInput.Email).First(&userInput).Error;err == nil{
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"message":"email telah terdaftar"})
		return
	}

	// hash password
	hashPassword,_ := bcrypt.GenerateFromPassword([]byte(userInput.Password),bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	if err := models.DB.Create(&userInput).Error; err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"message":err.Error()})
		return
	}

	c.JSON(http.StatusOK,gin.H{"message":"Successfully"})
}

func Logout(c *gin.Context) {
	 c.SetCookie("token","",-1,"/","localhost",false,true)
		c.JSON(http.StatusOK,gin.H{"message":"Logout"})
}