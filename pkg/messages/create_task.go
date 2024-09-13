package messages

type CreateTask struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
