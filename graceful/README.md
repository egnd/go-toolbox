# graceful shutdown

### Example:
```golang
package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/egnd/go-toolbox/graceful"
)

func main() {
	// init listener and subscribe to os signals
	ctx, cancel := graceful.Init(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	serv := http.Server{
		BaseContext: func(l net.Listener) context.Context { return ctx },
	}

	// start http server and register its stop callback
	graceful.Register(serv.ListenAndServe, func() error {
		return serv.Shutdown(ctx)
	})

	// start other service without stop callback
	graceful.Register(func() error {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		select {
		case <-ticker.C:
		case <-ctx.Done():
			return nil
		}

		return nil
	})

	// wait for os signal or error from one of the working services
	if err := graceful.Wait(); err != nil {
		log.Fatal(err)
	}

	log.Println("service is correctly stopped")
}
```