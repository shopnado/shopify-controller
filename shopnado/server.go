package shopnado

import (
	"fmt"
	"sync"

	"github.com/shopnado/shopify-controller/shopnado/signals"
	"github.com/sirupsen/logrus"
)

type Service interface {
	Run(s *Server) error
	Stop() error
	Name() string
}

type Server struct {
	Services    []Service
	WaitGroup   sync.WaitGroup
	StopChannel <-chan struct{}
}

func NewServer() *Server {
	server := &Server{
		StopChannel: signals.SignalHandler(),
	}

	go func() {
		<-server.StopChannel
		server.Shutdown()
	}()

	return server
}

func (server *Server) Register(services ...Service) {
	for _, service := range services {
		server.Services = append(server.Services, service)
	}
}

func (server *Server) Run() error {
	for _, service := range server.Services {
		server.WaitGroup.Add(1)
		s := service
		errs := make(chan error, 1)

		go func() {
			if err := s.Run(server); err != nil {
				errs <- fmt.Errorf("error while running: %s", err)
			}
		}()

		go func() {
			select {
			case err := <-errs:
				logrus.Error(err)
				logrus.Infof("[%s] shutting down", s.Name())
				s.Stop()
			}
		}()
	}

	server.WaitGroup.Wait()
	return nil
}

func (server *Server) Shutdown() error {
	for _, service := range server.Services {
		err := service.Stop()
		if err != nil {
			return err
		}
		server.WaitGroup.Done()
	}
	return nil
}

func (server *Server) Done() {
	server.WaitGroup.Done()
}
