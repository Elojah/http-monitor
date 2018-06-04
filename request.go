package monitor

import (
	"time"
)

// Request represents a read request.
type Request struct {
	IP         string
	UserID     string
	TS         time.Time
	URL        string
	StatusCode int
	SizeCode   int
}

// RequestMapper is a data interface for request object.
type RequestMapper interface {
	CreateRequest(Request) error
	ListRequest(RequestSubset) ([]Request, error)
	CountRequest(RequestSubset) (int, error)
}

// RequestSubset targets part of stored requests per date.
type RequestSubset struct {
	FromTS     time.Time
	ToTS       time.Time
	StatusCode *int
	Section    *string
}
