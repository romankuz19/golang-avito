package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SectionRequestBody struct {
	Slug string
}

type AddDeleteSections struct {
	SectionsAdd    []string
	SectionsDelete []string
}

// @Summary Create new section
// @Tags sections
// @Description create new section
// @ID create-section
// @Accept  json
// @Produce  json
// @Param requestBody body SectionRequestBody true "section (slug) name"
// @Success 200 {integer} json "message: Section created"
// @Failure 400 {integer} json "message: Failed"
// @Router /api/sections [post]
func (h *Handler) createSection(c *gin.Context) {

	var requestBody SectionRequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect input",
		})
	}

	err := h.services.Section.Create(requestBody.Slug)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Section created",
		})
	}
}

// @Summary Delete section
// @Tags sections
// @Description delete section
// @ID delete-section
// @Accept  json
// @Produce  json
// @Param requestBody body SectionRequestBody true "section (slug) name"
// @Success 200 {integer} json "message: Section deleted"
// @Failure 400 {integer} json "message: Failed"
// @Router /api/sections [delete]
func (h *Handler) deleteSection(c *gin.Context) {

	var requestBody SectionRequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect input",
		})
	}

	err := h.services.Section.Delete(requestBody.Slug)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Section deleted",
		})
	}
}

// @Summary Add user in section
// @Tags sections-users
// @Description add or delete user in sections
// @ID add-user
// @Accept  json
// @Produce  json
// @Param requestBody body AddDeleteSections true "list of sections to add or delete"
// @Success 200 {integer} json "message: Success"
// @Failure 400 {integer} json "message: Failed"
// @Router /api/sections/users/:id [put]
func (h *Handler) addUser(c *gin.Context) {

	var requestBody AddDeleteSections
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect input",
		})
	}

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad query params",
		})
	} else {
		err = h.services.Section.AddUser(requestBody.SectionsAdd, requestBody.SectionsDelete, userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Success",
			})
		}
	}

}

// @Summary Get user sections
// @Tags sections-users
// @Description get user sections
// @ID get-user-sections
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Failure 400 {integer} json "message: Failed"
// @Router /api/sections/users/:id [get]
func (h *Handler) getUserSections(c *gin.Context) {

	var sections []string

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad query params",
		})
	} else {
		sections = h.services.Section.GetUserSections(userId)
		c.JSON(http.StatusOK, sections)
	}
}
