package domain

import "time"

type Event struct {
	UserID int64     `json:"user_id"`
	Id     int64     `json:"id"`
	Name   string    `json:"name"`
	Desc   string    `json:"desc"`
	Date   time.Time `json:"date"`
}

type DeleteReq struct {
	UserID int64 `json:"user_id"`
	Id     int64 `json:"id"`
}

type AnswerOk struct {
	Result int64 `json:"result"`
}

type AnswerEvents struct {
	Result []Event `json:"result"`
}

type AnswerErr struct {
	Err string `json:"error"`
}
