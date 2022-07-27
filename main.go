package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/WojtekTomaszewski/sdclient/sdclient"
)

func main() {
	sc := sdclient.New().WithAPIKey(os.Getenv("APIKEY")).WithEndpoint(os.Getenv("ENDPOINT"))

	alerts := &sdclient.Alerts{
		Alerts: []sdclient.AlertItem{
			{
				Name:                   "Test Alert",
				Description:            "This is a test alert",
				Enabled:                true,
				Type:                   sdclient.ALERT_TYPE_PROMETHEUS,
				Timespan:               900000000,
				Condition:              "(kube_workload_status_desired{kube_cluster_name=~\".*\",kube_namespace_name=~\"calico-system|istio-system.*|kube-system\",kube_workload_name=~\".*\"}\n- kube_workload_status_running{kube_cluster_name=~\".*\",kube_namespace_name=~\"calico-system|istio-system.*|kube-system\",kube_workload_name=~\".*\"})\n> 0\n",
				SeverityLabel:          sdclient.ALERT_SERVERITY_HIGH,
				NotificationChannelIds: []int{1721},
			},
			{
				Name:                   "Test Alert 2",
				Description:            "This is a test alert 2",
				Enabled:                true,
				Type:                   sdclient.ALERT_TYPE_MANUAL,
				Timespan:               900000000,
				Condition:              "avg(avg(sysdig_host_fs_inodes_used_percent)) > 90.0",
				SeverityLabel:          sdclient.ALERT_SERVERITY_LOW,
				Filter:                 "host.hostName contains \"c00lde2f06ub5uji6rhg\"",
				NotificationChannelIds: []int{1721},
				SegmentBy: []string{
					"agent.tag.cluster",
					"host.hostName",
				},
				SegmentCondition: &sdclient.SegmentConditionObject{
					Type: "ANY",
				},
			},
		},
	}

	a, err := sc.CreateAlerts(alerts)
	if err != nil {
		log.Fatal(err)
	}

	b, _ := json.MarshalIndent(a, "", "  ")

	fmt.Println(string(b))
}
