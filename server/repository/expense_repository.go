package repository

import (
	"github.com/talhaunal7/expense-tracker/server/entity"
	"gorm.io/gorm"
)

type ExpenseRepository interface {
	Add(expense entity.Expense) error
	FindByID(id int) (entity.Expense, error)
	FindAllByUserID(userID int) ([]entity.Expense, error)
}

type ExpenseRepositoryImpl struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) ExpenseRepository {
	return ExpenseRepositoryImpl{db: db}
}

func (expenseRepo ExpenseRepositoryImpl) Add(expense entity.Expense) error {
	result := expenseRepo.db.Save(&expense)
	return result.Error
}

func (expenseRepo ExpenseRepositoryImpl) FindByID(id int) (entity.Expense, error) {
	var expense entity.Expense
	result := expenseRepo.db.Preload("Category").First(&expense, id)
	//result := expenseRepo.db.First(&expense, id)
	return expense, result.Error
}

func (expenseRepo ExpenseRepositoryImpl) FindAllByUserID(userID int) ([]entity.Expense, error) {
	var expenses []entity.Expense
	//result := expenseRepo.db.Preload("Category").Find(&expenses, userID)
	result := expenseRepo.db.Preload("Category").Where("user_id = ?", userID).Find(&expenses)
	return expenses, result.Error
}
