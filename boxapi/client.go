package box

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// API - Base interface type for Box API. Allows us to mock/test.
type API interface {
	SetPublicKeyID(string)
	GetPublicKeyID() string
	SetClientID(string)
	GetClientID() string
	SetClaimSub(string)
	GetClaimSub() string
	CreateJWTAssertion(string, string, string) (string, error)
}

// Client -
type Client struct {
	_PublicKeyID string
	_ClientID    string
	_ClaimSub    string
}

// SetPublicKeyID - Set the Box Public Key ID
func (p *Client) SetPublicKeyID(newKey string) {
	p._PublicKeyID = newKey
}

// GetPublicKeyID - Get the Box Public Key ID
func (p *Client) GetPublicKeyID() string {
	return p._PublicKeyID
}

// SetClientID - Set the Box ClientID
func (p *Client) SetClientID(newClientID string) {
	p._ClientID = newClientID
}

// GetClientID - Get the Box ClientID
func (p *Client) GetClientID() string {
	return p._ClientID
}

// SetClaimSub - Set the Box ClientID
func (p *Client) SetClaimSub(newClaimSub string) {
	p._ClaimSub = newClaimSub
}

// GetClaimSub - Get the ClaimSub ID
func (p *Client) GetClaimSub() string {
	return p._ClaimSub
}

// CreateJWTAssertion - build up the JSON Web Token for oAuth
func (p *Client) CreateJWTAssertion(PublicKeyID string, ClientID string, ClaimSub string) (string, error) {

	var signingKey []byte
	var err error
	var msg string
	var tokenString string

	signingKey, err = ioutil.ReadFile("./keys/private_key.pem")

	// Generate JTI Value
	jti, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		msg = "Unable to read signing key. Please ensure your signing key 'private_key.pem' is in the ./keys/ directory"
		return msg, err
	}

	token := jwt.New(jwt.GetSigningMethod("RS256"))
	// Build JWT Header - https://docs.box.com/v2.0/docs/app-auth
	token.Header["alg"] = "RS256"
	token.Header["typ"] = "JWT"
	token.Header["kid"] = PublicKeyID

	// Build JWT Claims - https://docs.box.com/v2.0/docs/app-auth
	token.Claims["iss"] = ClientID
	token.Claims["sub"] = ClaimSub
	token.Claims["box_sub_type"] = "enterprise"
	token.Claims["aud"] = JWTAUTHURL
	token.Claims["jti"] = jti
	token.Claims["exp"] = time.Now().Unix() + 36000

	// Sign the JWT
	tokenString, err = token.SignedString(signingKey)

	fmt.Println(err)
	fmt.Println(tokenString)

	if err != nil {
		msg = "Unable to sign token, please check that you have a signing key in ./keys/"
		return msg, err
	}

	return msg, err
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
