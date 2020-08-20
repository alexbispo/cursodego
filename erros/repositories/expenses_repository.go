package repositories

import (
	"cursodego/erros/models"
	"errors"
	"strconv"
)

// ErrExpenseNotFound when expense is not found
var ErrExpenseNotFound = errors.New("Record Not Found")

// ExpensesRepository a repository operations to Expenses
type ExpensesRepository struct {
	cache      map[string]*models.Expense
	sequencyID int
}

// NewExpensesRepository creates a ExpensesRepository reference
func NewExpensesRepository() *ExpensesRepository {
	repo := &ExpensesRepository{
		cache:      make(map[string]*models.Expense, 0),
		sequencyID: 0,
	}

	return repo
}

// Save a record
func (repo *ExpensesRepository) Save(expense *models.Expense) *models.Expense {
	repo.sequencyID++
	expense.ID = strconv.Itoa(repo.sequencyID)
	repo.cache[expense.ID] = expense
	return expense
}

// Get retrieve a record by id
func (repo *ExpensesRepository) Get(id string) (expense *models.Expense, err error) {
	expense, found := repo.cache[id]

	if !found {
		err = ErrExpenseNotFound
		return
	}

	return
}
