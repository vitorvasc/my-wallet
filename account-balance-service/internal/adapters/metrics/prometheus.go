package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "The HTTP request latencies in seconds.",
		},
		[]string{"method", "path"},
	)
)

func MapRoutes(r *gin.Engine) {
	prometheus.MustRegister(RequestDuration)

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
}

func RecordMetrics(handler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		timer := prometheus.NewTimer(RequestDuration.WithLabelValues(c.Request.Method, c.Request.URL.Path))
		defer timer.ObserveDuration()

		handler(c)
	}
}
