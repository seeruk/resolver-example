package rslvex

import (
	"context"
	"net/http"

	"github.com/seeruk/resolver-example/greeter"
	"github.com/seeruk/resolver-example/handlers"
)

// Resolver is a type that can resolve the dependencies of the rslvex application. It is the primary
// runtime resolver (i.e. not a testing resolver). Where it makes sense, some dependencies are
// singletons.
//
// This is a very small example. I typically build micro-services, and my resolver types are usually
// a few hundred lines long. If it gets too long, you can always split your resolver into multiple
// resolvers. It's just plain Go, do what works!
type Resolver struct {
	config Config
}

// NewResolver returns a new Resolver instance.
func NewResolver(config Config) *Resolver {
	return &Resolver{
		config: config,
	}
}

// ResolveHTTPServer resolves a new HTTP server instance, along with a function to shutdown the
// server gracefully, within the configured timeout.
//
// In a realistic HTTP application, you might choose to use a router like Chi, or Gorilla Mux. When
// I use a router I usually create a RouterFactory type that can build the resulting http.Handler
// type from the router. I pass each of my application's handlers to that factory, via the Resolver.
// You can use a few different patterns there, like you could make it a builder instead if you
// wanted if you have tons of handlers and using the constructor becomes a bit... unwieldy. Using
// the factory approach also means each handler has a resolver function (as they should do!)
func (r *Resolver) ResolveHTTPServer() (*http.Server, func() error) {
	// Building the HTTP server like this allows us to specify timeouts, important!
	server := &http.Server{
		Addr:         r.config.HTTP.Address,
		ReadTimeout:  r.config.HTTP.ReadTimeout,
		WriteTimeout: r.config.HTTP.WriteTimeout,

		// Here is where we start to use our dependency graph. We resolve the GreetingHandler which
		// itself has it's own dependencies. Not that that's visible here, keeping the dependency
		// resolution nice and flat, and easy to understand. Right here you probably don't care
		// about the GreetingHandler's dependencies, but if you want to see them, you can go look at
		// it's resolver function.
		Handler: r.ResolveGreetingHandler(),
	}

	// Closer can be used to gracefully shutdown this HTTP server.
	closer := func() error {
		ctx, cfn := context.WithTimeout(context.Background(), r.config.HTTP.GracePeriod)
		defer cfn()

		return server.Shutdown(ctx)
	}

	return server, closer
}

// ResolveGreetingHandler returns a new handlers.GreetingHandler instance.
func (r *Resolver) ResolveGreetingHandler() *handlers.GreetingHandler {
	return handlers.NewGreetingHandler(
		r.ResolveEnglishGreeter(),
	)
}

// ResolveEnglishGreeter returns a new EnglishGreeter instance.
func (r *Resolver) ResolveEnglishGreeter() *greeter.EnglishGreeter {
	return greeter.NewEnglishGreeter()
}
