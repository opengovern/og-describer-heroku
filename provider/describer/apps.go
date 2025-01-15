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

func ListApps(ctx context.Context, handler *resilientbridge.ResilientBridge, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	herokuChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(herokuChan)
		defer close(errorChan)
		if err := processApps(ctx, handler, herokuChan, &wg); err != nil {
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

func GetApp(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*models.Resource, error) {
	app, err := processApp(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	buildStack := model.Stack{
		ID:   app.BuildStack.ID,
		Name: app.BuildStack.Name,
	}
	generation := model.Stack{
		ID:   app.Generation.ID,
		Name: app.Generation.Name,
	}
	organization := model.Organization{
		ID:   app.Organization.ID,
		Name: app.Organization.Name,
	}
	owner := model.Owner{
		Email: app.Owner.Email,
		ID:    app.Owner.ID,
	}
	region := model.Region{
		ID:   app.Region.ID,
		Name: app.Region.Name,
	}
	space := model.Space{
		ID:     app.Space.ID,
		Name:   app.Space.Name,
		Shield: app.Space.Shield,
	}
	stack := model.Stack{
		ID:   app.Stack.ID,
		Name: app.Stack.Name,
	}
	team := model.Organization{
		ID:   app.Team.ID,
		Name: app.Team.Name,
	}
	value := models.Resource{
		ID:   app.ID,
		Name: app.Name,
		Description: model.AppDescription{
			ACM:                   app.ACM,
			ArchivedAt:            app.ArchivedAt,
			BuildStack:            buildStack,
			BuildpackProvidedDesc: app.BuildpackProvidedDesc,
			CreatedAt:             app.CreatedAt,
			Generation:            generation,
			GitURL:                app.GitURL,
			ID:                    app.ID,
			InternalRouting:       app.InternalRouting,
			Maintenance:           app.Maintenance,
			Name:                  app.Name,
			Organization:          &organization,
			Owner:                 owner,
			Region:                region,
			ReleasedAt:            app.ReleasedAt,
			RepoSize:              app.RepoSize,
			SlugSize:              app.SlugSize,
			Space:                 &space,
			Stack:                 stack,
			Team:                  &team,
			UpdatedAt:             app.UpdatedAt,
			WebURL:                app.WebURL,
		},
	}
	return &value, nil
}

func processApps(ctx context.Context, handler *resilientbridge.ResilientBridge, herokuChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var apps []model.AppJSON
	baseURL := "/apps"

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

	if err = json.Unmarshal(resp.Data, &apps); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}
	for _, app := range apps {
		wg.Add(1)
		go func(app model.AppJSON) {
			defer wg.Done()
			buildStack := model.Stack{
				ID:   app.BuildStack.ID,
				Name: app.BuildStack.Name,
			}
			generation := model.Stack{
				ID:   app.Generation.ID,
				Name: app.Generation.Name,
			}
			organization := model.Organization{
				ID:   app.Organization.ID,
				Name: app.Organization.Name,
			}
			owner := model.Owner{
				Email: app.Owner.Email,
				ID:    app.Owner.ID,
			}
			region := model.Region{
				ID:   app.Region.ID,
				Name: app.Region.Name,
			}
			space := model.Space{
				ID:     app.Space.ID,
				Name:   app.Space.Name,
				Shield: app.Space.Shield,
			}
			stack := model.Stack{
				ID:   app.Stack.ID,
				Name: app.Stack.Name,
			}
			team := model.Organization{
				ID:   app.Team.ID,
				Name: app.Team.Name,
			}
			value := models.Resource{
				ID:   app.ID,
				Name: app.Name,
				Description: model.AppDescription{
					ACM:                   app.ACM,
					ArchivedAt:            app.ArchivedAt,
					BuildStack:            buildStack,
					BuildpackProvidedDesc: app.BuildpackProvidedDesc,
					CreatedAt:             app.CreatedAt,
					Generation:            generation,
					GitURL:                app.GitURL,
					ID:                    app.ID,
					InternalRouting:       app.InternalRouting,
					Maintenance:           app.Maintenance,
					Name:                  app.Name,
					Organization:          &organization,
					Owner:                 owner,
					Region:                region,
					ReleasedAt:            app.ReleasedAt,
					RepoSize:              app.RepoSize,
					SlugSize:              app.SlugSize,
					Space:                 &space,
					Stack:                 stack,
					Team:                  &team,
					UpdatedAt:             app.UpdatedAt,
					WebURL:                app.WebURL,
				},
			}
			herokuChan <- value
		}(app)
	}
	return nil
}

func processApp(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*model.AppJSON, error) {
	var app model.AppJSON
	baseURL := "/apps/"

	finalURL := fmt.Sprintf("%s%s", baseURL, resourceID)

	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: finalURL,
		Headers:  map[string]string{"accept": "application/vnd.heroku+json; version=3"},
	}

	resp, err := handler.Request("fly", req)
	if err != nil {
		return nil, fmt.Errorf("request execution failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("error %d: %s", resp.StatusCode, string(resp.Data))
	}

	if err = json.Unmarshal(resp.Data, &app); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	return &app, nil
}
