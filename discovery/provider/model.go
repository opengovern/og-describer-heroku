//go:generate go run ../../pkg/sdk/runable/steampipe_es_client_generator/main.go -pluginPath ../../steampipe-plugin-REPLACEME/REPLACEME -file $GOFILE -output ../../pkg/sdk/es/resources_clients.go -resourceTypesFile ../resource_types/resource-types.json

// Implement types for each resource

package provider

import "time"

type Metadata struct{}

type OrganizationJSON struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Organization struct {
	ID   string
	Name string
}

type IdentityProviderOwnerJSON struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type IdentityProviderOwner struct {
	ID   string
	Name string
	Type string
}

type IdentityProviderJSON struct {
	ID           string                     `json:"id"`
	Name         string                     `json:"name"`
	Organization *OrganizationJSON          `json:"organization"`
	Owner        *IdentityProviderOwnerJSON `json:"owner"`
	Team         *OrganizationJSON          `json:"team"`
}

type IdentityProvider struct {
	ID           string
	Name         string
	Organization *Organization
	Owner        *IdentityProviderOwner
	Team         *Organization
}

type AccountJSON struct {
	AllowTracking           bool                  `json:"allow_tracking"`
	Beta                    bool                  `json:"beta"`
	CountryOfResidence      *string               `json:"country_of_residence"`
	CreatedAt               time.Time             `json:"created_at"`
	DefaultOrganization     *OrganizationJSON     `json:"default_organization"`
	DefaultTeam             *OrganizationJSON     `json:"default_team"`
	DelinquentAt            *time.Time            `json:"delinquent_at"`
	Email                   string                `json:"email"`
	Federated               bool                  `json:"federated"`
	ID                      string                `json:"id"`
	IdentityProvider        *IdentityProviderJSON `json:"identity_provider"`
	LastLogin               *time.Time            `json:"last_login"`
	Name                    *string               `json:"name"`
	SMSNumber               *string               `json:"sms_number"`
	SuspendedAt             *time.Time            `json:"suspended_at"`
	TwoFactorAuthentication bool                  `json:"two_factor_authentication"`
	UpdatedAt               time.Time             `json:"updated_at"`
	Verified                bool                  `json:"verified"`
}

type AccountDescription struct {
	AllowTracking           bool
	Beta                    bool
	CountryOfResidence      *string
	CreatedAt               time.Time
	DefaultOrganization     *Organization
	DefaultTeam             *Organization
	DelinquentAt            *time.Time
	Email                   string
	Federated               bool
	ID                      string
	IdentityProvider        *IdentityProvider
	LastLogin               *time.Time
	Name                    *string
	SMSNumber               *string
	SuspendedAt             *time.Time
	TwoFactorAuthentication bool
	UpdatedAt               time.Time
	Verified                bool
}

type StackJSON struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Stack struct {
	ID   string
	Name string
}

type OwnerJSON struct {
	Email string `json:"email"`
	ID    string `json:"id"`
}

type Owner struct {
	Email string
	ID    string
}

type RegionJSON struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Region struct {
	ID   string
	Name string
}

type SpaceJSON struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Shield bool   `json:"shield"`
}

type Space struct {
	ID     string
	Name   string
	Shield bool
}

type AppJSON struct {
	ACM                   bool              `json:"acm"`
	ArchivedAt            *time.Time        `json:"archived_at"`
	BuildStack            StackJSON         `json:"build_stack"`
	BuildpackProvidedDesc *string           `json:"buildpack_provided_description"`
	CreatedAt             time.Time         `json:"created_at"`
	Generation            StackJSON         `json:"generation"`
	GitURL                string            `json:"git_url"`
	ID                    string            `json:"id"`
	InternalRouting       *bool             `json:"internal_routing"`
	Maintenance           bool              `json:"maintenance"`
	Name                  string            `json:"name"`
	Organization          *OrganizationJSON `json:"organization"`
	Owner                 OwnerJSON         `json:"owner"`
	Region                RegionJSON        `json:"region"`
	ReleasedAt            *time.Time        `json:"released_at"`
	RepoSize              *int              `json:"repo_size"`
	SlugSize              *int              `json:"slug_size"`
	Space                 *SpaceJSON        `json:"space"`
	Stack                 StackJSON         `json:"stack"`
	Team                  *OrganizationJSON `json:"team"`
	UpdatedAt             time.Time         `json:"updated_at"`
	WebURL                *string           `json:"web_url"`
}

type AppDescription struct {
	ACM                   bool
	ArchivedAt            *time.Time
	BuildStack            Stack
	BuildpackProvidedDesc *string
	CreatedAt             time.Time
	Generation            Stack
	GitURL                string
	ID                    string
	InternalRouting       *bool
	Maintenance           bool
	Name                  string
	Organization          *Organization
	Owner                 Owner
	Region                Region
	ReleasedAt            *time.Time
	RepoSize              *int
	SlugSize              *int
	Space                 *Space
	Stack                 Stack
	Team                  *Organization
	UpdatedAt             time.Time
	WebURL                *string
}

