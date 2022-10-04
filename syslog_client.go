package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {

	networkPtr := flag.String("network", "udp", "syslog server udp or tcp")
	addressPtr := flag.String("address", "", "syslog server IP address")
	priorityPtr := flag.Int("priority", 5, "syslog message priority")
	messagePtr := flag.String("message", "Testing, testing, 1, 2, 3", "syslog message")
	flag.Parse()

	addressPort := fmt.Sprintf("%s:514", *addressPtr)

	conn, err := SetupClient(*networkPtr, addressPort)
	defer CloseClient(conn)
	if err == nil {
		Send(conn, *priorityPtr, *messagePtr)
	}
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
