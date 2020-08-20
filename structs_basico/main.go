package main

import (
	"cursodego/structs_basico/models"
	"fmt"
)

func main() {

	casa := models.Imovel{}

	fmt.Printf("A Casa é: %+v\r\n", casa)

	apartamento := models.Imovel{
		X:     10,
		Y:     15,
		Nome:  "Meu ap!",
		Valor: 1_000_000,
	}

	fmt.Printf("Apartamento: %+v\r\n", apartamento)

	gatoA := models.Gato{}
	gatoA.Idade = 2
	gatoA.Nome = "Mozart"

	fmt.Printf("Gato A: %+v\r\n", gatoA)

	casa.Valor = 2_000_000
	fmt.Printf("Valor da casa: %d\r\n", casa.Valor)

	teste := newImove()

	fmt.Println("Teste", teste)
	if teste == nil {
		fmt.Println("Teste null")
	}

}

func newImove() (imodel *models.Imovel) {
	return
}
