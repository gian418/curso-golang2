package main

import (
	"fmt"

	"github.com/gian418/cursogo/pacotes/operadora"
	"github.com/gian418/cursogo/pacotes/prefixo"
)

//NomeDoUsuario é o nome do usuário do sistema
var NomeDoUsuario = "Jeff"

func main() {
	fmt.Printf("Nome do usuario: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da operadora: %s\r\n", operadora.NomeOperadora)
}
