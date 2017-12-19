// Package client provides resources for configuring a chef client
package client

import (
	"bytes"
	"text/template"
)

type Settings struct {
	LogLevel             string
	LogLocation          string
	ValidationClientName string
	ChefServerURL        string
	NodeName             string
}

// https://golang.org/pkg/text/template/
// TODO: Currently an empty value is passed a newline is created
var client = `# https://docs.chef.io/config_rb_client.html
{{if .LogLevel}}log_level              {{.LogLevel}}{{end}}
{{if .LogLocation}}log_location           {{.LogLocation}}{{end}}
validation_client_name '{{.ValidationClientName}}'
validation_key         File.expand_path('/etc/chef/validation.pem')
chef_server_url        '{{.ChefServerURL}}'
json_attribs           '/etc/chef/run-list.json'
ssl_verify_mode        :verify_peer
local_key_generation   true
rest_timeout           30
http_retry_count       3
no_lazy_load           false

whitelist = []
automatic_attribute_whitelist whitelist
default_attribute_whitelist []
normal_attribute_whitelist []
override_attribute_whitelist []

ohai.disabled_plugins = [
    :Passwd
]
ohai.plugin_path += [
  '/etc/chef/ohai_plugins'
]

{{if .NodeName}}node_name "{{.NodeName}}"{{- end}}`

// Config is the formated client.rb file
func Config() string {
	var out bytes.Buffer
	settings := Settings{":info",
		"STDOUT",
		"corp-validator",
		"https://chef.example.com/organizations/MyOrg",
		"AAXXXYYYZZZ"}
	tmpl, err := template.New("client.rb").Parse(client)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(&out, settings)
	if err != nil {
		panic(err)
	}
	// return out.String()
	return "testing\n"
}
