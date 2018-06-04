package mem

// Service is a mem service to store data directly in memory.
type Service struct{}

// NewService returns a new valid service.
func NewService() *Service {
	return &Service{}
}
