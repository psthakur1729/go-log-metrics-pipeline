package tailer

import (
	"time"
)

var sampleLogs = []string{
	`{"timestamp":"2025-06-13T14:00:00Z", "level":"info", "message":"all good"}`,
	`{"timestamp":"2025-06-13T14:00:01Z", "level":"error", "message":"db down"}`,
	`{"timestamp":"2025-06-13T14:00:02Z", "level":"debug", "message":"retrying..."}`,
}

func SimulateTail(out chan<- string) {
	go func() {
		for {
			for _, log := range sampleLogs {
				out <- log
				time.Sleep(2 * time.Second)
			}
		}
	}()
}
