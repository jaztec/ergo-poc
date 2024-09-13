package db

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	db_gen "github.com/jaztec/ergo-poc/pkg/db/gen"
	"github.com/jaztec/ergo-poc/pkg/messages"
)

func convertNilStringType(in *string) pgtype.Text {
	if in == nil {
		return pgtype.Text{
			String: "",
		}
	}

	return pgtype.Text{
		Valid:  true,
		String: *in,
	}
}

func taskListToModel(tasks []db_gen.Task) []messages.Task {
	result := make([]messages.Task, 0, len(tasks))

	for _, task := range tasks {
		result = append(result, taskToModel(task))
	}

	return result
}

func taskToModel(task db_gen.Task) messages.Task {
	encoded, _ := task.ID.Value()
	id, _ := uuid.Parse(encoded.(string))

	var description *string

	if task.Description.Valid {
		description = &task.Description.String
	} else {
		description = nil
	}

	return messages.Task{
		ID:          id,
		Name:        task.Name,
		Description: description,
		CreatedAt:   task.CreatedAt.Time,
		UpdatedAt:   task.UpdatedAt.Time,
		Done:        task.Done.Bool,
	}
}
