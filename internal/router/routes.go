package router

import (
	"todo/internal/database"
	managetodo "todo/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(db *database.Connection) *gin.Engine {
	r := gin.Default()
	h := managetodo.Handler{
		Conn: db,
	}

	r.GET("/todo/:id", h.ReadTasks)
	r.POST("/todo/:id", h.AddTask)
	r.DELETE("/todo/:id", h.DeleteTask)
	r.PUT("/todo/:id", h.UpdateTask)

	return r
}
