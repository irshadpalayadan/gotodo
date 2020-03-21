package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetTodo(ctx *gin.Context) {

	rData := &ReturnData{Status: "success", Msg: ""} // ReturnData defined in TodoGlobal.go

	if len(todos) == 0 { // todos defined in TodoGlobal.go
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
