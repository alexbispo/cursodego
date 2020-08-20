package main

import (
	"cursodego/go_build/models"
	"cursodego/go_build/repositories"
	"encoding/json"
	"fmt"
	"strconv"
)

func main() {
	expensesRepository := repositories.NewExpensesRepository()

	expensesRepository.Save(&models.Expense{
		Name:  "Arrendamento",
		Value: 68_000,
	})

	expensesRepository.Save(&models.Expense{
		Name:  "EDP",
		Value: 2_000,
	})

	for i := 1; i < 4; i++ {
		id := strconv.Itoa(i)

		expense, err := expensesRepository.Get(id)

		if err == repositories.ErrExpenseNotFound {
			fmt.Println("[main] Ocorreu um erro 404", err.Error())
			break
		}

		expenseJSON, err := json.Marshal(expense)

		if err != nil {
			fmt.Println("[main] Ocorreu um erro ", err.Error())
			break
		}

		fmt.Println("Despesa:", string(expenseJSON))
	}
}
