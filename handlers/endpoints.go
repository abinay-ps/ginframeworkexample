package handlers

import (
	"github.com/abinay-ps/ginframeworkexample/functions"
	"github.com/gin-gonic/gin"
)

func CallHandlers(router *gin.Engine) {
	router.GET("/movie", functions.GetMovies)             //Get all movies
	router.GET("/movie/:id", functions.GetMovieByID)      //Get a movie by its ID
	router.POST("/movie", functions.CreateMovie)          // Create a new movie data
	router.PATCH("movie/:id", functions.UpdateMoviePrice) //Update a particular movie price
	router.DELETE("movie/:id", functions.DeleteMovie)     //Delete a particular movie by its ID
}
