package box

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// API - Base interface type for Box API. Allows us to mock/test.
type API interface {
	SetPublicKeyID(string)
	GetPublicKeyID() string
	SetClientID(string)
	GetClientID() string
	CreateJWTAssertion(string, string) error
}

// Client -
type Client struct {
	_PublicKeyID string
	_ClientID    string
}

// SetPublicKeyID - Set the Box Public Key ID
func (p *Client) SetPublicKeyID(newKey string) {
	p._PublicKeyID = newKey
}

// GetPublicKeyID - Get the Box Public Key ID
func (p *Client) GetPublicKeyID() string {
	return p._PublicKeyID
}

// SetClientKeyID - Set the Box ClientID
func (p *Client) SetClientID(newClientID string) {
	p._ClientID = newClientID
}

// GetClientID - Get the Box ClientID
func (p *Client) GetClientID() string {
	return p._ClientID
}

// CreateJWTAssertion - build up the JSON Web Token for oAuth
func (p *Client) CreateJWTAssertion(PublicKeyID string, ClientID string) error {

	var privateKey []byte
	var err error
	var tokenString string

	privateKey, _ = ioutil.ReadFile("/keys/private_key.pem")

	token := jwt.New(jwt.GetSigningMethod("RS256"))

	// Build JWT Header - https://docs.box.com/v2.0/docs/app-auth
	token.Header["alg"] = "RS256"
	token.Header["typ"] = "JWT"
	token.Header["kid"] = PublicKeyID

	// Build JWT Claims - https://docs.box.com/v2.0/docs/app-auth
	token.Claims["iss"] = ClientID
	token.Claims["sub"] = "This is my super fake ID"
	token.Claims["box_sub_type"] = "This is my super fake ID"
	token.Claims["aud"] = "This is my super fake ID"
	token.Claims["jti"] = "This is my super fake ID"
	token.Claims["exp"] = "This is my super fake ID"
	token.Claims["iat"] = "This is my super fake ID"
	token.Claims["nbf"] = "This is my super fake ID"
	token.Claims["exp"] = time.Now().Unix() + 36000

	// Sign the JWT
	tokenString, _ = token.SignedString(privateKey)

	println(tokenString)
	return err
}

// Generic post method, url and data are incoming. Response is a  base interface
// that we can use to return many response types.
func postJSON(action string, url string, data []byte, resp interface{}) (err error) {

	req, err := http.NewRequest(action, url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("charset", "UTF8")
	req.Header.Set("X-Accept", "application/json")

	client := &http.Client{}
	jsonResp, err := client.Do(req)

	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.NewDecoder(jsonResp.Body).Decode(resp)

	return err
}
