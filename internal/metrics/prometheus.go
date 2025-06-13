package metrics

import (
	"go-log-metrics-pipeline/internal/parser"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	logCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "log_entries_total",
			Help: "Total log entries by level",
		},
		[]string{"level"},
	)
)

func init() {
	prometheus.MustRegister(logCounter)
}

func Collect(in <-chan parser.LogEntry) {
	go func() {
		for entry := range in {
			logCounter.WithLabelValues(entry.Level).Inc()
		}
	}()
}

func StartHTTPServer() {
	http.Handle("/metrics", promhttp.Handler())
	log.Println("Prometheus metrics exposed on :2112/metrics")
	log.Fatal(http.ListenAndServe(":2112", nil))
}
