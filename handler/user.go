package handler

import (
	"net/http"
	"start-up/helper"
	"start-up/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewServiceHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// get input from user
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	// token, err := h.jwtService.GenerateToken()

	formatter := user.FormatUser(newUser, "toktoktoktok")

	response := helper.APIResponse("Account has been created !", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}
