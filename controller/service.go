package controller

import (
	"github.com/shopnado/shopify-controller/shopnado"
	"github.com/sirupsen/logrus"
)

type Service struct {
}

func NewService() shopnado.Service {
	return &Service{}
}

func (api *Service) Run(server *shopnado.Server) {
	logrus.Info("[controller] starting service")

	<-server.StopChannel // wait for stop signal
}

func (api *Service) Stop() error {
	logrus.Info("[controller] stopping service")
	return nil
}
