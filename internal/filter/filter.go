package filter

import (
	"go-log-metrics-pipeline/internal/parser"
)

func Filter(in <-chan parser.LogEntry, out chan<- parser.LogEntry, levels []string) {
	allowed := make(map[string]bool)
	for _, l := range levels {
		allowed[l] = true
	}

	go func() {
		for entry := range in {
			if allowed[entry.Level] {
				out <- entry
			}
		}
	}()
}
