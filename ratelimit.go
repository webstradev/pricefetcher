package main

import (
	"context"
	"fmt"
	"math/rand"
)

type rateLimitService struct {
	next PriceFetcher

	maxRequestsPerSecond int
}

func NewRateLimitService(next PriceFetcher, maxRequests int) PriceFetcher {
	return &rateLimitService{
		next:                 next,
		maxRequestsPerSecond: maxRequests,
	}
}

func (s *rateLimitService) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	blocked := mockRateLimit()

	if !blocked {
		return s.next.FetchPrice(ctx, ticker)
	}

	return 0.0, fmt.Errorf("ratelimit exceeded")
}

// Just mocking a ratelimit by blocking half of requests
func mockRateLimit() bool {
	return rand.Intn(2) == 1
}
