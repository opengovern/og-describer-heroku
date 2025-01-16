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

func ListBuilds(ctx context.Context, handler *resilientbridge.ResilientBridge, appName string, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	herokuChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(herokuChan)
		defer close(errorChan)
		if err := processBuilds(ctx, handler, appName, herokuChan, &wg); err != nil {
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

func GetBuild(ctx context.Context, handler *resilientbridge.ResilientBridge, appName, resourceID string) (*models.Resource, error) {
	build, err := processBuild(ctx, handler, appName, resourceID)
	if err != nil {
		return nil, err
	}
	var buildPacks []model.Buildpack
	if build.Buildpacks != nil {
		for _, buildPack := range *build.Buildpacks {
			buildPacks = append(buildPacks, model.Buildpack{
				Name: buildPack.Name,
				URL:  buildPack.URL,
			})
		}
	}
	release := model.Release{
		ID: build.Release.ID,
	}
	slug := model.Slug{
		ID: build.Slug.ID,
	}
	sourceBlob := model.SourceBlob{
		Checksum:           build.SourceBlob.Checksum,
		URL:                build.SourceBlob.URL,
		Version:            build.SourceBlob.Version,
		VersionDescription: build.SourceBlob.VersionDescription,
	}
	user := model.User{
		Email: build.User.Email,
		ID:    build.User.ID,
	}
	value := models.Resource{
		ID:   build.ID,
		Name: build.ID,
		Description: model.BuildDescription{
			AppID:           build.AppID,
			Buildpacks:      &buildPacks,
			CreatedAt:       build.CreatedAt,
			ID:              build.Stack,
			OutputStreamURL: build.OutputStreamURL,
			Release:         &release,
			Slug:            &slug,
			SourceBlob:      sourceBlob,
			Stack:           build.Stack,
			Status:          build.Status,
			UpdatedAt:       build.UpdatedAt,
			User:            user,
		},
	}
	return &value, nil
}

func processBuilds(ctx context.Context, handler *resilientbridge.ResilientBridge, appName string, herokuChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var builds []model.BuildJSON
	baseURL := "/apps/"

	finalURL := fmt.Sprintf("%s%s/builds", baseURL, appName)

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

	if err = json.Unmarshal(resp.Data, &builds); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}
	for _, build := range builds {
		wg.Add(1)
		go func(build model.BuildJSON) {
			defer wg.Done()
			var buildPacks []model.Buildpack
			if build.Buildpacks != nil {
				for _, buildPack := range *build.Buildpacks {
					buildPacks = append(buildPacks, model.Buildpack{
						Name: buildPack.Name,
						URL:  buildPack.URL,
					})
				}
			}
			release := model.Release{
				ID: build.Release.ID,
			}
			slug := model.Slug{
				ID: build.Slug.ID,
			}
			sourceBlob := model.SourceBlob{
				Checksum:           build.SourceBlob.Checksum,
				URL:                build.SourceBlob.URL,
				Version:            build.SourceBlob.Version,
				VersionDescription: build.SourceBlob.VersionDescription,
			}
			user := model.User{
				Email: build.User.Email,
				ID:    build.User.ID,
			}
			value := models.Resource{
				ID:   build.ID,
				Name: build.ID,
				Description: model.BuildDescription{
					AppID:           build.AppID,
					Buildpacks:      &buildPacks,
					CreatedAt:       build.CreatedAt,
					ID:              build.Stack,
					OutputStreamURL: build.OutputStreamURL,
					Release:         &release,
					Slug:            &slug,
					SourceBlob:      sourceBlob,
					Stack:           build.Stack,
					Status:          build.Status,
					UpdatedAt:       build.UpdatedAt,
					User:            user,
				},
			}
			herokuChan <- value
		}(build)
	}
	return nil
}

func processBuild(ctx context.Context, handler *resilientbridge.ResilientBridge, appName, resourceID string) (*model.BuildJSON, error) {
	var build model.BuildJSON
	baseURL := "/apps/"

	finalURL := fmt.Sprintf("%s%s/builds/%s", baseURL, appName, resourceID)

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

	if err = json.Unmarshal(resp.Data, &build); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	return &build, nil
}
