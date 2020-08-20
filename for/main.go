package main

import "fmt"

func main() {
	for i := 1; i <= 3; i++ {
		fmt.Println("i: ", i)
	}

	contador := 1
	for contador <= 3 {
		fmt.Println("contador: ", contador)
		contador++
	}

	contador = 3
	for {
		fmt.Println("contador inverso: ", contador)
		if contador == 1 {
			break
		}
		contador--
	}

	texto := "O rato roeu a roupa do rei de roma"
	var inicio, final int
	iniciar := true
	for indice, caracter := range texto {
		if caracter == ' ' {
			final = indice - 1
			if inicio == 0 && final < 2 {
				fmt.Printf(" %d", inicio)
			} else {
				fmt.Printf(" %d - %d", inicio, final)
			}
			fmt.Printf("\r\n")
			iniciar = true

		} else {
			if iniciar {
				inicio = indice
				iniciar = false
			}
			fmt.Printf("%q", caracter)
		}
	}
	fmt.Printf(" %d - %d\r\n", inicio, len(texto)-1)
}
