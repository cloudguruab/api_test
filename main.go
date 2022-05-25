package main

import (
    "github.com/cloudguruab/api_test/models"
    "github.com/gin-gonic/gin"
    "github.com/cloudguruab/api_test/controllers"
)

func main() {
    r := gin.Default()

    models.ConnectDatabase()

    r.POST("/books", controllers.CreateBook)
    r.GET("/books", controllers.FindBooks)
    r.GET("/books/:id", controllers.FindBook) // left off here fetching rout at /books/1

    r.Run()
}
