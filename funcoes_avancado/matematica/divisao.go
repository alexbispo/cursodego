package matematica

// Dividir divide dois numeros
func Dividir(a int, b int) (resultado int) {
	resultado = a / b
	return
}

// DividirComResto doois numeros e retorna o resultado e o resto
func DividirComResto(a int, b int) (resultado int, resto int) {
	resultado = a / b
	resto = a % b
	return
}
