package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shopnado/shopify-controller/api/routes"
	"github.com/shopnado/shopify-controller/shopnado"
	"github.com/sirupsen/logrus"
)

type Service struct {
	httpserver *http.Server
}

func NewService() shopnado.Service {
	return &Service{
		httpserver: &http.Server{Handler: setupRouter()},
	}
}

func (api *Service) Run(server *shopnado.Server) {
	logrus.Info("[api] starting service")
	addr := "0.0.0.0:8080"
	api.serve(addr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	logrus.Infof("[api] service started, http available at %s", addr)
	<-server.StopChannel // wait for stop signal
}

func (api *Service) Stop() error {
	logrus.Info("[api] stopping service")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := api.httpserver.Shutdown(ctx); err != nil {
		logrus.Errorf("[api] server forced to shutdown: %s", err)
	}

	logrus.Info("[api] http server closed")
	return nil
}

func (api *Service) serve(addr string) {
	api.httpserver.Addr = addr

	go func() {
		if err := api.httpserver.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Errorf("[api] listen and server error %s", err)
		}
	}()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	routes.Register(r)
	return r
}
