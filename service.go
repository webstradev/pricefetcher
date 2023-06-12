package main

import (
	"context"
	"fmt"
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
	"WBS": 200_000.0,
	"BTC": 20_000.0,
	"ETH": 200.0,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("the given ticker (%s) is not supported", ticker)
	}

	return price, nil
}
