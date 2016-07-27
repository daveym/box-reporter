package api

import "github.com/daveym/box-reporter/box"

// Authenticate against box API. Interface used to allow mock to be passed in.
func Authenticate(bc box.API) string {

	msg := ""

	if len(bc.GetClientID()) == 0 {
		msg = box.CLIENTIDNOTFOUNDEn
		return msg
	}

	err := bc.UserAuthorise(box.UserAuthorisationURL, bc.GetClientID(), box.RedirectURI)
	if err != nil {
		msg = box.ERRORAPPROVINGLINTEn
		return msg
	}

	return msg
}
