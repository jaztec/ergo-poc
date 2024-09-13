package messages

type UpdateTask struct {
	ID          string
	Name        *string
	Description *string
	Done        *bool
}
