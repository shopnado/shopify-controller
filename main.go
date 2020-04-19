//go:generate go run controller/codegen/cleanup/main.go
//go:generate rm -rf controller/generated
//go:generate go run controller/codegen/main.go

package main

import (
	"github.com/shopnado/shopify-controller/api"
	"github.com/shopnado/shopify-controller/controller"
	"github.com/shopnado/shopify-controller/shopnado"
	"github.com/shopnado/shopify-controller/ui"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("[shopnado] starting services")
	server := shopnado.NewServer()
	server.Register(api.NewService(), controller.NewService(), ui.NewService())
	server.Run()
}
