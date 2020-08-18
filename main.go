package main

import (
	"github.com/cache/controller"
	"github.com/cache/model"
)

var server model.Server

func main() {
	controller.InitializeApp()
	go controller.GetFromDB()
	controller.Mycontroller()

}
