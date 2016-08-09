package box

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os/exec"
	"strings"
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
	SendOAuthRequest(string, string, string) (string, error)
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
func (p *Client) CreateJWTAssertion(PublicKeyID string, ClientID string, Sub string) (string, error) {

	var signingKey []byte
	var err error
	var msg, tokenString string

	signingKey, err = ioutil.ReadFile("./keys/private_key.pem")

	// Generate JTI Value
	jti, err := exec.Command("uuidgen").Output()
	if err != nil {
		fmt.Println(err.Error())
		return msg, err
	}

	if err != nil {
		msg = "Unable to read signing key. Please ensure your private signing key is in the ./keys/ directory"
		return msg, err
	}

	/* Keys generated using the following. Note, use a recent version of OpenSSL!
	./openssl genpkey -algorithm RSA -out private_key.pem -pkeyopt rsa_keygen_bits:2048
	./openssl rsa -pubout -in private_key.pem -out public_key.pem
	*/
	token := jwt.New(jwt.GetSigningMethod("RS256"))

	// Build JWT Header - https://docs.box.com/v2.0/docs/app-auth
	token.Header["alg"] = "RS256"
	token.Header["typ"] = "JWT"
	token.Header["kid"] = PublicKeyID

	// Build JWT Claims - https://docs.box.com/v2.0/docs/app-auth
	token.Claims["iss"] = ClientID
	token.Claims["sub"] = Sub
	token.Claims["box_sub_type"] = "enterprise"
	token.Claims["aud"] = JWTAUTHURL
	token.Claims["jti"] = jti
	token.Claims["exp"] = time.Now().Unix() + 36000

	// Sign the JWT
	tokenString, err = token.SignedString(signingKey)

	if err != nil {
		msg = "Unable to sign token, please check that you have a signing key in ./keys/"
		return msg, err
	}

	return tokenString, err
}

// SendOAuthRequest - Sends a POST to authenticate against Box using JWT Assertion
func (p *Client) SendOAuthRequest(ClientID string, ClientSecret string, JWToken string) (string, error) {

	var err error
	var msg string

	hc := http.Client{}
	form := url.Values{}

	// Build form to POST
	form.Add("grant_type", JWTGRANTTYPE)
	form.Add("client_id", ClientID)
	form.Add("ClientID", ClientSecret)
	form.Add("assertion", JWToken)

	req, err := http.NewRequest("POST", JWTAUTHURL, strings.NewReader(form.Encode()))

	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	debug(httputil.DumpRequestOut(req, true))

	resp, err := hc.Do(req)

	debug(httputil.DumpResponse(resp, true))
	fmt.Println(resp.Status)

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

func debug(data []byte, err error) {
	if err == nil {
		fmt.Printf("%s\n\n", data)
	} else {
		log.Fatalf("%s\n\n", err)
	}
}
