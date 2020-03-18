package main

import(
	log "github.com/sirupsen/logrus"
	"fmt"
)

func main() {
	fmt.Print("hello world")
	log.Info("hello world 2")
}