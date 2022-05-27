package main

import (
    "github.com/cloudguruab/api_test/models"
    "github.com/gin-gonic/gin"
    "github.com/cloudguruab/api_test/controllers"
)

func main() {
    r := gin.Default()

    models.ConnectDatabase()

    r.DELETE("/books/:id", controllers.DeleteBook)
    r.PATCH("/books/:id", controllers.UpdateBook)
    r.POST("/books", controllers.CreateBook)
    r.GET("/books", controllers.FindBooks)
    r.GET("/books/:id", controllers.FindBook)

    r.Run()
}
