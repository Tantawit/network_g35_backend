package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	OnlineUserCounter = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "cuwander_online_user_total",
			Help: "count total online user",
		},
		[]string{},
	)
)
