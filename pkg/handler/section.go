package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createSection(c *gin.Context) {
	type SectionRequestBody struct {
		Slug string
	}
	var requestBody SectionRequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		panic(err)
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

func (h *Handler) deleteSection(c *gin.Context) {
	type SectionRequestBody struct {
		Slug string
	}
	var requestBody SectionRequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		panic(err)
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

func (h *Handler) addUser(c *gin.Context) {

	type Data struct {
		SectionsAdd    []string
		SectionsDelete []string
	}

	var requestBody Data

	if err := c.BindJSON(&requestBody); err != nil {
		panic(err)
	}

	userId, err := strconv.Atoi(c.Param("id"))
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
func (h *Handler) getUserSections(c *gin.Context) {
	var sections []string
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	sections = h.services.Section.GetUserSections(userId)
	c.JSON(http.StatusOK, sections)
}
