package monitor

// Mappers wraps all monitor services.
type Mappers struct {
	LogAlertMapper
	LogSectionMapper
	SectionMapper
	TickMapper
}
