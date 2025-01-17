package maps

import (
	"github.com/opengovern/og-describer-heroku/discovery/describers"
	model "github.com/opengovern/og-describer-heroku/discovery/pkg/models"
	"github.com/opengovern/og-describer-heroku/discovery/provider"
	"github.com/opengovern/og-describer-heroku/global"
)

var ResourceTypes = map[string]model.ResourceType{

	"Heroku/Account": {
		IntegrationType: global.IntegrationName,
		ResourceName:    "Heroku/Account",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByHeroku(describers.ListAccounts),
		GetDescriber:    nil,
	},

	"Heroku/App": {
		IntegrationType: global.IntegrationName,
		ResourceName:    "Heroku/App",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByHeroku(describers.ListApps),
		GetDescriber:    provider.DescribeSingleByHeroku(describers.GetApp),
	},

	"Heroku/Build": {
		IntegrationType: global.IntegrationName,
		ResourceName:    "Heroku/Build",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByHeroku(describers.ListBuilds),
		GetDescriber:    provider.DescribeSingleByHeroku(describers.GetBuild),
	},

	"Heroku/ConfigVars": {
		IntegrationType: global.IntegrationName,
		ResourceName:    "Heroku/ConfigVars",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByHeroku(describers.ListConfigVars),
		GetDescriber:    nil,
	},

	"Heroku/Domain": {
		IntegrationType: global.IntegrationName,
		ResourceName:    "Heroku/Domain",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByHeroku(describers.ListDomains),
		GetDescriber:    provider.DescribeSingleByHeroku(describers.GetDomain),
	},

	"Heroku/Dyno": {
		IntegrationType: global.IntegrationName,
		ResourceName:    "Heroku/Dyno",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByHeroku(describers.ListDynos),
		GetDescriber:    provider.DescribeSingleByHeroku(describers.GetDyno),
	},

	"Heroku/Dyno/Size": {
		IntegrationType: global.IntegrationName,
		ResourceName:    "Heroku/Dyno/Size",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByHeroku(describers.ListDynoSizes),
		GetDescriber:    provider.DescribeSingleByHeroku(describers.GetDynoSize),
	},
}
