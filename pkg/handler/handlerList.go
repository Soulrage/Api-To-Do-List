package handler

import (
	"To-Do/internal/models"
	"To-Do/internal/modelRequests"
	"To-Do/internal/service"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


// CreateTask
// @Summary      Create a new task
// @Description  Creates a new task in the system based on the provided JSON payload.
// @Produce      application/json
// @Tags         tasks
// @Accept       json
// @Param        task body models.CreateTaskRequest true "Task details"
// @Success      200  {object}  map[string]interface{}  "Task successfully created"
// @Failure      400  {object}  map[string]interface{}  "Invalid input"
// @Failure      500  {object}  map[string]interface{}  "Could not create task"
// @Router       /api/CreateTask [post]
func (h *Handler) CreateTask(c *gin.Context) {
	var task models.CreateTaskRequest
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "data": task})
		return
	}

	err := service.CreateTask(h.DBConnect, task.Title, task.Description, task.DueDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": "successful", "message": "Task created"})
}





// GetTasks
// @Summary      Retrieve tasks
// @Description  Gets a list of tasks from the system. If an ID is provided, retrieves a single task by ID. Otherwise, retrieves all tasks.
// @Produce      application/json
// @Tags         tasks
// @Param        id  query     int  false  "Task ID" // Измените "path" на "query" для строки запроса
// @Success      200  {object}   models.Tasks  "List of tasks or single task"
// @Failure      404  {object}  map[string]interface{}  "Task not found"
// @Failure      500  {object}  map[string]interface{}  "Internal server error"
// @Router       /api/tasks [get]
func (h *Handler) GetTasks(c *gin.Context) {
	idParam := c.Query("id")
	if idParam != "" {
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		task, err := service.GetTasksById(h.DBConnect, uint(id))
		if err != nil {
			if err == service.ErrTaskNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve task"})
			}
			return
		}
		c.JSON(http.StatusOK, task)
	} else {
		tasks, err := service.GetAllTasks(h.DBConnect)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve tasks"})
			return
		}
		c.JSON(http.StatusOK, tasks)
	}
}





// UpdTask
// @Summary      Update a task
// @Description  Updates an existing task in the system based on the provided JSON payload.
// @Produce      application/json
// @Tags         tasks
// @Accept       json
// @Param        updatedTask body modelRequests.UpdateTaskRequest true "Task details"
// @Success      200  {object}  map[string]interface{}  "Task successfully updated"
// @Failure      400  {object}  map[string]interface{}  "Invalid input"
// @Failure      500  {object}  map[string]interface{}  "Could not update task"
// @Router       /api/UpdTask [put]
func (h *Handler) UpdTask(c *gin.Context) {
	var updatedTask models.Tasks
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedTask.UpdateAt = service.GetCurrentTimeRFC3339()
	err := service.UpdateTask(h.DBConnect, updatedTask.ID, updatedTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}





// DeleteTask
// @Summary      Delete a task
// @Description  Deletes an existing task from the system by ID.
// @Produce      application/json
// @Tags         tasks
// @Param        id    path      int    true  "Task ID"
// @Success      204   {object}  map[string]interface{}  "Task successfully deleted"
// @Failure      404   {object}  map[string]interface{}  "Task not found"
// @Failure      500   {object}  map[string]interface{}  "Could not delete task"
// @Router       /api/DeleteTask/{id} [delete]
func (h *Handler) DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	err = service.DeleteTaskById(h.DBConnect, uint(id))
	if err != nil {
		if errors.Is(err, service.ErrTaskNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete task"})
		return
	}

	c.Status(http.StatusNoContent)
}


// Auth
// @Summary      User Authentication
// @Description  Authenticates user and generates JWT token.
// @Produce      application/json
// @Tags         auth
// @Accept       json
// @Param        user body modelRequests.AuthUserRequest true "User credentials"
// @Success      200  {object}  map[string]string  "JWT token"
// @Failure      400  {object}  map[string]interface{}  "Invalid input"
// @Failure      500  {object}  map[string]interface{}  "User does not exist"
// @Router       /api/Auth [post]
func (h *Handler) Auth(c *gin.Context) {
	var user modelRequests.AuthUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := service.GenerateToken(h.DBConnect, user.Login, user.Password)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Such user does not exist"})
		return
	}
	c.JSON(http.StatusOK, token)
}



// Registration
// @Summary      User Registration
// @Description  Registers a new user in the system.
// @Produce      application/json
// @Tags         auth
// @Accept       json
// @Param        user body modelRequests.RegisterUserRequests true "User credentials"
// @Success      200  {object}  map[string]string  "User successfully created"
// @Failure      400  {object}  map[string]interface{}  "Invalid input"
// @Failure      500  {object}  map[string]interface{}  "User already exists"
// @Router       /api/Registration [post]
func (h *Handler) Registration(c *gin.Context) {
	var user modelRequests.RegisterUserRequests
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := service.RegistrationUsers(h.DBConnect, user.Login, user.Password, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "This user already exists"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"create_user": "successful", "message": "User created"})
}

