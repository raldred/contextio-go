package ciolite

// Api functions that support: https://context.io/docs/lite/users/email_accounts/folders/messages

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/contextio/contextio-go/cioutil"
)

// GetUserEmailAccountsFolderMessageParams query values data struct.
// Optional: Delimiter, IncludeBody, BodyType, IncludeHeaders, IncludeFlags,
// and (for GetUserEmailAccountsFolderMessages only) Limit, Offset.
// 	https://context.io/docs/lite/users/email_accounts/folders/messages#get
// 	https://context.io/docs/lite/users/email_accounts/folders/messages#id-get
type GetUserEmailAccountsFolderMessageParams struct {
	// Optional:
	Delimiter    string `json:"delimiter,omitempty"`
	BodyType     string `json:"body_type,omitempty"`
	IncludeBody  bool   `json:"include_body,omitempty"`
	IncludeFlags bool   `json:"include_flags,omitempty"`

	// IncludeHeaders can be "0", "1", or "raw"
	IncludeHeaders string `json:"include_headers,omitempty"`

	// Optional for GetUserEmailAccountsFolderMessages (not used by GetUserEmailAccountFolderMessage):
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
}

// GetUsersEmailAccountFolderMessagesResponse data struct
// 	https://context.io/docs/lite/users/email_accounts/folders/messages#get
// 	https://context.io/docs/lite/users/email_accounts/folders/messages#id-get
type GetUsersEmailAccountFolderMessagesResponse struct {
	EmailMessageID string `json:"email_message_id,omitempty"`
	Subject        string `json:"subject,omitempty"`
	MessageID      string `json:"message_id,omitempty"`
	InReplyTo      string `json:"in_reply_to,omitempty"`
	ResourceURL    string `json:"resource_url,omitempty"`

	Folders         []string `json:"folders,omitempty"`
	ListHeaders     []string `json:"list_headers,omitempty"`
	References      []string `json:"references,omitempty"`
	ReceivedHeaders []string `json:"received_headers,omitempty"`

	Addresses GetUsersEmailAccountFolderMessageAddresses `json:"addresses,omitempty"`

	PersonInfo PersonInfo `json:"person_info,omitempty"`

	Attachments []struct {
		Type               string `json:"type,omitempty"`
		FileName           string `json:"file_name,omitempty"`
		BodySection        string `json:"body_section,omitempty"`
		ContentDisposition string `json:"content_disposition,omitempty"`
		EmailMessageID     string `json:"email_message_id,omitempty"`
		XAttachmentID      string `json:"x_attachment_id,omitempty"`

		Size         int `json:"size,omitempty"`
		AttachmentID int `json:"attachment_id,omitempty"`
	} `json:"attachments,omitempty"`

	Bodies []struct {
		BodySection string `json:"body_section,omitempty"`
		Type        string `json:"type,omitempty"`
		Encoding    string `json:"encoding,omitempty"`

		Size int `json:"size,omitempty"`
	} `json:"bodies,omitempty"`

	SentAt     int `json:"sent_at,omitempty"`
	ReceivedAt int `json:"received_at,omitempty"`
}

// PersonInfo data struct within GetUsersEmailAccountFolderMessagesResponse and WebhookMessageData
// 	https://context.io/docs/lite/users/email_accounts/folders/messages#get
// 	https://context.io/docs/lite/users/email_accounts/folders/messages#id-get
type PersonInfo map[string]map[string]string

// UnmarshalJSON is here because the empty state is an array in the json, and is a object/map when populated
func (m *PersonInfo) UnmarshalJSON(b []byte) error {
	if bytes.Equal([]byte(`[]`), b) {
		// its the empty array, set an empty map
		*m = make(map[string]map[string]string)
		return nil
	}
	mp := make(map[string]map[string]string)
	err := json.Unmarshal(b, &mp)
	if err != nil {
		return err
	}
	*m = mp
	return nil
}

