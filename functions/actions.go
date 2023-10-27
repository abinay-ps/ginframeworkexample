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
	id := c.Param("id")
	for _, val := range entity.Movies {
		if id == val.ID {
			c.JSON(http.StatusOK, val)
			return
		}
	}
	c.JSON(http.StatusOK, "No movie exists for the inputed ID")
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
