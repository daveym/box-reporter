package box

// MockClient - Used for mocking
type MockClient struct {
	_PublicKeyID string
	_ClientID    string
	_ClaimSub    string
}

// SetPublicKeyID - Set the Box Public Key ID
func (p *MockClient) SetPublicKeyID(newKey string) {
	p._PublicKeyID = newKey
}

// GetPublicKeyID - Get the Box Public Key ID
func (p *MockClient) GetPublicKeyID() string {
	return p._PublicKeyID
}

// SetClientKeyID - Set the Box ClientID
func (p *MockClient) SetClientKeyID(newClientID string) {
	p._ClientID = newClientID
}

// GetClientID - Get the Box ClientID
func (p *MockClient) GetClientID() string {
	return p._ClientID
}

// CreateJWTAssertion - build up the JSON Web Token for oAuth
func (p *MockClient) CreateJWTAssertion(PublicKeyID string) error {
	var err error
	return err
}

// SetClaimSub - Set the Box ClientID
func (p *MockClient) SetClaimSub(newClaimSub string) {
	p._ClaimSub = newClaimSub
}

// GetClaimSub - Get the ClaimSub ID
func (p *MockClient) GetClaimSub() string {
	return p._ClaimSub
}
