package todo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

// this is some what incorrect if multiple go files uses
// same global object/type it should be specified in the
// global constant file with same package

type TodoItem struct {
	ID           string `json:"id"`
	USERNAME     string `json:"user"`
	CreatedTime  int64  `json:"createdat`
	Content      string `json:"content"`
	Reminder     bool   `json:"reminder"`
	ReminderTime int64  `json:"remindat"`
}

type ReturnData struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
}

var todos []TodoItem

func GetTodo(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	rData := &ReturnData{Status: "success", Msg: ""}

	if len(todos) == 0 {
		rData.Data = struct{}{}
	} else {
		rData.Data = todos
	}

	todoJson, err := json.Marshal(rData)
	if err != nil {

		logrus.WithFields(logrus.Fields{
			"file": "GetTodo.go",
			"line": "40",
		}).Fatal(" error while parsing the json data")

		rData.Status = "failure"
		rData.Msg = "error while json parsing"
		rData.Data = "{ }"
		todoJson, err := json.Marshal(rData)
		if err != nil {
			fmt.Fprintln(res, string(todoJson))
			return
		}
	}
	fmt.Fprint(res, string(todoJson))
}
