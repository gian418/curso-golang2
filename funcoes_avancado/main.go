package main

import (
	"fmt"

	"github.com/gian418/funcoes_avancado/matematica"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 2, 2)
	fmt.Printf("O resultado da multiplicação foi: %d\r\n", resultado)

	resultado = matematica.Soma(3, 3)
	fmt.Printf("O resultado da soma foi: %d\r\n", resultado)

	resultado = matematica.Calculo(matematica.Divisor, 12, 3)
	fmt.Printf("O resultado da divisão foi: %d\r\n", resultado)

	resultado, resto := matematica.DivisorComResto(12, 5)
	fmt.Printf("O resultado da divisão com resto foi: %d %d\r\n", resultado, resto)

}
