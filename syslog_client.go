package syslog_client

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

	syslogMessage := message.NewSyslogMessage(protocol, addressPort, priority, *messagePtr)

	conn, err := SetupClient(syslogMessage)
	defer CloseClient(conn)
	if err == nil {
		Send(syslogMessage, conn)
	}
}

func SetupProtocol(protocol string) (validProtocol string) {

	if strings.Compare(protocol, message.TCP) != 0 && strings.Compare(protocol, message.UDP) != 0 {
		return message.UDP
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
