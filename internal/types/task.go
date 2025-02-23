package types

type Task struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	IsDone     bool   `json:"is_done"`
	TaskListID int    `json:"task_list_id"`
}
