package box

// BoxHome - Box base url
const BoxHome string = "https://api.box.com"

// CfgFile - Name of viper config file. Lint will store its values in the working directory.
const CfgFile string = "box-reporter.yaml"

// URLS
const (
	// JWTAuthenticationURL - URL for authenticating an oAuth 2.0 JWT. See https://docs.box.com/v2.0/docs/app-auth
	JWTAUTHURL string = "https://api.box.com/oauth2/token"
)
