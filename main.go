package main

import (
    "github.com/cloudguruab/api_test/models"
    "github.com/gin-gonic/gin"
    "github.com/cloudguruab/api_test/controllers"
)

func main() {
    r := gin.Default()

  // r.GET("/", func(c *gin.Context) {
  //   c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
  // })

    
    models.ConnectDatabase()

    r.GET("/books", controllers.FindBooks)

    r.Run()
}
