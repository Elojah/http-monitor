package dto

import (
	"errors"
	"net"
	"regexp"
	"strconv"
	"time"

	monitor "github.com/elojah/http-monitor"
)

const (
	timeFormat = "[02/Jan/2006:15:04:05 -0700]"
)

var (
	clfRgx = regexp.MustCompile(`(\S+)\s+(\S+)\s+(\S+)\s+(\[.*?\])\s+(".*?")\s+(\S+)\s+(\S+)`)
)

// CLF represents a common log format line.
type CLF struct {
	ip         string
	identity   string
	userID     string
	ts         string
	url        string
	statusCode string
	sizeCode   string
}

// NewCLF returns a new CLF from s. Returns an error if the string format is not clf.
func NewCLF(s string) (CLF, error) {
	parts := clfRgx.FindStringSubmatch(s)
	if len(parts) != 8 {
		return CLF{}, errors.New("invalid common log format")
	}
	return CLF{
		ip:         parts[1],
		identity:   parts[2],
		userID:     parts[3],
		ts:         parts[4],
		url:        parts[5],
		statusCode: parts[6],
		sizeCode:   parts[7],
	}, nil
}

// NewRequest returns a request from a CLF.
func (clf CLF) NewRequest() (monitor.Request, error) {
	var req monitor.Request
	var err error
	req.IP, err = net.ResolveIPAddr("ip", clf.ip)
	if err != nil {
		return req, err
	}
	req.UserID = clf.userID
	if req.TS, err = time.Parse(timeFormat, clf.ts); err != nil {
		return req, err
	}
	req.URL = clf.url
	if req.StatusCode, err = strconv.Atoi(clf.statusCode); err != nil {
		return req, err
	}
	if req.SizeCode, err = strconv.Atoi(clf.sizeCode); err != nil {
		return req, err
	}
	return req, nil
}
