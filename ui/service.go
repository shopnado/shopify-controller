package ui

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
	logrus.Info("[ui] starting service")

	<-server.StopChannel // wait for stop signal
}

func (api *Service) Stop() error {
	logrus.Info("[ui] stopping service")
	return nil
}
