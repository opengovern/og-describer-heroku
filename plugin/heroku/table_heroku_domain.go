package heroku

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-heroku/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableHerokuDomain(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "heroku_domain",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDomain,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetDomain,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "acm_status", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ACMStatus"), Description: "The ACM status of the domain."},
			{Name: "acm_status_reason", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ACMStatusReason"), Description: "Reason for the ACM status."},
			{Name: "app_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.AppID"), Description: "The unique identifier of the app."},
			{Name: "app_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.AppName"), Description: "The name of the app."},
			{Name: "cname", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.CName"), Description: "Canonical name record."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.CreatedAt"), Description: "When the domain was created."},
			{Name: "hostname", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Hostname"), Description: "Full hostname of the domain."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ID"), Description: "The unique identifier of the domain."},
			{Name: "kind", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Kind"), Description: "Type of domain (heroku or custom)."},
			{Name: "sni_endpoint", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.SNIEndpoint"), Description: "SNI endpoint associated with the domain."},
			{Name: "status", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Status"), Description: "The current status of the domain."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.UpdatedAt"), Description: "When the domain was last updated."},
		}),
	}
}
