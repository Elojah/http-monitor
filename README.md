# http-monitor
HTTP monitoring daemon

## How to start
```
> make dep #install dependencies
> make #compile monitor
> ./bin/monitor -c bin/config.json
```

## Configuration
```
{
    // log file
    "log_file": "/var/log/access.log",
    // tick interval to print stats in seconds.
    "stats_interval": 10,
    // Number of request per second required to trigger alert.
    "alert_req_per_sec":10,
    // time to consider to trigger an alert in seconds.
    "alert_trigger_time": 120,
    // time to consider after an alert to report in seconds.
    "alert_report_time": 120,
    // local/remote redis informations
    "redis": {
        "addr": "127.0.0.1:9851",
        "password": "secret",
        "db": 0
    }
}
```
