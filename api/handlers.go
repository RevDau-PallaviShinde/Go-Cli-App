package api

import (
	"net/http"
	"strconv"

	"github.com/RevDau-PallaviShinde/Go-Cli-App/handlers"
	"github.com/RevDau-PallaviShinde/Go-Cli-App/models"
	"github.com/gin-gonic/gin"
)

// CreateTaskRequest holds the JSON body for creating a task.
type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Priority    string `json:"priority"`
}

// UpdateTaskRequest holds the JSON body for updating a task.
type UpdateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
}

// ErrorResponse is a standard error reply.
type ErrorResponse struct {
	Error string `json:"error"`
}

// MessageResponse is a standard success reply.
type MessageResponse struct {
	Message string `json:"message"`
}

// SummaryResponse holds the task summary data.
type SummaryResponse struct {
	Total     int `json:"total"`
	Completed int `json:"completed"`
	Pending   int `json:"pending"`
}

// CreateTask godoc
// @Summary      Create Task
// @Description  Add a new task
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        task  body      CreateTaskRequest  true  "Task data"
// @Success      201   {object}  MessageResponse
// @Failure      400   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /tasks [post]
func CreateTask(c *gin.Context) {
	var req CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	if err := handlers.AddTask(req.Title, req.Description, req.DueDate, req.Priority); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, MessageResponse{Message: "task created successfully"})
}

// GetTasks godoc
// @Summary      List Tasks
// @Description  Get all tasks with optional filters
// @Tags         tasks
// @Produce      json
// @Param        status    query     string  false  "Filter by status (pending, in-progress, completed)"
// @Param        priority  query     string  false  "Filter by priority (low, medium, high)"
// @Success      200       {array}   models.Task
// @Failure      500       {object}  ErrorResponse
// @Router       /tasks [get]
func GetTasks(c *gin.Context) {
	status := c.Query("status")
	priority := c.Query("priority")

	tasks, err := handlers.ListTasks(status, priority)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	// Return empty array instead of null
	if tasks == nil {
		tasks = []models.Task{}
	}

	c.JSON(http.StatusOK, tasks)
}

// UpdateTask godoc
// @Summary      Update Task
// @Description  Update an existing task by ID
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id    path      int                true  "Task ID"
// @Param        task  body      UpdateTaskRequest  true  "Fields to update"
// @Success      200   {object}  MessageResponse
// @Failure      400   {object}  ErrorResponse
// @Failure      404   {object}  ErrorResponse
// @Router       /tasks/{id} [put]
func UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid task ID"})
		return
	}

	var req UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	if err := handlers.UpdateTask(id, req.Title, req.Description, req.DueDate, req.Status, req.Priority); err != nil {
		if err.Error() == "task not found" {
			c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, MessageResponse{Message: "task updated successfully"})
}

// DeleteTask godoc
// @Summary      Delete Task
// @Description  Delete a task by ID
// @Tags         tasks
// @Produce      json
// @Param        id   path      int  true  "Task ID"
// @Success      200  {object}  MessageResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Router       /tasks/{id} [delete]
func DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid task ID"})
		return
	}

	if err := handlers.DeleteTask(id); err != nil {
		if err.Error() == "task not found" {
			c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, MessageResponse{Message: "task deleted successfully"})
}

// SearchTasks godoc
// @Summary      Search Tasks
// @Description  Search tasks by keyword in title or description
// @Tags         tasks
// @Produce      json
// @Param        q    query     string  true  "Search keyword"
// @Success      200  {array}   models.Task
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /tasks/search [get]
func SearchTasks(c *gin.Context) {
	keyword := c.Query("q")

	tasks, err := handlers.SearchTasks(keyword)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	if tasks == nil {
		tasks = []models.Task{}
	}

	c.JSON(http.StatusOK, tasks)
}

// GetSummary godoc
// @Summary      Task Summary
// @Description  Get counts of total, completed, and pending tasks
// @Tags         summary
// @Produce      json
// @Success      200  {object}  SummaryResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /tasks/summary [get]
func GetSummary(c *gin.Context) {
	total, completed, pending, err := handlers.GetSummary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SummaryResponse{
		Total:     total,
		Completed: completed,
		Pending:   pending,
	})
}

// ExportTasks godoc
// @Summary      Export Tasks
// @Description  Export all tasks to a CSV file
// @Tags         tasks
// @Produce      json
// @Param        file  query     string  false  "Output file name"  default(tasks.csv)
// @Success      200   {object}  MessageResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /tasks/export [get]
func ExportTasks(c *gin.Context) {
	fileName := c.DefaultQuery("file", "tasks.csv")

	if err := handlers.ExportTasksToCSV(fileName); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, MessageResponse{Message: "tasks exported to " + fileName})
}
