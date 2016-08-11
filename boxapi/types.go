package box

// oAuthResponse holds decoded JSON response from Box
type oAuthResponse struct {
	AccessToken  string   `json:"access_token"`
	Expires      int      `json:"expires_in"`
	RestrictedTo []string `json:"restricted_to"`
	TokenType    string   `json:"token_type"`
}
