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
	ValidationKey        string
	ChefServerURL        string
	JSONAttribs          string
	SSLVerifyMode        string
	LocalKeyGeneration   bool
	RestTimeout          int
	HTTPRetryCount       int
	NoLazyLoad           bool
	OhaiDirectory        string
	OhaiDisabledPlugins  []string
	NodeName             string
}

// https://golang.org/pkg/text/template/
// TODO: If an empty value is passed a newline is created
// TODO: OhaiDisabledPlugins - we should add a comma + new line if len > 1
var client = `# https://docs.chef.io/config_rb_client.html
{{if .LogLevel}}log_level              {{.LogLevel}}{{end}}
{{if .LogLocation}}log_location           {{.LogLocation}}{{end}}
{{if .ValidationClientName}}validation_client_name '{{.ValidationClientName}}'{{end}}
{{if .ValidationKey}}validation_key         File.expand_path('{{.ValidationKey}}'){{end}}
{{if .ChefServerURL}}chef_server_url        '{{.ChefServerURL}}'{{end}}
{{if .JSONAttribs}}json_attribs           '{{.JSONAttribs}}'{{end}}
{{if .SSLVerifyMode}}ssl_verify_mode        {{.SSLVerifyMode}}{{end}}
{{if .LocalKeyGeneration}}local_key_generation   {{.LocalKeyGeneration}}{{end}}
{{if .RestTimeout}}rest_timeout           {{.RestTimeout}}{{end}}
{{if .HTTPRetryCount}}http_retry_count       {{.HTTPRetryCount}}{{end}}
{{if .NoLazyLoad}}no_lazy_load           {{.NoLazyLoad}}{{end}}

automatic_attribute_whitelist []
default_attribute_whitelist []
normal_attribute_whitelist []
override_attribute_whitelist []

{{if .OhaiDirectory}}ohai.directory = '{{.OhaiDirectory}}'{{end}}
{{- if .OhaiDisabledPlugins}}
{{ $disabled_plugins := .OhaiDisabledPlugins }}
ohai.disabled_plugins = [
{{ range $disabled_plugins }}
     {{- . -}}
{{ end }}
]
{{end}}

{{if .NodeName}}node_name "{{.NodeName}}"{{- end}}
`

// Config is the formated client.rb file
func Config(set Settings) (config string, err error) {
	var out bytes.Buffer
	tmpl, err := template.New("client.rb").Parse(client)
	if err != nil {
		return "", err
	}
	err = tmpl.Execute(&out, set)
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
