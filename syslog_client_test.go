package main

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
		{"UDP", message.UDP},
		{"tcp", message.TCP},
	}

	for _, tt := range tests {

		protocol := setupProtocol(tt.protocol)
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

		protocol := setupProtocol(tt.protocol)
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

		addressPort := setupAddressPort(tt.address)
		if addressPort != tt.validAddressPort {
			t.Errorf("AddressPort function with a valid %s didn't return with valid %s", tt.address, tt.validAddressPort)
		}
	}
}

func TestCalculatePriotyValidArgs(t *testing.T) {

	priority := calculatePriority(3)
	if priority != 11 {
		t.Errorf("CalculatePriority function with severity of 3 didn't return priority of 11 (user error).")
	}
}

func TestCalculatePriotyInvalidArgs(t *testing.T) {

	priority := calculatePriority(20)
	if priority != 15 {
		t.Errorf("CalculatePriority function with severity of 20 didn't return priority of 15 (user debug).")
	}
}

func TestSetupClientValidArgs(t *testing.T) {

	sm := message.NewSyslogMessage("udp", "192.168.48.10:514", "", 0)

	conn, err := setupClient(sm)
	defer closeClient(conn)
	if err != nil {
		t.Errorf("SetupClient function didn't error with valid args.")
	}
}

func TestSetupClientUnknownAddress(t *testing.T) {

	sm := message.NewSyslogMessage("udp", ":514", "", 0)

	conn, err := setupClient(sm)
	defer closeClient(conn)
	if err != nil {
		t.Errorf("SetupClient function didn't error with unknown address.")
	}
}

func TestSendValidArgs(t *testing.T) {

	sm := message.NewSyslogMessage("udp", "192.168.48.10:514", "Testing", 3)

	conn, _ := setupClient(sm)
	defer closeClient(conn)
	err := send(sm, conn)
	if err != nil {
		t.Errorf("Send function didn't error with valid args.")
	}
}
