package monitor

import (
	"net"
	"regexp"
	"time"
)

var (
	sectionRgx = regexp.MustCompile(`\/(.*?)(\/|s)`)
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

// Section returns the URL section.
func (r Request) Section() string {
	return sectionRgx.FindString(r.URL)
}
