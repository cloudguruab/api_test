package controllers

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/cloudguruab/api_test/models"
)

func FindBooks(c *gin.Context) {
  var books []models.Book
  models.DB.Find(&books)

  c.JSON(http.StatusOK, gin.H{"data": books})
}

func CreateBook(c *gin.Context) {
    // struct value
    var input models.CreateBookInput
    
    // validation
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    // create book in db 
    book := models.Book{Title: input.Title, Author: input.Author, Rating: input.Rating}
    models.DB.Create(&book)
    
    // create status response 200 for creation of book 
    c.JSON(http.StatusOK, gin.H{"data": book})
}

func FindBook(c *gin.Context) {
    var book models.Book
    
    if err := models.DB.Where("id = ?", c.Param("id")).First(&Book).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"data": book})
}
