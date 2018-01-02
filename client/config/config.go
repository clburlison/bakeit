package config

var (
	// ChefClientURL map - Accepts a full URL to a chef-client. Can be blank to download from chef.io.
	ChefClientURL = map[string]string{
		"darwin":  "",
		"windows": "",
		"linux":   "",
	}

	// ChefClientVersion string - Accepts "latest" or a specific version IE - 13.6.4
	ChefClientVersion = "13.6.0"

	// ChefClientPreRelease string - Download pre-release chef client versions. String of false or true.
	ChefClientPreRelease = "false"

	// Force bool - Remove old chef files before running
	Force = false

	// Verbose bool - set standard output verbosity
	Verbose bool

	// UserShortName string - When set bakeit will check for the user and bail if they are the current user.
	// Useful if you use chef to manage a specific service account.
	UserShortName = "admin"
)
