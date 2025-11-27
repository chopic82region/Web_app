package models

import (
	"encoding/json"
	"log"
	"time"
)

type ErrMessage struct {
	Message string
	Time    time.Time
}

func NewErrMessage(err string) *ErrMessage {
	return &ErrMessage{
		Message: err,
		Time:    time.Now(),
	}
}

func (e *ErrMessage) MessageToString() string {
	b, err := json.MarshalIndent(e, " ", "	")
	if err != nil {
		log.Fatal("Error of message to json")
	}

	return string(b)
}
