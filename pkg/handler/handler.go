package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/romankuz19/avito-proj/docs"
	"github.com/romankuz19/avito-proj/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		users := api.Group("users")
		{
			users.POST("/", h.createUser)
		}
		sections := api.Group("sections")
		{
			sections.POST("/", h.createSection)
			sections.DELETE("/", h.deleteSection)

			users := sections.Group("users")
			{
				users.GET("/:id", h.getUserSections)
				users.PUT("/:id", h.addUser)
			}

		}
	}

	return router
}
