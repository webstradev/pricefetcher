package main

import (
	"flag"
)

func main() {
	var (
		jsonAddr = flag.String("jsonaddr", ":3000", "The listen address for the JSON Server")
		grpcAddr = flag.String("grpcaddr", ":4000", "The listen address for the gRPC Server")
	)

	flag.Parse()

	svc := &priceFetcher{}
	svcWithLoggingAndMetrics := NewLoggingService(NewMetricService(svc))

	jsonServer := NewJSONAPIServer(*jsonAddr, svcWithLoggingAndMetrics)
	jsonServer.Run()

	makeGRPCServerAndRun(*grpcAddr, svcWithLoggingAndMetrics)
}
