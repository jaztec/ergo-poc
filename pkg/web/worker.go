package web

import (
	"encoding/json"
	"ergo.services/ergo/act"
	"ergo.services/ergo/gen"
	"net/http"
)

const (
	databaseName = gen.Atom("database")
)

type Worker struct {
	act.WebWorker
}

func (w *Worker) Init(_ ...any) error {
	return nil
}

func (w *Worker) HandleGet(_ gen.PID, writer http.ResponseWriter, req *http.Request) error {

	writer.Header().Set("Content-Type", "application/json")

	msg, err := parseGetRequest(req)
	if err != nil {
		return handleError(writer, err)
	}

	res, err := w.Call(databaseName, msg)
	if err != nil {
		return handleError(writer, err)
	}

	b, err := json.Marshal(res)
	if err != nil {
		return handleError(writer, err)
	}

	_, err = writer.Write(b)

	return handleError(writer, err)
}

func (w *Worker) HandlePost(_ gen.PID, writer http.ResponseWriter, req *http.Request) error {
	writer.Header().Set("Content-Type", "application/json")

	create, err := parsePostRequest(req)
	if err != nil {
		return handleError(writer, err)
	}

	res, err := w.Call(databaseName, create)
	if err != nil {
		return handleError(writer, err)
	}

	b, err := json.Marshal(res)
	if err != nil {
		return handleError(writer, err)
	}

	writer.WriteHeader(http.StatusCreated)
	_, err = writer.Write(b)

	return nil
}

func (w *Worker) HandlePut(_ gen.PID, writer http.ResponseWriter, req *http.Request) error {
	writer.Header().Set("Content-Type", "application/json")

	update, err := parsePutRequest(req)
	if err != nil {
		return handleError(writer, err)
	}

	err = w.Send(databaseName, update)
	if err != nil {
		return handleError(writer, err)
	}

	writer.WriteHeader(http.StatusAccepted)

	return nil
}

func NewWorker() gen.ProcessBehavior {
	return &Worker{}
}
