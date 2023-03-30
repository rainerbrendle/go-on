package server

import (
	"go-on.com/server/controller"
	"go-on.com/server/model"
)

func Run() {
   model.Init()
   controller.Start()
}
