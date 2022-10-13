package message

const TCP string = "tcp"
const UDP string = "udp"

type SyslogMessage struct {
	Protocol    string
	AddressPort string
	Priority    int
	Message     string
}

func NewSyslogMessage(protocol, addressPort string, priority int, message string) *SyslogMessage {
	syslogMessage := SyslogMessage{Protocol: protocol, AddressPort: addressPort, Priority: priority, Message: message}
	return &syslogMessage
}
