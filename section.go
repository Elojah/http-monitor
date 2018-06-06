package monitor

// Section represents a section and its number of hits.
type Section struct {
	Name string
	Hit  int
}

// SectionMapper is a data interface for request hit object.
type SectionMapper interface {
	IncrSection(string) error
	ListSection(SectionSubset) ([]Section, error)
	ResetSection() error
}

// SectionSubset targets part of stored requests per date.
type SectionSubset struct {
	TopHits *uint
}
