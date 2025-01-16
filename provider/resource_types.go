package provider

import (
	model "github.com/opengovern/og-describer-heroku/pkg/sdk/models"
	"github.com/opengovern/og-describer-heroku/provider/configs"
	"github.com/opengovern/og-describer-heroku/provider/describer"
)

var ResourceTypes = map[string]model.ResourceType{

	"Heroku/Account": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Heroku/Account",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByHeroku(describer.ListAccounts),
		GetDescriber:    nil,
	},

	"Heroku/App": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Heroku/App",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByHeroku(describer.ListApps),
		GetDescriber:    DescribeSingleByHeroku(describer.GetApp),
	},

	"Heroku/Build": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Heroku/Build",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByHeroku(describer.ListBuilds),
		GetDescriber:    DescribeSingleByHeroku(describer.GetBuild),
	},

	"Heroku/ConfigVars": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Heroku/ConfigVars",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByHeroku(describer.ListConfigVars),
		GetDescriber:    nil,
	},

	"Heroku/Domain": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Heroku/Domain",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByHeroku(describer.ListDomains),
		GetDescriber:    DescribeSingleByHeroku(describer.GetDomain),
	},

	"Heroku/Dyno": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Heroku/Dyno",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByHeroku(describer.ListDynos),
		GetDescriber:    DescribeSingleByHeroku(describer.GetDyno),
	},

	"Heroku/Dyno/Size": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Heroku/Dyno/Size",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByHeroku(describer.ListDynoSizes),
		GetDescriber:    DescribeSingleByHeroku(describer.GetDynoSize),
	},
}
