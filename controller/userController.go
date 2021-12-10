package controller

import (
	"example/login/model"

	"example/login/service"

	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func AddUser(c *gin.Context) {
	user := model.User{}
	err := c.BindJSON(&user)
	hashedPassword, err := HashPassword(user.Password)
	user, err = service.AddUser(user.Email, hashedPassword)
	exception := ""
	if err != nil {
		exception = err.Error()
	}
	c.JSON(200, gin.H{"exception": exception, "data": user})
}

func GetUser(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Params.ByName("userID"), 10, 64)
	user, err := service.GetUser(uint(userID))
	password := c.Params.ByName("password")
	exception := ""
	if err != nil {
		exception = err.Error()
		c.JSON(500, gin.H{"Error": exception})
		return
	}
	if !CheckPasswordHash(password, user.Password) {
		c.JSON(403, gin.H{"exception": "Password is not valid for User with id " + string(userID)})
		return
	}
	c.JSON(200, gin.H{"exception": exception, "data": user})
}
