package sdclient

const (
	URI_CHANNELS  = "/api/notificationChannels"
	URI_TEAMS     = "/api/teams"
	URI_ALERTS    = "/api/alerts"
	URI_ALERTS_V2 = "/api/v2/alerts"

	ALERT_SERVERITY_LOW    = "low"
	ALERT_SERVERITY_MEDIUM = "medium"
	ALERT_SERVERITY_HIGH   = "high"
	ALERT_TYPE_PROMETHEUS  = "PROMETHEUS"
	ALERT_TYPE_MANUAL      = "MANUAL"
)

var Regions = map[string]string{
	"eu-de":    "https://eu-de.monitoring.cloud.ibm.com",
	"us-south": "https://us-south.monitoring.cloud.ibm.com",
	"us-east":  "https://us-east.monitoring.cloud.ibm.com",
}
