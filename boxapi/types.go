package box

// oAuthResponse holds decoded JSON response from Box
type oAuthResponse struct {
	AccessToken  string   `json:"access_token"`
	Expires      int      `json:"expires_in"`
	RestrictedTo []string `json:"restricted_to"`
	TokenType    string   `json:"token_type"`
}

// AppUserRequest holds  JSON request to create an AppUser
type AppUserRequest struct {
	IsPlatformAccess bool   `json:"is_platform_access_only,bool"`
	Name             string `json:"name"`
}

// AppUser holds the newly created AppUser
type AppUserResponse struct {
	Usertype      string  `json:"type"`
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Login         string  `json:"login"`
	CreatedAt     string  `json:"created_at"`
	ModifiedAt    string  `json:"modified_at"`
	Language      string  `json:"language"`
	Timezone      string  `json:"timezone"`
	SpaceAmount   float32 `json:"space_amount"`
	SpaceUsed     float32 `json:"space_used"`
	MaxUploadSize float32 `json:"max_upload_size"`
	Status        string  `json:"status"`
	JobTitle      string  `json:"job_title"`
	Phone         string  `json:"phone"`
	Address       string  `json:"address"`
	AvatarURL     string  `json:"avatar_url"`
}
