package http

import (
	"github.com/polaris1119/go-web-layout/demo/logic"
	"github.com/polaris1119/go-web-layout/demo/repository"
	"github.com/polaris1119/go-web-layout/http/controller"
)

// 存放路由
func Route() {
	file := "demo.json"
	controller.NewDemoHttpHandler(logic.NewDemoLogic(repository.NewFileRepository(file)))
}
