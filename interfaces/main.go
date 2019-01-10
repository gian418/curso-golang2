package main

import (
	"fmt"

	"github.com/gian418/interfaces/model"
)

func main() {
	jojo := model.Ave{}
	jojo.Nome = "Jojo da Silva"

	queroAcordarComUmCacarejo(jojo)
	queroOuvirUmPataNoLago(jojo)

}

func queroAcordarComUmCacarejo(g model.Galinha) {
	fmt.Println(g.Cacareja())
}

func queroOuvirUmPataNoLago(p model.Pata) {
	fmt.Println(p.Quack())
}
