package describers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-heroku/discovery/pkg/models"
	"github.com/opengovern/og-describer-heroku/discovery/provider"
	resilientbridge "github.com/opengovern/resilient-bridge"
)

func ListConfigVars(ctx context.Context, handler *resilientbridge.ResilientBridge, appName string, stream *models.StreamSender) ([]models.Resource, error) {
	values, err := processConfigVars(ctx, handler, appName, stream)
	if err != nil {
		return nil, err
	}
	return values, nil
}

func processConfigVars(ctx context.Context, handler *resilientbridge.ResilientBridge, appName string, stream *models.StreamSender) ([]models.Resource, error) {
	baseURL := "/apps/"

	finalURL := fmt.Sprintf("%s%s/config-vars", baseURL, appName)

	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: finalURL,
		Headers:  map[string]string{"accept": "application/vnd.heroku+json; version=3"},
	}

	resp, err := handler.Request("heroku", req)
	if err != nil {
		return nil, fmt.Errorf("request execution failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("error %d: %s", resp.StatusCode, string(resp.Data))
	}

	var configVarsJSON map[string]interface{}
	if err = json.Unmarshal(resp.Data, &configVarsJSON); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	var values []models.Resource
	value := models.Resource{
		ID:   appName,
		Name: appName,
		Description: provider.ConfigVarsDescription{
			AppName:   appName,
			Variables: configVarsJSON,
		},
	}
	if stream != nil {
		if err := (*stream)(value); err != nil {
			return nil, err
		}
	} else {
		values = append(values, value)
	}
	return values, nil
}
