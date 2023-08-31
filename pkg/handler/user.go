package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createUser(c *gin.Context) {
	type UserRequestBody struct {
		Name string
	}
	var requestBody UserRequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		panic(err)
	}

	err := h.services.User.Create(requestBody.Name)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "User created",
		})
	}

}
