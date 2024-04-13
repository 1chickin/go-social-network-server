package controller

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"os"
	"regexp"

	"github.com/1chickin/go-social-network-server/config"
	"github.com/1chickin/go-social-network-server/internal/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// get username & password from request
	var requestBody struct {
		Username string
		Password string
		Email    string
	}

	if c.Bind(&requestBody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to load request body!",
		})
		return
	}

	// check username exist
	var existingUser model.User
	resultCheckExist := config.DB.Where("username = ?", requestBody.Username).First(&existingUser)
	if resultCheckExist.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username already exists!",
		})
		return
	}

	// check email format valid
	emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, requestBody.Email)
	if !match {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email format!",
		})
		return
	}

	// check email exist
	resultCheckExist = config.DB.Where("email = ?", requestBody.Email).First(&existingUser)
	if resultCheckExist.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email already exists!",
		})
		return
	}

	// generate salt
	newSalt := make([]byte, 10)
	_, err := rand.Read(newSalt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate new salt"})
		return
	}

	// get Pepper from env
	pepper := os.Getenv("PEPPER")
	if len(pepper) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get pepper from env"})
		return
	}

	// hash password with salt and pepper by bcrypt
	// ref: https://pkg.go.dev/golang.org/x/crypto/bcrypt#GenerateFromPassword
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password+hex.EncodeToString(newSalt)+pepper), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password!",
		})
		return
	}

	// create user
	user := &model.User{Username: requestBody.Username, HashedPassword: string(hashedPassword), Salt: hex.EncodeToString(newSalt), Email: requestBody.Email}
	result := config.DB.Create(&user) // pass pointer of data to Create
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user!",
		})
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{
		"msg": "Success!",
	})
}
