package message

const TCP string = "tcp"
const UDP string = "udp"

type SyslogMessage struct {
	Protocol    string
	AddressPort string
	Message     string
	Priority    uint8
}

func NewSyslogMessage(protocol, addressPort, message string, priority uint8) *SyslogMessage {
	syslogMessage := SyslogMessage{Protocol: protocol, AddressPort: addressPort, Message: message, Priority: priority}
	return &syslogMessage
}
