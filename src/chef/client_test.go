package chef

import (
	"io/ioutil"
	"testing"
)

func TestClient(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/client_config")
	if err != nil {
		t.Errorf("'client_config' is missing!\n")
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
		"/etc/chef/ohai_plugins",
		[]string{":Passwd"},
		"AAXXXYYYZZZ"}
	config, err := BuildConfig(settings)
	// fmt.Printf(config)
	if err != nil {
		t.Errorf("Error getting a config: %s\n", err)
	}
	if config != testConfig {
		t.Errorf("Config did not match valid config\n")
	}
}
