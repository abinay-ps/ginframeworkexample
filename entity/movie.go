package entity

type Movie struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
	Price    string `json:"price"`
}

var Movies = []Movie{
	{ID: "1", Title: "The Dark Knight", Director: "Christopher Nolan", Price: "100"},
	{ID: "2", Title: "Tommy Boy", Director: "Peter Segal", Price: "110"},
	{ID: "3", Title: "Titanic", Director: "James Cameron", Price: "200"},
}
