package rslvex

import "github.com/seeruk/resolver-example/httpsrv"

// Config contains all application configuration. The config sits alongside the resolver in this
// case. In applications with multiple binaries (i.e. main.go's) you could make a package for each,
// like how I've made this package the same name as the rslvex command.
//
// Configuration is pushed into sub-packages, and pulled into here. This also means you can look at
// this type and determine what configuration is relevant to the application you're running.
//
// You could also use this type of struct to load something like JSON or YAML. In other applications
// I've used Consul's K/V store and put JSON blobs under a key, and then unmarshalled it onto a
// config struct similar to this one.
type Config struct {
	HTTP httpsrv.Config
}
