package helpers

import (
	"encoding/json"

	"dev11/server/domain"
)

func SerealizeErr(errStr string) ([]byte, error) {
	data, err := json.Marshal(domain.AnswerErr{errStr})
	if err != nil {
		return nil, err
	}

	return data, nil
}

func SerealizeOk(id int64) ([]byte, error) {
	data, err := json.Marshal(domain.AnswerOk{id})
	if err != nil {
		return nil, err
	}

	return data, nil
}

func SerealizeEvents(e []domain.Event) ([]byte, error) {
	data, err := json.Marshal(domain.AnswerEvents{e})
	if err != nil {
		return nil, err
	}

	return data, nil
}
