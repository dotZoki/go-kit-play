package main

import (
	"context"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

// ----

// Endpoint tpye
type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

// ------------------------------------------------

func main() {
	svc := stringService{}

	uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)

	countHandler := httptransport.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
