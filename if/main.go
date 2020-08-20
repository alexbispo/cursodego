package main

import "fmt"

func main() {
	if desc, deuCerto := processaPagamento("123"); deuCerto {
		fmt.Println(desc)
	}

	// Não compila porque desc apenas é valdo no escopo de if
	// fmt.Println(desc)
}

func processaPagamento(numeroDoCartao string) (descricao string, status bool) {
	if numeroDoCartao == "123" {
		status = true
		descricao = "Pagamento realizado com sucesso"
		return
	}

	descricao = "Cartão inválido"

	return
}
