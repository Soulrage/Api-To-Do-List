package modelRequests


type UpdateTaskRequest struct {
	ID 			uint `json:"ID"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	DueDate     string `json:"due_date" binding:"required"`
}