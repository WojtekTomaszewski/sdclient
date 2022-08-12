package sdclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type SilencingRules []SilencingRule

type SilencingRule struct {
	ID                     int    `json:"id,omitempty"`
	Version                int    `json:"version,omitempty"`
	CreatedOn              int64  `json:"createdOn,omitempty"`
	ModifiedOn             int64  `json:"modifiedOn,omitempty"`
	CustomerID             int    `json:"customerId,omitempty"`
	TeamID                 int    `json:"teamId,omitempty"`
	Name                   string `json:"name"`
	Enabled                bool   `json:"enabled"`
	StartTs                int64  `json:"startTs"`
	DurationInSec          int    `json:"durationInSec"`
	Scope                  string `json:"scope"`
	NotificationChannelIds []int  `json:"notificationChannelIds,omitempty"`
}

func (c *Client) ListSilencingRules() (*SilencingRules, error) {
	return c.ListSilencingRulesWithContext(context.Background())
}

func (c *Client) ListSilencingRulesWithContext(ctx context.Context) (*SilencingRules, error) {

	fullURL := fmt.Sprintf("%s%s", c.Endpoint, URI_SILENCERULES)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, err
	}

	var res = new(SilencingRules)

	if err := c.sendRequest(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) GetSilencingRule(id int) (*SilencingRule, error) {
	return c.GetSilencingRuleWithContext(context.Background(), id)
}

func (c *Client) GetSilencingRuleWithContext(ctx context.Context, id int) (*SilencingRule, error) {

	fullURL := fmt.Sprintf("%s%s/%d", c.Endpoint, URI_SILENCERULES, id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, err
	}

	var res = new(SilencingRule)

	if err := c.sendRequest(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) DeleteSilencingRule(id int) error {
	return c.DeleteSilencingRuleWithContext(context.Background(), id)
}

func (c *Client) DeleteSilencingRuleWithContext(ctx context.Context, id int) error {

	fullURL := fmt.Sprintf("%s%s/%d", c.Endpoint, URI_SILENCERULES, id)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, fullURL, nil)
	if err != nil {
		return err
	}

	if err := c.sendRequest(req, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) CreateSilencingRule(rule *SilencingRule) (*SilencingRule, error) {
	return c.CreateSilencingRuleWithContext(context.Background(), rule)
}

func (c *Client) CreateSilencingRuleWithContext(ctx context.Context, rule *SilencingRule) (*SilencingRule, error) {
	fullURL := fmt.Sprintf("%s%s", c.Endpoint, URI_SILENCERULES)

	body, err := json.Marshal(rule)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fullURL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	var res = new(SilencingRule)

	if err := c.sendRequest(req, res); err != nil {
		return nil, err
	}

	return res, nil
}
