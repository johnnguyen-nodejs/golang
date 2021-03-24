package main
import (
	"time"
	"net/http"
	"log"
	"context"
	"github.com/johnnguyen-nodejs/golang/handlers"
	"github.com/nicholasjackson/env"
	"os"
	"os/signal"
)
var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {
	env.Parse()

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	// Create the handlers
	ph := handlers.NewProducts(l)
	// Create a new serve mux and register the handlers
	sm := http.NewServeMux()
	sm.Handle("/", ph)
	// Create a new server
	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	// Start the server
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <- sigChan
	l.Println("recieved terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
