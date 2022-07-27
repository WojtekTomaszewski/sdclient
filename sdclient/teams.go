package sdclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Teams struct {
	Teams []TeamItem `json:"teams,omitempty"`
}

type Team struct {
	Team TeamItem `json:"team,omitempty"`
}

type TeamItem struct {
	Origin              string                  `json:"origin,omitempty"`
	Description         string                  `json:"description,omitempty"`
	Version             int                     `json:"version,omitempty"`
	Products            []string                `json:"products,omitempty"`
	Immutable           bool                    `json:"immutable,omitempty"`
	DateCreated         int64                   `json:"dateCreated,omitempty"`
	LastUpdated         int64                   `json:"lastUpdated,omitempty"`
	CustomerID          int                     `json:"customerId,omitempty"`
	ID                  int                     `json:"id,omitempty"`
	Filter              string                  `json:"filterId,omitempty"`
	NamespaceFilters    *NamespaceFiltersObject `json:"namespaceFilters,omitempty"`
	DefaultTeamRole     string                  `json:"defaultTeamRole,omitempty"`
	Theme               string                  `json:"theme,omitempty"`
	Show                string                  `json:"show,omitempty"`
	EntryPoint          *EnrtyPointObject       `json:"entryPoint,omitempty"`
	CanUseSysdigCapture bool                    `json:"canUseSysdigCapture"`
	CanUseAgentCli      bool                    `json:"canUseAgentCli"`
	CanUseCustomEvents  bool                    `json:"canUseCustomEvents"`
	CanUseAwsMetrics    bool                    `json:"canUseAwsMetrics"`
	CanUseBeaconMetrics bool                    `json:"canUseBeaconMetrics"`
	CanUseRapidResponse bool                    `json:"canUseRapidResponse"`
	UserCount           int                     `json:"userCount,omitempty"`
	Name                string                  `json:"name,omitempty"`
	Default             bool                    `json:"default,omitempty"`
	UsersRole           *UserRoleObject         `json:"usersRole,omitempty"`
	Users               []string                `json:"users"`
}

type NamespaceFiltersObject struct {
	PrometheusRemoteWrite string `json:"prometheusRemoteWrite,omitempty"`
	IbmPlatformMetrics    string `json:"ibmPlatformMetrics,omitempty"`
}

type EnrtyPointObject struct {
	Module string `json:"module,omitempty"`
}

type UserRoleObject struct {
	UserId         int    `json:"userId,omitempty"`
	Role           string `json:"role,omitempty"`
	TeamId         int    `json:"teamId,omitempty"`
	TeamName       string `json:"teamName,omitempty"`
	TeamTheme      string `json:"teamTheme,omitempty"`
	UserName       string `json:"userName,omitempty"`
	Admin          bool   `json:"admin,omitempty"`
	RemovalWarning string `json:"removalWarning,omitempty"`
}

func (c *Client) ListTeams() (*Teams, error) {
	return c.ListTeamsWithContext(context.Background())
}

func (c *Client) ListTeamsWithContext(ctx context.Context) (*Teams, error) {

	fullURL := fmt.Sprintf("%s%s", c.Endpoint, URI_TEAMS)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, err
	}

	var res = new(Teams)

	if err := c.sendRequest(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) GetTeam(id int) (*Team, error) {
	return c.GetTeamWithContext(context.Background(), id)
}

func (c *Client) GetTeamWithContext(ctx context.Context, id int) (*Team, error) {

	fullURL := fmt.Sprintf("%s%s/%d", c.Endpoint, URI_TEAMS, id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, err
	}

	var res = new(Team)

	if err := c.sendRequest(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) CreateTeam(team *TeamItem) (*Team, error) {
	return c.CreateTeamWithContext(context.Background(), team)
}

func (c *Client) CreateTeamWithContext(ctx context.Context, team *TeamItem) (*Team, error) {

	fullURL := fmt.Sprintf("%s%s", c.Endpoint, URI_TEAMS)

	body, err := json.Marshal(team)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fullURL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	var res = new(Team)

	if err := c.sendRequest(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) DeleteTeam(teamID int) error {
	return c.DeleteTeamWithContext(context.Background(), teamID)
}

func (c *Client) DeleteTeamWithContext(ctx context.Context, teamID int) error {

	fullURL := fmt.Sprintf("%s%s/%d", c.Endpoint, URI_TEAMS, teamID)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, fullURL, nil)
	if err != nil {
		return err
	}

	if err := c.sendRequest(req, nil); err != nil {
		return err
	}

	return nil
}

func NewTeam(name, description string) *TeamItem {
	return &TeamItem{
		Name:                name,
		Description:         description,
		Show:                "host",
		Theme:               "#7BB0B2",
		CanUseSysdigCapture: true,
		CanUseAgentCli:      true,
		CanUseCustomEvents:  true,
		CanUseAwsMetrics:    false,
		CanUseBeaconMetrics: false,
		CanUseRapidResponse: false,
		Users:               []string{},
	}
}
