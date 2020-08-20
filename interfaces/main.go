package main

import (
	"cursodego/interfaces/models"
	"fmt"
)

func main() {
	meuPet := models.Pet{
		Nome: "Mozart",
	}

	latir(meuPet)
	miar(meuPet)
}

func latir(p models.Cachorro) {
	fmt.Println(p.Late())
}

func miar(p models.Gato) {
	fmt.Println(p.Mia())
}
