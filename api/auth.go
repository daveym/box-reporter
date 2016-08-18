package api

import (
	"fmt"

	box "github.com/daveym/box-reporter-go/boxapi"
)

// Authenticate against box API. Interface used to allow mock to be passed in.
func Authenticate(bc box.API) string {

	var msg, JWToken string
	var EnterpriseAccessToken, UserAccessToken string
	var AppUser box.AppUserResponse
	var err error

	// Construct the initial JWT Assertion - https://docs.box.com/v2.0/docs/app-auth
	JWToken, err = bc.CreateJWTAssertion(bc.GetPublicKeyID(), bc.GetClientID(), bc.GetClaimSub(), &AppUser)
	if err != nil {
		return msg
	}

	// Submit an oAuth 2.0 request to obtain an access token
	EnterpriseAccessToken, err = bc.SendOAuthRequest(bc.GetClientID(), bc.GetClientSecret(), JWToken)
	if err != nil {
		return msg
	}

	// Using the access token, create an AppUser
	AppUser, err = bc.CreateAppUser(EnterpriseAccessToken)
	if err != nil {
		return msg
	}

	// for the app user, get a user access token
	UserAccessToken, err = bc.CreateJWTAssertion(bc.GetPublicKeyID(), bc.GetClientID(), bc.GetClaimSub(), &AppUser)
	if err != nil {
		return msg
	}

	fmt.Println("User Access Token: " + UserAccessToken)

	return msg
}
