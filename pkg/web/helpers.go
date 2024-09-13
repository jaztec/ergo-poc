package web

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/jaztec/ergo-poc/pkg/messages"
	"net/http"
	"strconv"
)

func parseGetRequest(req *http.Request) (any, error) {
	var msg any

	if req.URL.Query().Has("id") {
		id := req.URL.Query().Get("id")

		// Validate
		taskId, err := uuid.Parse(id)
		if err != nil {
			return nil, err
		}

		msg = messages.TaskById{ID: taskId.String()}

	} else {
		var limit int32 = 10
		var page int32 = 1

		if req.URL.Query().Has("limit") {
			q, err := strconv.Atoi(req.URL.Query().Get("limit"))
			if err == nil {
				limit = int32(q)
			}
		}

		if req.URL.Query().Has("page") {
			q, err := strconv.Atoi(req.URL.Query().Get("page"))
			if err == nil {
				page = int32(q)
			}

			if page < 1 {
				return nil, errors.New("page must be greater than 0")
			}
		}

		msg = messages.TaskList{
			Limit: limit,
			Page:  page,
		}

	}

	return msg, nil
}

func handleError(w http.ResponseWriter, err error) error {
	if err == nil {
		return nil
	}

	content := map[string]string{
		"error": err.Error(),
	}
	b, _ := json.Marshal(content)

	_, _ = w.Write(b)

	return err
}
