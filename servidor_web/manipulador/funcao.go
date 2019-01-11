package manipulador

import (
	"fmt"
	"net/http"
)

//Funcao é uma manipulador de requisição http na rota /funcao
func Funcao(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Aqui é um manipulador usando função em um pacote")
}
