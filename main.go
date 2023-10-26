package main

import "github.com/gin-gonic/gin"

var Movies = []Movie{
	{ID: "1", Title: "The Dark Knight", Director: "Christopher Nolan", Price: "100"},
	{ID: "2", Title: "Tommy Boy", Director: "Peter Segal", Price: "110"},
	{ID: "3", Title: "Titanic", Director: "James Cameron", Price: "200"},
}

func main() {
	router := gin.Default()

}
