package api

import "github.com/daveym/box-reporter-go/boxapi"

// Authenticate against box API. Interface used to allow mock to be passed in.
func Authenticate(bc box.API) string {

	var err error

	err = bc.CreateJWTAssertion(bc.GetPublicKeyID(), bc.GetClientID())

	return err.Error()
}
