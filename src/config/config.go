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
		"windows": `{"run_list": ["role[cpe_base]"]}`,
		"linux":   "",
	}

	// ChefClientCertPath map - Client cert path
	ChefClientCertPath = map[string]string{
		"darwin":  "/etc/chef/client.pem",
		"windows": "",
		"linux":   "",
	}

	// ChefClientValidationKey map - Client cert path
	ChefClientValidationKey = map[string]string{
		"darwin":  "/etc/chef/validation.pem",
		"windows": "C:\\chef\\validation.pem",
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

	// ChefClientJSONAttribs map - Path to json runlist file
	ChefClientJSONAttribs = map[string]string{
		"darwin":  "/etc/chef/run-list.json",
		"windows": "C:\\chef\\first-boot.json",
		"linux":   "",
	}

	// ChefClientExecPath map - Path to the chef-client executable
	ChefClientExecPath = map[string]string{
		"darwin":  "/usr/local/bin/chef-client",
		"windows": "C:\\opscode\\chef\\bin\\chef-client",
		"linux":   "",
	}

	ChefClientLogLevel             = ":info"
	ChefClientLogLocation          = "STDOUT"
	ChefClientValidationClientName = "corp-validator"
	ChefClientChefServerURL        = "https://chef.example.com/organizations/MyOrg"
	ChefClientSSLVerifyMode        = ":verify_peer"
	ChefClientLocalKeyGeneration   = true
	ChefClientRestTimeout          = 30
	ChefClientHTTPRetryCount       = 3
	ChefClientNoLazyLoad           = false

	// FirstRunLogFile map - Path to the logfile for first chef run
	FirstRunLogFile = map[string]string{
		"darwin":  "/Library/Chef/Logs/first_chef_run.log",
		"windows": "C:\\chef\\logs\\first_chef_run.txt",
		"linux":   "",
	}

	// PreFlight - command and arguments to run before bootstrap

	// PreFlightCommand map - Example: "msiexec.exe"
	PreFlightCommand = map[string]string{
		"darwin":  "",
		"windows": "",
		"linux":   "",
	}

	// PreFlightArguments map - Example: {"/qn", "/i", "c:\\installer.msi"}
	PreFlightArguments = map[string][]string{
		"darwin":  {},
		"windows": {},
		"linux":   {},
	}

	// PreFlightRequired map - Set to true to quit if the preflight fails.
	PreFlightRequired = map[string]bool{
		"darwin":  false,
		"windows": false,
		"linux":   false,
	}

	// PostFlight - command and arguments to run before bootstrap

	// PostFlightCommand map - Example: "msiexec.exe"
	PostFlightCommand = map[string]string{
		"darwin":  "",
		"windows": "",
		"linux":   "",
	}

	// PostFlightArguments map - Example: {"/qn", "/i", "c:\\installer.msi"}
	PostFlightArguments = map[string][]string{
		"darwin":  {},
		"windows": {},
		"linux":   {},
	}

	// PostFlightRequired map - Set to true to quit if the postflight fails.
	PostFlightRequired = map[string]bool{
		"darwin":  false,
		"windows": false,
		"linux":   false,
	}

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
