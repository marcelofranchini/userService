package handler

// import (
// 	"net/http"
// 	"your_project/internal/user/app/useCase"
// 	"your_project/internal/user/domain"

// 	"github.com/gin-gonic/gin"
// )

// // LoginHandler é responsável por lidar com o login de um usuário
// type LoginHandler struct {
// 	loginUseCase *useCase.LoginUseCase
// }

// // NewLoginHandler cria uma nova instância do LoginHandler
// func NewLoginHandler(loginUseCase *useCase.LoginUseCase) *LoginHandler {
// 	return &LoginHandler{loginUseCase: loginUseCase}
// }

// // Login é o método que trata a requisição HTTP para realizar o login do usuário
// func (h *LoginHandler) Login(c *gin.Context) {
// 	var loginDetails domain.LoginDetails
// 	if err := c.ShouldBindJSON(&loginDetails); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Chama o caso de uso do login
// 	token, err := h.loginUseCase.Execute(loginDetails)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Retorna o token de autenticação
// 	c.JSON(http.StatusOK, gin.H{"token": token})
// }
