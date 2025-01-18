package main

import (
	"github.com/opengovern/og-describer-heroku/cloudql/heroku"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: heroku.Plugin})
}
