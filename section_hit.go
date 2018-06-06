package monitor

// SectionHit represents a section request and its number of hits.
type SectionHit struct {
	Section string
	Hit     int
}

// SectionHitMapper is a data interface for request hit object.
type SectionHitMapper interface {
	AddSectionHit(string) error
	ListSectionHit(SectionHitSubset) ([]SectionHit, error)
	ResetSectionHit() error
}

// SectionHitSubset targets part of stored requests per date.
type SectionHitSubset struct {
	TopHits *uint
}
