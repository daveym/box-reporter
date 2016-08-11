package box

// APPUSERNAME - used to name the user when creating an App User
const APPUSERNAME string = "box-reporter"

// JWTAUTHURL - URL for oAUTH for Box
const JWTAUTHURL string = "https://api.box.com/oauth2/token"

// JWTUSERURL - URL for creating an App User
const JWTUSERURL string = "https://api.box.com/2.0/users"

// JWTGRANTTYPE - mandatory parameter for box oAuth
const JWTGRANTTYPE string = "urn:ietf:params:oauth:grant-type:jwt-bearer"
