package main

import "fmt"

//Imovel é uma struct que armazena dados de um imovel
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {

	casa := Imovel{}
	fmt.Printf("A casa é: %+v\r\n", casa)

	apartamento := Imovel{17, 56, "Meu AP", 760000}
	fmt.Printf("O Apartamento é: %+v\r\n", apartamento)

	chacara := Imovel{
		Y:     85,
		Nome:  "Chacara",
		X:     22,
		valor: 55,
	}
	fmt.Printf("A chacara é: %+v\r\n", chacara)

	casa.Nome = "Lar Doce Lar"
	casa.valor = 450000
	casa.X = 18
	casa.Y = 31
	fmt.Printf("A casa é: %+v\r\n", casa)

}
