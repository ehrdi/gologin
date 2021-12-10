package controller

import (
	"example/login/dto"
	"example/login/service"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	loginDto := dto.LoginDto{}
	err := c.ShouldBindJSON(&loginDto)
	if err != nil {
		c.JSON(403, "Invalid arguments")
		return
	}
	user, err := service.GetUserByEmail(loginDto.Email)
	if err != nil {
		c.JSON(404, "User with Email "+loginDto.Email+" not found.")
		return
	}
	if !CheckPasswordHash(loginDto.Password, user.Password) {
		c.JSON(403, gin.H{"exception": "Password is not valid for User with Email " + loginDto.Email})
		return
	}
	c.JSON(200, gin.H{"message": "Login successful", "data": user})
}
