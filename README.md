## Syslog Client

[![Go Build and Test](https://github.com/andybowskill/syslog_client/actions/workflows/go.yml/badge.svg)](https://github.com/andybowskill/syslog_client/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/andybowskill/syslog_client)](https://goreportcard.com/report/github.com/andybowskill/syslog_client)

A simple command line tool to check the syslog server is running by sending syslog messages to it.

Acceptable command line flags are:

* `p` - Syslog server protocol, udp or tcp. Defaults to "udp".
* `a` - Syslog server address IPv4 address. Defaults to "".
* `s` - Syslog message severity. Defaults to 5, (Notice).
* `m` - Syslog message. Defaults to "Testing, testing, 1, 2, 3".

### Example
```
syslog-client>go run ./ -p="udp" -a="172.16.30.2" -s=0 -m="Testing the syslog server by sending alert message."
```
```
syslog-client>go run ./ -a="192.168.10.1" -m="Testing the syslog server by sending notice message."
```
```
syslog-client>go run ./ -p="tcp" -a="10.0.0.2" -s=1
```