package api

import (
	"fmt"
	"testing"

	"github.com/daveym/box-reporter/box"
)

func TestRetrieveNoConsumerKey(t *testing.T) {

	t.Log("Executing: TestRetrieveNoConsumerKey")

	var searchVal string
	var domainVal string
	var tagVal string
	var countVal = 1

	mc := &box.MockClient{}
	mc.SetConsumerKey("")

	expectedmsg := box.CONSUMERKEYNOTVALIDEn
	actualmsg := Retrieve(mc, searchVal, domainVal, tagVal, countVal)

	if actualmsg != expectedmsg {
		t.Log("Expected: " + expectedmsg)
		t.Log("Actual: " + actualmsg)
		t.Fatal("TestValidConsumerKey failed")
	}
}

func TestRetrieveNoCountCriteria(t *testing.T) {

	t.Log("Executing: TestRetrieveNoCountCriteria")

	var searchVal, domainVal, tagVal string
	var countVal int

	mc := &box.MockClient{}
	mc.SetConsumerKey("45678")
	mc.SetAccessToken("SUCCESS")

	searchVal = ""
	domainVal = ""
	tagVal = ""
	countVal = 0

	expectedmsg := box.COUNTGREATERTHANZEROen
	actualmsg := Retrieve(mc, searchVal, domainVal, tagVal, countVal)

	if actualmsg != expectedmsg {
		t.Log("Expected: " + expectedmsg)
		t.Log("Actual: " + actualmsg)
		t.Fatal("TestRetrieveNoSearchCriteria failed")
	}
}

func TestRetrieveNoCriteria(t *testing.T) {

	t.Log("Executing: TestRetrieveNoCriteria")

	var searchVal, domainVal, tagVal string
	var countVal int

	mc := &box.MockClient{}
	mc.SetConsumerKey("45678")
	mc.SetAccessToken("SUCCESS")

	searchVal = ""
	domainVal = ""
	tagVal = ""
	countVal = 1

	expectedmsg := box.SPECIFYSEARCHen
	actualmsg := Retrieve(mc, searchVal, domainVal, tagVal, countVal)

	if actualmsg != expectedmsg {
		t.Log("Expected: " + expectedmsg)
		t.Log("Actual: " + actualmsg)
		t.Fatal("TestRetrieveNoCriteria failed")
	}
}

func TestRetrieveWithNoItemsFound(t *testing.T) {

	t.Log("Executing: TestRetrieveWithSearchCriteria")

	var searchVal, domainVal, tagVal string
	var countVal int

	mc := &box.MockClient{}
	mc.SetConsumerKey("45678")
	mc.SetAccessToken("SUCCESS")

	searchVal = "nothing"
	domainVal = ""
	tagVal = ""
	countVal = 1

	expectedmsg := box.NOMATCHINGVALUESen
	actualmsg := Retrieve(mc, searchVal, domainVal, tagVal, countVal)

	if actualmsg != expectedmsg {
		t.Log("Expected: " + expectedmsg)
		t.Log("Actual: " + actualmsg)
		t.Fatal("TestRetrieveWithSearchCriteria failed")
	}
}

func TestRetrieveWithSearchCriteria(t *testing.T) {

	t.Log("Executing: TestRetrieveWithSearchCriteria")

	var searchVal, domainVal, tagVal string
	var countVal int

	mc := &box.MockClient{}
	mc.SetConsumerKey("45678")
	mc.SetAccessToken("SUCCESS")

	searchVal = "docker"
	domainVal = ""
	tagVal = ""
	countVal = 1

	expectedmsg := fmt.Sprintf("%v,%v,%v\n", 11111, "Docker", "http://docker.com")
	actualmsg := Retrieve(mc, searchVal, domainVal, tagVal, countVal)

	if actualmsg != expectedmsg {
		t.Log("Expected: " + expectedmsg)
		t.Log("Actual: " + actualmsg)
		t.Fatal("TestRetrieveWithSearchCriteria failed")
	}
}
