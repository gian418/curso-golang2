package model

//Movie representa um filme
type Movie struct {
	ID     int    `db:"id"`
	Title  string `db:"title"`
	Winner bool   `db:"winner"`
	Year   int    `db:"year"`
}
