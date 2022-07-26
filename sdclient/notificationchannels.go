package sdclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// NotificationChannel represents a notification channel request/response object
type NotificationChannel struct {
	NotificationChannel NotificationChannelItem `json:"notificationChannel"`
}

// NotificationChannels represents a list of notification channels request/response object
type NotificationChannels struct {
	NotificationChannels []NotificationChannelItem `json:"notificationChannels"`
}

// NotificationChannelItem represents a single notification channel
type NotificationChannelItem struct {
	ID                   int              `json:"id,omitempty"`
	Version              int              `json:"version,omitempty"`
	CreatedOn            int64            `json:"createdOn,omitempty"`
	ModifiedOn           int64            `json:"modifiedOn,omitempty"`
	TeamID               int              `json:"teamId,omitempty"`
	Type                 string           `json:"type,omitempty"`
	Enabled              bool             `json:"enabled,omitempty"`
	SendTestNotification bool             `json:"sendTestNotification,omitempty"`
	Name                 string           `json:"name,omitempty"`
	SettingsID           int              `json:"settingsId,omitempty"`
	Options              *json.RawMessage `json:"options,omitempty"`
}

// EmailNotificationChannel represents options for an email notification channel
type EmailNotificationChannelOptions struct {
	NotifyOnResolve bool     `json:"notifyOnResolve,omitempty"`
	NotifyOnOk      bool     `json:"notifyOnOk,omitempty"`
	EmailRecipients []string `json:"emailRecipients,omitempty"`
}

// PagerDutyNotificationChannelOptions represents options for a PagerDuty notification channel
type PagerDutyNotificationChannelOptions struct {
	NotifyOnResolve bool   `json:"notifyOnResolve,omitempty"`
	NotifyOnOk      bool   `json:"notifyOnOk,omitempty"`
	Account         string `json:"account,omitempty"`
	ServiceKey      string `json:"serviceKey,omitempty"`
	ServiceName     string `json:"serviceName,omitempty"`
}

// SlackNotificationChannelOptions represents options for a Slack notification channel
type SlackNotificationChannelOptions struct {
	NotifyOnResolve bool   `json:"notifyOnResolve,omitempty"`
	NotifyOnOk      bool   `json:"notifyOnOk,omitempty"`
	Channel         string `json:"channel,omitempty"`
	URL             string `json:"url,omitempty"`
}

// ListNotificationChannels returns a list of all notification channels
func (c *Client) ListNotificationChannels() (*NotificationChannels, error) {
	return c.ListNotificationChannelsWithContext(context.Background())
}

// ListNotificationChannelsWithContext returns a list of all notification channels
func (c *Client) ListNotificationChannelsWithContext(ctx context.Context) (*NotificationChannels, error) {

	fullURL := fmt.Sprintf("%s%s", c.Endpoint, URI_CHANNELS)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, err
	}

	var res = new(NotificationChannels)

	if err := c.sendRequest(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetNotificationChannel returns a single notification channel
func (c *Client) GetNotificationChannel(id int) (*NotificationChannel, error) {
	return c.GetNotificationChannelWithContext(context.Background(), id)
}

// GetNotificationChannelWithContext returns a single notification channel
func (c *Client) GetNotificationChannelWithContext(ctx context.Context, id int) (*NotificationChannel, error) {

	fullURL := fmt.Sprintf("%s%s/%d", c.Endpoint, URI_CHANNELS, id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, err
	}

	var res = new(NotificationChannel)

	if err := c.sendRequest(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreateEmailNotificationChannel creates an notification channel
func (c *Client) CreateNotificationChannel(channel *NotificationChannel) (*NotificationChannel, error) {
	return c.CreateNotificationChannelWithContext(context.Background(), channel)
}

// CreateEmailNotificationChannelWithContext creates an notification channel
func (c *Client) CreateNotificationChannelWithContext(ctx context.Context, channel *NotificationChannel) (*NotificationChannel, error) {

	fullURL := fmt.Sprintf("%s%s", c.Endpoint, URI_CHANNELS)

	body, err := json.Marshal(channel)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fullURL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	var res = new(NotificationChannel)

	if err := c.sendRequest(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// DeleteNotificationChannel deletes an notification channel
func (c *Client) DeleteNotificationChannel(id int) error {
	return c.DeleteNotificationChannelWithContext(context.Background(), id)
}

// DeleteNotificationChannel deletes an notification channel
func (c *Client) DeleteNotificationChannelWithContext(ctx context.Context, id int) error {

	fullURL := fmt.Sprintf("%s%s/%d", c.Endpoint, URI_CHANNELS, id)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, fullURL, nil)
	if err != nil {
		return err
	}

	if err := c.sendRequest(req, nil); err != nil {
		return err
	}

	return nil
}

// NewNotificationChannel creates a new notification channel of type channelType and provided options
func NewNotificationChannel(channelName, channelType string, options interface{}) (*NotificationChannel, error) {
	byteOptions, err := json.Marshal(options)
	if err != nil {
		return nil, err
	}

	rawOptions := json.RawMessage(byteOptions)

	return &NotificationChannel{
		NotificationChannel: NotificationChannelItem{
			Type:                 channelType,
			Name:                 channelName,
			SendTestNotification: false,
			Enabled:              true,
			Options:              &rawOptions,
		},
	}, nil
}
