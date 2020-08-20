package main

import (
	"fmt"
	"strconv"
)

type despesa struct {
	name           string
	valor          int
	valorFormatado string
}

func main() {
	despesaUm := despesa{
		name:  "Arrendamento",
		valor: 68_000,
	}

	fmt.Printf("Despesa um: %p - %+v\r\n", &despesaUm, despesaUm)

	despesaDois := new(despesa)
	despesaDois.name = "EDP"
	despesaDois.valor = 2_000

	fmt.Printf("Despesa dois: %p -  %+v\r\n", &despesaDois, despesaDois)

	formataValorDaDespesa(despesaUm)

	fmt.Printf("(formataValorDaDespesa) Despesa um: %+v\r\n", despesaUm)

	formataValorDaDespesaPonteiro(&despesaUm)

	fmt.Printf("(formataValorDaDespesaPonteiro) Despesa um: %+v\r\n", despesaUm)

	formataValorDaDespesa(*despesaDois)

	fmt.Printf("(formataValorDaDespesa) Despesa dois: %+v\r\n", despesaDois)

	formataValorDaDespesaPonteiro(despesaDois)

	fmt.Printf("(formataValorDaDespesaPonteiro) Despesa dois: %+v\r\n", despesaDois)
}

func formataValorDaDespesa(desp despesa) {
	valorConvertido := desp.valor / 100
	desp.valorFormatado = strconv.Itoa(valorConvertido)
}

func formataValorDaDespesaPonteiro(desp *despesa) {
	valorConvertido := desp.valor / 100
	desp.valorFormatado = strconv.Itoa(valorConvertido)
}
