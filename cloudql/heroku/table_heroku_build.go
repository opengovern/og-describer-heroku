package heroku

import (
	"context"
	"github.com/opengovern/og-describer-heroku/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableHerokuBuild(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "heroku_build",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListBuild,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetBuild,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "app_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.AppID"), Description: "The unique identifier for the app."},
			{Name: "buildpacks", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Buildpacks"), Description: "List of buildpacks used in the build."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.CreatedAt"), Description: "When the build was created."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ID"), Description: "The unique identifier of the build."},
			{Name: "output_stream_url", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.OutputStreamURL"), Description: "URL to stream build process output."},
			{Name: "release", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Release"), Description: "Release object associated with the build."},
			{Name: "slug", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Slug"), Description: "Slug created by the build."},
			{Name: "source_blob", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.SourceBlob"), Description: "Source blob details for the build."},
			{Name: "stack", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Stack"), Description: "Stack used for the build."},
			{Name: "status", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Status"), Description: "Build status."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.UpdatedAt"), Description: "When the build was last updated."},
			{Name: "user", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.User"), Description: "User details associated with the build."},
		}),
	}
}
