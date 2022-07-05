package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/motorheads/catalog_service/models"
	"github.com/motorheads/catalog_service/storage"
)

func GetAllProducts(c *gin.Context) {
	resp, err := storage.GetAllProducts()
	if err != nil {
		fmt.Println("Error while retrieving products from the database")
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products": resp,
	})
}

func GetProductByID(c *gin.Context) {
	var product_id int
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		if key == "product_id" {
			product_id, _ = strconv.Atoi(queryValue)
		}
	}

	resp, err := storage.GetProductByID(product_id)
	if err != nil {
		fmt.Println("Error while retrieving the product from the database")
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product": resp,
	})
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		fmt.Println("Error while binding JSON to product")
		fmt.Println(err)
		return
	}
	err = storage.CreateProduct(&product)
	if err != nil {
		fmt.Println("Error while creating product in the database")
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "the product was created succesfully",
	})
}

func UpdateProductByID(c *gin.Context) {
	var product models.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		fmt.Println("Error while binding JSON to product")
		fmt.Println(err)
		return
	}
	err = storage.UpdateProduct(product)
	if err != nil {
		fmt.Println("Error while updating product in the database")
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "the product was updated succesfully",
	})
}

func DeleteProductByID(c *gin.Context) {
	var product_id int
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		if key == "product_id" {
			product_id, _ = strconv.Atoi(queryValue)
		}
	}

	err := storage.DeleteProductByID(product_id)
	if err != nil {
		fmt.Println("Error while deleting the product from the database")
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "the product was deleted succesfully",
	})
}
