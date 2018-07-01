package admin

import (
	"context"
	"net/http"
)

// App is the admin app
type App struct {
	ctx     context.Context
	handler http.Handler
}

// Serve starts the admin app
func (a App) Serve() error {
	server := &http.Server{
		Addr:    "0.0.0.0:9090",
		Handler: a.handler,
	}

	errCh := make(chan error)
	go func() {
		errCh <- server.ListenAndServe()
	}()
	select {
	case err := <-errCh:
		// if the server errors, it doesn't need to be shutdown - it's already dead
		//
		// pass the error up the stack so that the caller can handle it
		// and shut down the other servers
		return err
	case <-a.ctx.Done():
		// prob should move this up the stack. we should time out on shutdown
		// too
		shutdownCtx := context.Background()
		return server.Shutdown(shutdownCtx)
	}
	return nil
}
