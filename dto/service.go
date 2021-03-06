package dto

import (
	monitor "github.com/elojah/http-monitor"
)

var _ monitor.LogAlertMapper = (*Service)(nil)
var _ monitor.LogSectionMapper = (*Service)(nil)

// Service is the data transfer object service mockable.
type Service struct{}

// NewService returns a new valid dto service.
func NewService() *Service {
	return &Service{}
}
