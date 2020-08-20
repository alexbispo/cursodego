package main

import "fmt"

func main() {
	var gatos [3]string
	gatos[0] = "Mozart"
	gatos[1] = "Tempestade"
	gatos[2] = "Caramelo"

	fmt.Println("Quantidade de gatos", len(gatos))

	for _, gato := range gatos {
		fmt.Println(gato)
	}

	paisesQueVisitei := [2]string{"Chile", "Portugal"}

	fmt.Println("Quantidade de paises visitados", len(paisesQueVisitei))
	fmt.Println(paisesQueVisitei)

	livros := [...]string{"Genesis", "Êxodo", "Levítico", "Números", "Deuteronômico"}

	fmt.Println("Quantidade de licros escritos por Moisés", len(livros))
	fmt.Println(livros)

}
