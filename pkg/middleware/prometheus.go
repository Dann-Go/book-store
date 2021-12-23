package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"strconv"
	"strings"
	"time"
)

var (
	dflBuckets = []float64{1.0, 2.5, 5.0, 10.0, 30.0, 60.0}
)

// Opts specifies options how to create new PrometheusMiddleware.
type Opts struct {
	// Buckets specifies an custom buckets to be used in request histogram.
	Buckets []float64
}

const (
	requestCountName    = "http_requests_total"
	requestDurationName = "http_request_duration_seconds"
	requestReservedName = "reserved_books_count"
)

// PrometheusMiddleware represents webint metrics with its service name
type PrometheusMiddleware struct {
	serviceName     string
	requestCount    *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
	queries         []string
}

//var (
//	Books_reserved = prometheus.NewCounterVec(
//		prometheus.CounterOpts{
//			Name: requestReservedName,
//			Help: "Number of reserved books",
//		}, []string{"id", "status_code"})
//)

// Metrics registers new metrics and wraps it in gin.HandlerFunc
func (p *PrometheusMiddleware) Metrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		begin := time.Now()
		c.Next()
		go p.requestCount.WithLabelValues(
			p.serviceName,
			sanitizeMethod(c.Request.Method),
			c.Request.URL.String(),
			sanitizeCode(c.Writer.Status()),
		).Inc()
		go p.requestDuration.WithLabelValues(
			sanitizeCode(c.Writer.Status()),
			sanitizeMethod(c.Request.Method),
			c.Request.URL.Path,
		).Observe(float64(time.Since(begin)) / float64(time.Second))
	}
}

// NewPrometheusMiddleware return new instance of PrometheusMiddleware with given service name and bucket for histogram
func NewPrometheusMiddleware(serviceName string, opts Opts, queries ...string) *PrometheusMiddleware {
	var p PrometheusMiddleware
	p.serviceName = serviceName
	p.requestCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: serviceName,
			Name:      requestCountName,
			Help:      "HTTP requests statistics",
		}, []string{"service", "method", "url", "status_code"})
	buckets := opts.Buckets
	if len(buckets) == 0 {
		buckets = dflBuckets
	}
	p.requestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: serviceName,
		Name:      requestDurationName,
		Help:      "Duration of HTTP requests.",
		Buckets:   buckets,
	}, []string{"status_code", "method", "url"})
	for _, q := range queries {
		p.queries = append(p.queries, q)
	}
	return &p
}

func sanitizeMethod(m string) string {
	return strings.ToLower(m)
}

func sanitizeCode(s int) string {
	return strconv.Itoa(s)
}
