package main

import (
	"log"
	"net/http"

	"github.com/MehulxBuilds/go-ratelimiter/internal/service"
)

func main() {

	// Request Handler
	handler := service.EndpointHandler

	// Middlewares for specific rate limiting strategies
	bucketMiddleware := service.BucketRateLimiter
	perClientMiddleware := service.PerClientRateLimiter
	tollboothMiddleware := service.TollboothRateLimiter

	// Impl Token Bucket
	http.Handle("/ping/bucket", bucketMiddleware(handler))

	// Impl Per-Client Limiting
	http.Handle("/ping/per-client", perClientMiddleware(handler))

	// Impl Tollbooth Limiting
	http.Handle("/ping/tollbooth", tollboothMiddleware(handler))
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("There was an error listening on port :8080", err)
	}
}
