package main

import (
	"flag"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "The listen address for the JSON Server")
	flag.Parse()

	svc := &priceFetcher{}
	svcWithLoggingAndMetrics := NewLoggingService(NewMetricService(svc))

	server := NewJSONAPIServer(*listenAddr, svcWithLoggingAndMetrics)

	server.Run()
}
