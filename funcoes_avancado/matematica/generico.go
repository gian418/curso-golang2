package matematica

//Calculo executa qualquer tipo de calculo
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}

//Multiplicador retorna a multiplicação de dois numeros inteiros
func Multiplicador(x int, y int) int {
	return x * y
}

//Divisor retorna a divisão de dois numeros inteiros
func Divisor(numeroA int, numeroB int) (resultado int) {
	resultado = numeroA / numeroB
	return
}

//DivisorComResto divisão com resto
func DivisorComResto(numeroA int, numeroB int) (resultado int, resto int) {
	resultado = numeroA / numeroB
	resto = numeroA % numeroB
	return
}
