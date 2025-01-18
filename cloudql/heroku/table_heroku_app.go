package heroku

import (
	"context"
	"github.com/opengovern/og-describer-heroku/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableHerokuApp(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "heroku_app",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListApp,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetApp,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "acm", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.ACM"), Description: "ACM status of this app."},
			{Name: "archived_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.ArchivedAt"), Description: "When the app was archived."},
			{Name: "build_stack", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.BuildStack"), Description: "Build stack details of the app."},
			{Name: "buildpack_provided_desc", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.BuildpackProvidedDesc"), Description: "Description from buildpack of the app."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.CreatedAt"), Description: "When the app was created."},
			{Name: "generation", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Generation"), Description: "Generation details of the app."},
			{Name: "git_url", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.GitURL"), Description: "Git repository URL of the app."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ID"), Description: "The unique identifier of the app."},
			{Name: "internal_routing", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.InternalRouting"), Description: "Whether the app has internal routing enabled."},
			{Name: "maintenance", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.Maintenance"), Description: "Maintenance status of the app."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Name"), Description: "The name of the app."},
			{Name: "organization", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Organization"), Description: "Organization associated with the app."},
			{Name: "owner", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Owner"), Description: "Owner details of the app."},
			{Name: "region", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Region"), Description: "Region details of the app."},
			{Name: "released_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.ReleasedAt"), Description: "When the app was released."},
			{Name: "repo_size", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.RepoSize"), Description: "Git repository size of the app in bytes."},
			{Name: "slug_size", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.SlugSize"), Description: "Slug size of the app in bytes."},
			{Name: "space", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Space"), Description: "Space details associated with the app."},
			{Name: "stack", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Stack"), Description: "Stack details of the app."},
			{Name: "team", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Team"), Description: "Team associated with the app."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.UpdatedAt"), Description: "When the app was last updated."},
			{Name: "web_url", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.WebURL"), Description: "Web URL of the app."},
		}),
	}
}
