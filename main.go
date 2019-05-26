package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	"html"
	"log"
	"net/http"
	"time"
)

func main() {

	mdlw := middleware.New(middleware.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})

	myHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//emulate workload
		time.Sleep(200 * time.Millisecond)
		fmt.Fprintf(w, "{ \"path\": %q }", html.EscapeString(r.URL.Path))
	})

	h := mdlw.Handler("", myHandler)

	// Serve metrics.
	log.Printf("serving metrics at: %s", ":9090")
	go http.ListenAndServe(":9090", promhttp.Handler())

	log.Printf("listening at: %s", ":8080")
	if err := http.ListenAndServe(":8080", h); err != nil {
		log.Panicf("error while serving: %s", err)
	}
}
