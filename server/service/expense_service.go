package service

import (
	"errors"
	"github.com/talhaunal7/expense-tracker/server/entity"
	"github.com/talhaunal7/expense-tracker/server/model/dto"
	"github.com/talhaunal7/expense-tracker/server/model/request"
	"github.com/talhaunal7/expense-tracker/server/repository"
	"strconv"
)

type ExpenseService interface {
	Add(expenseAddRequest request.ExpenseAdd, userId int) error
	Get(expenseId string, userId int) (*dto.ExpenseDto, error)
	GetAll(userId int) ([]*dto.ExpenseDto, error)
}

type ExpenseServiceImpl struct {
	expenseRepository repository.ExpenseRepository
}

func NewExpenseService(expenseRepository repository.ExpenseRepository) ExpenseService {
	return &ExpenseServiceImpl{
		expenseRepository: expenseRepository,
	}
}

func (exp ExpenseServiceImpl) Add(expenseAddRequest request.ExpenseAdd, userId int) error {
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

func (exp ExpenseServiceImpl) Get(expenseId string, userId int) (*dto.ExpenseDto, error) {

	expId, err := strconv.Atoi(expenseId)
	if err != nil {
		return nil, err
	}
	expense, err := exp.expenseRepository.FindByID(expId)
	if err != nil {
		return nil, err
	}
	if expense.UserID != userId {
		return nil, errors.New("expense not found")
	}
	expenseDto := dto.ExpenseDto{
		Amount:      expense.Amount,
		Description: expense.Description,
		Category:    expense.Category.Name, //TODO load ?
	}
	return &expenseDto, nil
}

func (exp ExpenseServiceImpl) GetAll(userId int) ([]*dto.ExpenseDto, error) {
	expenses, err := exp.expenseRepository.FindAllByUserID(userId)
	if err != nil {
		return nil, err
	}
	expenseDtos := make([]*dto.ExpenseDto, len(expenses))
	for i, expense := range expenses {
		expenseDtos[i] = &dto.ExpenseDto{
			Amount:      expense.Amount,
			Description: expense.Description,
			Category:    expense.Category.Name,
		}
	}
	return expenseDtos, nil
}
