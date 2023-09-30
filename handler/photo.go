package handler

import (
	"crowdfunding/helper"
	"crowdfunding/photo"
	"crowdfunding/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type photoHandler struct {
	service photo.Service
}

func NewPhotoHandler(service photo.Service) *photoHandler {
	return &photoHandler{service}
}

func (h *photoHandler) CreatePhoto(c *gin.Context) {
	var input photo.CreatePhotoInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors":errors}

		response := helper.APIResponse("failed to create photo", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User) //ngambil user yang sedang login
	input.User = currentUser
	userID := currentUser.ID

	file, err := c.FormFile("images")
	if err != nil {

		emptyPath := ""
		_, err := h.service.CreatePhoto(input, emptyPath)
		if err != nil {
			data := gin.H{"is_uploaded": false}
			response := helper.APIResponse("Failed to create photo", http.StatusBadRequest, "error", data)

			c.JSON(http.StatusBadRequest, response)
			return
		}

		data := gin.H{"is_uploaded": true}
		response := helper.APIResponse("Photo successfully created (without image)", http.StatusOK, "success", data)

		c.JSON(http.StatusOK, response)
		return
	}

	path := fmt.Sprint("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload photo", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.CreatePhoto(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload photo", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("photo successfuly uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}

func (h *photoHandler) GetPhotos(c *gin.Context) {

	photos, err := h.service.GetPhotos()
	if err != nil {
		response := helper.APIResponse("Error to get photos", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("List of photo", http.StatusOK, "success", photo.FormatPhotos(photos))
	c.JSON(http.StatusOK, response)

}

func (h *photoHandler) UpdatePhoto(c *gin.Context) {

	var id photo.GetId
	err := c.ShouldBindUri(&id)
	if err != nil {
		response := helper.APIResponse("Failed to update photo", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData photo.UpdatePhotoInput
	err = c.ShouldBind(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update photo", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser
	userID := currentUser.ID

	file, err := c.FormFile("images")
	if err != nil {
		emptyPath := ""
		_, err := h.service.UpdatePhoto(id, inputData, emptyPath)
		if err != nil {
			data := gin.H{"is_uploaded": false}
			response := helper.APIResponse("Failed to create photo", http.StatusBadRequest, "error", data)

			c.JSON(http.StatusBadRequest, response)
			return
		}

		data := gin.H{"is_uploaded": true}
		response := helper.APIResponse("Photo successfully created (without image)", http.StatusOK, "success", data)

		c.JSON(http.StatusOK, response)
		return
	}

	path := fmt.Sprint("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload photo", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.UpdatePhoto(id, inputData, path)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to update user", http.StatusForbidden, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Success to update photo", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)


}

func (h *photoHandler) Delete(c *gin.Context) {

	var id photo.GetId
	err := c.ShouldBindUri(&id)
	if err != nil {
		response := helper.APIResponse("Failed to delete photo", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	_, err = h.service.Delete(id.ID, currentUser.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed to delete photo", http.StatusForbidden, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to delete photo", http.StatusOK, "success", "Success deleted photo")
	c.JSON(http.StatusOK, response)


}