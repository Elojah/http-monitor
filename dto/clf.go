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
	timeFormat = "2006/Jan/02:15:04:05 -0700"
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
	re := regexp.MustCompile(`(\S+)\s+(\S+)\s+(\S+)\s+(\[.*?\])\s+(".*?")\s+(\S+)\s+(\S+)`)
	parts := re.SubexpNames()
	if len(parts) != 7 {
		return CLF{}, errors.New("invalid common log format")
	}
	return CLF{
		ip:         parts[0],
		identity:   parts[1],
		userID:     parts[2],
		ts:         parts[3],
		url:        parts[4],
		statusCode: parts[5],
		sizeCode:   parts[6],
	}, nil
}

// NewRequest returns a request from a CLF.
func (clf CLF) NewRequest() (monitor.Request, error) {
	var req monitor.Request
	var err error
	req.IP, err = net.ResolveIPAddr("tcp", clf.ip)
	if err != nil {
		return req, err
	}
	req.Identity = clf.identity
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
