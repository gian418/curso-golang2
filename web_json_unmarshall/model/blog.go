package model

//Blog apenas uma estrutura para receber a resposta do servidor
type Blog struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
