package repository

import (
	"github.com/talhaunal7/expense-tracker/server/entity"
	"gorm.io/gorm"
)

type ExpenseRepository interface {
	Add(expense entity.Expense) error
}

type ExpenseRepositoryImpl struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) ExpenseRepository {
	return ExpenseRepositoryImpl{db: db}
}

func (expenseRepo ExpenseRepositoryImpl) Add(expense entity.Expense) error {
	result := expenseRepo.db.Create(&expense)
	return result.Error
}
