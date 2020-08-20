package main

import (
	"cursodego/funcoes_avancado/matematica"
	"fmt"
)

func main() {
	x, y := 5, 8
	resultado := matematica.Calculo(mutilicador, x, y)

	fmt.Printf("%d x %d = %d\r\n", x, y, resultado)

	resultado = matematica.Soma(x, y)
	fmt.Printf("%d + %d = %d\r\n", x, y, resultado)

	x, y = 40, 5
	resultado = matematica.Dividir(x, y)
	fmt.Printf("%d / %d = %d\r\n", x, y, resultado)

	x, y = 5, 8
	resultado, resto := matematica.DividirComResto(x, y)
	fmt.Printf("%d / %d = %d. Resto: %d\r\n", x, y, resultado, resto)
}

func mutilicador(x int, y int) int {
	return x * y
}
