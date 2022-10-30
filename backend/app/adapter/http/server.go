package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gold-kou/cfn-practice/backend/app/adapter/http/controller"

	"github.com/gorilla/mux"
)

const gracefulShutdownTimeoutDefault = 5

var gracefulShutdownTimeout time.Duration

func init() {
	t, e := time.ParseDuration(os.Getenv("GRACEFUL_SHUTDOWN_TIMEOUT_SECOND") + "s")
	if e != nil {
		gracefulShutdownTimeout = gracefulShutdownTimeoutDefault
	} else {
		gracefulShutdownTimeout = t
	}
}

func Serve() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/health/liveness", controller.HealthController)
	r.HandleFunc("/health/readiness", controller.HealthController)
	r.HandleFunc("/messages", controller.MessageController)
	r.HandleFunc("/messages/{id}", controller.MessageController)

	// graceful shutdown
	server := &http.Server{Addr: fmt.Sprintf(":%v", 80), Handler: r}
	idleConnsClosed := make(chan struct{})
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGTERM)
		<-sigCh

		ctx, cancel := context.WithTimeout(context.Background(), gracefulShutdownTimeout)
		defer cancel()
		if e := server.Shutdown(ctx); e != nil {
			// Error from closing listeners, or context gracefulShutdownTimeout:
			log.Panic("Failed to gracefully shutdown ", e)
		}
		close(idleConnsClosed)
	}()

	// launch
	log.Println("Server started!")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Panic(err)
	}
	<-idleConnsClosed
}
