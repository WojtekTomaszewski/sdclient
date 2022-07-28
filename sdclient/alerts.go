package sdclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Alert struct {
	Alert AlertItem `json:"alert,omitempty"`
}

type Alerts struct {
	Alerts []AlertItem `json:"alerts,omitempty"`
}

type AlertItem struct {
	ID                     int                       `json:"id,omitempty"`
	Version                int                       `json:"version,omitempty"`
	CreatedOn              int64                     `json:"createdOn,omitempty"`
	ModifiedOn             int64                     `json:"modifiedOn,omitempty"`
	Type                   string                    `json:"type,omitempty"`
	Name                   string                    `json:"name,omitempty"`
	Description            string                    `json:"description,omitempty"`
	Enabled                bool                      `json:"enabled,omitempty"`
	NotificationChannelIds []int                     `json:"notificationChannelIds,omitempty"`
	Filter                 string                    `json:"filter,omitempty"`
	Severity               int                       `json:"severity,omitempty"`
	Timespan               int                       `json:"timespan,omitempty"`
	CustomNotification     *CustomNotificationObject `json:"customNotification,omitempty"`
	NotificationCount      int                       `json:"notificationCount,omitempty"`
	TeamID                 int                       `json:"teamId,omitempty"`
	AutoCreated            bool                      `json:"autoCreated,omitempty"`
	SysdigCapture          *SysdigCaptureObject      `json:"sysdigCapture,omitempty"`
	RateOfChange           bool                      `json:"rateOfChange,omitempty"`
	ReNotifyMinutes        int                       `json:"reNotifyMinutes,omitempty"`
	ReNotify               bool                      `json:"reNotify,omitempty"`
	InvalidMetrics         []string                  `json:"invalidMetrics,omitempty"`
	GroupName              string                    `json:"groupName,omitempty"`
	AlertTemplateName      string                    `json:"alertTemplateName,omitempty"`
	AlertTemplateVersion   int                       `json:"alertTemplateVersion,omitempty"`
	Links                  []string                  `json:"links,omitempty"`
	Valid                  bool                      `json:"valid,omitempty"`
	SeverityLabel          string                    `json:"severityLabel,omitempty"`
	SegmentBy              []string                  `json:"segmentBy,omitempty"`
	SegmentCondition       *SegmentConditionObject   `json:"segmentCondition,omitempty"`
	Condition              string                    `json:"condition,omitempty"`
	CustomerID             int                       `json:"customerId,omitempty"`
	LastCheckTimeInMs      int64                     `json:"lastCheckTimeInMs,omitempty"`
}

type CustomNotificationObject struct {
	AppendText     string `json:"appendText,omitempty"`
	PrependText    string `json:"prependText,omitempty"`
	TitleTemplate  string `json:"titleTemplate,omitempty"`
	UseNewTemplate bool   `json:"useNewTemplate,omitempty"`
	Subject        string `json:"subject,omitempty"`
}

type SysdigCaptureObject struct {
	Name       string      `json:"name,omitempty"`
	Filters    string      `json:"filters,omitempty"`
	Duration   int         `json:"duration,omitempty"`
	Type       string      `json:"type,omitempty"`
	BucketName interface{} `json:"bucketName,omitempty"`
	Folder     string      `json:"folder,omitempty"`
	Enabled    bool        `json:"enabled,omitempty"`
	StorageID  interface{} `json:"storageId,omitempty"`
}

type SegmentConditionObject struct {
	Type  string  `json:"type,omitempty"`
	Value float64 `json:"value,omitempty"`
}

// ListAlerts returns a list of alerts
func (c *Client) ListAlerts() (*Alerts, error) {
	return c.ListAlertsWithContext(context.Background())
}

// ListAlertsWithContext returns a list of alerts
func (c *Client) ListAlertsWithContext(ctx context.Context) (*Alerts, error) {

	fullURL := fmt.Sprintf("%s%s", c.Endpoint, URI_ALERTS)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, err
	}

	var res = new(Alerts)

	if err := c.sendRequest(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetAlert returns an alert by ID
func (c *Client) GetAlert(id int) (*Alert, error) {
	return c.GetAlertWithContext(context.Background(), id)
}

// GetAlertWithContext returns an alert by ID
func (c *Client) GetAlertWithContext(ctx context.Context, id int) (*Alert, error) {

	fullURL := fmt.Sprintf("%s%s/%d", c.Endpoint, URI_ALERTS, id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, err
	}

	var res = new(Alert)

	if err := c.sendRequest(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreateAlerts creates a new alerts from provided alerts object
func (c *Client) CreateAlerts(alerts *Alerts) (*Alerts, error) {
	return c.CreateAlertsWithContext(context.Background(), alerts)
}

// CreateAlertsWithContext creates a new alerts from provided alerts object
func (c *Client) CreateAlertsWithContext(ctx context.Context, alerts *Alerts) (*Alerts, error) {

	fullURL := fmt.Sprintf("%s%s", c.Endpoint, URI_ALERTS_V2)

	byteBody, err := json.MarshalIndent(alerts, "", "  ")
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fullURL, bytes.NewReader(byteBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	var res = new(Alerts)

	if err := c.sendRequest(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateAlerts updates an alerts from provided alerts object
func (c *Client) UpdateAlerts(alerts *Alerts) (*Alerts, error) {
	return c.UpdateAlertsWithContext(context.Background(), alerts)
}

// UpdateAlertsWithContext creates a new alerts from provided alerts object
func (c *Client) UpdateAlertsWithContext(ctx context.Context, alerts *Alerts) (*Alerts, error) {

	fullURL := fmt.Sprintf("%s%s", c.Endpoint, URI_ALERTS_V2)

	byteBody, err := json.MarshalIndent(alerts, "", "  ")
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, fullURL, bytes.NewReader(byteBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	var res = new(Alerts)

	if err := c.sendRequest(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreateAlert creates a new alert from provided alert object
func (c *Client) CreateAlert(alert *Alert) (*Alert, error) {
	return c.CreateAlertWithContext(context.Background(), alert)
}

// CreateAlertWithContext creates a new alerts from provided alerts object
func (c *Client) CreateAlertWithContext(ctx context.Context, alert *Alert) (*Alert, error) {

	fullURL := fmt.Sprintf("%s%s", c.Endpoint, URI_ALERTS)

	byteBody, err := json.MarshalIndent(alert, "", "  ")
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fullURL, bytes.NewReader(byteBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	var res = new(Alert)

	if err := c.sendRequest(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateAlert updates an alert
func (c *Client) UpdateAlert(alert *Alert) (*Alert, error) {
	return c.UpdateAlertWithContext(context.Background(), alert)
}

// UpdateAlertWithContext updates an alert
func (c *Client) UpdateAlertWithContext(ctx context.Context, alert *Alert) (*Alert, error) {

	fullURL := fmt.Sprintf("%s%s", c.Endpoint, URI_ALERTS)

	byteBody, err := json.MarshalIndent(alert, "", "  ")
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, fullURL, bytes.NewReader(byteBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	var res = new(Alert)

	if err := c.sendRequest(req, res); err != nil {
		return nil, err
	}

	return res, nil
}
