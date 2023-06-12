package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	svc := &priceFetcher{}
	svcWithLoggingAndMetrics := NewLoggingService(NewMetricService(svc))

	price, err := svcWithLoggingAndMetrics.FetchPrice(context.Background(), "ETH")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(price)
}
