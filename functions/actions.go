package functions

import (
	"net/http"

	"github.com/abinay-ps/ginframeworkexample/entity"
	"github.com/gin-gonic/gin"
)

func getMovies(c *gin.Context) {
	c.JSON(http.StatusOK, entity.Movies)
}

func getMovieByID(c *gin.Context) {
	c.JSON(http.StatusOK, entity.Movies)
}

func createMovie(c *gin.Context) {
	c.JSON(http.StatusOK, entity.Movies)
}

func pdateMoviePrice(c *gin.Context) {
	c.JSON(http.StatusOK, entity.Movies)
}

func deleteMovie(c *gin.Context) {
	c.JSON(http.StatusOK, entity.Movies)
}
