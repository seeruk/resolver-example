package httpsrv

import "time"

// Config contains all configuration relevant to the HTTP server.
type Config struct {
	Address        string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	HandlerTimeout time.Duration
	GracePeriod    time.Duration
}
