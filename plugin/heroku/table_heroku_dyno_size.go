package heroku

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-heroku/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableHerokuDynoSize(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "heroku_dyno_size",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDynoSize,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetDynoSize,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "architecture", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Architecture"), Description: "CPU architecture of this dyno size."},
			{Name: "compute", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.Compute"), Description: "Minimum vCPUs."},
			{Name: "cost", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Cost"), Description: "Price information for this dyno size."},
			{Name: "dedicated", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.Dedicated"), Description: "Whether this dyno will be dedicated to one user."},
			{Name: "generation", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Generation"), Description: "Generation of the Heroku platform for this dyno size."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ID"), Description: "The unique identifier of this dyno size."},
			{Name: "memory", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("Description.Memory"), Description: "Amount of RAM in GB."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Name"), Description: "The name of this dyno size."},
			{Name: "precise_dyno_units", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("Description.PreciseDynoUnits"), Description: "Unit of consumption for Heroku Enterprise customers."},
			{Name: "private_space_only", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.PrivateSpaceOnly"), Description: "Whether this dyno can only be provisioned in a private space."},
		}),
	}
}
