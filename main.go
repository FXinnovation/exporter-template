package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
)

func init() {
	prometheus.MustRegister(version.NewCollector("my_exporter"))
}

func main() {
	collector := NewCollector()
	prometheus.MustRegister(collector)

	http.Handle("/metrics", promhttp.Handler())
	log.Println("Beginning to serve on port :1984")
	log.Fatal(http.ListenAndServe(":1984", nil))
}
