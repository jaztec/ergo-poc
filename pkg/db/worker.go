package db

import (
	"context"
	"ergo.services/ergo/act"
	"ergo.services/ergo/gen"
	db_gen "github.com/jaztec/ergo-poc/pkg/db/gen"
	"github.com/jaztec/ergo-poc/pkg/messages"
	"time"
)

const timeout = 50 * time.Second

type Worker struct {
	act.Actor

	queries *db_gen.Queries
}

func (w *Worker) Init(args ...any) error {
	w.queries = args[0].(*db_gen.Queries)

	return nil
}

func (w *Worker) HandleCall(from gen.PID, _ gen.Ref, request any) (any, error) {
	ctx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	switch req := request.(type) {

	case messages.TaskList:
		return w.listTasks(ctx, req)

	case messages.TaskById:
		return w.taskById(ctx, req)

	case messages.CreateTask:
		return w.createTask(ctx, req)

	case messages.UpdateTask:
		return w.updateTask(ctx, req)

	}

	return messages.Task{}, gen.ErrUnknown
}

func (w *Worker) HandleMessage(_ gen.PID, message any) error {
	ctx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	switch req := message.(type) {

	case messages.CreateTask:
		_, err := w.createTask(ctx, req)

		return err

	case messages.UpdateTask:
		_, err := w.updateTask(ctx, req)

		return err

	}

	return gen.ErrUnknown
}

func NewWorker() gen.ProcessBehavior {
	return &Worker{}
}
