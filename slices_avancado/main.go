package main

import (
	"fmt"
	"sort"
)

func main() {
	cidades := []string{"Espírito Santo", "Rio de Janeiro", "Santa Catarina"}
	fmt.Println(cidades, len(cidades), cap(cidades))

	cidades = append(cidades, "Minas Gerais", "São Paulo")
	fmt.Println(cidades, len(cidades), cap(cidades))

	sort.Strings(cidades)
	fmt.Println(cidades, len(cidades), cap(cidades))

	fmt.Println("Sudeste", cidades[:3], cidades[4:])

	fmt.Println("Sul", cidades[3:4])

	cidades = deletarCidadeByIndice(cidades, 2)
	fmt.Println(cidades, len(cidades), cap(cidades))
}

func deletarCidadeByIndice(list []string, indice int) []string {
	temp := list[:indice]
	temp = append(temp, list[indice+1:]...)
	return temp
}
