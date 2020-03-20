package todo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func AddTodo(res http.ResponseWriter, req *http.Request) {
	var tempTodo TodoItem
	if req.Body == nil {
		http.Error(res, "please send request body", 400)
		return
	}

	err := json.NewDecoder(req.Body).Decode(&tempTodo)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}

	todos = append(todos, tempTodo)

	res.Header().Set("Content-Type", "application/json")
	rData := &ReturnData{Status: "success", Msg: "todo inserted successfully"}
	rDataJson, err := json.Marshal(rData)
	if err != nil {

		logrus.WithFields(logrus.Fields{
			"file": "GetTodo.go",
			"line": "40",
		}).Fatal(" error while parsing the json data")

		rData.Status = "failure"
		rData.Msg = "error while json parsing"
		rData.Data = "{ }"
		rDataJson, err := json.Marshal(rData)
		if err != nil {
			fmt.Fprintln(res, string(rDataJson))
			return
		}
	}
	fmt.Fprint(res, string(rDataJson))
}
