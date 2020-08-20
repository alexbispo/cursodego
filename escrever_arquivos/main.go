package main

import (
	"bufio"
	"cursodego/ler_arquivos/models"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func main() {
	arquivoDeDespesas, err := os.Open("despesas.csv")
	check(err)
	defer arquivoDeDespesas.Close()

	// lerComScanner(arquivoDeDespesas)

	leitorCsv := csv.NewReader(arquivoDeDespesas)
	conteudo, err := leitorCsv.ReadAll()
	check(err)

	despesas := make([]models.Despesa, 10)
	for numeroDaLinha, linha := range conteudo {
		if numeroDaLinha == 0 {
			continue
		}

		valor, err := strconv.Atoi(linha[1])
		check(err)

		despesa := models.Despesa{
			Nome:  linha[0],
			Valor: valor,
		}
		despesas = append(despesas, despesa)
	}

	arquivoJSON, err := os.Create("despesas.json")
	check(err)
	defer arquivoJSON.Close()

	escritor := bufio.NewWriter(arquivoJSON)

	_, err = escritor.WriteString("[")
	check(err)

	quantidadeDeDespesas := len(despesas)
	for _, despesa := range despesas[:quantidadeDeDespesas-1] {
		despesaJSON, err := json.Marshal(despesa)
		check(err)

		_, err = escritor.WriteString(string(despesaJSON) + ",")
		check(err)
	}

	ultimaDespesa := despesas[quantidadeDeDespesas-1:][0]
	ultimaDespesaJSON, err := json.Marshal(ultimaDespesa)
	check(err)

	_, err = escritor.WriteString(string(ultimaDespesaJSON))
	check(err)

	_, err = escritor.WriteString("]")
	check(err)

	err = escritor.Flush()
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println("Ocorreu um erro", err.Error())
		panic(err)
	}
}

func lerComScanner(arquivo *os.File) {
	scanner := bufio.NewScanner(arquivo)
	// jump the header
	scanner.Scan()
	for scanner.Scan() {
		linha := scanner.Text()
		fmt.Println("Despesa:", linha)
	}

}
