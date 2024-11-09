package service

import (
	"github.com/talhaunal7/expense-tracker/server/entity"
	"github.com/talhaunal7/expense-tracker/server/model"
	"github.com/talhaunal7/expense-tracker/server/repository"
)

type ExpenseService interface {
	Add(expenseAddRequest model.ExpenseAdd, userId int) error
}

type ExpenseServiceImpl struct {
	expenseRepository repository.ExpenseRepository
}

func NewExpenseService(expenseRepository repository.ExpenseRepository) ExpenseService {
	return &ExpenseServiceImpl{
		expenseRepository: expenseRepository,
	}
}

func (exp ExpenseServiceImpl) Add(expenseAddRequest model.ExpenseAdd, userId int) error {

	expense := entity.Expense{
		Amount:      expenseAddRequest.Amount,
		Description: expenseAddRequest.Description,
		CategoryID:  expenseAddRequest.CategoryID,
		UserID:      userId,
	}

	if err := exp.expenseRepository.Add(expense); err != nil {
		return err
	}

	return nil
}
