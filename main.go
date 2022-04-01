package main

import (
	// "fmt"
	// "log"

	"github.com/gin-gonic/gin"
	// "google.golang.org/protobuf/internal/encoding/text"
	// "github.com/swaggo/files"
	// "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "map/docs"

	// "map/handler"
	"map/model"
	"net/http"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var Articles []model.Article

// var a model.Article
var reqBody model.Article

type DefaultError struct {
	Message string `json:"message"`
}

type ResponseError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
type ResponseSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Create Article godoc
// @Summary     Create an Article
// @Description it create article based on on input data
// @ID create-handler
// @Tags   Article
// @Accept       json
// @Produce      json
// @Param        data    body  model.Article true "Article data"
// @Success      200   {object}      ResponseSuccess
// @Failure      400,404  {object}  ResponseError
// @Failure      500  {object}  DefaultError
// @Router       /a [POST]
func CreateArticle(c *gin.Context) {

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error": true,
			"body":  "invalid request body",
		})
		return
	}

	// reqBody.CreatedAt = &time.Time{}
	Articles = append(Articles, reqBody)
	c.JSON(200, gin.H{
		"error": false,
	})

}

// @title           Article API Docs
// @version         2.0
// @description     This is a sample article server celler server.
// @termsOfService  http://github.com/Golang_Tutorial/Article_Array
// @contact.name   MY
// @contact.email  babdusalom@gmail.com
// @host      localhost:8080

// Show all articles godoc
// @Summary      Show  all article
// @Description  get all articles in memory
// @Tags   Article
// @ID                get-all-article-handler
// @Produce      json
// @Accept        json
// @Success      200  {array}  model.Article
// @Failure      500  {string}  DefaultError
// @Router       /a [GET]
func GetAll(c *gin.Context) {

	c.JSON(200, Articles)
}

// Get Article by Id  godoc
// @Summary   get article by id
// @Description it return Article which you send id
// @ID getarticle-by-id-handler
// @Tags   Article
// @Accept       json
// @Produce      json
// @Param        id   path  int     true "Param ID"
// @Success      200   {object}      model.Article
// @Failure      400,404  {object}  ResponseError
// @Failure      500  {object}  DefaultError
// @Router       /a/id{id} [GET]
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

// Delete Article by Id  godoc
// @Summary   delete article by id
// @Description it delete Article which you send id of article
// @ID delete-handler
// @Tags   Article
// @Accept       json
// @Produce      json
// @Param        id   path  int     true "Param ID"
// @Success      200   {object}    ResponseSuccess
// @Failure      400,404  {object}  ResponseError
// @Failure      500  {object}  DefaultError
// @Router       /a/id{id} [DELETE]
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

// Update Article godoc
// @Summary     Update an Article
// @Description it update article based on on input data
// @ID update-handler
// @Tags   Article
// @Accept       json
// @Produce      json
// @Param        data    body  model.Article true "Article data"
// @Success      200   {object}      ResponseSuccess
// @Failure      400,404  {object}  ResponseError
// @Failure      500  {object}  DefaultError
// @Router       /a/id{id} [PUT]
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
func main() {

	r := gin.Default()
	// url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"ssssss": "jjjj",
		})
	})
	// articleRoutes := r.Group("/a")
	{

		r.POST("/a", CreateArticle)
		r.GET("/a", GetAll)
		r.GET("/a/id:id", GetItById)
		r.DELETE("/a/id:id", DeleteitById)
		r.PUT("/a/id:id", UpdateItById)

		// url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
		// articleRoutes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		r.Run(":8080")

	}

}
