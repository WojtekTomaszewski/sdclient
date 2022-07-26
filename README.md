# Sysdig Monitor API wrapper to manage notification channels and alerts

## Setup

```bash
go get -u https://github.com/WojtekTomaszewski/sysdig-client/sdclient
```

## Examples

### Create Sysdig client

```go
sc := sdclient.New().WithEndpoint(os.Getenv("ENDPOINT")).WithAPIKey(os.Getenv("APIKEY"))
````

### List all alerts

```go
a, _ := sc.ListAlerts()
b, _ := json.MarshalIndent(a, "", "  ")
fmt.Println(string(b))
```

### Create Slack alert channel

```go
opt := sdclient.EmailNotificationChannelOptions{
		NotifyOnResolve: true,
		NotifyOnOk:      true,
		EmailRecipients: []string{"user@example.com"},
	}
	
ch, _ := sdclient.NewNotificationChannel("test", "EMAIL", opt)
```