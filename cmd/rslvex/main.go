package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/seeruk/resolver-example/rslvex"
)

// rslvex is the resolver example server command.
func main() {
	var conf rslvex.Config

	// These lines get a bit long... normally I'd probably not be using the built-in flags, but I
	// didn't want to be pulling in loads of libraries for this example.
	flag.StringVar(&conf.HTTP.Address, "http-addr", "0.0.0.0:8080", "HTTP: Address to listen on, def: 0.0.0.0:8080")
	flag.DurationVar(&conf.HTTP.ReadTimeout, "http-read-timeout", 5*time.Second, "HTTP: Server read timeout, def: 5s")
	flag.DurationVar(&conf.HTTP.WriteTimeout, "http-write-timeout", 55*time.Second, "HTTP: Server write timeout, def: 55s")
	flag.DurationVar(&conf.HTTP.HandlerTimeout, "http-handler-timeout", 30*time.Second, "HTTP: Server handler timeout, def: 30s")
	flag.DurationVar(&conf.HTTP.GracePeriod, "http-grace-period", 5*time.Second, "HTTP: Server shutdown grace period, def: 5s")
	flag.Parse()

	resolver := rslvex.NewResolver(conf)

	// Catch interrupt and kill signals, so we can trigger the HTTP server's shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	// Resolve HTTP server. Note that this is the only thing using the resolver in main, and that
	// the resolver is never passed to any other types. In a more realistic application you would
	// likely need to ask for a few things in main, but if you're asking for loads then you might
	// want to consider if you could group some things up in another type.
	httpServer, closer := resolver.ResolveHTTPServer()

	go func() {
		err := httpServer.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("HTTP server listening at %q\n", conf.HTTP.Address)
	log.Printf("Caught signal %q. Shutting down...\n", <-signals)

	err := closer()
	if err != nil {
		log.Fatal(err)
	}
}
