package graceful

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type StartFunc func() error

type ShutdownFunc func(ctx context.Context) error

var DefaultGracefulShutdownTimeout = 5 * time.Second

// Graceful sets up graceful handling, typically for an HTTP server
func Graceful(start StartFunc, shutdown ShutdownFunc) error {
	var (
		stopChan = make(chan os.Signal)
		errChan  = make(chan error)
	)

	go graceful(stopChan, errChan, shutdown)

	// Start the server
	if err := start(); err != http.ErrServerClosed {
		return err
	}

	return <-errChan
}

// graceful setup the graceful shutdown handler
func graceful(stopChan chan os.Signal, errChan chan error, shutdown ShutdownFunc) {
	signal.Notify(
		stopChan,
		// The SIGINT signal is sent when the user at the controlling terminal presses the interrupt character,
		// which by default is ^C (Control-C)
		syscall.SIGINT,
		// SIGTERM signal causes the program to exit
		syscall.SIGTERM,
	)

	<-stopChan

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), DefaultGracefulShutdownTimeout)
	defer cancel()

	if err := shutdown(ctxWithTimeout); err != nil {
		errChan <- err
		return
	}

	errChan <- nil
}
