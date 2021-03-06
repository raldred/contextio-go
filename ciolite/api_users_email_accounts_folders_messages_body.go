package ciolite

// Api functions that support: https://context.io/docs/lite/users/email_accounts/folders/messages/body

import (
	"fmt"
	"net/url"

	"github.com/contextio/contextio-go/cioutil"
)

// GetUserEmailAccountsFolderMessageBodyParams query values data struct.
// Optional: Delimiter, Type.
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/body#get
type GetUserEmailAccountsFolderMessageBodyParams struct {
	// Optional:
	Delimiter string `json:"delimiter,omitempty"`
	Type      string `json:"type,omitempty"`
}

// GetUserEmailAccountsFolderMessageBodyResponse data struct
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/body#get
type GetUserEmailAccountsFolderMessageBodyResponse struct {
	Type        string `json:"type,omitempty"`
	Charset     string `json:"charset,omitempty"`
	Content     string `json:"content,omitempty"`
	BodySection string `json:"body_section,omitempty"`
}

// GetUserEmailAccountsFolderMessageBody fetches the message body of a given email.
// queryValues may optionally contain Delimiter, Type
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/body#get
func (cioLite CioLite) GetUserEmailAccountsFolderMessageBody(userID string, label string, folder string, messageID string, queryValues GetUserEmailAccountsFolderMessageBodyParams) ([]GetUserEmailAccountsFolderMessageBodyResponse, error) {

	// Make request
	request := cioutil.ClientRequest{
		Method:      "GET",
		Path:        fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s/messages/%s/body", userID, label, url.QueryEscape(folder), url.QueryEscape(messageID)),
		QueryValues: queryValues,
	}

	// Make response
	var response []GetUserEmailAccountsFolderMessageBodyResponse

	// Request
	err := cioLite.DoFormRequest(request, &response)

	return response, err
}
