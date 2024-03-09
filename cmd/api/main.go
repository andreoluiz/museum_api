package main

import (
	"net/http"

	"github.com/VirtualArtExplore/api/internal"
	"github.com/VirtualArtExplore/api/internal/database"
	"github.com/VirtualArtExplore/api/internal/post"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	connectionString := "postgresql://postgres:3834@localhost:5432/posts"

	conn, err := database.NewConnection(connectionString)
	if err != nil {
		panic(err)
	}

	repo := post.Repository{
		Conn: conn,
	}

	service := post.Service{
		Repository: repo,
	}

	g := gin.Default()
	g.POST("/posts", func(ctx *gin.Context) {
		var post internal.Post
		if err := ctx.BindJSON(&post); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := service.Create(post); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}

	})

	defer conn.Close()

	g.DELETE("posts/:id", func(ctx *gin.Context) {

		param := ctx.Param("id")
		id, err := uuid.Parse(param)
		if err != nil {
			ctx.JSON(http.StatusNotFound, nil)
		}

		if err := service.Delete(id); err != nil {
			statusCode := http.StatusInternalServerError
			if err == post.ErrPostNotFound {
				statusCode = http.StatusNotFound
			}

			ctx.JSON(statusCode, gin.H{
				"error": err.Error(),
			})
			return
		}

	})

	g.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Teste Acampamento AEA, A Aggy é gay!",
		})
	})
	g.Run()
}
