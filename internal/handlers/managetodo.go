package managetodo

import (
	"log"
	"net/http"
	"todo/internal/database"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Conn *database.Connection
}

type NewTask struct {
	Text   string `json:"new_task"`
	IdUser string `json:"id_user"`
}

type ReadTasks struct {
	Iduser string `json:"id_user"`
}

type DeleteTask struct {
	TaskId string `json:"task_id"`
}

type UpdateTask struct {
	Task   string `json:"task"`
	TaskId string `json:"task_id"`
}

// Read Task Handler
func (h *Handler) ReadTasks(c *gin.Context) {
	var readtasks ReadTasks
	if err := c.BindJSON(&readtasks); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	log.Println(readtasks.Iduser)
	taskList := h.Conn.ReadToDo(readtasks.Iduser)

	c.JSON(http.StatusOK, gin.H{
		"tasks": taskList,
	})
}

// Add new task handler
func (h *Handler) AddTask(c *gin.Context) {
	var newTask NewTask
	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	h.Conn.AddTask(newTask.Text, newTask.IdUser)
	c.JSON(200, gin.H{"message": "Task was added successfully"})
}

// Delete task handler
func (h *Handler) DeleteTask(c *gin.Context) {
	var deleteTask DeleteTask
	if err := c.BindJSON(&deleteTask); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	h.Conn.DeleteTask(deleteTask.TaskId)
	c.JSON(http.StatusOK, gin.H{
		"message": "task was deleted successfully"})
}

// Update task handler
func (h *Handler) UpdateTask(c *gin.Context) {
	var updateTask UpdateTask
	if err := c.BindJSON(&updateTask); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	h.Conn.UpdateTask(updateTask.Task, updateTask.TaskId)
	c.JSON(http.StatusOK, gin.H{
		"message": "Task was updated successfully",
	})
}
