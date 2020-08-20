package main

import "fmt"

func main() {
	// slice com elementos e capcidade não inicializados
	var gatos []string
	fmt.Println(gatos, len(gatos), cap(gatos))

	gatos = append(gatos, "Mozart")
	fmt.Println(gatos, len(gatos), cap(gatos))

	gatos = append(gatos, "Tempestade")
	fmt.Println(gatos, len(gatos), cap(gatos))

	nums := make([]int, 3)
	fmt.Println(nums, len(nums), cap(nums))

	// slice criado com uma capacidade (3) e elements pré-definidos
	livros := []string{"Eclesiastes", "Salmos", "Provérbios"}
	fmt.Println(livros, len(livros), cap(livros))

	// aqui observe que a capacidade do slice irá dobrar
	// mesmo adicionando apenas mais um emelento
	// isto acontece, porque o slice já foi criado com uma capacidade definida (3)
	livros = append(livros, "Cantares de Salomão")
	fmt.Println(livros, len(livros), cap(livros))

	nome := []byte("Rei ")
	fmt.Println(string(nome), len(nome), cap(nome))

	nome = append(nome, "Salomão"...)

	fmt.Println(string(nome), len(nome), cap(nome))
}
