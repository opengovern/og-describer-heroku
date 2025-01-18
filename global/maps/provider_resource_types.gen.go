package maps

import (
	"github.com/opengovern/og-describer-heroku/discovery/describers"
	model "github.com/opengovern/og-describer-heroku/discovery/pkg/models"
	"github.com/opengovern/og-describer-heroku/discovery/provider"
	"github.com/opengovern/og-describer-heroku/platform/constants"
	"github.com/opengovern/og-util/pkg/integration/interfaces"
)

var ResourceTypes = map[string]model.ResourceType{

	"Heroku/Account": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Heroku/Account",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByHeroku(describers.ListAccounts),
		GetDescriber:    nil,
	},

	"Heroku/App": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Heroku/App",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByHeroku(describers.ListApps),
		GetDescriber:    provider.DescribeSingleByHeroku(describers.GetApp),
	},

	"Heroku/Build": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Heroku/Build",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByHeroku(describers.ListBuilds),
		GetDescriber:    provider.DescribeSingleByHeroku(describers.GetBuild),
	},

	"Heroku/ConfigVars": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Heroku/ConfigVars",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByHeroku(describers.ListConfigVars),
		GetDescriber:    nil,
	},

	"Heroku/Domain": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Heroku/Domain",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByHeroku(describers.ListDomains),
		GetDescriber:    provider.DescribeSingleByHeroku(describers.GetDomain),
	},

	"Heroku/Dyno": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Heroku/Dyno",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByHeroku(describers.ListDynos),
		GetDescriber:    provider.DescribeSingleByHeroku(describers.GetDyno),
	},

	"Heroku/Dyno/Size": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Heroku/Dyno/Size",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByHeroku(describers.ListDynoSizes),
		GetDescriber:    provider.DescribeSingleByHeroku(describers.GetDynoSize),
	},
}

var ResourceTypeConfigs = map[string]*interfaces.ResourceTypeConfiguration{

	"Heroku/Account": {
		Name:            "Heroku/Account",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Heroku/App": {
		Name:            "Heroku/App",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Heroku/Build": {
		Name:            "Heroku/Build",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Heroku/ConfigVars": {
		Name:            "Heroku/ConfigVars",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Heroku/Domain": {
		Name:            "Heroku/Domain",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Heroku/Dyno": {
		Name:            "Heroku/Dyno",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Heroku/Dyno/Size": {
		Name:            "Heroku/Dyno/Size",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},
}

var ResourceTypesList = []string{
	"Heroku/Account",
	"Heroku/App",
	"Heroku/Build",
	"Heroku/ConfigVars",
	"Heroku/Domain",
	"Heroku/Dyno",
	"Heroku/Dyno/Size",
}
