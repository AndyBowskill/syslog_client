package syslog_client

import (
	"testing"

	"github.com/AndyBowskill/syslog_client/message"
)

func TestSetupProtocolValidArgs(t *testing.T) {

	var tests = []struct {
		protocol      string
		validProtocol string
	}{
		{"udp", message.UDP},
		{"tcp", message.TCP},
	}

	for _, tt := range tests {

		protocol := SetupProtocol(tt.protocol)
		if protocol != tt.validProtocol {
			t.Errorf("SetupProtocol function with a %s didn't return with %s", tt.protocol, tt.validProtocol)
		}
	}
}

func TestSetupProtocolInvalidArgs(t *testing.T) {

	var tests = []struct {
		protocol      string
		validProtocol string
	}{
		{"", message.UDP},
		{"123 Test", message.UDP},
	}

	for _, tt := range tests {

		protocol := SetupProtocol(tt.protocol)
		if protocol != tt.validProtocol {
			t.Errorf("SetupProtocol function with a %s didn't return with %s", tt.protocol, tt.validProtocol)
		}
	}
}

func TestAddressPortValidArgs(t *testing.T) {

	var tests = []struct {
		address          string
		validAddressPort string
	}{
		{"172.16.20.3", "172.16.20.3:514"},
		{"", ":514"},
	}

	for _, tt := range tests {

		addressPort := SetupAddressPort(tt.address)
		if addressPort != tt.validAddressPort {
			t.Errorf("AddressPort function with a valid %s didn't return with valid %s", tt.address, tt.validAddressPort)
		}
	}
}

func TestCalculatePriotyValidArgs(t *testing.T) {

	priority := CalculatePriority(3)
	if priority != 11 {
		t.Errorf("CalculatePriority function with severity of 3 didn't return priority of 11 (user error).")
	}
}

func TestCalculatePriotyInvalidArgs(t *testing.T) {

	priority := CalculatePriority(20)
	if priority != 15 {
		t.Errorf("CalculatePriority function with severity of 20 didn't return priority of 15 (user debug).")
	}
}

func TestSetupClientValidArgs(t *testing.T) {

	syslogMessage := message.NewSyslogMessage("udp", "192.168.48.10:514", 0, "")

	conn, err := SetupClient(syslogMessage)
	defer CloseClient(conn)
	if err != nil {
		t.Errorf("SetupClient function didn't error with valid args.")
	}
}

func TestSetupClientUnknownAddress(t *testing.T) {

	syslogMessage := message.NewSyslogMessage("udp", ":514", 0, "")

	conn, err := SetupClient(syslogMessage)
	defer CloseClient(conn)
	if err != nil {
		t.Errorf("SetupClient function didn't error with unknown address.")
	}
}

func TestSendValidArgs(t *testing.T) {

	syslogMessage := message.NewSyslogMessage("udp", "192.168.48.10:514", 3, "Testing")

	conn, _ := SetupClient(syslogMessage)
	defer CloseClient(conn)
	err := Send(syslogMessage, conn)
	if err != nil {
		t.Errorf("Send function didn't error with valid args.")
	}
}
