package api

import "github.com/daveym/box-reporter-go/boxapi"

// Authenticate against box API. Interface used to allow mock to be passed in.
func Authenticate(bc box.API) string {

	var msg, JWToken string
	var err error

	JWToken, err = bc.CreateJWTAssertion(bc.GetPublicKeyID(), bc.GetClientID(), bc.GetClaimSub())

	if err != nil {
		return msg
	}

	msg, err = bc.SendOAuthRequest(bc.GetClientID(), "2vnt7KVZeVnJCQGoRuuPHOVBD370vGUj", JWToken)

	return msg
}
