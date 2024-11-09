package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/talhaunal7/expense-tracker/server/middleware"
	"github.com/talhaunal7/expense-tracker/server/model"
	"github.com/talhaunal7/expense-tracker/server/service"
	"log"
	"net/http"
)

type ExpenseController struct {
	ExpenseService service.ExpenseService
}

func NewExpenseController(expenseService service.ExpenseService) ExpenseController {
	return ExpenseController{
		ExpenseService: expenseService,
	}
}

func (exp *ExpenseController) AddExpense(ctx *gin.Context) {

	var expenseAdd model.ExpenseAdd
	if err := ctx.ShouldBindJSON(&expenseAdd); err != nil {
		log.Print(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
	}
	userId := middleware.GetUserIdFromContext(ctx)
	if err := exp.ExpenseService.Add(expenseAdd, int(userId)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "successfully added"})

}
