package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	StepCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "cuwander_step_total",
			Help: "count total user step",
		},
		[]string{"event_type", "faculty", "enrolled_year"},
	)
)
