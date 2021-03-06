package ciolite

// Api functions that support: https://context.io/docs/lite/connect_tokens

import (
	"fmt"

	"github.com/contextio/contextio-go/cioutil"
)

// GetOAuthProvidersResponse data struct
// 	https://context.io/docs/lite/oauth_providers#get
// 	https://context.io/docs/lite/oauth_providers#id-get
type GetOAuthProvidersResponse struct {
	Type                   string `json:"type,omitempty"`
	ProviderConsumerKey    string `json:"provider_consumer_key,omitempty"`
	ProviderConsumerSecret string `json:"provider_consumer_secret,omitempty"`
	ResourceURL            string `json:"resource_url,omitempty"`
}

// CreateOAuthProviderParams form values data struct.
// Requires Type, ProviderConsumerKey, ProviderConsumerSecret.
// 	https://context.io/docs/lite/oauth_providers#post
type CreateOAuthProviderParams struct {
	// Requires:
	Type                   string `json:"type"`
	ProviderConsumerKey    string `json:"provider_consumer_key"`
	ProviderConsumerSecret string `json:"provider_consumer_secret"`
}

// CreateOAuthProviderResponse data struct
// 	https://context.io/docs/lite/oauth_providers#post
type CreateOAuthProviderResponse struct {
	Success             bool   `json:"success,omitempty"`
	ProviderConsumerKey string `json:"provider_consumer_key,omitempty"`
	ResourceURL         string `json:"resource_url,omitempty"`
}

// DeleteOAuthProviderResponse data struct
// 	https://context.io/docs/lite/oauth_providers#id-delete
type DeleteOAuthProviderResponse struct {
	Success bool `json:"success,omitempty"`
}

// GetOAuthProviders get the list of OAuth providers configured.
// 	https://context.io/docs/lite/oauth_providers#get
func (cioLite CioLite) GetOAuthProviders() ([]GetOAuthProvidersResponse, error) {

	// Make request
	request := cioutil.ClientRequest{
		Method: "GET",
		Path:   "/oauth_providers",
	}

	// Make response
	var response []GetOAuthProvidersResponse

	// Request
	err := cioLite.DoFormRequest(request, &response)

	return response, err
}

// GetOAuthProvider gets information about a given OAuth provider.
// 	https://context.io/docs/lite/oauth_providers#id-get
func (cioLite CioLite) GetOAuthProvider(key string) (GetOAuthProvidersResponse, error) {

	// Make request
	request := cioutil.ClientRequest{
		Method: "GET",
		Path:   fmt.Sprintf("/oauth_providers/%s", key),
	}

	// Make response
	var response GetOAuthProvidersResponse

	// Request
	err := cioLite.DoFormRequest(request, &response)

	return response, err
}

// CreateOAuthProvider adds a new OAuth2 provider.
// formValues requires Type, ProviderConsumerKey, and ProviderConsumerSecret
// 	https://context.io/docs/lite/oauth_providers#post
func (cioLite CioLite) CreateOAuthProvider(formValues CreateOAuthProviderParams) (CreateOAuthProviderResponse, error) {

	// Make request
	request := cioutil.ClientRequest{
		Method:     "POST",
		Path:       "/oauth_providers",
		FormValues: formValues,
	}

	// Make response
	var response CreateOAuthProviderResponse

	// Request
	err := cioLite.DoFormRequest(request, &response)

	return response, err
}

// DeleteOAuthProvider removes a given OAuth provider.
// 	https://context.io/docs/lite/oauth_providers#id-delete
func (cioLite CioLite) DeleteOAuthProvider(key string) (DeleteOAuthProviderResponse, error) {

	// Make request
	request := cioutil.ClientRequest{
		Method: "DELETE",
		Path:   fmt.Sprintf("/oauth_providers/%s", key),
	}

	// Make response
	var response DeleteOAuthProviderResponse

	// Request
	err := cioLite.DoFormRequest(request, &response)

	return response, err
}
