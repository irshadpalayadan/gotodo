package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func GetTodo(ctx *gin.Context) {

	rData := &ReturnData{Status: "success", Msg: ""}

	if len(todos) == 0 {
		rData.Data = struct{}{}
	} else {
		rData.Data = todos
	}

	if err := ctx.ShouldBind(&rData); err == nil {
		ctx.JSON(http.StatusOK, rData)
		return
	}

	logrus.WithFields(logrus.Fields{
		"file": "GetTodo.go",
		"line": "40",
	}).Fatal(" error while parsing the json data")

	rData.Status = "failure"
	rData.Msg = "error while json parsing"
	rData.Data = "{ }"
	ctx.JSON(http.StatusOK, rData)
}
