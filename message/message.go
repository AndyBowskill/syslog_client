package message

const TCP string = "tcp"
const UDP string = "udp"

type SyslogMessage struct {
	Protocol    string
	AddressPort string
	Message     string
	Priority    int
}

func NewSyslogMessage(protocol, addressPort, message string, priority int) *SyslogMessage {
	syslogMessage := SyslogMessage{Protocol: protocol, AddressPort: addressPort, Message: message, Priority: priority}
	return &syslogMessage
}
