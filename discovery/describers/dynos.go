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

func ListDynos(ctx context.Context, handler *resilientbridge.ResilientBridge, appName string, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	herokuChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(herokuChan)
		defer close(errorChan)
		if err := processDynos(ctx, handler, appName, herokuChan, &wg); err != nil {
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

func GetDyno(ctx context.Context, handler *resilientbridge.ResilientBridge, appName, resourceID string) (*models.Resource, error) {
	dyno, err := processDyno(ctx, handler, appName, resourceID)
	if err != nil {
		return nil, err
	}
	release := provider.DynoRelease{
		ID:      dyno.Release.ID,
		Version: dyno.Release.Version,
	}
	value := models.Resource{
		ID:   dyno.ID,
		Name: dyno.Name,
		Description: provider.DynoDescription{
			AppID:     dyno.AppID,
			AppName:   dyno.AppName,
			AttachURL: dyno.AttachURL,
			Command:   dyno.Command,
			CreatedAt: dyno.CreatedAt,
			ID:        dyno.ID,
			Name:      dyno.Name,
			Release:   release,
			Size:      dyno.Size,
			State:     dyno.State,
			Type:      dyno.Type,
			UpdatedAt: dyno.UpdatedAt,
		},
	}
	return &value, nil
}

func processDynos(ctx context.Context, handler *resilientbridge.ResilientBridge, appName string, herokuChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var dynos []provider.DynoJSON
	baseURL := "/apps/"

	finalURL := fmt.Sprintf("%s%s/dynos", baseURL, appName)

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

	if err = json.Unmarshal(resp.Data, &dynos); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}
	for _, dyno := range dynos {
		wg.Add(1)
		go func(dyno provider.DynoJSON) {
			defer wg.Done()
			release := provider.DynoRelease{
				ID:      dyno.Release.ID,
				Version: dyno.Release.Version,
			}
			value := models.Resource{
				ID:   dyno.ID,
				Name: dyno.Name,
				Description: provider.DynoDescription{
					AppID:     dyno.AppID,
					AppName:   dyno.AppName,
					AttachURL: dyno.AttachURL,
					Command:   dyno.Command,
					CreatedAt: dyno.CreatedAt,
					ID:        dyno.ID,
					Name:      dyno.Name,
					Release:   release,
					Size:      dyno.Size,
					State:     dyno.State,
					Type:      dyno.Type,
					UpdatedAt: dyno.UpdatedAt,
				},
			}
			herokuChan <- value
		}(dyno)
	}
	return nil
}

func processDyno(ctx context.Context, handler *resilientbridge.ResilientBridge, appName, resourceID string) (*provider.DynoJSON, error) {
	var dyno provider.DynoJSON
	baseURL := "/apps/"

	finalURL := fmt.Sprintf("%s%s/dynos/%s", baseURL, appName, resourceID)

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

	if err = json.Unmarshal(resp.Data, &dyno); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	return &dyno, nil
}
