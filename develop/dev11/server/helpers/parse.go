package helpers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"dev11/server/domain"
)

func ParseUserDate(r *http.Request) (int64, time.Time, error) {
	userId := r.URL.Query().Get("user_id")
	date := r.URL.Query().Get("date")

	if userId == "" || date == "" {
		return 0, time.Now(), errors.New("bad query")
	}

	uid, err := strconv.Atoi(userId)
	if err != nil {
		return 0, time.Now(), err
	}

	tm, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0, time.Now(), err
	}

	return int64(uid), tm, nil
}

func ParseEvent(r *http.Request) (*domain.Event, error) {
	var t domain.Event

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func ParseDelete(r *http.Request) (int64, int64, error) {
	var t domain.DeleteReq

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		return 0, 0, err
	}

	return t.UserID, t.Id, nil
}
