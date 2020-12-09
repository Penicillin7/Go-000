package main

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	server http.Server
}

func (s *Server) start() error {
	return s.server.ListenAndServe()
}

func (s *Server) shutDown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func NewServer(addr string) *Server {
	return &Server{
		server: http.Server{
			Addr: addr,
		},
	}
}


func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	g, _ := errgroup.WithContext(ctx)
	errChan := make(chan error)

	// start server1
	g.Go(func() error {
		s1 := NewServer(":8888")
		go func() {
			if err := s1.start(); err != nil {
				errChan <- err
			}
		}()
		log.Printf("start server_1")
		return nil
	})

	// start server1
	g.Go(func() error {
		s2 := NewServer(":9999")
		go func() {
			if err := s2.start(); err != nil {
				errChan <- err
			}
		}()
		log.Printf("start server_2")
		return nil
	})

	// 监听
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGPIPE)

	g.Go(func() error {
		for {
			select {
			case sig := <-quit:
				return errors.New("exit signal " + sig.String())
			case <-ctx.Done():
				signal.Stop(quit)
				log.Printf("shutDown signal")
				return nil
			}
		}
	})
	if err := g.Wait(); err != nil {
		log.Printf("Error:%v\n", err)
	}
}
