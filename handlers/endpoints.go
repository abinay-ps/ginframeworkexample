package handlers

import "github.com/gin-gonic/gin"

func callhandlers(router *gin.Engine) {
	router.GET("/movie", functions.getMovies)             //Get all movies
	router.GET("/movie/:id", functions.getMovieByID)      //Get a movie by its ID
	router.POST("/movie", functions.createMovie)          // Create a new movie data
	router.PATCH("movie/:id", functions.updateMoviePrice) //Update a particular movie price
	router.DELETE("movie/:id", functions.deleteMovie)     //Delete a particular movie by its ID
}
