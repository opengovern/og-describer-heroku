package describers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-heroku/discovery/pkg/models"
	"github.com/opengovern/og-describer-heroku/discovery/provider"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"sync"
)

func ListDomains(ctx context.Context, handler *resilientbridge.ResilientBridge, appName string, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	herokuChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(herokuChan)
		defer close(errorChan)
		if err := processDomains(ctx, handler, appName, herokuChan, &wg); err != nil {
			errorChan <- err // Send error to the error channel
		}
		wg.Wait()
	}()

	var values []models.Resource
	for {
		select {
		case value, ok := <-herokuChan:
			if !ok {
				return values, nil
			}
			if stream != nil {
				if err := (*stream)(value); err != nil {
					return nil, err
				}
			} else {
				values = append(values, value)
			}
		case err := <-errorChan:
			return nil, err
		}
	}
}

func GetDomain(ctx context.Context, handler *resilientbridge.ResilientBridge, appName, resourceID string) (*models.Resource, error) {
	domain, err := processDomain(ctx, handler, appName, resourceID)
	if err != nil {
		return nil, err
	}
	sniEndpoint := provider.SNIEndpoint{
		ID:   domain.SNIEndpoint.ID,
		Name: domain.SNIEndpoint.Name,
	}
	value := models.Resource{
		ID:   domain.ID,
		Name: domain.ID,
		Description: provider.DomainDescription{
			ACMStatus:       domain.ACMStatus,
			ACMStatusReason: domain.ACMStatusReason,
			AppID:           domain.AppID,
			AppName:         domain.AppName,
			CName:           domain.CName,
			CreatedAt:       domain.CreatedAt,
			Hostname:        domain.Hostname,
			ID:              domain.ID,
			Kind:            domain.Kind,
			SNIEndpoint:     &sniEndpoint,
			Status:          domain.Status,
			UpdatedAt:       domain.UpdatedAt,
		},
	}
	return &value, nil
}

func processDomains(ctx context.Context, handler *resilientbridge.ResilientBridge, appName string, herokuChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var domains []provider.DomainJSON
	baseURL := "/apps/"

	finalURL := fmt.Sprintf("%s%s/domains", baseURL, appName)

	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: finalURL,
		Headers:  map[string]string{"accept": "application/vnd.heroku+json; version=3"},
	}

	resp, err := handler.Request("heroku", req)
	if err != nil {
		return fmt.Errorf("request execution failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("error %d: %s", resp.StatusCode, string(resp.Data))
	}

	if err = json.Unmarshal(resp.Data, &domains); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}
	for _, domain := range domains {
		wg.Add(1)
		go func(domain provider.DomainJSON) {
			defer wg.Done()
			sniEndpoint := provider.SNIEndpoint{
				ID:   domain.SNIEndpoint.ID,
				Name: domain.SNIEndpoint.Name,
			}
			value := models.Resource{
				ID:   domain.ID,
				Name: domain.ID,
				Description: provider.DomainDescription{
					ACMStatus:       domain.ACMStatus,
					ACMStatusReason: domain.ACMStatusReason,
					AppID:           domain.AppID,
					AppName:         domain.AppName,
					CName:           domain.CName,
					CreatedAt:       domain.CreatedAt,
					Hostname:        domain.Hostname,
					ID:              domain.ID,
					Kind:            domain.Kind,
					SNIEndpoint:     &sniEndpoint,
					Status:          domain.Status,
					UpdatedAt:       domain.UpdatedAt,
				},
			}
			herokuChan <- value
		}(domain)
	}
	return nil
}

func processDomain(ctx context.Context, handler *resilientbridge.ResilientBridge, appName, resourceID string) (*provider.DomainJSON, error) {
	var domain provider.DomainJSON
	baseURL := "/apps/"

	finalURL := fmt.Sprintf("%s%s/domains/%s", baseURL, appName, resourceID)

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

	if err = json.Unmarshal(resp.Data, &domain); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	return &domain, nil
}
