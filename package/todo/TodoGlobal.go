package todo

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// this is some what incorrect if multiple go files uses
// same global object/type it should be specified in the
// global constant file with same package

type TodoItem struct {
	ID           uuid.UUID `json:"id"`
	UserName     string    `json:"user"`
	CreatedTime  int64     `json:"createdat"`
	Content      string    `json:"content"`
	Reminder     bool      `json:"reminder"`
	ReminderTime int64     `json:"remindat"`
}

type ReturnData struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
}

func (todoItem *TodoItem) updateId() {
	id, err := uuid.NewUUID()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file": "TodoGlobal.go",
		}).Fatal("Error while creating new uuid for todo item id", err)
	}
	todoItem.ID = id
}

var todos []TodoItem
