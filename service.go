package main

import (
	"context"
	"fmt"
	"time"
)

// PriceFetcher is an interface that can fetch a price given a ticker
type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

// implements PriceFetcher interface
type priceFetcher struct{}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetcher(ctx, ticker)
}

var priceMocks = map[string]float64{
	"BTC": 20_000.0,
	"ETH": 200.0,
	"WBS": 200_000.0,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	// Fake round trip to some external API
	time.Sleep(100 * time.Millisecond)
	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("the given ticker (%s) is not supported", ticker)
	}

	return price, nil

}
