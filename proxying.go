package main

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"

	httptransport "github.com/go-kit/kit/transport/http"
)

type proxymw struct {
	next      StringService     // Serve most requests via this service...
	uppercase endpoint.Endpoint // ...except Uppercase, which gets served by this endpoint
}

func (mw proxymw) Uppercase(ctx context.Context, s string) (string, error) {
	response, err := mw.uppercase(ctx, uppercaseRequest{S: s})
	if err != nil {
		return "", err
	}
	resp := response.(uppercaseResponse)
	if resp.Err != "" {
		return resp.V, errors.New(resp.Err)
	}
	return resp.V, nil
}

func proxyingMiddleware(proxyURL string) ServiceMiddleware {
	return func(next StringService) StringService {
		return proxymw{next, makeUppercaseEndpoint(proxyURL)}
	}
}

func makeUppercaseEndpoint(proxyURL string) endpoint.Endpoint {
	return httptransport.NewClient(
		"GET",
		mustParseURL(proxyURL),
		encodeUppercaseRequest,
		decodeUppercaseResponse,
	).Endpoint()
}
