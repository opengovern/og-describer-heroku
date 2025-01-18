package heroku

import (
	"context"
	"github.com/opengovern/og-describer-heroku/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableHerokuAccount(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "heroku_account",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAccount,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetAccount,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "allow_tracking", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.AllowTracking"), Description: "Whether to allow third-party web activity tracking."},
			{Name: "beta", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.Beta"), Description: "Whether allowed to utilize beta Heroku features."},
			{Name: "country_of_residence", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.CountryOfResidence"), Description: "Country where account owner resides."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.CreatedAt"), Description: "When account was created."},
			{Name: "default_organization", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.DefaultOrganization"), Description: "Team selected by default."},
			{Name: "default_team", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.DefaultTeam"), Description: "Team selected by default."},
			{Name: "delinquent_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.DelinquentAt"), Description: "When account became delinquent."},
			{Name: "email", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Email"), Description: "Unique email address of the account."},
			{Name: "federated", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.Federated"), Description: "Whether the user is federated and belongs to an Identity Provider."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ID"), Description: "The unique identifier of the account."},
			{Name: "identity_provider", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.IdentityProvider"), Description: "Identity Provider details for federated users."},
			{Name: "last_login", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.LastLogin"), Description: "When the account last authorized with Heroku."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Name"), Description: "Full name of the account owner."},
			{Name: "sms_number", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.SMSNumber"), Description: "SMS number of the account."},
			{Name: "suspended_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.SuspendedAt"), Description: "When the account was suspended."},
			{Name: "two_factor_authentication", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.TwoFactorAuthentication"), Description: "Whether two-factor authentication is enabled on the account."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.UpdatedAt"), Description: "When the account was updated."},
			{Name: "verified", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.Verified"), Description: "Whether the account has been verified with billing information."},
		}),
	}
}
