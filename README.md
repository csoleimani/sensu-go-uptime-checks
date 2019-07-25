# Sensu Go Uptime Checks
[![Bonsai Asset Badge](https://img.shields.io/badge/Sensu%20Go%20Uptime%20Checks-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/asachs01/sensu-go-uptime-checks) [![TravisCI Build Status](https://travis-ci.org/asachs01/sensu-go-uptime-checks.svg?branch=master)](https://travis-ci.org/asachs01/sensu-go-uptime-checks)

This plugin provides two separate checks for system uptime for Sensu Go.

* sensu-go-uptime-status
* sensu-go-uptime-metrics

The `sensu-go-uptime-status` check takes the flags `-w` (warning) and `-c` (critical) and a time duration after each flag (e.g., 24h). The time values can be represented in seconds(s), minutes(m), or hours(h). By default the plugin will use a warning value of 72h and a critical value of 168h (1 week). 

The `sensu-go-uptime-metrics` plugin outputs metric data in the `nagios_perfdata`format. See the example below for what a metric check using this plugin would look like.

## Installation

While it's generally recommended to use an asset, you can download a copy of the handler plugin from [releases][1],
or create an executable script from this source.

From the local path of the sensu-go-uptime-checks repository:

**sensu-go-uptime-status**
```
go build -o /usr/local/bin/sensu-go-uptime-status uptime_status_check/main.go
```

**sensu-go-uptime-metrics**
```
go build -o /usr/local/bin/sensu-go-uptime-metrics uptime_metric_check/main.go
```

## Configuration

### Asset Registration

Assets are the best way to make use of this check. If you're not using an asset, please consider doing so! You can find this asset on the [Bonsai Asset Index](https://bonsai.sensu.io/assets/asachs01/sensu-go-uptime-checks).

You can download the asset definition there, or you can do a little bit of copy/pasta and use the one below:

```yml
---
type: Asset
api_version: core/v2
metadata:
  name: sensu-go-uptime-checks
  namespace: CHANGEME
  labels: {}
  annotations: {}
spec:
  url: https://github.com/asachs01/sensu-go-uptime-checks/releases/download/0.0.1/sensu-go-uptime-checks_0.0.1_linux_amd64.tar.gz
  sha512: 
  filters:
  - entity.system.os == 'linux'
  - entity.system.arch == 'amd64'
```

**NOTE**: PLEASE ENSURE YOU UPDATE YOUR URL AND SHA512 BEFORE USING THE ASSET. If you don't, you might just be stuck on a super old version. Don't say I didn't warn you ¯\\_(ツ)_/¯

### Handler Configuration

Example Sensu Go definition:

**sensu-go-uptime-status**
```yml
type: CheckConfig
api_version: core/v2
metadata:
  name: sensu-go-uptime-status
  namespace: CHANGEME
spec:
  command: sensu-go-uptime-status
  runtime_assets:
  - sensu-go-uptime-checks
  interval: 60
  publish: true
  handlers:
  - slack
  subscriptions:
  - system
```

**sensu-go-uptime-metics**
```yml
type: CheckConfig
api_version: core/v2
metadata:
  name: sensu-go-uptime-metrics
  namespace: CHANGEME
spec:
  command: sensu-go-uptime-metrics
  runtime_assets:
  - sensu-go-uptime-checks
  interval: 60
  publish: true
  output_metric_format: nagios_perfdata
  output_metric_handlers:
  - influxdb
  subscriptions:
  - system
```

## Usage Examples

### Command line help

```
The Sensu Go check for system uptime

Usage:
  sensu-go-uptime-status [flags]

Flags:
  -w, --warning (time in s,m,h)   Warning value in seconds, minutes, or hours, default is 72 hours (72h)
  -c, --critical (time in s,m,h)   Warning value in seconds, minutes, or hours default is 1 week (168h)
  -h, --help         help for sensu-go-uptime-status
```

## Supported Operating Systems

This project uses `gopsutil`, and is thus largely dependent on the systems that it supports. For this plugin, the following operating systems are supported:

* Linux
* FreeBSD
* OpenBSD
* Mac OS X
* Windows (states as not supported, but I've confirmed that it is)
* Solaris

:)

## Contributing

See https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md

[1]: https://github.com/asachs01/sensu-go-uptime-checks/releases
