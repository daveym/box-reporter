package box

// oAuthResponse holds decoded JSON response from Box
type oAuthResponse struct {
	AccessToken  string   `json:"access_token"`
	Expires      int      `json:"expires_in"`
	RestrictedTo []string `json:"restricted_to"`
	TokenType    string   `json:"token_type"`
}

// AppUser holds decoded JSON response from Box for an AppUser
type AppUserResponse struct {
	Usertype      string `json:"type"`
	ID            string `json:"id"`
	Name          string `json:"name"`
	Login         string `json:"login"`
	CreatedAt     string `json:"created_at"`
	ModifiedAt    string `json:"modified_at"`
	Language      string `json:"language"`
	Timezone      string `json:"timezone"`
	SpaceAmount   string `json:"space_amount"`
	SpaceUsed     string `json:"space_used"`
	MaxUploadSize string `json:"max_upload_size"`
	Status        string `json:"status"`
	JobTitle      string `json:"job_title"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	AvatarURL     string `json:"avatar_url"`
}
