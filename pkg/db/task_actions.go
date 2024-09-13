package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	db_gen "github.com/jaztec/ergo-poc/pkg/db/gen"
	"github.com/jaztec/ergo-poc/pkg/messages"
)

func (w *Worker) listTasks(ctx context.Context, req messages.TaskList) ([]messages.Task, error) {
	page := req.Page
	if page < 1 {
		page = 1
	}

	params := db_gen.ListTasksParams{
		Offset: page - 1,
		Limit:  req.Limit,
	}

	tasks, err := w.queries.ListTasks(ctx, params)

	return taskListToModel(tasks), err
}

func (w *Worker) taskById(ctx context.Context, req messages.TaskById) (messages.Task, error) {
	id := pgtype.UUID{}
	if err := id.Scan(req.ID); err != nil {
		return messages.Task{}, err
	}

	task, err := w.queries.GetTask(ctx, id)

	return taskToModel(task), err
}

func (w *Worker) createTask(ctx context.Context, req messages.CreateTask) (messages.Task, error) {
	task, err := w.queries.InsertTask(ctx, db_gen.InsertTaskParams{
		Name:        req.Name,
		Description: convertNilStringType(req.Description),
	})

	return taskToModel(task), err
}

func (w *Worker) updateTask(ctx context.Context, req messages.UpdateTask) (messages.Task, error) {
	id := pgtype.UUID{}
	if err := id.Scan(req.ID); err != nil {
		return messages.Task{}, err
	}

	task, err := w.queries.GetTask(ctx, id)
	if err != nil {
		return messages.Task{}, err
	}

	var updateParams db_gen.UpdateTaskParams
	updateParams.ID = id

	if req.Name != nil {
		updateParams.Name = *req.Name
	} else {
		updateParams.Name = task.Name
	}

	if req.Description != nil {
		updateParams.Description = convertNilStringType(req.Description)
	} else {
		updateParams.Description = task.Description
	}

	if req.Done != nil {
		updateParams.Done = pgtype.Bool{Bool: *req.Done, Valid: true}
	} else {
		updateParams.Done = task.Done
	}

	updatedTask, err := w.queries.UpdateTask(ctx, updateParams)

	return taskToModel(updatedTask), err
}
