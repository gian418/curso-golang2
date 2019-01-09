package matematica

//Calculo executa qualquer tipo de calculo
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}

//Multiplicador retorna a multiplicação de dois numeros inteiros
func Multiplicador(x int, y int) int {
	return x * y
}
