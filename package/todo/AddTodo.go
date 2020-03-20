package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AddTodo(ctx *gin.Context) {

	var tempTodo TodoItem

	if err := ctx.BindJSON(&tempTodo); err != nil {

		logrus.WithFields(logrus.Fields{
			"file": "AddTodo.go",
		}).Fatal("Error while parsing the body")

		ctx.JSON(http.StatusBadRequest, &ReturnData{ Status: "failure", Msg: "invalid data passed", Data: err })
		return
	}

	todos = append(todos, tempTodo)
	ctx.JSON(http.StatusOK, &ReturnData{ Status: "success", Msg: "todo inserted successfully"})
}
