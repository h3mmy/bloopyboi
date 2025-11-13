package services

import (
	httptransport "github.com/go-openapi/runtime/client"
	"net/http"
)

const (
	// ServiceLoggerFieldKey is the key for the service name in the logger.
	ServiceLoggerFieldKey = "service_name"
	// DefaultUUID is the default UUID for a bloopyboi instance.
	DefaultUUID = ""
)

// GetTLSTransport returns a TLS transport that can be configured to skip verification.
// TODO: This should not panic.
func GetTLSTransport(insecure bool) http.RoundTripper {
	tlsTransport, err := httptransport.TLSTransport(httptransport.TLSClientOptions{
		InsecureSkipVerify: insecure,
	})
	if err != nil {
		panic(err)
	}
	return tlsTransport
}
