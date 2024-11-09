package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/talhaunal7/expense-tracker/server/middleware"
	"github.com/talhaunal7/expense-tracker/server/model/request"
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

	var expenseAdd request.ExpenseAdd
	if err := ctx.ShouldBindJSON(&expenseAdd); err != nil {
		log.Print(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	userId := middleware.GetUserIdFromContext(ctx)
	if err := exp.ExpenseService.Add(expenseAdd, userId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "successfully added"})
}

func (exp *ExpenseController) GetExpense(ctx *gin.Context) {
	expenseId := ctx.Param("id")
	userId := middleware.GetUserIdFromContext(ctx)

	if expenseDto, err := exp.ExpenseService.Get(expenseId, userId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"expense": expenseDto})
	}
	return
}

func (exp *ExpenseController) GetAll(ctx *gin.Context) {
	userId := middleware.GetUserIdFromContext(ctx)
	if allExpenses, err := exp.ExpenseService.GetAll(userId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"expenses": allExpenses})
	}

}

func (exp *ExpenseController) RegisterExpenseRoutes(rg *gin.RouterGroup) {
	expenseRoute := rg.Group("/expenses")
	expenseRoute.POST("/", middleware.ValidateToken(), exp.AddExpense)
	expenseRoute.GET("/:id", middleware.ValidateToken(), exp.GetExpense)
	expenseRoute.GET("/", middleware.ValidateToken(), exp.GetAll)
}
