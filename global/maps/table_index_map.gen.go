package maps

import (
	"github.com/opengovern/og-describer-heroku/discovery/pkg/es"
)

var Map = map[string]string{
	"Heroku/Account":    "heroku_account",
	"Heroku/App":        "heroku_app",
	"Heroku/Build":      "heroku_build",
	"Heroku/ConfigVars": "heroku_config_vars",
	"Heroku/Domain":     "heroku_domain",
	"Heroku/Dyno":       "heroku_dyno",
	"Heroku/Dyno/Size":  "heroku_dyno_size",
}

var DescriptionMap = map[string]interface{}{
	"Heroku/Account":    opengovernance.Account{},
	"Heroku/App":        opengovernance.App{},
	"Heroku/Build":      opengovernance.Build{},
	"Heroku/ConfigVars": opengovernance.ConfigVars{},
	"Heroku/Domain":     opengovernance.Domain{},
	"Heroku/Dyno":       opengovernance.Dyno{},
	"Heroku/Dyno/Size":  opengovernance.DynoSize{},
}

var ReverseMap = map[string]string{
	"heroku_account":     "Heroku/Account",
	"heroku_app":         "Heroku/App",
	"heroku_build":       "Heroku/Build",
	"heroku_config_vars": "Heroku/ConfigVars",
	"heroku_domain":      "Heroku/Domain",
	"heroku_dyno":        "Heroku/Dyno",
	"heroku_dyno_size":   "Heroku/Dyno/Size",
}
