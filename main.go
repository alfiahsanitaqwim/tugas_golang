package main

import (
	"tugas_golang_alfi_ahsani/database"
	"tugas_golang_alfi_ahsani/entity"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading %s file: %s", ".env", err)
	}
	database.Connect()
	r := gin.Default()

	// show all
	r.GET("/books", func(c *gin.Context) {
		var ent []entity.Books
		err := database.DB.Find(&ent).Error
		if err != nil {
		}
		c.JSON(200, gin.H{
			"data": ent,
		})
	})

	// show one
	r.GET("/book/:id", func(c *gin.Context) {
		var id = c.Param("id")
		var ent []entity.Books
		err := database.DB.Model(ent).Where("ID = ?", id).Find(&ent).Error
		if err != nil {
		}
		c.JSON(200, gin.H{
			"data": ent,
		})
	})

	type InvitationRequest struct {
		Title  string `json:"title"`
		Author string `json:"author"`
		Desc   string `json:"desc"`
	}

	// save data
	r.POST("/books/add", func(c *gin.Context) {
		var req InvitationRequest
		err := c.ShouldBindJSON(&req)
		if err != nil {
		}
		data := entity.Books{
			Title:  req.Title,
			Author: req.Author,
			Desc:   req.Desc,
		}
		err = database.DB.Create(&data).Error
		if err != nil {
		}
		c.JSON(200, gin.H{
			"message": "save berhasil",
		})
	})

	r.PATCH("/books/update/:id", func(c *gin.Context) {
		var id = c.Param("id")
		var ent []entity.Books
		err := database.DB.Model(ent).Where("ID = ?", id).Find(&ent).Error
		if err != nil {
		}

		var req InvitationRequest
		err2 := c.ShouldBindJSON(&req)
		if err2 != nil {
		}
		data := entity.Books{
			Title:  req.Title,
			Author: req.Author,
			Desc:   req.Desc,
		}

		c.JSON(200, gin.H{
			"message": "save berhasil di upate",
			"data":    data,
			"id":      id,
			"ent":     ent,
		})

	})

	r.DELETE("/books/delete/{id}", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
