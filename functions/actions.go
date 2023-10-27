package functions

import (
	"net/http"

	"github.com/abinay-ps/ginframeworkexample/entity"
	"github.com/gin-gonic/gin"
)

func GetMovies(c *gin.Context) {
	c.JSON(http.StatusOK, entity.Movies)
}

func GetMovieByID(c *gin.Context) {
	c.JSON(http.StatusOK, entity.Movies)
}

func CreateMovie(c *gin.Context) {
	c.JSON(http.StatusOK, entity.Movies)
}

func UpdateMoviePrice(c *gin.Context) {
	c.JSON(http.StatusOK, entity.Movies)
}

func DeleteMovie(c *gin.Context) {
	c.JSON(http.StatusOK, entity.Movies)
}
