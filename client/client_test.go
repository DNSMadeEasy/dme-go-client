package client

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetClientWithRequiredFields(t *testing.T) {
	apiKey := "testApiKey"
	secretKey := "testSecretKey"

	cl := GetClient(apiKey, secretKey)

	assert.NotNil(t, cl)
	assert.NotNil(t, cl.httpclient)
	assert.Equal(t, apiKey, cl.apiKey)
	assert.Equal(t, secretKey, cl.secretKey)
	assert.Equal(t, baseURL, cl.baseURL)
	assert.False(t, cl.insecure)
	assert.Empty(t, cl.proxyURL)

	// Verify that the proxy is not set
	transport := cl.httpclient.Transport.(*http.Transport)
	assert.Nil(t, transport.Proxy)
}

func TestGetClientWithOptionalFields(t *testing.T) {
	apiKey := "testApiKey"
	secretKey := "testSecretKey"
	customBaseURL := "https://custom.api.test.com"
	proxyURL := "http://proxy.test.com"

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
}
