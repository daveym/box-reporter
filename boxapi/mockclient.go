package box

// MockClient - Used for mocking
type MockClient struct {
	_ClientID string
}

// SetClientID - Set the ClientID
func (p *MockClient) SetClientID(newKey string) {
	p._ClientID = newKey
}

// GetClientID - Get the ClientID
func (p *MockClient) GetClientID() string {
	return p._ClientID
}

// Authorise - Mock instance
func (p *MockClient) Authorise(url string, code string, uri string) error {

	var err error

	return err
}
