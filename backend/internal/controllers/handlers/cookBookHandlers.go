package handlers

import (
	"cook_book/backend/internal/model"
	"cook_book/backend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllRecipes(c *gin.Context) {
	result, err := h.cookService.GetAll()
	if err != nil {
		utils.SendResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendResponseJSON(c, http.StatusOK, result)
}

func (h *Handler) GetRecipeByID(c *gin.Context) {
	id, err := utils.CheckCorrectID(c)
	if err != nil {
		return
	}

	result, err := h.cookService.GetByID(id)
	if err != nil {
		utils.SendResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendResponseJSON(c, http.StatusOK, result)
}

func (h *Handler) CreateRecipe(c *gin.Context) {
	var createdRecipe model.CreateCookBook

	if !utils.BindJSON(c, &createdRecipe) {
		return
	}

	recipe := model.CookBook{
		Title:       createdRecipe.Title,
		Description: createdRecipe.Description,
	}

	if err := h.cookService.Create(&recipe); err != nil {
		utils.SendResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendResponseJSON(c, http.StatusCreated, "Successfully created")
}

func (h *Handler) UpdateRecipe(c *gin.Context) {
	id, err := utils.CheckCorrectID(c)
	if err != nil {
		return
	}

	var updatedRecipe model.UpdateCookBook

	if !utils.BindJSON(c, &updatedRecipe) {
		return
	}

	recipe := model.CookBook{
		Title:       updatedRecipe.Title,
		Description: updatedRecipe.Description,
	}

	if err := h.cookService.Update(&recipe, id); err != nil {
		utils.SendResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendResponseJSON(c, http.StatusOK, "Seccessfully updated")
}

func (h *Handler) DeleteRecipe(c *gin.Context) {
	id, err := utils.CheckCorrectID(c)
	if err != nil {
		return
	}

	if err := h.cookService.Delete(id); err != nil {
		utils.SendResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendResponseJSON(c, http.StatusOK, "Successfully deleted")
}