// GetUsersEmailAccountFolderMessageAddresses data struct within GetUsersEmailAccountFolderMessagesResponse
// 	https://context.io/docs/lite/users/email_accounts/folders/messages#get
// 	https://context.io/docs/lite/users/email_accounts/folders/messages#id-get
type GetUsersEmailAccountFolderMessageAddresses struct {
	From []struct {
		Email string `json:"email,omitempty"`
		Name  string `json:"name,omitempty"`
	} `json:"from,omitempty"`

	To []struct {
		Email string `json:"email,omitempty"`
		Name  string `json:"name,omitempty"`
	} `json:"to,omitempty"`

	Cc []struct {
		Email string `json:"email,omitempty"`
		Name  string `json:"name,omitempty"`
	} `json:"cc,omitempty"`

	Bcc []struct {
		Email string `json:"email,omitempty"`
		Name  string `json:"name,omitempty"`
	} `json:"bcc,omitempty"`

	Sender []struct {
		Email string `json:"email,omitempty"`
		Name  string `json:"name,omitempty"`
	} `json:"sender,omitempty"`

	ReplyTo []struct {
		Email string `json:"email,omitempty"`
		Name  string `json:"name,omitempty"`
	} `json:"reply_to,omitempty"`
}

// MoveUserEmailAccountFolderMessageParams form values data struct.
// Requires: NewFolderID, and may optionally contain Delimiter.
// 	https://context.io/docs/lite/users/email_accounts/folders/messages#id-put
type MoveUserEmailAccountFolderMessageParams struct {
	// Required:
	NewFolderID string `json:"new_folder_id"`
	// Optional:
	Delimiter string `json:"delimiter,omitempty"`
}

// MoveUserEmailAccountFolderMessageResponse data struct
// 	https://context.io/docs/lite/users/email_accounts/folders/messages#id-put
type MoveUserEmailAccountFolderMessageResponse struct {
	Success bool `json:"success,omitempty"`
}

// GetUserEmailAccountsFolderMessages gets listings of email messages for a user.
// queryValues may optionally contain Delimiter, IncludeBody, BodyType,
// IncludeHeaders, IncludeFlags, Limit, Offset
// 	https://context.io/docs/lite/users/email_accounts/folders/messages#get
func (cioLite CioLite) GetUserEmailAccountsFolderMessages(userID string, label string, folder string, queryValues GetUserEmailAccountsFolderMessageParams) ([]GetUsersEmailAccountFolderMessagesResponse, error) {

	// Make request
	request := cioutil.ClientRequest{
		Method:      "GET",
		Path:        fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s/messages", userID, label, url.QueryEscape(folder)),
		QueryValues: queryValues,
	}

	// Make response
	var response []GetUsersEmailAccountFolderMessagesResponse

	// Request
	err := cioLite.DoFormRequest(request, &response)

	return response, err
}

// GetUserEmailAccountFolderMessage gets file, contact and other information about a given email message.
// queryValues may optionally contain Delimiter, IncludeBody, BodyType, IncludeHeaders, IncludeFlags
// 	https://context.io/docs/lite/users/email_accounts/folders/messages#id-get
func (cioLite CioLite) GetUserEmailAccountFolderMessage(userID string, label string, folder string, messageID string, queryValues GetUserEmailAccountsFolderMessageParams) (GetUsersEmailAccountFolderMessagesResponse, error) {

	// Make request
	request := cioutil.ClientRequest{
		Method:      "GET",
		Path:        fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s/messages/%s", userID, label, url.QueryEscape(folder), url.QueryEscape(messageID)),
		QueryValues: queryValues,
	}

	// Make response
	var response GetUsersEmailAccountFolderMessagesResponse

	// Request
	err := cioLite.DoFormRequest(request, &response)

	return response, err
}

// MoveUserEmailAccountFolderMessage moves a message.
// formValues requires NewFolderID, and may optionally contain Delimiter
// 	https://context.io/docs/lite/users/email_accounts/folders/messages#id-put
func (cioLite CioLite) MoveUserEmailAccountFolderMessage(userID string, label string, folder string, messageID string, queryValues MoveUserEmailAccountFolderMessageParams) (MoveUserEmailAccountFolderMessageResponse, error) {

	// Make request
	request := cioutil.ClientRequest{
		Method:      "PUT",
		Path:        fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s/messages/%s", userID, label, url.QueryEscape(folder), url.QueryEscape(messageID)),
		QueryValues: queryValues,
	}

	// Make response
	var response MoveUserEmailAccountFolderMessageResponse

	// Request
	err := cioLite.DoFormRequest(request, &response)

	return response, err
}
