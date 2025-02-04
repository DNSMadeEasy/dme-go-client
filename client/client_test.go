package client

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetClient(t *testing.T) {
	t.Run("should get client with required fields", func(t *testing.T) {
		resetClient()
		apiKey := "testApiKey"
		secretKey := "testSecretKey"

		cl := GetClient(apiKey, secretKey)

		assert.NotNil(t, cl)
		assert.NotNil(t, cl.httpclient)
		assert.Equal(t, apiKey, cl.apiKey)
		assert.Equal(t, secretKey, cl.secretKey)
		assert.Equal(t, "https://api.dnsmadeeasy.com/V2.0", cl.baseURL)
		assert.False(t, cl.insecure)
		assert.Empty(t, cl.proxyURL)

		// Verify that the proxy is not set
		transport := cl.httpclient.Transport.(*http.Transport)
		assert.Nil(t, transport.Proxy)
	})

	t.Run("should get client with optional fields", func(t *testing.T) {
		resetClient()
		apiKey := "testApiKey"
		secretKey := "testSecretKey"
		customBaseURL := "https://custom.api.test.com"
		proxyURL := "https://proxy.test.com"

		cl := GetClient(
			apiKey,
			secretKey,
			BaseURL(customBaseURL),
			Insecure(true),
			ProxyURL(proxyURL),
		)

		assert.NotNil(t, cl)
		assert.Equal(t, apiKey, cl.apiKey)
		assert.Equal(t, secretKey, cl.secretKey)
		assert.Equal(t, customBaseURL, cl.baseURL)
		assert.True(t, cl.insecure)
		assert.Equal(t, proxyURL, cl.proxyURL)
		assert.NotNil(t, cl.httpclient)

		// Verify that the proxy is set correctly
		transport := cl.httpclient.Transport.(*http.Transport)
		assert.NotNil(t, transport.Proxy)
		proxy, err := transport.Proxy(&http.Request{})
		assert.NoError(t, err)
		assert.Equal(t, proxyURL, proxy.String())
	})

	t.Run("should use default base url if custom base url is '/' only", func(t *testing.T) {
		resetClient()
		apiKey := "testApiKey"
		secretKey := "testSecretKey"
		customBaseURL := "/"

		cl := GetClient(
			apiKey,
			secretKey,
			BaseURL(customBaseURL),
		)

		assert.NotNil(t, cl)
		assert.Equal(t, "https://api.dnsmadeeasy.com/V2.0", cl.baseURL)
	})

	t.Run("should sanitize base url with single trailing '/'", func(t *testing.T) {
		resetClient()
		apiKey := "testApiKey"
		secretKey := "testSecretKey"
		customBaseURL := "https://custom.api.with-trailing-slash.com/"

		cl := GetClient(
			apiKey,
			secretKey,
			BaseURL(customBaseURL),
		)

		assert.NotNil(t, cl)
		assert.Equal(t, "https://custom.api.with-trailing-slash.com", cl.baseURL)
	})

	t.Run("should sanitize base url with multiple trailing '////'", func(t *testing.T) {
		resetClient()
		apiKey := "testApiKey"
		secretKey := "testSecretKey"
		customBaseURL := "https://custom.api.with-trailing-slash.com////"

		cl := GetClient(
			apiKey,
			secretKey,
			BaseURL(customBaseURL),
		)

		assert.NotNil(t, cl)
		assert.Equal(t, "https://custom.api.with-trailing-slash.com", cl.baseURL)
	})

	t.Run("should sanitize url with leading/trailing white spaces, missing protocol and trailing slash", func(t *testing.T) {
		resetClient()
		apiKey := "testApiKey"
		secretKey := "testSecretKey"
		customBaseURL := "   custom.api.com////   "

		cl := GetClient(
			apiKey,
			secretKey,
			BaseURL(customBaseURL),
		)

		assert.NotNil(t, cl)
		assert.Equal(t, "https://custom.api.com", cl.baseURL)
	})
}
