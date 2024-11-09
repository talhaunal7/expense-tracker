package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/talhaunal7/expense-tracker/server/middleware"
	"github.com/talhaunal7/expense-tracker/server/model"
	"github.com/talhaunal7/expense-tracker/server/service"
	"log"
	"net/http"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

func (uc *UserController) Register(ctx *gin.Context) {

	var user model.UserRegister

	if err := ctx.ShouldBindJSON(&user); err != nil {
		log.Print(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	err := uc.UserService.Register(&user)

	if err != nil {
		log.Print(err.Error())
		ctx.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "successfully registered"})

}

func (uc *UserController) Login(ctx *gin.Context) {

	var user model.UserLogin

	if err := ctx.ShouldBindJSON(&user); err != nil {
		log.Print(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	userResponse, token, err := uc.UserService.Login(&user)

	if err != nil {
		log.Print(err.Error())
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	tokenString := "Bearer " + *token
	ctx.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{"username": userResponse.FirstName + " " + userResponse.LastName, "id": userResponse.ID})
}

func (uc *UserController) Logout(ctx *gin.Context) {
	userId := middleware.GetUserIdFromContext(ctx)
	err := uc.UserService.Logout(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/users")
	userRoute.POST("/register", uc.Register)
	userRoute.POST("/login", uc.Login)
	userRoute.POST("/logout", middleware.ValidateToken(), uc.Logout)
}
