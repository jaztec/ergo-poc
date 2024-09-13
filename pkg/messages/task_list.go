package messages

type TaskList struct {
	Page  int32 `json:"page"`
	Limit int32 `json:"limit"`
}
