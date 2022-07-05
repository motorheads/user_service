package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/motorheads/user_service/models"
	"github.com/motorheads/user_service/storage"
)

func GetAllUsers(c *gin.Context) {
	resp, err := storage.GetAllUsers()
	if err != nil {
		fmt.Println("Error while retrieving users from the database")
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": resp,
	})
}

func GetUser(c *gin.Context) {
	var user_id int
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		if key == "user_id" {
			user_id, _ = strconv.Atoi(queryValue)
		}
	}

	resp, err := storage.GetUser(user_id)
	if err != nil {
		fmt.Println("Error while retrieving the user from the database")
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": resp,
	})
}

func CreateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println("Error while binding JSON to user")
		fmt.Println(err)
		return
	}
	err = storage.CreateUser(&user)
	if err != nil {
		fmt.Println("Error while creating user in the database")
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "the user was created succesfully",
	})
}

func UpdateUser(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println("Error while binding JSON to user")
		fmt.Println(err)
		return
	}

	var user_id int
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		if key == "user_id" {
			user_id, _ = strconv.Atoi(queryValue)
		}
	}
	user.ID = user_id

	err = storage.UpdateUser(user)
	if err != nil {
		fmt.Println("Error while updating user in the database")
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "the user was updated succesfully",
	})
}

func DeleteUser(c *gin.Context) {
	var user_id int
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		if key == "user_id" {
			user_id, _ = strconv.Atoi(queryValue)
		}
	}

	err := storage.DeleteUser(user_id)
	if err != nil {
		fmt.Println("Error while deleting the user from the database")
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "the user was deleted succesfully",
	})
}
