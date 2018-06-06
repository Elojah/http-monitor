package monitor

import (
	"net"
	"regexp"
	"time"
)

var (
	sectionRgx = regexp.MustCompile(`(/.*?)(?:/|\s)`)
)

// Request represents a read request.
type Request struct {
	ID         ID
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
	matches := sectionRgx.FindStringSubmatch(r.URL)
	if len(matches) < 2 {
		return ""
	}
	return matches[1]
}
