package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	avitoproj "github.com/romankuz19/avito-proj"
)

type HistoryRequestBody struct {
	Date string
}

// @Summary Get user history
// @Tags users
// @Description get user history
// @ID get-user-history
// @Accept  json
// @Produce  json
// @Param requestBody body HistoryRequestBody true "date ('YY-MM')"
// @Success 200 {array} avitoproj.UserHistory
// @Failure 400 {integer} json "message: Failed"
// @Router /api/users/:id/history [post]
func (h *Handler) getUserHistory(c *gin.Context) {

	var requestBody HistoryRequestBody
	var history []avitoproj.UserHistory
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
		history = h.services.History.GetUserHistory(userId, requestBody.Date)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed",
			})
		} else {
			c.JSON(http.StatusOK, history)
		}
	}

}
