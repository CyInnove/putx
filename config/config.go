package config

import (
	"crypto/tls"
	"net/http"
)

// GetHTTPClient returns an HTTP client with custom settings.
func GetHTTPClient() *http.Client {
	// Create an HTTP client with TLS settings to skip certificate verification
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
	}

	return client
}
