package sdclient

const (
	URI_CHANNELS     = "/api/notificationChannels"
	URI_TEAMS        = "/api/teams"
	URI_ALERTS       = "/api/alerts"
	URI_ALERTS_V2    = "/api/v2/alerts"
	URI_SILENCERULES = "/api/v1/silencingRules"

	ALERT_SERVERITY_LOW    = "low"
	ALERT_SERVERITY_MEDIUM = "medium"
	ALERT_SERVERITY_HIGH   = "high"
	ALERT_TYPE_PROMETHEUS  = "PROMETHEUS"
	ALERT_TYPE_MANUAL      = "MANUAL"
)

var Regions = map[string]string{
	"eu-de":    "https://eu-de.monitoring.cloud.ibm.com",
	"eu-gb":    "https://eu-gb.monitoring.cloud.ibm.com",
	"us-south": "https://us-south.monitoring.cloud.ibm.com",
	"us-east":  "https://us-east.monitoring.cloud.ibm.com",
	"ca-tor":   "https://ca-tor.monitoring.cloud.ibm.com",
	"au-syd":   "https://au-syd.monitoring.cloud.ibm.com",
	"jp-osa":   "https://jp-osa.monitoring.cloud.ibm.com",
	"jp-tok":   "https://jp-tok.monitoring.cloud.ibm.com",
	"br-sao":   "https://br-sao.monitoring.cloud.ibm.com",
}
