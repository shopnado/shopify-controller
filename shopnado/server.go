package shopnado

import (
	"sync"

	"github.com/shopnado/shopify-controller/shopnado/signals"
)

type Service interface {
	Run(s *Server)
	Stop() error
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

func (s *Server) Register(services ...Service) {
	for _, service := range services {
		s.Services = append(s.Services, service)
	}
}

func (s *Server) Run() {
	for _, service := range s.Services {
		s.WaitGroup.Add(1)
		go service.Run(s)

	}
	s.WaitGroup.Wait()
}

func (s *Server) Shutdown() error {
	for _, service := range s.Services {
		err := service.Stop()
		if err != nil {
			return err
		}
		s.WaitGroup.Done()
	}
	return nil
}

func (s *Server) Done() {
	s.WaitGroup.Done()
}
