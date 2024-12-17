package handler

import (
	"net/http"
	"userService/internal/user/app/useCase"
	"userService/internal/user/domain"

	"github.com/gin-gonic/gin"
)

type CreateUserHandler struct {
	createUserUseCase *useCase.CreateUserUseCase
}

func NewCreateUserHandler(createUserUseCase *useCase.CreateUserUseCase) *CreateUserHandler {
	return &CreateUserHandler{createUserUseCase: createUserUseCase}
}

func (h *CreateUserHandler) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userCreateRequest := domain.UserCreateRequest{
		Email:    user.Email,
		Password: user.Password,
	}

	if err := h.createUserUseCase.Execute(user, userCreateRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created"})
}
