package message

const TCP string = "tcp"
const UDP string = "udp"

type SyslogMessage struct {
	Protocol    string
	AddressPort string
	Message     string
	Priority    uint
}

func NewSyslogMessage(protocol, addressPort, message string, priority uint) *SyslogMessage {
	syslogMessage := SyslogMessage{
		Protocol:    protocol,
		AddressPort: addressPort,
		Message:     message,
		Priority:    priority,
	}
	return &syslogMessage
}