type BuildpackJSON struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Buildpack struct {
	Name string
	URL  string
}

type ReleaseJSON struct {
	ID string `json:"id"`
}

type Release struct {
	ID string
}

type SlugJSON struct {
	ID string `json:"id"`
}

type Slug struct {
	ID string
}

type SourceBlobJSON struct {
	Checksum           *string `json:"checksum"`
	URL                string  `json:"url"`
	Version            *string `json:"version"`
	VersionDescription *string `json:"version_description"`
}

type SourceBlob struct {
	Checksum           *string
	URL                string
	Version            *string
	VersionDescription *string
}

type UserJSON struct {
	Email string `json:"email"`
	ID    string `json:"id"`
}

type User struct {
	Email string
	ID    string
}

type BuildJSON struct {
	AppID           string           `json:"app_id"`
	Buildpacks      *[]BuildpackJSON `json:"buildpacks"`
	CreatedAt       time.Time        `json:"created_at"`
	ID              string           `json:"id"`
	OutputStreamURL string           `json:"output_stream_url"`
	Release         *ReleaseJSON     `json:"release"`
	Slug            *SlugJSON        `json:"slug"`
	SourceBlob      SourceBlobJSON   `json:"source_blob"`
	Stack           string           `json:"stack"`
	Status          string           `json:"status"`
	UpdatedAt       time.Time        `json:"updated_at"`
	User            UserJSON         `json:"user"`
}

type BuildDescription struct {
	AppID           string
	Buildpacks      *[]Buildpack
	CreatedAt       time.Time
	ID              string
	OutputStreamURL string
	Release         *Release
	Slug            *Slug
	SourceBlob      SourceBlob
	Stack           string
	Status          string
	UpdatedAt       time.Time
	User            User
}

type ConfigVarsDescription struct {
	AppName   string
	Variables map[string]interface{}
}

type SNIEndpointJSON struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type SNIEndpoint struct {
	ID   string
	Name string
}

type DomainJSON struct {
	ACMStatus       *string          `json:"acm_status"`
	ACMStatusReason *string          `json:"acm_status_reason"`
	AppID           string           `json:"app_id"`
	AppName         string           `json:"app_name"`
	CName           *string          `json:"cname"`
	CreatedAt       time.Time        `json:"created_at"`
	Hostname        string           `json:"hostname"`
	ID              string           `json:"id"`
	Kind            string           `json:"kind"`
	SNIEndpoint     *SNIEndpointJSON `json:"sni_endpoint"`
	Status          string           `json:"status"`
	UpdatedAt       time.Time        `json:"updated_at"`
}

type DomainDescription struct {
	ACMStatus       *string
	ACMStatusReason *string
	AppID           string
	AppName         string
	CName           *string
	CreatedAt       time.Time
	Hostname        string
	ID              string
	Kind            string
	SNIEndpoint     *SNIEndpoint
	Status          string
	UpdatedAt       time.Time
}

type DynoReleaseJSON struct {
	ID      string `json:"id"`
	Version int    `json:"version"`
}

type DynoRelease struct {
	ID      string
	Version int
}

type DynoJSON struct {
	AppID     string          `json:"app_id"`
	AppName   string          `json:"app_name"`
	AttachURL *string         `json:"attach_url"`
	Command   string          `json:"command"`
	CreatedAt time.Time       `json:"created_at"`
	ID        string          `json:"id"`
	Name      string          `json:"name"`
	Release   DynoReleaseJSON `json:"release"`
	Size      string          `json:"size"`
	State     string          `json:"state"`
	Type      string          `json:"type"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type DynoDescription struct {
	AppID     string
	AppName   string
	AttachURL *string
	Command   string
	CreatedAt time.Time
	ID        string
	Name      string
	Release   DynoRelease
	Size      string
	State     string
	Type      string
	UpdatedAt time.Time
}

type GenerationJSON struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Generation struct {
	ID   string
	Name string
}

type DynoSizeJSON struct {
	Architecture     string         `json:"architecture"`
	Compute          int            `json:"compute"`
	Cost             any            `json:"cost"`
	Dedicated        bool           `json:"dedicated"`
	Generation       GenerationJSON `json:"generation"`
	ID               string         `json:"id"`
	Memory           float64        `json:"memory"`
	Name             string         `json:"name"`
	PreciseDynoUnits float64        `json:"precise_dyno_units"`
	PrivateSpaceOnly bool           `json:"private_space_only"`
}

type DynoSizeDescription struct {
	Architecture     string
	Compute          int
	Cost             any
	Dedicated        bool
	Generation       Generation
	ID               string
	Memory           float64
	Name             string
	PreciseDynoUnits float64
	PrivateSpaceOnly bool
}
