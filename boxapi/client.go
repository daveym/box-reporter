package box

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// API - Base interface type for Box API. Allows us to mock/test.
type API interface {
	SetClientID(string)
	GetClientID() string
	Authorise(string, string, string) error
}

// Client -
type Client struct {
	_ClientID string
}

// SetClientID - SSet the ClientID
func (p *Client) SetClientID(newKey string) {
	p._ClientID = newKey
}

// GetClientID - Set the new consumer Key
func (p *Client) GetClientID() string {
	return p._ClientID
}

// Authorise -  perform a Box Server to Server Authorisation
func (p *Client) Authorise(url string, clientID string, uri string) error {

	var err error

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
