package core_test

import (
	"crypto/rand"
	"io/ioutil"
	"os"
	"testing"

	mh "gx/ipfs/QmPnFwZ2JXKnXgMw8CdBPxn7FWh6LLdjUjxV1fKHuJnkr8/go-multihash"
	libp2pc "gx/ipfs/QmPvyPwuCgJ7pDmrKDxRtsScJgBaM5h4EpRL2qQJsmXf4n/go-libp2p-crypto"
	logger "gx/ipfs/QmcaSwFc5RBg8yCq54QURwEU4nwjfCpjbpmaAm4VbdGLKv/go-logging"

	"github.com/segmentio/ksuid"
	. "github.com/textileio/textile-go/core"
	"github.com/textileio/textile-go/keypair"
	"github.com/textileio/textile-go/mill"
	"github.com/textileio/textile-go/repo"
	"github.com/textileio/textile-go/schema/textile"
)

var repoPath = "testdata/.textile"
var node *Textile

var schemaHash mh.Multihash

func TestInitRepo(t *testing.T) {
	os.RemoveAll(repoPath)
	accnt := keypair.Random()
	if err := InitRepo(InitConfig{
		Account:  accnt,
		RepoPath: repoPath,
		LogLevel: logger.ERROR,
	}); err != nil {
		t.Errorf("init node failed: %s", err)
	}
}

func TestNewTextile(t *testing.T) {
	var err error
	node, err = NewTextile(RunConfig{
		RepoPath: repoPath,
	})
	if err != nil {
		t.Errorf("create node failed: %s", err)
	}
}

func TestTextile_Start(t *testing.T) {
	if err := node.Start(); err != nil {
		t.Errorf("start node failed: %s", err)
	}
	<-node.OnlineCh()
}

func TestTextile_Started(t *testing.T) {
	if !node.Started() {
		t.Errorf("should report started")
	}
}

func TestTextile_Online(t *testing.T) {
	if !node.Online() {
		t.Errorf("should report online")
	}
}

func TestTextile_GetMedia(t *testing.T) {
	f, err := os.Open("../mill/testdata/image.jpeg")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	media, err := node.GetMedia(f, &mill.ImageResize{})
	if err != nil {
		t.Fatal(err)
	}
	if media != "image/jpeg" {
		t.Errorf("wrong media type: %s", media)
	}
}

func TestTextile_AddSchema(t *testing.T) {
	file, err := node.AddSchema(textile.Photos, "test")
	if err != nil {
		t.Fatal(err)
	}
	schemaHash, err = mh.FromB58String(file.Hash)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTextile_AddThread(t *testing.T) {
	sk, _, err := libp2pc.GenerateEd25519Key(rand.Reader)
	if err != nil {
		t.Error(err)
	}
	config := AddThreadConfig{
		Key:       ksuid.New().String(),
		Name:      "test",
		Schema:    schemaHash,
		Initiator: node.Account().Address(),
		Type:      repo.OpenThread,
		Join:      true,
	}
	thrd, err := node.AddThread(sk, config)
	if err != nil {
		t.Errorf("add thread failed: %s", err)
		return
	}
	if thrd == nil {
		t.Error("add thread didn't return thread")
	}
}

func TestTextile_AddFile(t *testing.T) {
	f, err := os.Open("../mill/testdata/image.jpeg")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	m := &mill.ImageResize{
		Opts: mill.ImageResizeOpts{
			Width:   "200",
			Quality: "75",
		},
	}
	conf := AddFileConfig{
		Input: data,
		Name:  "image.jpeg",
		Media: "image/jpeg",
	}

	file, err := node.AddFile(m, conf)
	if err != nil {
		t.Errorf("add file failed: %s", err)
		return
	}

	if file.Mill != "/image/resize" {
		t.Error("wrong mill")
	}
	if file.Checksum != "6Em1qNFcc1B3mSToS2aVyubDi7R7w7QedY84eNbiWBW3" {
		t.Error("wrong checksum")
	}
}

func TestTextile_Stop(t *testing.T) {
	if err := node.Stop(); err != nil {
		t.Errorf("stop node failed: %s", err)
	}
}

func TestTextile_StartedAgain(t *testing.T) {
	if node.Started() {
		t.Errorf("should report stopped")
	}
}

func TestTextile_OnlineAgain(t *testing.T) {
	if node.Online() {
		t.Errorf("should report offline")
	}
}

func TestTextile_Teardown(t *testing.T) {
	node = nil
	os.RemoveAll(repoPath)
}
