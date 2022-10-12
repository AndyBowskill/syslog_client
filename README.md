## Syslog Client

[![Go build and test](https://github.com/andybowskill/syslog_client/actions/workflows/go.yml/badge.svg)](https://github.com/andybowskill/syslog_client/actions/workflows/go.yml)

A simple command line tool to check the syslog server is running by sending syslogs to it.

Acceptable command line flags are:

* `protocol` - Syslog server protocol, udp or tcp. Defaults to "udp".
* `address` - Syslog server address IPv4 address. Defaults to "".
* `severity` - Syslog message severity. Defaults to 5, (Notice).
* `message` - Syslog message. Defaults to "Testing, testing, 1, 2, 3".

### Example
```
syslog-client>go run ./ -network="udp" -address="172.16.30.2" -severity=0 -message="Testing the syslog server by sending alert message."
```
```
syslog-client>go run ./ -address="192.168.10.1" -message="Testing the syslog server by sending notice message."
```
```
syslog-client>go run ./ -protocol="tcp" -address="10.0.0.2" -severity=1
```