package global

import "github.com/opengovern/og-util/pkg/integration"

const (
	IntegrationTypeLower = "heroku"                                    // example: aws, azure
	IntegrationName      = integration.Type("heroku_account")          // example: AWS_ACCOUNT, AZURE_SUBSCRIPTION
	OGPluginRepoURL      = "github.com/opengovern/og-describer-heroku" // example: github.com/opengovern/og-describers-aws
)

type IntegrationCredentials struct {
	Token string `json:"token"`
}
