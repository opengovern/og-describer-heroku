package heroku

import (
	"context"
	"github.com/opengovern/og-describer-heroku/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableHerokuConfigVars(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "heroku_config_vars",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListConfigVars,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("app_name"),
			Hydrate:    opengovernance.GetConfigVars,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "app_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.AppName"), Description: "The name for the app."},
			{Name: "variables", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Variables"), Description: "A map of configuration variables for the app."},
		}),
	}
}
