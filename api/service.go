package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shopnado/shopify-controller/api/routes"
	"github.com/shopnado/shopify-controller/shopnado"
	"github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
)

type Service struct {
	ctx        *cli.Context
	httpserver *http.Server
}

func NewService(ctx *cli.Context) shopnado.Service {
	return &Service{
		ctx:        ctx,
		httpserver: &http.Server{Handler: setupRouter()},
	}
}

func (api *Service) Name() string {
	return "api"
}

func (api *Service) Run(server *shopnado.Server) error {
	logrus.Infof("[%s] starting service", api.Name())

	addr := "0.0.0.0:8080"
	api.serve(addr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	logrus.Infof("[%s] service started, http available at %s", api.Name(), addr)
	<-server.StopChannel // wait for stop signal

	return nil
}

func (api *Service) Stop() error {
	logrus.Infof("[%s] stopping service", api.Name())
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := api.httpserver.Shutdown(ctx); err != nil {
		logrus.Errorf("[%s] server forced to shutdown: %s", api.Name(), err)
	}

	logrus.Infof("[%s] http server closed", api.Name())
	return nil
}

func (api *Service) serve(addr string) {
	api.httpserver.Addr = addr

	go func() {
		if err := api.httpserver.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Errorf("[%s] listen and server error %s", api.Name(), err)
		}
	}()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	routes.Register(r)
	return r
}
