package monitor

// Mappers wraps all monitor services.
type Mappers struct {
	LogSectionMapper
	SectionMapper
	TickMapper
	AlertMapper
}
