package api

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/daveym/box-reporter/box"
)

// Authenticate against box API. Interface used to allow mock to be passed in.
func Authenticate(bc box.API) string {

	msg := ""

	if len(bc.GetConsumerKey()) == 0 {
		msg = box.CONSUMERKEYNOTFOUNDEn
		return msg
	}

	AuthNResp := &box.AuthenticationResponse{}
	err := bc.Authenticate(bc.GetConsumerKey(), AuthNResp)

	if err != nil {
		msg = box.CONSUMERKEYNOTVALIDEn
		return msg
	}

	err = bc.UserAuthorise(box.UserAuthorisationURL, AuthNResp.Code, box.RedirectURI)
	if err != nil {
		msg = box.ERRORAPPROVINGLINTEn
		return msg
	}

	fmt.Print(box.ACKNOWLEDGEAUTHen)
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	AuthRResp := &box.AuthorisationResponse{}
	err = bc.RetrieveAccessToken(bc.GetConsumerKey(), AuthNResp.Code, AuthRResp)

	if err != nil {
		msg = box.ERRORAUTHen
		return msg
	}

	cfgval := fmt.Sprintf("ConsumerKey: %v\nAccessToken: %v\nUsername: %v",
		bc.GetConsumerKey(), AuthRResp.AccessToken, AuthRResp.Username)

	err = ioutil.WriteFile(box.CfgFile, []byte(cfgval), 0644)

	if err != nil {
		msg = box.ERRORSAVINGCONSUMERKEYen
		return msg
	}

	msg = box.AUTHSUCCESSen

	return msg
}
