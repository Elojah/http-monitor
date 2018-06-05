package monitor

import (
	"net"
	"time"
)

// Request represents a read request.
type Request struct {
	IP         *net.IPAddr
	Identity   string
	UserID     string
	TS         time.Time
	URL        string
	StatusCode int
	SizeCode   int
}

// RequestHit represents a request and its number of hits
type RequestHit struct {
	Request
	Hit int
}

// RequestHitMapper is a data interface for request hit object.
type RequestHitMapper interface {
	AddRequestHit(Request) error
	ListRequestHit(RequestSubset) ([]RequestHit, error)
	ResetRequestHit() error
}

// RequestSubset targets part of stored requests per date.
type RequestSubset struct {
	TopHits *uint
}
