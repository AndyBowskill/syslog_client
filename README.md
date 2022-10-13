## Syslog Client

[![Go Build and Test](https://github.com/andybowskill/syslog_client/actions/workflows/go.yml/badge.svg)](https://github.com/andybowskill/syslog_client/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/andybowskill/syslog_client)](https://goreportcard.com/report/github.com/andybowskill/syslog_client)

A simple command line tool to check the syslog server is running by sending syslog messages to it.

Acceptable command line flags are:

| Flag | Description                                              |
|:-----|:---------------------------------------------------------|
| -p   | Syslog server protocol, udp or tcp. Defaults to "udp".   |
| -a   | Syslog server IPv4 address. Defaults to "".              |
| -s   | Syslog message severity. Defaults to 5, (Notice).        |
| -m   | Syslog message. Defaults to "Testing, testing, 1, 2, 3". |

### Install

```
$ go install github.com/AndyBowskill/syslog_client@latest
```

### Example
```
$ syslog_client -a="192.168.48.10" -m="Syslog message from Fedora" -s=4
```
```
$ syslog_client -p="tcp" -s=3 -a="192.168.48.10" -m="Syslog message from Ubuntu"
```