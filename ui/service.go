package ui

import (
	"github.com/shopnado/shopify-controller/shopnado"
	"github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
)

type Service struct {
	ctx *cli.Context
}

func NewService(ctx *cli.Context) shopnado.Service {
	return &Service{
		ctx: ctx,
	}
}

func (ui *Service) Name() string {
	return "ui"
}

func (ui *Service) Run(server *shopnado.Server) error {
	logrus.Infof("[%s] starting service", ui.Name())
	<-server.StopChannel // wait for stop signal
	return nil
}

func (ui *Service) Stop() error {
	logrus.Infof("[%s] stopping service", ui.Name())
	return nil
}
