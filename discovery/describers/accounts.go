package describers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-heroku/discovery/pkg/models"
	"github.com/opengovern/og-describer-heroku/discovery/provider"
	resilientbridge "github.com/opengovern/resilient-bridge"
)

func ListAccounts(ctx context.Context, handler *resilientbridge.ResilientBridge, appName string, stream *models.StreamSender) ([]models.Resource, error) {
	values, err := processAccounts(ctx, handler, stream)
	if err != nil {
		return nil, err
	}
	return values, nil
}

func processAccounts(ctx context.Context, handler *resilientbridge.ResilientBridge, stream *models.StreamSender) ([]models.Resource, error) {
	var account provider.AccountJSON
	baseURL := "/account"

	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: baseURL,
		Headers:  map[string]string{"accept": "application/vnd.heroku+json; version=3"},
	}

	resp, err := handler.Request("heroku", req)
	if err != nil {
		return nil, fmt.Errorf("request execution failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("error %d: %s", resp.StatusCode, string(resp.Data))
	}

	if err = json.Unmarshal(resp.Data, &account); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	var values []models.Resource
	var name string
	if account.Name != nil {
		name = *account.Name
	}
	var organization provider.Organization
	if account.DefaultOrganization != nil {
		organization = provider.Organization{
			ID:   account.DefaultOrganization.ID,
			Name: account.DefaultOrganization.Name,
		}
	}
	var team provider.Organization
	if account.DefaultTeam != nil {
		team = provider.Organization{
			ID:   account.DefaultTeam.ID,
			Name: account.DefaultTeam.Name,
		}
	}
	var identityProviderOwner provider.IdentityProviderOwner
	if account.IdentityProvider != nil {
		identityProviderOwner = provider.IdentityProviderOwner{
			ID:   account.IdentityProvider.Owner.ID,
			Name: account.IdentityProvider.Owner.Name,
			Type: account.IdentityProvider.Owner.Type,
		}
	}
	var providerOrganization provider.Organization
	if account.IdentityProvider != nil {
		providerOrganization = provider.Organization{
			ID:   account.IdentityProvider.Organization.ID,
			Name: account.IdentityProvider.Organization.Name,
		}
	}
	var providerTeam provider.Organization
	if account.IdentityProvider != nil {
		providerTeam = provider.Organization{
			ID:   account.IdentityProvider.Team.ID,
			Name: account.IdentityProvider.Team.Name,
		}
	}
	var identityProvider provider.IdentityProvider
	if account.IdentityProvider != nil {
		identityProvider = provider.IdentityProvider{
			ID:           account.IdentityProvider.ID,
			Name:         account.IdentityProvider.Name,
			Organization: &providerOrganization,
			Owner:        &identityProviderOwner,
			Team:         &providerTeam,
		}
	}
	value := models.Resource{
		ID:   account.ID,
		Name: name,
		Description: provider.AccountDescription{
			AllowTracking:           account.AllowTracking,
			Beta:                    account.Beta,
			CountryOfResidence:      account.CountryOfResidence,
			CreatedAt:               account.CreatedAt,
			DefaultOrganization:     &organization,
			DefaultTeam:             &team,
			DelinquentAt:            account.DelinquentAt,
			Email:                   account.Email,
			Federated:               account.Federated,
			ID:                      account.ID,
			IdentityProvider:        &identityProvider,
			LastLogin:               account.LastLogin,
			Name:                    account.Name,
			SMSNumber:               account.SMSNumber,
			SuspendedAt:             account.SuspendedAt,
			TwoFactorAuthentication: account.TwoFactorAuthentication,
			UpdatedAt:               account.UpdatedAt,
			Verified:                account.Verified,
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
