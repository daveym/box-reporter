package api

import (
	"testing"

	"github.com/daveym/box-reporter/box"
)

var err error

func TestAuthNoConsumerKey(t *testing.T) {

	t.Log("Executing: TestAuthNoConsumerKey")

	mc := &box.MockClient{}
	mc.SetConsumerKey("")

	expectedmsg := box.CONSUMERKEYNOTFOUNDEn
	actualmsg := Authenticate(mc)

	if actualmsg != expectedmsg {
		t.Log("Expected: " + expectedmsg)
		t.Log("Actual: " + actualmsg)
		t.Fatal("TestValidConsumerKey failed")
	}
}

func TestAuthInvalidConsumerKey(t *testing.T) {

	t.Log("Executing: TestAuthInvalidConsumerKey")

	mc := &box.MockClient{}
	mc.SetConsumerKey("INVALIDKEY")

	expectedmsg := box.CONSUMERKEYNOTVALIDEn
	actualmsg := Authenticate(mc)

	if actualmsg != expectedmsg {
		t.Log("Expected: " + expectedmsg)
		t.Log("Actual: " + actualmsg)
		t.Fatal("TestAuthInvalidConsumerKey failed")
	}
}

func TestBrowserAuthFail(t *testing.T) {

	t.Log("Executing: TestBrowserAuthFail")

	mc := &box.MockClient{}
	mc.SetConsumerKey("INVALIDBROWSER")

	expectedmsg := box.ERRORAPPROVINGLINTEn
	actualmsg := Authenticate(mc)

	if actualmsg != expectedmsg {
		t.Log("Expected: " + expectedmsg)
		t.Log("Actual: " + actualmsg)
		t.Fatal("TestBrowserAuthFail failed")
	}
}

func TestAuthGetAccessTokenFail(t *testing.T) {

	t.Log("Executing: TestAuthGetAccessTokenFail")

	mc := &box.MockClient{}
	mc.SetConsumerKey("FAIL")
	mc.SetAccessToken("FAIL")

	expectedmsg := box.ERRORAUTHen
	actualmsg := Authenticate(mc)

	if actualmsg != expectedmsg {
		t.Log("Expected: " + expectedmsg)
		t.Log("Actual: " + actualmsg)
		t.Fatal("TestAuthGetAccessTokenFail failed")
	}
}

func TestAuthGetAccessTokenSuccess(t *testing.T) {

	t.Log("Executing: TestAuthGetAccessTokenSuccess")

	mc := &box.MockClient{}
	mc.SetConsumerKey("45678")
	mc.SetAccessToken("SUCCESS")

	expectedmsg := box.AUTHSUCCESSen
	actualmsg := Authenticate(mc)

	if actualmsg != expectedmsg {
		t.Log("Expected: " + expectedmsg)
		t.Log("Actual: " + actualmsg)
		t.Fatal("TestAuthGetAccessTokenSuccess failed")
	}
}
