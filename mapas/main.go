package main

import (
	"cursodego/mapas/models"
	"fmt"
)

var cache map[string]models.Expense

func main() {
	cache = make(map[string]models.Expense, 0)

	expenseOne := models.Expense{
		Name:  "Arrendamento",
		Value: 68_000,
	}

	cache["arrendamento"] = expenseOne

	fmt.Printf("%+v\r\n", cache)

	if expense, found := cache["arrendamento"]; found {
		fmt.Printf("%+v\r\n", expense)
	}

	for key, value := range cache {
		fmt.Printf("%s = %+v\r\n", key, value.Value)
	}

	fmt.Printf("tamanho do mapa: %d\r\n", len(cache))

	delete(cache, "arrendamento")

	fmt.Printf("tamanho do mapa: %d\r\n", len(cache))
}
