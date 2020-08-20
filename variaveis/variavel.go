package main

import "fmt"

var (
	// Endereco é uma variável públlica
	Endereco string

	// Telefone tmabém é publica
	Telefone string

	quantidade, estoque int     // quantidade = 0
	comprou             bool    // comprou = false
	valor               float64 // valor = 0.00
	palavras            rune
)

func main() {
	teste := "Valor do teste"

	fmt.Printf("endereco: %s\r\n", Endereco)
	fmt.Printf("quantidade: %d\r\n", quantidade)
	fmt.Printf("comprou: %v\r\n", comprou)
	fmt.Printf("palavras: %v\r\n", palavras)
	fmt.Printf("estoque: %d\r\n", estoque)
	fmt.Printf("teste: %s\r\n", teste)

}
