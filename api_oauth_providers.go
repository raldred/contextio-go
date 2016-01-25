package ciolite

// Api functions that support: https://context.io/docs/lite/connect_tokens

// Imports
import (
	"fmt"
)

// GetOAuthProvidersResponse ...
type GetOAuthProvidersResponse struct {
	Type                   string `json:"type,omitempty"`
	ProviderConsumerKey    string `json:"provider_consumer_key,omitempty"`
	ProviderConsumerSecret string `json:"provider_consumer_secret,omitempty"`
	ResourceURL            string `json:"resource_url,omitempty"`
}

// CreateOAuthProviderResponse ...
type CreateOAuthProviderResponse struct {
	Success             string `json:"success,omitempty"`
	ProviderConsumerKey string `json:"provider_consumer_key,omitempty"`
	ResourceURL         string `json:"resource_url,omitempty"`
}

// DeleteOAuthProviderResponse ...
type DeleteOAuthProviderResponse struct {
	Success string `json:"success,omitempty"`
}

// GetOAuthProviders get the list of OAuth providers configured.
// 	https://context.io/docs/lite/oauth_providers#get
func (cioLite *CioLite) GetOAuthProviders() ([]GetOAuthProvidersResponse, error) {

	// Make request
	request := clientRequest{
		method: "GET",
		path:   "/oauth_providers",
	}

	// Make response
	var response []GetOAuthProvidersResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// GetOAuthProvider gets information about a given OAuth provider.
// 	https://context.io/docs/lite/oauth_providers#id-get
func (cioLite *CioLite) GetOAuthProvider(key string) (GetOAuthProvidersResponse, error) {

	// Make request
	request := clientRequest{
		method: "GET",
		path:   fmt.Sprintf("/oauth_providers/%s", key),
	}

	// Make response
	var response GetOAuthProvidersResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// CreateOAuthProvider adds a new OAuth2 provider.
// formValues requires CioParams.Type, CioParams.ProviderConsumerKey, and CioParams.ProviderConsumerSecret
// 	https://context.io/docs/lite/oauth_providers#post
func (cioLite *CioLite) CreateOAuthProvider(formValues CioParams) (CreateOAuthProviderResponse, error) {

	// Make request
	request := clientRequest{
		method:     "POST",
		path:       "/oauth_providers",
		formValues: formValues,
	}

	// Make response
	var response CreateOAuthProviderResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// DeleteOAuthProvider removes a given OAuth provider.
// 	https://context.io/docs/lite/oauth_providers#id-delete
func (cioLite *CioLite) DeleteOAuthProvider(key string) (DeleteOAuthProviderResponse, error) {

	// Make request
	request := clientRequest{
		method: "DELETE",
		path:   fmt.Sprintf("/oauth_providers/%s", key),
	}

	// Make response
	var response DeleteOAuthProviderResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}
