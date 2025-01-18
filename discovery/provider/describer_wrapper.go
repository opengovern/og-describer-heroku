package provider

import (
	"errors"
	"github.com/opengovern/og-describer-heroku/discovery/describers"
	"github.com/opengovern/og-describer-heroku/discovery/pkg/models"
	"github.com/opengovern/og-util/pkg/describe/enums"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"github.com/opengovern/resilient-bridge/adapters"
	"golang.org/x/net/context"
	"time"
)

// DescribeListByHeroku A wrapper to pass Heroku authorization to describers functions
func DescribeListByHeroku(describe func(context.Context, *resilientbridge.ResilientBridge, string, *models.StreamSender) ([]models.Resource, error)) models.ResourceDescriber {
	return func(ctx context.Context, cfg models.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, stream *models.StreamSender) ([]models.Resource, error) {
		ctx = describers.WithTriggerType(ctx, triggerType)

		var err error
		// Check for the token
		if cfg.Token == "" {
			return nil, errors.New("token must be configured")
		}

		resilientBridge := resilientbridge.NewResilientBridge()

		restMaxRequests := 500
		restWindowSecs := int64(60)

		// Register TailScale provider
		resilientBridge.RegisterProvider("heroku", &adapters.DopplerAdapter{APIToken: cfg.Token}, &resilientbridge.ProviderConfig{
			UseProviderLimits:   true,
			MaxRequestsOverride: &restMaxRequests,
			WindowSecsOverride:  &restWindowSecs,
			MaxRetries:          3,
			BaseBackoff:         200 * time.Millisecond,
		})

		appName := additionalParameters["AppName"]
		// Get values from describers
		var values []models.Resource
		result, err := describe(ctx, resilientBridge, appName, stream)
		if err != nil {
			return nil, err
		}
		values = append(values, result...)
		return values, nil
	}
}

// DescribeSingleByHeroku A wrapper to pass Heroku authorization to describers functions
func DescribeSingleByHeroku(describe func(context.Context, *resilientbridge.ResilientBridge, string, string) (*models.Resource, error)) models.SingleResourceDescriber {
	return func(ctx context.Context, cfg models.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, resourceID string, stream *models.StreamSender) (*models.Resource, error) {
		ctx = describers.WithTriggerType(ctx, triggerType)

		var err error
		// Check for the token
		if cfg.Token == "" {
			return nil, errors.New("token must be configured")
		}

		resilientBridge := resilientbridge.NewResilientBridge()

		restMaxRequests := 500
		restWindowSecs := int64(60)

		// Register TailScale provider
		resilientBridge.RegisterProvider("heroku", &adapters.DopplerAdapter{APIToken: cfg.Token}, &resilientbridge.ProviderConfig{
			UseProviderLimits:   true,
			MaxRequestsOverride: &restMaxRequests,
			WindowSecsOverride:  &restWindowSecs,
			MaxRetries:          3,
			BaseBackoff:         200 * time.Millisecond,
		})

		appName := additionalParameters["AppName"]
		// Get value from describers
		value, err := describe(ctx, resilientBridge, appName, resourceID)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
}
