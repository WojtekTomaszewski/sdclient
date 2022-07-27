package sdclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type SingleAlert struct {
	Alert AlertItem `json:"alert,omitempty"`
}

type ListAlerts struct {
	Alerts      []AlertsItem `json:"alerts,omitempty"`
	AlertsCount int          `json:"alertsCount,omitempty"`
}

type AlertItem struct {
	ID                            int                              `json:"id,omitempty"`
	Name                          string                           `json:"name,omitempty"`
	Description                   string                           `json:"description,omitempty"`
	DurationSec                   int                              `json:"durationSec,omitempty"`
	Type                          string                           `json:"type,omitempty"`
	Group                         string                           `json:"group,omitempty"`
	Severity                      string                           `json:"severity,omitempty"`
	TeamID                        int                              `json:"teamId,omitempty"`
	Enabled                       bool                             `json:"enabled,omitempty"`
	NotificationChannelConfigList []NotificationChannelItem        `json:"notificationChannelConfigList,omitempty"`
	CustomNotificationTemplate    CustomNotificationTemplateObject `json:"customNotificationTemplate,omitempty"`
	Config                        *json.RawMessage                 `json:"config,omitempty"`
	Links                         []interface{}                    `json:"links,omitempty"`
	Version                       int                              `json:"version,omitempty"`
	CreatedOn                     int64                            `json:"createdOn,omitempty"`
	ModifiedOn                    int64                            `json:"modifiedOn,omitempty"`
}

type AlertsItem struct {
	ID                        int                       `json:"id,omitempty"`
	Condition                 string                    `json:"condition,omitempty"`
	CreatedOn                 int64                     `json:"createdOn,omitempty"`
	CustomNotification        *CustomNotificationObject `json:"customNotification,omitempty"`
	Description               string                    `json:"description,omitempty"`
	Enabled                   bool                      `json:"enabled,omitempty"`
	Filter                    string                    `json:"filter,omitempty"`
	ModifiedOn                int64                     `json:"modifiedOn,omitempty"`
	Name                      string                    `json:"name,omitempty"`
	NotificationChannels      []NotificationChannelItem `json:"notificationChannels,omitempty"`
	NotificationChannelIds    []int                     `json:"notificationChannelIds,omitempty"`
	EventsCount               []interface{}             `json:"eventsCount,omitempty"`
	NotificationCount         interface{}               `json:"notificationCount,omitempty"`
	ReNotify                  bool                      `json:"reNotify,omitempty"`
	Severity                  *json.RawMessage          `json:"severity,omitempty"`
	SeverityLabel             string                    `json:"severityLabel,omitempty"`
	SegmentBy                 []string                  `json:"segmentBy,omitempty"`
	SegmentCondition          *SegmentConditionOptions  `json:"segmentCondition,omitempty"`
	SysdigCapture             interface{}               `json:"sysdigCapture,omitempty"`
	TeamID                    int                       `json:"teamId,omitempty"`
	Timespan                  int                       `json:"timespan,omitempty"`
	Type                      string                    `json:"type,omitempty"`
	Valid                     bool                      `json:"valid,omitempty"`
	Version                   int                       `json:"version,omitempty"`
	Triggering                bool                      `json:"triggering,omitempty"`
	LastTriggeredNotification interface{}               `json:"lastTriggeredNotification,omitempty"`
	GroupName                 string                    `json:"groupName,omitempty"`
	AlertTemplateName         string                    `json:"alertTemplateName,omitempty"`
}

type SegmentConditionOptions struct {
	Type  string  `json:"type,omitempty"`
	Value float64 `json:"value,omitempty"`
}

type CustomNotificationObject struct {
	AppendText    string `json:"appendText,omitempty"`
	PrependText   string `json:"prependText,omitempty"`
	TitleTemplate string `json:"titleTemplate,omitempty"`
}

type CustomNotificationTemplateObject struct {
	AppendText  string `json:"appendText,omitempty"`
	PrependText string `json:"prependText,omitempty"`
	Subject     string `json:"subject,omitempty"`
}

type ManualAlertItemConfig struct {
	Metric struct {
		ID                string   `json:"id,omitempty"`
		PublicID          string   `json:"publicId,omitempty"`
		MetricType        string   `json:"metricType,omitempty"`
		Type              string   `json:"type,omitempty"`
		Scale             int      `json:"scale,omitempty"`
		GroupAggregations []string `json:"groupAggregations,omitempty"`
		TimeAggregations  []string `json:"timeAggregations,omitempty"`
	} `json:"metric,omitempty"`
	GroupAggregation string `json:"groupAggregation,omitempty"`
	TimeAggregation  string `json:"timeAggregation,omitempty"`
	Scope            struct {
		Expressions []interface{} `json:"expressions,omitempty"`
	} `json:"scope,omitempty"`
	Threshold         int    `json:"threshold,omitempty"`
	ConditionOperator string `json:"conditionOperator,omitempty"`
	SegmentBy         []struct {
		ID       string `json:"id,omitempty"`
		PublicID string `json:"publicId,omitempty"`
	} `json:"segmentBy,omitempty"`
	NotificationGroupingCondition struct {
		Type string `json:"type,omitempty"`
	} `json:"notificationGroupingCondition,omitempty"`
}

type PrometheusAlertItemConfig struct {
	Query string `json:"query,omitempty"`
}

// ListAlerts returns a list of alerts
func (c *Client) ListAlerts() (*ListAlerts, error) {
	return c.ListAlertsWithContext(context.Background())
}

// ListAlertsWithContext returns a list of alerts
func (c *Client) ListAlertsWithContext(ctx context.Context) (*ListAlerts, error) {

	fullURL := fmt.Sprintf("%s%s", c.Endpoint, URI_ALERTS)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, err
	}

	var res = new(ListAlerts)

	if err := c.sendRequest(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetAlert returns an alert by ID
func (c *Client) GetAlert(id int) (*SingleAlert, error) {
	return c.GetAlertWithContext(context.Background(), id)
}

// GetAlertWithContext returns an alert by ID
func (c *Client) GetAlertWithContext(ctx context.Context, id int) (*SingleAlert, error) {

	fullURL := fmt.Sprintf("%s%s/%d", c.Endpoint, URI_ALERTS, id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, err
	}

	var res = new(SingleAlert)

	if err := c.sendRequest(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// BulkCreateAlerts creates multiple alerts in one request.
// https://app.sysdigcloud.com/api/public/docs/index.html#operation/bulkCreateAlertUsingPOST_1
func (c *Client) BulkCreateAlerts(alerts *ListAlerts) (*ListAlerts, error) {
	return c.BulkCreateAlertsWithContext(context.Background(), alerts)
}

// BulkCreateAlertsWithContext creates multiple alerts in one request.
// https://app.sysdigcloud.com/api/public/docs/index.html#operation/bulkCreateAlertUsingPOST_1
func (c *Client) BulkCreateAlertsWithContext(ctx context.Context, alerts *ListAlerts) (*ListAlerts, error) {

	fullURL := fmt.Sprintf("%s%s", c.Endpoint, URI_ALERTS)

	byteBody, err := json.MarshalIndent(alerts, "", "  ")
	if err != nil {
		return nil, err
	}

	fmt.Println(string(byteBody))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fullURL, bytes.NewReader(byteBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	var res = new(ListAlerts)

	if err := c.sendRequest(req, res); err != nil {
		return nil, err
	}

	return res, nil
}
