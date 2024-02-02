package services

import (
	httptransport "github.com/go-openapi/runtime/client"
	"net/http"
)

const (
	ServiceLoggerFieldKey = "service_name"
	DefaultUUID = ""
)

// GetTLSTransport Get a TLS transport instance, that skips verification if configured via environment variables.
func GetTLSTransport(insecure bool) http.RoundTripper {
	tlsTransport, err := httptransport.TLSTransport(httptransport.TLSClientOptions{
		InsecureSkipVerify: insecure,
	})
	if err != nil {
		panic(err)
	}
	return tlsTransport
}
