package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
)

const tcp string = "tcp"
const udp string = "udp"

type syslogMessage struct {
	protocol    string
	addressPort string
	priority    int
	message     string
}

func newSyslogMessage(protocol, addressPort string, priority int, message string) *syslogMessage {
	syslogMessage := syslogMessage{protocol: protocol, addressPort: addressPort, priority: priority, message: message}
	return &syslogMessage
}

func main() {

	protocolPtr := flag.String("protocol", udp, "syslog server protocol, udp or tcp")
	addressPtr := flag.String("address", "", "syslog server IP address")
	severityPtr := flag.Int("severity", 5, "syslog message severity")
	messagePtr := flag.String("message", "Testing, testing, 1, 2, 3", "syslog message")
	flag.Parse()

	protocol := SetupProtocol(*protocolPtr)
	addressPort := SetupAddressPort(*addressPtr)
	priority := CalculatePriority(*severityPtr)

	syslogMessage := newSyslogMessage(protocol, addressPort, priority, *messagePtr)

	conn, err := (*syslogMessage).SetupClient()
	defer CloseClient(conn)
	if err == nil {
		(*syslogMessage).Send(conn)
	}
}

func SetupProtocol(protocol string) (validProtocol string) {

	if strings.Compare(protocol, tcp) != 0 && strings.Compare(protocol, udp) != 0 {
		return udp
	}

	return protocol
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

func (sm *syslogMessage) SetupClient() (net.Conn, error) {

	conn, err := net.Dial(sm.protocol, sm.addressPort)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (sm *syslogMessage) Send(conn net.Conn) error {

	_, err := fmt.Fprintf(conn, "<%d> %s", sm.priority, sm.message)
	if err != nil {
		return err
	}

	return nil
}

func CloseClient(conn net.Conn) {
	conn.Close()
}
