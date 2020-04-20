package controller

import (
	"context"
	"fmt"
	"os"

	"github.com/rancher/wrangler/pkg/resolvehome"
	"github.com/rancher/wrangler/pkg/start"
	shopnadocontroller "github.com/shopnado/shopify-controller/controller/generated/controllers/shopnado.xyz"
	"github.com/shopnado/shopify-controller/shopnado"
	"github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
	"k8s.io/client-go/tools/clientcmd"
)

type Service struct {
	ctx *cli.Context
}

func NewService(ctx *cli.Context) shopnado.Service {
	return &Service{
		ctx: ctx,
	}
}

func (controller *Service) Name() string {
	return "controller"
}

func (controller *Service) Run(server *shopnado.Server) error {
	logrus.Infof("[%s] starting service", controller.Name())

	kubeconfig, err := resolvehome.Resolve(controller.ctx.String("kubeconfig"))
	if err != nil {
		return fmt.Errorf("[%s] resolving home dir for kubeconfig file failed", controller.Name())
	}

	stat, err := os.Stat(kubeconfig)
	if _, err := stat, err; os.IsNotExist(err) {
		kubeconfig = ""
	}

	threadiness := controller.ctx.Int("threads")
	masterurl := controller.ctx.String("masterurl")
	ns := controller.ctx.String("namespace")

	cfg, err := clientcmd.BuildConfigFromFlags(masterurl, kubeconfig)
	if err != nil {
		return fmt.Errorf("[%s] error building kubeconfig: %s", controller.Name(), err.Error())
	}

	shopnadoFactory, err := shopnadocontroller.NewFactoryFromConfigWithNamespace(cfg, ns)
	if err != nil {
		return fmt.Errorf("[%s] error building terraform controllers: %s", controller.Name(), err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())
	Register(ctx, shopnadoFactory.Shopnado().V1().Event())

	if err := start.All(ctx, threadiness, shopnadoFactory); err != nil {
		return fmt.Errorf("[%s] error starting controller: %s", controller.Name(), err.Error())
	}

	<-server.StopChannel // wait for stop signal
	cancel()
	return nil
}

func (controller *Service) Stop() error {
	logrus.Infof("[%s] stopping service", controller.Name())
	return nil
}
