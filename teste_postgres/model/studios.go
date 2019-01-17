package model

//Studios representa um estudio
type Studios struct {
	Movie    Movie  `db:"movie_id"`
	Producer string `db:"producers"`
}
