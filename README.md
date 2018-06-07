# http-monitor
HTTP monitoring daemon

## How to start
### local
```
> make dep #install dependencies
> make #compile monitor
> docker-compose -d redis #start redis
> ./bin/monitor -c bin/config.local.json
```
### container
```
> docker-compose -d #start redis and monitor
```
docker `http_monitor_app` need the permission to mount a volume on /var/log/access.log
```
    volumes:
      - /var/log/access.log:/var/log/access.log
```
If you want to track another file, you will need to change in `log_file` configuration AND in this volume. e.g:
```
    volumes:
      - bin/access.log:/var/log/access.log
```
```
    ...
    "log_reader": {
        ...
        "log_file": "bin/access.log",
        ...
```

## Configuration
```
{
    "alerter":{
        // Number of request per second required to trigger alert.
        "treshold": 1200,
        // time to consider to trigger an alert in seconds.
        "trigger_range": "2m",
        // time to consider to trigger a recover.
        "trigger_recover": "2m",
        // minimum time between two alerts.
        "rebound_gap": "2m",
        // time between 2 alert checks
        "reccur_gap": "1s"
    },
    "log_reader": {
        // log file
        "log_file": "bin/access.log",
        // tick interval to print stats in seconds.
        "stats_gap": "5s",
        // number of top request hit to show in stats
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
