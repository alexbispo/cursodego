package matematica

// Calculo faz um calculo genérico
func Calculo(funcao func(int, int) int, a int, b int) int {
	return funcao(a, b)
}
