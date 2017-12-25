package client

import (
	"io/ioutil"
	"testing"
)

func TestClient(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/client_config.rb")
	if err != nil {
		t.Errorf("Test client_config.rb is missing!\n")
	}
	testConfig := string(b)
	settings := Settings{":info",
		"STDOUT",
		"corp-validator",
		"/etc/chef/validation.pem",
		"https://chef.example.com/organizations/MyOrg",
		"/etc/chef/run-list.json",
		":verify_peer",
		true,
		30,
		3,
		false,
		"AAXXXYYYZZZ"}
	config, err := Config(settings)
	// fmt.Printf(config)
	if err != nil {
		t.Errorf("Error getting a config: %s\n", err)
	}
	if config != testConfig {
		t.Errorf("Config did not match valid config\n")
	}
}
