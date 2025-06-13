package parser

import (
	"encoding/json"
	"log"
)

type LogEntry struct {
	Timestamp string `json:"timestamp"`
	Level     string `json:"level"`
	Message   string `json:"message"`
}

func Parse(in <-chan string, out chan<- LogEntry) {
	go func() {
		for line := range in {
			var entry LogEntry
			if err := json.Unmarshal([]byte(line), &entry); err == nil {
				out <- entry
			} else {
				log.Println("parse error:", err)
			}
		}
	}()
}
