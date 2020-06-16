package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

var (
	submitTime = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name:      "any_number",
			Help:      "A number to test gauge",
	})

	requestLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:      "request_latency_seconds",
			Help:      "Time spent in this service.",
			Buckets:   []float64{0.01, 0.02, 0.05, 0.1, 0.2, 0.5, 1.0, 2.0, 5.0, 10.0, 20.0, 30.0, 60.0, 120.0, 300.0},
		}, []string{},
	)
)

type RequestLatency struct {
	histo *prometheus.HistogramVec
	start time.Time
}

func Register() {
	prometheus.MustRegister(submitTime)
	prometheus.MustRegister(requestLatency)
}

func NewAdmissionLatency() *RequestLatency {
	return &RequestLatency{
		histo: requestLatency,
		start: time.Now(),
	}
}

func (t *RequestLatency) Observe() {
	submitTime.Set(3)
	(*t.histo).WithLabelValues().Observe(time.Now().Sub(t.start).Seconds())
}