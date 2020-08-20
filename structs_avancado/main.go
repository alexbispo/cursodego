package main

import (
	"cursodego/structs_avancado/models"
	"encoding/json"
	"fmt"
)

func main() {
	meuGato := models.Gato{
		Nome:  "Mozart",
		Idade: 2,
	}

	// Não compila, numeroDeVacinas não foi exportado
	// meuGato.numeroDeVacinas = 1

	fmt.Printf("%+v\r\n", meuGato)

	meuGato.AdicionaVacina()

	fmt.Printf("Nome do gato: %s\r\n", meuGato.Nome)
	fmt.Printf("Idade do gato: %d\r\n", meuGato.Idade)
	fmt.Printf("Numero de vacinas: %d\r\n", meuGato.NumeroDeVacinas())

	objJSON, _ := json.Marshal(meuGato)

	fmt.Printf("JSON: %+v\r\n", string(objJSON))

	cachorro := models.Cachorro{
		Nome: "Cachorrinho",
	}

	fmt.Println("[cachorro]", cachorro)
	cachorro.SetIdade(2)
	fmt.Println("[cachorro]", cachorro, cachorro.GetIdade())

	cachorro2 := models.NovoCachorro("Cachorrinho")

	fmt.Println("[cachorro2]", cachorro2)
	cachorro2.SetIdade(2)
	fmt.Println("[cachorro2]", cachorro2, cachorro2.GetIdade())
}
