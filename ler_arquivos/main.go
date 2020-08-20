package main

import (
	"bufio"
	"cursodego/ler_arquivos/models"
	"encoding/csv"
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

		fmt.Println(despesa)

		despesas = append(despesas, despesa)

	}

	fmt.Println(despesas, cap(despesas), len(despesas))
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
