package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/webstradev/pricefetcher/types"
)

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

type JSONAPIServer struct {
	listenAddr string

	svc PriceFetcher
}

func NewJSONAPIServer(listenAddr string, svc PriceFetcher) *JSONAPIServer {
	return &JSONAPIServer{
		svc:        svc,
		listenAddr: listenAddr,
	}
}

func (s *JSONAPIServer) Run() {
	http.HandleFunc("/", makeHTTPHandlerFunc(s.handleFetchPrice))

	http.ListenAndServe(s.listenAddr, nil)
}

func makeHTTPHandlerFunc(apiFn APIFunc) http.HandlerFunc {
	uuid := uuid.New()
	ctx := context.Background()
	ctx = context.WithValue(ctx, keyRequestID, uuid.String())

	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFn(ctx, w, r); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		}
	}
}

func (s *JSONAPIServer) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ticker := r.URL.Query().Get("ticker")

	price, err := s.svc.FetchPrice(ctx, ticker)
	if err != nil {
		return err
	}

	priceResp := types.PriceResponse{
		Price:  price,
		Ticker: ticker,
	}

	return writeJSON(w, http.StatusOK, &priceResp)
}

func writeJSON(w http.ResponseWriter, status int, body any) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(body)
}
