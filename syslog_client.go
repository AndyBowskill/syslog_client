package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {

	networkPtr := flag.String("network", "udp", "syslog server udp or tcp")
	addressPtr := flag.String("address", "", "syslog server IP address")
	severityPtr := flag.Int("severity", 5, "syslog message severity")
	messagePtr := flag.String("message", "Testing, testing, 1, 2, 3", "syslog message")
	flag.Parse()

	addressPort := AddressPort(*addressPtr)
	priority := CalculatePriority(*severityPtr)

	conn, err := SetupClient(*networkPtr, addressPort)
	defer CloseClient(conn)
	if err == nil {
		Send(conn, priority, *messagePtr)
	}
}

func AddressPort(address string) (addressPort string) {

	addressPort = fmt.Sprintf("%s:514", address)
	return addressPort
}

func CalculatePriority(severity int) (priority int) {

	if severity > 7 {
		severity = 7
	}

	//Priority is user-level facility (1), add 8, then multplied by the severity
	priority = (8 + severity)
	return priority
}

func SetupClient(network, address string) (net.Conn, error) {

	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func Send(conn net.Conn, priority int, message string) error {

	_, err := fmt.Fprintf(conn, "<%d> %s", priority, message)
	if err != nil {
		return err
	}

	return nil
}

func CloseClient(conn net.Conn) {
	conn.Close()
}
