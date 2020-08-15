package main

import (
	"cache/controller"
	"cache/model"
)

func main() {
	model.DBSetup()
	controller.Mycontroller()
}
