package handler

import (
	"crowdfunding/auth"
	"crowdfunding/helper"
	"crowdfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service 
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var errors []string

		errors = helper.FormatValidationError(err)

		errorMessage := gin.H{ "errors" : errors }

		response := helper.APIResponse("Register user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return 
	}

	 _, err = h.userService.IsEmailAvailable(input.Email)
	 if err != nil {
		errorMessage := gin.H{"errors": err.Error()} //memanggil error yang ada di service
		response := helper.APIResponse("Register user failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response) 
		return
	 }

	 newUser, err := h.userService.RegisterUser(input)
	 if err != nil {
		response := helper.APIResponse("Register user failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return 
	 }

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.APIResponse("Register user failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, token)

	response := helper.APIResponse("Register User Successfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) Login(c *gin.Context) {

	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{ "errors" : errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.LoginUser(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()} //memanggil error yang ada di service
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response) 
		return
	}

	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	formatter := user.FormatUser(loggedinUser, token)

	response := helper.APIResponse("Login Success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)

	formatter := user.FormatUser(currentUser, "")

	response := helper.APIResponse("Successfully fetch user data", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) UpdateUser(c *gin.Context) {

	var id user.GetId
	err := c.ShouldBindUri(&id)
	if err != nil {
		response := helper.APIResponse("Failed to update user", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData user.UpdateUserInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update user", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	inputData.UserLogin = currentUser

	updatedUser, err := h.userService.UpdateUser(id, inputData)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed to update user", http.StatusForbidden, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update user", http.StatusOK, "success", updatedUser)
	c.JSON(http.StatusOK, response)


}

func (h *userHandler) DeleteUser(c *gin.Context) {

	var id user.GetId
	err := c.ShouldBindUri(&id)
	if err != nil {
		response := helper.APIResponse("Failed to delete user", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	_, err = h.userService.Delete(id.ID, currentUser.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed to delete user", http.StatusForbidden, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to delete user", http.StatusOK, "success", "Success deleted user")
	c.JSON(http.StatusOK, response)


}