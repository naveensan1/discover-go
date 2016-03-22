package main

//Movie - a struct of a movie
type Movie struct {
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	Metascore  string
	ImdbRating string
	ImdbVotes  string
	ImdbID     string
	Type       string
	Response   string
}

//MovieQuery - Search Results
type MovieQuery struct {
	Search       []Movie
	Response     string
	totalResults string
}
