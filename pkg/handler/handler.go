package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	swag "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Handler struct {
	DBConnect *gorm.DB

}

func (h *Handler) InitRoutes() * gin.Engine{
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Accept"},
	}))

	dsn := "host=db user=Oleg password=Oleg dbname=ToDo port=5432"
	var err error
	h.DBConnect, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("Open DB failed")
	}
	router.GET("/swagger/*any", swag.WrapHandler(swaggerFiles.Handler))
	r := router.Group("/api")
	{
		r.POST("/CreateTask", h.CreateTask)
		r.POST("/Registration", h.Registration)
		r.POST("/Auth", h.Auth)
		r.GET("/tasks", h.GetTasks)
		r.PUT("/UpdTask", h.UpdTask)
		r.DELETE("/DeleteTask/:id", h.DeleteTask)}

	return router
}