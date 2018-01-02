package config

var (
	// Verbose - set standard output verbosity
	Verbose bool

	// ChefClientURL map - Accepts a full URL to a chef-client. Can be blank to download from chef.io.
	ChefClientURL = map[string]string{
		"darwin": "",
		// "darwin": "https://packages.chef.io/files/stable/" +
		// 	"chef/13.6.4/mac_os_x/10.13/chef-13.6.4-1.dmg",
		"windows": "",
		"linux":   "",
	}

	// ChefClientVersion string - Accepts "latest" or a specific version IE - 13.6.4
	ChefClientVersion = "13.6.0"

	// ChefClientPreRelease string - Download pre-release chef client versions. String of false or true.
	ChefClientPreRelease = "false"
)
