package ciolite

// Api functions that support: https://context.io/docs/lite/users/email_accounts/folders

import (
	"errors"
	"fmt"
)

// GetUserEmailAcountsFoldersParams query values data struct.
// Optional: IncludeNamesOnly.
// 	https://context.io/docs/lite/users/email_accounts/folders#get
type GetUserEmailAcountsFoldersParams struct {
	// Optional:
	IncludeNamesOnly bool `json:"include_names_only,omitempty"`
}

// GetUsersEmailAccountFoldersResponse data struct
// 	https://context.io/docs/lite/users/email_accounts/folders#get
// 	https://context.io/docs/lite/users/email_accounts/folders#id-get
type GetUsersEmailAccountFoldersResponse struct {
	Name             string `json:"name,omitempty"`
	SymbolicName     string `json:"symbolic_name,omitempty"`
	NbMessages       int    `json:"nb_messages,omitempty"`
	NbUnseenMessages int    `json:"nb_unseen_messages,omitempty"`
	Delimiter        string `json:"delimiter,omitempty"`
	ResourceURL      string `json:"resource_url,omitempty"`
}

// EmailAccountFolderParams query values data struct.
// Optional: Delimiter
// 	https://context.io/docs/lite/users/email_accounts/folders#id-get
// 	https://context.io/docs/lite/users/email_accounts/folders#id-post
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/attachments#get
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/attachments#id-get
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/flags#get
//  https://context.io/docs/lite/users/email_accounts/folders/messages/raw#get
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/read#post
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/read#delete
type EmailAccountFolderParams struct {
	// Optional:
	Delimiter string `json:"delimiter,omitempty"`
}

// CreateEmailAccountFolderResponse data struct
// 	https://context.io/docs/lite/users/email_accounts/folders#id-post
type CreateEmailAccountFolderResponse struct {
	Success bool `json:"success,omitempty"`
}

// GetUserEmailAccountsFolders gets a list of folders in an email account.
// queryValues may optionally contain IncludeNamesOnly
// 	https://context.io/docs/lite/users/email_accounts/folders#get
func (cioLite CioLite) GetUserEmailAccountsFolders(userID string, label string, queryValues GetUserEmailAcountsFoldersParams) ([]GetUsersEmailAccountFoldersResponse, error) {

	// Make request
	request := clientRequest{
		method:      "GET",
		path:        fmt.Sprintf("/users/%s/email_accounts/%s/folders", userID, label),
		queryValues: queryValues,
	}

	// Make response
	var response []GetUsersEmailAccountFoldersResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// GetUserEmailAccountFolder gets information about a given folder.
// queryValues may optionally contain Delimiter
// 	https://context.io/docs/lite/users/email_accounts/folders#id-get
func (cioLite CioLite) GetUserEmailAccountFolder(userID string, label string, folder string, queryValues EmailAccountFolderParams) (GetUsersEmailAccountFoldersResponse, error) {

	// Make request
	request := clientRequest{
		method:      "GET",
		path:        fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s", userID, label, folder),
		queryValues: queryValues,
	}

	// Make response
	var response GetUsersEmailAccountFoldersResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// CreateUserEmailAccountFolder create a folder on an email account.
// This call will fail if the folder already exists.
// queryValues may optionally contain Delimiter
// 	https://context.io/docs/lite/users/email_accounts/folders#id-post
func (cioLite CioLite) CreateUserEmailAccountFolder(userID string, label string, folder string, formValues EmailAccountFolderParams) (CreateEmailAccountFolderResponse, error) {

	// Make request
	request := clientRequest{
		method:     "POST",
		path:       fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s", userID, label, folder),
		formValues: formValues,
	}

	// Make response
	var response CreateEmailAccountFolderResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// SafeCreateUserEmailAccountFolder will safely check if a folder exists, and create it if it does not.
// This function returns a bool representing whether it had to create a folder, and any errors it received.
// queryValues may optionally contain Delimiter
func (cioLite CioLite) SafeCreateUserEmailAccountFolder(userID string, label string, folder string, formValues EmailAccountFolderParams) (bool, error) {

	existsResponse, err := cioLite.GetUserEmailAccountFolder(userID, label, folder, formValues)
	if err == nil && existsResponse.Name == folder {
		// It exists already, so return false and no error
		return false, nil
	}

	createResponse, err := cioLite.CreateUserEmailAccountFolder(userID, label, folder, formValues)
	if err != nil {
		return true, err
	}
	if !createResponse.Success {
		return true, errors.New("Unable to create folder")
	}
	// Created successfully
	return true, nil
}
