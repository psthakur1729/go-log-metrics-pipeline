package main

import (
	"go-log-metrics-pipeline/internal/filter"
	"go-log-metrics-pipeline/internal/metrics"
	"go-log-metrics-pipeline/internal/parser"
	"go-log-metrics-pipeline/internal/tailer"
)

func main() {
	logChan := make(chan string)
	parsedChan := make(chan parser.LogEntry)
	filteredChan := make(chan parser.LogEntry)

	go tailer.SimulateTail(logChan)
	go parser.Parse(logChan, parsedChan)
	go filter.Filter(parsedChan, filteredChan, []string{"error", "warn"})
	go metrics.Collect(filteredChan)

	metrics.StartHTTPServer() // Exposes /metrics
}
