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
	SetClientID(string)
	GetClientID() string
	CreateJWTAssertion() error
}

// Client -
type Client struct {
	_ClientID string
}

// SetClientID - Set the ClientID
func (p *Client) SetClientID(newKey string) {
	p._ClientID = newKey
}

// GetClientID - Get the ClientID
func (p *Client) GetClientID() string {
	return p._ClientID
}

// CreateJWTAssertion - build up the JSON Web Token for oAuth
func (p *Client) CreateJWTAssertion() error {

	var privateKey []byte
	var err error
	var tokenString string

	privateKey, _ = ioutil.ReadFile("Location of your demo.rsa")

	token := jwt.New(jwt.GetSigningMethod("RS256"))

	token.Header["kid"] = "Public Key ID"

	token.Claims["ID"] = "This is my super fake ID"
	token.Claims["exp"] = time.Now().Unix() + 36000
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
