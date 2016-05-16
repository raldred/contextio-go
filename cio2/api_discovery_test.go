package cio2

import (
	"reflect"
	"testing"
)

// TestActualDiscoveryRequestToCioForGoogle tests sending an actual
// GetDiscovery request to CIO, for gmail and googleapps accounts
// (internet connection required, real CIO key/secret required).
func TestActualDiscoveryRequestToCioForGoogle(t *testing.T) {
	t.Parallel()

	cio2, logger := NewTestCio2WithLogger(t)

	expected := GetDiscoveryResponse{
		Email: "test@gmail.com",
		Type:  "gmail",
		Found: true,
		IMAP: GetDiscoveryIMAPResponse{
			Server:   "imap.gmail.com",
			Username: "test@gmail.com",
			UseSSL:   true,
			OAuth:    true,
			Port:     993,
		},
	}

	// gmail
	response, err := cio2.GetDiscovery(GetDiscoveryParams{Email: "test@gmail.com", SourceType: "IMAP"})
	expected.Documentation = response.Documentation
	expected.ResourceURL = response.ResourceURL

	if err != nil || !reflect.DeepEqual(expected, response) {
		t.Error("Expected GetDiscovery Response: ", expected, "; Got: ", response, "; With Error: ", err, "; With Log: ", logger.String())
	}

	// googleapps
	expected.Email = "test@google.com"
	expected.IMAP.Username = "test@google.com"
	expected.Type = "googleapps"
	expected.IMAP.Server = "imap.googlemail.com"

	response, err = cio2.GetDiscovery(GetDiscoveryParams{Email: "test@google.com", SourceType: "IMAP"})
	expected.Documentation = response.Documentation
	expected.ResourceURL = response.ResourceURL

	if err != nil || !reflect.DeepEqual(expected, response) {
		t.Error("Expected GetDiscovery Response: ", expected, "; Got: ", response, "; With Error: ", err, "; With Log: ", logger.String())
	}

	if len(logger.String()) < 20 {
		t.Error("Expected some output from logger; Got: ", logger.String())
	}
}

// TestActualDiscoveryRequestToCioForMicrosoft tests sending an actual
// GetDiscovery request to CIO, for outlook and hotmail accounts
// (internet connection required, real CIO key/secret required).
func TestActualDiscoveryRequestToCioForMicrosoft(t *testing.T) {
	t.Parallel()

	cio2 := NewTestCio2(t)

	expected := GetDiscoveryResponse{
		Email: "test@hotmail.com",
		Type:  "msliveconnect",
		Found: true,
		IMAP: GetDiscoveryIMAPResponse{
			Server:   "imap-mail.outlook.com",
			Username: "test@hotmail.com",
			UseSSL:   true,
			OAuth:    true,
			Port:     993,
		},
	}

	// hotmail
	response, err := cio2.GetDiscovery(GetDiscoveryParams{Email: "test@hotmail.com", SourceType: "IMAP"})
	expected.Documentation = response.Documentation
	expected.ResourceURL = response.ResourceURL

	if err != nil || !reflect.DeepEqual(expected, response) {
		t.Error("Expected GetDiscovery Response: ", expected, "; Got: ", response, "; With Error: ", err)
	}

	// outlook
	expected.Email = "test@outlook.com"
	expected.IMAP.Username = "test@outlook.com"

	response, err = cio2.GetDiscovery(GetDiscoveryParams{Email: "test@outlook.com", SourceType: "IMAP"})
	expected.Documentation = response.Documentation
	expected.ResourceURL = response.ResourceURL

	if err != nil || !reflect.DeepEqual(expected, response) {
		t.Error("Expected GetDiscovery Response: ", expected, "; Got: ", response, "; With Error: ", err)
	}
}

// TestActualDiscoveryRequestToCioForYahoo tests sending an actual
// GetDiscovery request to CIO, for yahoo accounts
// (internet connection required, real CIO key/secret required).
func TestActualDiscoveryRequestToCioForYahoo(t *testing.T) {
	t.Parallel()

	cio2 := NewTestCio2(t)

	expected := GetDiscoveryResponse{
		Email: "test@yahoo.com",
		Type:  "generic",
		Found: true,
		IMAP: GetDiscoveryIMAPResponse{
			Server:   "imap.mail.yahoo.com",
			Username: "test@yahoo.com",
			UseSSL:   true,
			OAuth:    false,
			Port:     993,
		},
	}

	// yahoo
	response, err := cio2.GetDiscovery(GetDiscoveryParams{Email: "test@yahoo.com", SourceType: "IMAP"})
	expected.Documentation = response.Documentation
	expected.ResourceURL = response.ResourceURL

	if err != nil || !reflect.DeepEqual(expected, response) {
		t.Error("Expected GetDiscovery Response: ", expected, "; Got: ", response, "; With Error: ", err)
	}

	// yahoo.com.cn
	expected.Email = "test@yahoo.com.cn"
	expected.IMAP.Username = "test@yahoo.com.cn"

	response, err = cio2.GetDiscovery(GetDiscoveryParams{Email: "test@yahoo.com.cn", SourceType: "IMAP"})
	expected.Documentation = response.Documentation
	expected.ResourceURL = response.ResourceURL

	if err != nil || !reflect.DeepEqual(expected, response) {
		t.Error("Expected GetDiscovery Response: ", expected, "; Got: ", response, "; With Error: ", err)
	}
}

// TestActualDiscoveryRequestToCioForAol tests sending an actual
// GetDiscovery request to CIO, for an AOL account
// (internet connection required, real CIO key/secret required).
func TestActualDiscoveryRequestToCioForAol(t *testing.T) {
	t.Parallel()

	cio2 := NewTestCio2(t)

	expected := GetDiscoveryResponse{
		Email: "test@aol.com",
		Type:  "generic",
		Found: true,
		IMAP: GetDiscoveryIMAPResponse{
			Server:   "imap.aol.com",
			Username: "test@aol.com",
			UseSSL:   true,
			OAuth:    false,
			Port:     993,
		},
	}

	// aol
	response, err := cio2.GetDiscovery(GetDiscoveryParams{Email: "test@aol.com", SourceType: "IMAP"})
	expected.Documentation = response.Documentation
	expected.ResourceURL = response.ResourceURL

	if err != nil || !reflect.DeepEqual(expected, response) {
		t.Error("Expected GetDiscovery Response: ", expected, "; Got: ", response, "; With Error: ", err)
	}
}

// TestActualDiscoveryRequestToCioForNonExistent tests sending an actual
// GetDiscovery request to CIO, for a non-existent email service provider
// (internet connection required, real CIO key/secret required).
func TestActualDiscoveryRequestToCioForNonExistent(t *testing.T) {
	t.Parallel()

	cio2 := NewTestCio2(t)

	expected := GetDiscoveryResponse{
		Email: "test@bogusblahblahfoobar.com",
		Type:  "",
		Found: false,
		IMAP: GetDiscoveryIMAPResponse{
			Server:   "",
			Username: "",
			UseSSL:   false,
			OAuth:    false,
			Port:     0,
		},
	}

	// aol
	response, err := cio2.GetDiscovery(GetDiscoveryParams{Email: "test@bogusblahblahfoobar.com", SourceType: "IMAP"})
	expected.Documentation = response.Documentation
	expected.ResourceURL = response.ResourceURL

	if err != nil || !reflect.DeepEqual(expected, response) {
		t.Error("Expected GetDiscovery Response: ", expected, "; Got: ", response, "; With Error: ", err)
	}
}
