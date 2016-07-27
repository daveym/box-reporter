package box

// BoxHome - Box base url
const BoxHome string = "https://api.box.com"

// CfgFile - Name of viper config file. Lint will store its values in the working directory.
const CfgFile string = "box-reporter.yaml"

// URLS
const (
	// AuthorisationURL - API address to Authorise and recieve an access Token
	AuthorisationURL string = "https://account.box.com/api/oauth2/authorize"
	//UserAuthorisationURL - Address that a user must enter into their browser to Authorise access to Box
	UserAuthorisationURL string = "https://account.box.com/api/oauth2/authorize?"
	//RedirectURI - Link back location after Authorisation has been granted
	RedirectURI string = "https://github.com/daveym/box-reporter/blob/master/AUTHCOMPLETE.md"
)
