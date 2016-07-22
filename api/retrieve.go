package api

import (
	"fmt"

	"github.com/daveym/box-reporter/box"
)

// Retrieve against box API. Interface used to allow mock to be passed in.
func Retrieve(pc box.API, searchVal string, domainVal string, tagVal string, countVal int) string {

	msg := ""

	if len(pc.GetConsumerKey()) == 0 {
		msg = box.CONSUMERKEYNOTVALIDEn
		return msg
	}

	if countVal == 0 {
		msg = box.COUNTGREATERTHANZEROen
		return msg
	}

	if len(searchVal) == 0 && len(domainVal) == 0 && len(tagVal) == 0 {
		msg = box.SPECIFYSEARCHen
		return msg
	}

	itemreq := box.RetrieveRequest{}
	itemreq.ConsumerKey = pc.GetConsumerKey()
	itemreq.AccessToken = pc.GetAccessToken()

	if searchVal != "all" {
		itemreq.Search = searchVal
	}

	itemreq.Domain = domainVal
	itemreq.Tag = tagVal
	itemreq.Count = countVal
	itemreq.Sort = box.SortNewest

	itemresp := &box.RetrieveResponse{}

	err := pc.Retrieve(itemreq, itemresp)

	if err != nil {
		msg = box.ERRORRETRIEVINGen + err.Error()
		return msg
	}

	items := itemresp.List

	if len(items) == 0 {
		msg = box.NOMATCHINGVALUESen
		return msg
	}

	for _, item := range items {
		msg = msg + fmt.Sprintf("%v,%v,%v\n", item.ItemID, item.GivenTitle, item.GivenURL)
	}

	return msg
}
