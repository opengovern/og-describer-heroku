package describer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-heroku/pkg/sdk/models"
	"github.com/opengovern/og-describer-heroku/provider/model"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"sync"
)

func ListDynoSizes(ctx context.Context, handler *resilientbridge.ResilientBridge, appName string, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	herokuChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(herokuChan)
		defer close(errorChan)
		if err := processDynoSizes(ctx, handler, appName, herokuChan, &wg); err != nil {
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

func GetDynoSize(ctx context.Context, handler *resilientbridge.ResilientBridge, appName, resourceID string) (*models.Resource, error) {
	dynoSize, err := processDynoSize(ctx, handler, appName, resourceID)
	if err != nil {
		return nil, err
	}
	generation := model.Generation{
		ID:   dynoSize.Generation.ID,
		Name: dynoSize.Generation.Name,
	}
	value := models.Resource{
		ID:   dynoSize.ID,
		Name: dynoSize.Name,
		Description: model.DynoSizeDescription{
			Architecture:     dynoSize.Architecture,
			Compute:          dynoSize.Compute,
			Cost:             dynoSize.Cost,
			Dedicated:        dynoSize.Dedicated,
			Generation:       generation,
			ID:               dynoSize.ID,
			Memory:           dynoSize.Memory,
			Name:             dynoSize.Name,
			PreciseDynoUnits: dynoSize.PreciseDynoUnits,
			PrivateSpaceOnly: dynoSize.PrivateSpaceOnly,
		},
	}
	return &value, nil
}

func processDynoSizes(ctx context.Context, handler *resilientbridge.ResilientBridge, appName string, herokuChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var dynoSizes []model.DynoSizeJSON
	baseURL := "/dyno-sizes"

	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: baseURL,
		Headers:  map[string]string{"accept": "application/vnd.heroku+json; version=3"},
	}

	resp, err := handler.Request("heroku", req)
	if err != nil {
		return fmt.Errorf("request execution failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("error %d: %s", resp.StatusCode, string(resp.Data))
	}

	if err = json.Unmarshal(resp.Data, &dynoSizes); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}
	for _, dynoSize := range dynoSizes {
		wg.Add(1)
		go func(dynoSize model.DynoSizeJSON) {
			defer wg.Done()
			generation := model.Generation{
				ID:   dynoSize.Generation.ID,
				Name: dynoSize.Generation.Name,
			}
			value := models.Resource{
				ID:   dynoSize.ID,
				Name: dynoSize.Name,
				Description: model.DynoSizeDescription{
					Architecture:     dynoSize.Architecture,
					Compute:          dynoSize.Compute,
					Cost:             dynoSize.Cost,
					Dedicated:        dynoSize.Dedicated,
					Generation:       generation,
					ID:               dynoSize.ID,
					Memory:           dynoSize.Memory,
					Name:             dynoSize.Name,
					PreciseDynoUnits: dynoSize.PreciseDynoUnits,
					PrivateSpaceOnly: dynoSize.PrivateSpaceOnly,
				},
			}
			herokuChan <- value
		}(dynoSize)
	}
	return nil
}

func processDynoSize(ctx context.Context, handler *resilientbridge.ResilientBridge, appName, resourceID string) (*model.DynoSizeJSON, error) {
	var dynoSize model.DynoSizeJSON
	baseURL := "/dyno-sizes/"

	finalURL := fmt.Sprintf("%s%s", baseURL, resourceID)

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

	if err = json.Unmarshal(resp.Data, &dynoSize); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	return &dynoSize, nil
}
