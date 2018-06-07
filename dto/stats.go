package dto

import (
	"fmt"

	monitor "github.com/elojah/http-monitor"
)

// LogSection is the dto implementation of LogSectionMapper.
func (s *Service) LogSection(sections []monitor.Section, top uint) {
	var count int
	for _, section := range sections {
		count += section.Hit
	}
	fmt.Printf("Hits: %d, Different sections: %d\n", count, len(sections))
	for i := 0; i < int(top) && i < len(sections); i++ {
		section := sections[i]
		fmt.Printf("%s: %d - %d%%\n", section.Name, section.Hit, (section.Hit*100)/count)
	}
}
