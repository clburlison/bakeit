package config

var (
	// ChefClientURL map - Accepts a full URL to a chef-client. Can be blank to download from chef.io.
	ChefClientURL = map[string]string{
		"darwin":  "",
		"windows": "",
		"linux":   "",
	}

	// ChefClientVersion string - Accepts "latest" or a specific version IE - 13.6.4
	ChefClientVersion = "latest"

	// ChefClientPreRelease string - Download pre-release chef client versions. String of false or true.
	ChefClientPreRelease = "false"

	// ChefClientRunListJSON map - Chef Run list
	ChefClientRunListJSON = map[string]string{
		"darwin":  `{"run_list": ["role[cpe_base]"]}`,
		"windows": "",
		"linux":   "",
	}

	// ChefClientOhaiDirectory map - Ohai plugin directory
	ChefClientOhaiDirectory = map[string]string{
		"darwin":  "/etc/chef/ohai_plugins",
		"windows": "",
		"linux":   "",
	}

	// ChefClientOhaiDisabledPlugins map - Plugins to disable with Ohai
	ChefClientOhaiDisabledPlugins = map[string][]string{
		"darwin":  {":Passwd"},
		"windows": {},
		"linux":   {},
	}

	ChefClientLogLevel             = ":info"
	ChefClientLogLocation          = "STDOUT"
	ChefClientValidationClientName = "corp-validator"
	ChefClientValidationKey        = "/etc/chef/validation.pem"
	ChefClientChefServerURL        = "https://chef.example.com/organizations/MyOrg"
	ChefClientJSONAttribs          = "/etc/chef/run-list.json"
	ChefClientSSLVerifyMode        = ":verify_peer"
	ChefClientLocalKeyGeneration   = true
	ChefClientRestTimeout          = 30
	ChefClientHTTPRetryCount       = 3
	ChefClientNoLazyLoad           = false

	// Force bool - Remove old chef files before running
	Force = false

	// Verbose bool - set standard output verbosity
	Verbose bool

	// UserShortName string - When set bakeit will check for the user and bail if they are the current user.
	// Useful if you use chef to manage a specific service account.
	UserShortName = "admin"
)

// ValidationPEM - The validation certificate from a chef server.
var ValidationPEM = `-----BEGIN RSA PRIVATE KEY-----
validation pem goes here
-----END RSA PRIVATE KEY-----
`

// OrgCert - The organization certificate. Required if using a self signed cert from your chef server.
// If left unmodified no cert is written.
var OrgCert = `-----BEGIN RSA PRIVATE KEY-----
org cert goes here
-----END RSA PRIVATE KEY-----
`
