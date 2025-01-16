package heroku

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-heroku/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableHerokuDyno(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "heroku_dyno",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDyno,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetDyno,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "app_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.AppID"), Description: "The unique identifier of the app."},
			{Name: "app_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.AppName"), Description: "The name of the app."},
			{Name: "attach_url", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.AttachURL"), Description: "A URL to stream output from attached processes."},
			{Name: "command", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Command"), Description: "Command used to start this process."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.CreatedAt"), Description: "When the dyno was created."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ID"), Description: "The unique identifier of the dyno."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Name"), Description: "The name of this process on the dyno."},
			{Name: "release", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Release"), Description: "Release associated with the dyno."},
			{Name: "size", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Size"), Description: "Size of the dyno."},
			{Name: "state", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.State"), Description: "Current status of the dyno (e.g., crashed, down, idle, starting, up)."},
			{Name: "type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Type"), Description: "Type of process."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.UpdatedAt"), Description: "When the dyno was last updated."},
		}),
	}
}
