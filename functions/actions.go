package functions

import (
	"fmt"
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
	c.String(http.StatusOK, "No movie exists for the inputed ID: %v", id)
}

func CreateMovie(c *gin.Context) {
	var newMovie entity.Movie
	if err := c.BindJSON(&newMovie); err != nil {
		fmt.Println("err: ", err)
		return
	}
	entity.Movies = append(entity.Movies, newMovie)
	c.String(http.StatusCreated, "New Movie with the details: %v is created", newMovie)
}

func UpdateMoviePrice(c *gin.Context) {
	id := c.Param("id")
	price := c.Param("price")
	for i, val := range entity.Movies {
		if id == val.ID {
			entity.Movies[i].Price = price
			c.String(http.StatusOK, "The price of the movie %v is updated as %v", val.Title, price)
			return
		}
	}
	c.String(http.StatusOK, "No movie exists for the inputed id: %v", id)
}

func DeleteMovie(c *gin.Context) {
	id := c.Param("id")
	for i, val := range entity.Movies {
		if id == val.ID {
			entity.Movies = append(entity.Movies[:i], entity.Movies[i+1:]...)
			c.String(http.StatusOK, "The movie with id: %v is deleted", id)
			return
		}
	}
}
