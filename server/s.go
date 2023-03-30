package server

import (
	"go-on.com/server/controller"
	"go-on.com/server/model"
	"go-on.com/server/data"
)

func Run() {
   data.DbConnect()
   model.Init()
   controller.Start()
}
