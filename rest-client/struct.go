package main

//Movie - a struct of a movie
type Movie struct {
	Title      string
	Year       string
	ImdbRating string
	ImdbVotes  float32 `json:"ImdbVotes,string"`
	ImdbID     string
}

//MovieQuery - Search Results
type MovieQuery struct {
	Search []Movie
}
