package mobile_test

import (
	"testing"

	. "github.com/textileio/textile-go/mobile"
	"encoding/json"
)

var textile *Node
var hash string

func TestNewTextile(t *testing.T) {
	textile = NewTextile("testdata/.ipfs", "")
}

func TestNode_Start(t *testing.T) {
	err := textile.Start()
	if err != nil {
		t.Errorf("start mobile node failed: %s", err)
	}
}

func TestNode_PinPhoto(t *testing.T) {
	hash, err := textile.PinPhoto("testdata/test.jpg")
	if err != nil {
		t.Errorf("pin photo failed: %s", err)
		return
	}
	if len(hash) == 0 {
		t.Errorf("pin photo got bad hash: %s", hash)
	}
}

func TestNode_GetPhotos(t *testing.T) {
	res, err := textile.GetPhotos("", -1)
	if err != nil {
		t.Errorf("get photos failed: %s", err)
		return
	}
	list := PhotoList{}
	json.Unmarshal([]byte(res), &list)
	if len(list.Hashes) == 0 {
		t.Errorf("get photos bad result")
	}
	hash = list.Hashes[0]
}

func TestNode_GetPhotoBase64String(t *testing.T) {
	res, err := textile.GetPhotoBase64String(hash + "/thumb.jpg")
	if err != nil {
		t.Errorf("get photo base64 string failed: %s", err)
		return
	}
	if len(res) == 0 {
		t.Errorf("get photo base64 string bad result")
	}
}

func TestNode_Stop(t *testing.T) {
	err := textile.Stop()
	if err != nil {
		t.Errorf("stop mobile node failed: %s", err)
	}
}