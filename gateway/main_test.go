package gateway_test

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/textileio/go-textile/core"
	. "github.com/textileio/go-textile/gateway"
	"github.com/textileio/go-textile/keypair"
)

var repoPath = "testdata/.textile"

func TestGateway_Creation(t *testing.T) {
	os.RemoveAll(repoPath)

	err := core.InitRepo(core.InitConfig{
		Account:     keypair.Random(),
		RepoPath:    repoPath,
		GatewayAddr: fmt.Sprintf("127.0.0.1:%s", core.GetRandomPort()),
	})
	if err != nil {
		t.Errorf("init node failed: %s", err)
		return
	}

	node, err := core.NewTextile(core.RunConfig{
		RepoPath: repoPath,
	})
	if err != nil {
		t.Errorf("create node failed: %s", err)
		return
	}

	Host = &Gateway{Node: node}
	Host.Start(node.Config().Addresses.Gateway)
}

func TestGateway_Addr(t *testing.T) {
	if len(Host.Addr()) == 0 {
		t.Error("get gateway address failed")
	}
}

func TestGateway_Cors(t *testing.T) {
	// Prepare the URL
	addr := "http://" + Host.Addr() + "/health"

	// Make the OPTIONS request, which initialises CORS
	req, err := http.NewRequest("OPTIONS", addr, nil)
	if err != nil {
		t.Error(err)
		return
	}

	// Simulate an external origin
	req.Header.Add("Origin", "http://example.com/")
	req.Header.Add("Access-Control-Request-Method", "POST")

	// Send the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
		return
	}

	// Close everything
	defer res.Body.Close()

	// Check its result
	if res.StatusCode != http.StatusOK {
		t.Error(fmt.Errorf("%s returned unexpected status: %d", addr, res.StatusCode))
		return
	}
}

func TestGateway_Stop(t *testing.T) {
	err := Host.Stop()
	if err != nil {
		t.Errorf("stop gateway failed: %s", err)
	}
	os.RemoveAll(repoPath)
}
