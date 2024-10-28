package http

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/superdev/ecommerce/gateway/internal/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewHTTPServer(
	router *gin.Engine,
	conf *config.Config,
) (*http.Server, error) {

	port := conf.AppPort
	if 0 == port {
		port = 8080
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	return srv, nil
}

func NewHTTPServerRunner(
	logger zap.Logger,
	lf fx.Lifecycle,
	srv *http.Server,
) HttpServerRunner {
	runner := &httpServerRunner{
		srv: srv,
	}
	// Start the server when the lifecycle starts
	lf.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			logger.Debug(fmt.Sprintf("Starting HTTP server at %s", srv.Addr))
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Debug(fmt.Sprintf("Shutting down HTTP server at %s", srv.Addr))
			return srv.Shutdown(ctx)
		},
	})

	return runner
}

type HttpServerRunner interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type httpServerRunner struct {
	srv *http.Server
}

func (s *httpServerRunner) Start(ctx context.Context) (err error) {
	go func() {
		err = s.srv.ListenAndServe()
	}()
	time.Sleep(time.Second)
	return err
}

func (s *httpServerRunner) Stop(ctx context.Context) (err error) {
	err = s.srv.Shutdown(ctx)
	return err
}
