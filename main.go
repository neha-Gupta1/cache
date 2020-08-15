package main

import (
	"github.com/cache/controller"
	"github.com/cache/model"
)

func main() {
	model.DBSetup()
	controller.Mycontroller()
}
