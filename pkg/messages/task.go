package messages

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Done        bool      `json:"done"`
}
