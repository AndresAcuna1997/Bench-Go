package routes

import (
	"fmt"
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {

	var newUser models.User

	err := c.ShouldBindJSON(&newUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request data, there is a missing required field"})
		return
	}

	fmt.Println(newUser)

	err = newUser.Save()

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created", "user": newUser})
}

func ListUsers(c *gin.Context) {
	result, err := models.ListUsers()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not retreive users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": result})
}

func Login(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request data, there is a missing required field"})
		return
	}

	err = user.ValidateCredential()

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": token})

}
