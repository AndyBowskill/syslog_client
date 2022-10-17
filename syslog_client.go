package main

import (
	"flag"
	"fmt"
	"net"
	"strings"

	"github.com/AndyBowskill/syslog_client/message"
)

func main() {

	protocolPtr := flag.String("p", message.UDP, "syslog server protocol, udp or tcp")
	addressPtr := flag.String("a", "", "syslog server IP address")
	severityPtr := flag.Int("s", 5, "syslog message severity")
	messagePtr := flag.String("m", "Testing, testing, 1, 2, 3", "syslog message")
	flag.Parse()

	protocol := SetupProtocol(*protocolPtr)
	addressPort := SetupAddressPort(*addressPtr)
	priority := CalculatePriority(*severityPtr)

	sm := message.NewSyslogMessage(protocol, addressPort, *messagePtr, priority)

	conn, err := SetupClient(sm)
	defer CloseClient(conn)
	if err == nil {
		Send(sm, conn)
	}
}

func SetupProtocol(protocol string) (validProtocol string) {

	lowerProtocol := strings.ToLower(protocol)

	if strings.Compare(lowerProtocol, message.TCP) != 0 && strings.Compare(lowerProtocol, message.UDP) != 0 {
		return message.UDP
	}

	return lowerProtocol
}

func SetupAddressPort(address string) (addressPort string) {

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

func SetupClient(sm *message.SyslogMessage) (net.Conn, error) {

	conn, err := net.Dial(sm.Protocol, sm.AddressPort)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func Send(sm *message.SyslogMessage, conn net.Conn) error {

	_, err := fmt.Fprintf(conn, "<%d> %s", sm.Priority, sm.Message)
	if err != nil {
		return err
	}

	return nil
}

func CloseClient(conn net.Conn) {
	conn.Close()
}
