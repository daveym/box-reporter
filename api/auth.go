package api

import (
	"fmt"

	"github.com/daveym/box-reporter-go/boxapi"
)

// Authenticate against box API. Interface used to allow mock to be passed in.
func Authenticate(bc box.API) string {

	var msg, JWToken, AccessToken string
	var err error

	JWToken, err = bc.CreateJWTAssertion(bc.GetPublicKeyID(), bc.GetClientID(), bc.GetClaimSub())

	if err != nil {
		return msg
	}

	AccessToken, err = bc.SendOAuthRequest(bc.GetClientID(), bc.GetClientSecret(), JWToken)

	fmt.Println(AccessToken)
	return msg
}
