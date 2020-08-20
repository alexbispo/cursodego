package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	diaDaSemana := 7

	switch diaDaSemana {
	case 1:
		fmt.Println("Domingo")
	case 2:
		fmt.Println("Segunda")
	case 3:
		fmt.Println("Terça")
	case 4:
		fmt.Println("Quarta")
	case 5:
		fmt.Println("Quinta")
	case 6:
		fmt.Println("Sexta")
	case 7:
		fmt.Println("Sábado")
	default:
		fmt.Println("A semana somente possui 7 dias.")
	}

	fmt.Println("Computador você é da família do Unix?")

	switch runtime.GOOS {
	case "linux":
		fallthrough
	case "darwin":
		fmt.Println("Computador: Sim, eu sou!")
	default:
		fmt.Println("Computador: Prefiro não comentar!")
	}

	hoje := time.Now()

	switch hoje.Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Hoje é dia de folga")
	default:
		fmt.Println("Hoje é dia de trabalhar")
	}

	switch {
	case hoje.Weekday() == time.Saturday && hoje.Day() == 18:
		fmt.Println("Eba! Hoje é dia do lixo")
	case hoje.Day() > 10:
		fmt.Println("Falta somente uma semana para o dia do lixo")
	default:
		fmt.Println("Quando será o nosso dia do lixo?")
	}

}
