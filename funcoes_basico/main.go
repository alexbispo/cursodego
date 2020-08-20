package main

import (
	"cursodego/funcoes_basico/matematica"
	"fmt"
)

func main() {
	x, y := 5, 8
	resultado := matematica.Calculo(mutilicador, x, y)

	fmt.Printf("%d x %d = %d\r\n", x, y, resultado)

	resultado = matematica.Soma(x, y)
	fmt.Printf("%d + %d = %d\r\n", x, y, resultado)
}

func mutilicador(x int, y int) int {
	return x * y
}
