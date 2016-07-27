package box

import "errors"

// MockClient - Used for mocking
type MockClient struct {
	_consumerKey string
	_accessToken string
}

// SetConsumerKey - Set the new consumer Key
func (p *MockClient) SetConsumerKey(newKey string) {
	p._consumerKey = newKey
}

// GetConsumerKey - Set the new consumer Key
func (p *MockClient) GetConsumerKey() string {
	return p._consumerKey
}

// SetAccessToken - Set the new access token
func (p *MockClient) SetAccessToken(newToken string) {
	p._accessToken = newToken
}

// GetAccessToken - Set the new access token
func (p *MockClient) GetAccessToken() string {
	return p._accessToken
}

// UserAuthorise - Mock instance
func (p *MockClient) UserAuthorise(url string, code string, uri string) error {

	var err error

	if p.GetConsumerKey() == "INVALIDBROWSER" {
		err = errors.New("Invalid Key")
	}
	return err
}

// RetrieveAccessToken -  Mock instance
func (p *MockClient) RetrieveAccessToken(consumerKey string, code string, resp *AuthorisationResponse) error {

	var err error
	if consumerKey == "FAIL" {
		err = errors.New("Invalid Key")
		return err
	}

	return nil
}
