package main

import (
	"cursodego/pacotes/operadora"
	"cursodego/pacotes/prefixo"
	"fmt"
)

// NomeDoUsuario é o nome do usuário do sistema
var NomeDoUsuario = "Alex"

func main() {
	fmt.Printf("Nome do usuário: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da Operadora: %s\r\n", operadora.NomeOperadora)
	fmt.Printf("Prefixo mais Nome da Operadora: %s\r\n", operadora.PrefixoMaisOperadora)

	// Não compila acessar variável não exportada.
	// fmt.Printf("Valor da variável teste: %s\r\n", prefixo.teste)
}
