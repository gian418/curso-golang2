package model

//Producers representa um produtor
type Producers struct {
	Movie    Movie  `db:"movie_id"`
	Producer string `db:"producers"`
}
