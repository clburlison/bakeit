# https://docs.chef.io/config_rb_client.html
log_level              :info
log_location           STDOUT
validation_client_name 'corp-validator'
validation_key         File.expand_path('/etc/chef/validation.pem')
chef_server_url        'https://chef.example.com/organizations/MyOrg'
json_attribs           '/etc/chef/run-list.json'
ssl_verify_mode        :verify_peer
local_key_generation   true
rest_timeout           30
http_retry_count       3


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

node_name "AAXXXYYYZZZ"
