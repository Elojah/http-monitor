# http-monitor
HTTP monitoring daemon

## How to start
```
> make dep #install dependencies
> make #compile monitor
> docker-compose -d #start redis
> ./bin/monitor -c bin/config.json
```

## Configuration
```
{
    // log file
    "log_file": "bin/access.log",
    "alerter":{
        // Number of request per second required to trigger alert.
        "treshold":120,
        // time to consider to trigger an alert in seconds.
        "trigger_range": "10s",
        // minimum time between two alerts.
        "rebound_gap": "2s",
        // time between 2 alert checks
        "reccur_gap": "1s"
    },
    "log_reader": {
        // tick interval to print stats in seconds.
        "stats_gap": "5s",
        // number of top reauest hit to show in stats
        "top_display": 10
    },
    // local/remote redis informations
    "redis": {
        "addr": "127.0.0.1:6379",
        "password": "secret",
        "db": 0
    }
}
```
