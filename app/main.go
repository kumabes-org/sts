package main

import (
	"log"
	"net/http"

	"github.com/pbnjay/memory"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func memoriaLivre() float64 {
	memoria_livre := memory.FreeMemory()
	return float64(memoria_livre)
}

func memoriaTotal() float64 {
	memoria_total := memory.TotalMemory()
	return float64(memoria_total)
}

var (
	memoriaLivreBytesGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "memoria_livre_bytes",
		Help: "Quantidade de memória livre em bytes",
	})

	memoriaLivreMegasGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "memoria_livre_megas",
		Help: "Quantidade de memória livre em megas",
	})

	memoriaTotalBytesGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "memoria_total_bytes",
		Help: "Quantidade de memória total em bytes",
	})

	memoriaTotalGigasGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "memoria_total_gigas",
		Help: "Quantidade de memória total em gigas",
	})
)

func init() {
	prometheus.MustRegister(memoriaLivreBytesGauge)
	prometheus.MustRegister(memoriaLivreMegasGauge)
	prometheus.MustRegister(memoriaTotalBytesGauge)
	prometheus.MustRegister(memoriaTotalGigasGauge)
}

func main() {
	memoriaLivreBytesGauge.Set(memoriaLivre())
	memoriaLivreMegasGauge.Set(memoriaLivre() / 1024 / 1024)
	memoriaTotalBytesGauge.Set(memoriaTotal())
	memoriaTotalGigasGauge.Set(memoriaTotal() / 1024 / 1024 / 1024)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":7788", nil))
}
