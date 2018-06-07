package monitor

// Section represents a section and its number of hits.
type Section struct {
	Name string
	Hit  int
}

// SectionMapper is a data interface for Section object.
type SectionMapper interface {
	IncrSection(string) error
	ListSection(SectionSubset) ([]Section, error)
	ResetSection() error
}

// LogSectionMapper is a log interface for sections.
type LogSectionMapper interface {
	LogSection([]Section, uint)
}

// SectionSubset targets part of stored sections per hits.
type SectionSubset struct {
	TopHits *uint
}
