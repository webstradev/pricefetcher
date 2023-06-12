package main

import (
	"flag"
)

func main() {
	var (
		jsonAddr = flag.String("jsonaddr", ":3000", "The listen address for the JSON Server")
		grpcAddr = flag.String("grpcaddr", ":4000", "The listen address for the gRPC Server")
		svc      = NewLoggingService(NewMetricService(&priceFetcher{}))
	)

	flag.Parse()

	go makeGRPCServerAndRun(*grpcAddr, svc)

	jsonServer := NewJSONAPIServer(*jsonAddr, svc)
	jsonServer.Run()
}
