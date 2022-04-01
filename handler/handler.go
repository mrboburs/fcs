package handler

import (
	"map/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// "github.com/gin-gonic/gin"
// "github.com/google/uuid"

var Articles []model.Article

var reqBody model.Article

func GetAll(c *gin.Context) {
	var a model.Article
	a.ID = "1"
	a.Title = "umidabonu"

	c.JSON(200, a)
}
func DeleteitById(c *gin.Context) {
	id := c.Param("id")
	for i, a := range Articles {
		if id == a.ID {
			Articles = append(Articles[:i], Articles[i+1:]...)
			c.JSON(200, gin.H{
				"error": false,
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"error":   true,
		"message": "invalid Article id",
	})
}
func UpdateItById(c *gin.Context) {
	id := c.Param("id")

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error": true,
			"body":  "invalid request body",
		})
		return
	}
	for i, a := range Articles {
		if a.ID == id {
			Articles[i].Title = reqBody.Title

			c.JSON(200, gin.H{
				"error": false,
			})
			return

		}

	}
	c.JSON(422, gin.H{
		"error": true,
		"body":  "invalid article id",
	})

}

func GetItById(c *gin.Context) {
	id := c.Param("id")
	for _, a := range Articles {
		if a.ID == id {
			c.JSON(200, a)
		}
	}

	c.JSON(404, gin.H{
		"error":   true,
		"message": "invalid id",
	})
}

func CreateArticle(c *gin.Context) {

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error": true,
			"body":  "invalid request body",
		})
		return
	}
	reqBody.ID = uuid.New().String()
	// reqBody.CreatedAt = &time.Time{}
	Articles = append(Articles, reqBody)
	c.JSON(200, gin.H{
		"error": false,
	})

}
