package main

import (
	"testing"

	"github.com/AndyBowskill/syslog_client/message"
)

func Test_setupProtocol(t *testing.T) {

	var tests = []struct {
		protocol      string
		validProtocol string
	}{
		{"udp", message.UDP},
		{"UDP", message.UDP},
		{"tcp", message.TCP},
		{"", message.UDP},
		{"123 Test", message.UDP},
	}

	for _, tt := range tests {
		protocol := setupProtocol(tt.protocol)
		if protocol != tt.validProtocol {
			t.Errorf("setupProtocol function with a %s didn't return with %s", tt.protocol, tt.validProtocol)
		}
	}
}

func Test_setupAddressPort(t *testing.T) {

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
			t.Errorf("addressPort function with a valid %s didn't return with valid %s", tt.address, tt.validAddressPort)
		}
	}
}

func Test_calculatePrioty(t *testing.T) {

	var tests = []struct {
		severity      uint
		validPriority uint
	}{
		{3, 11},
		{20, 15},
		{0, 8},
	}

	for _, tt := range tests {
		priority := calculatePriority(tt.severity)
		if priority != tt.validPriority {
			t.Errorf("calculatePriority function with a %d severity didn't return with %d priority", tt.severity, tt.validPriority)
		}
	}
}

func Test_setupClientValidArgs(t *testing.T) {

	sm := message.NewSyslogMessage("udp", "192.168.48.10:514", "", 0)

	conn, err := setupClient(sm)
	defer closeClient(conn)
	if err != nil {
		t.Errorf("setupClient function did error with valid args.")
	}
}

func Test_setupClientUnknownAddress(t *testing.T) {

	sm := message.NewSyslogMessage("udp", ":514", "", 0)

	conn, err := setupClient(sm)
	defer closeClient(conn)
	if err != nil {
		t.Errorf("setupClient function didn't error with unknown address.")
	}
}

func Test_setupClientInvalidArgs(t *testing.T) {

	sm := message.NewSyslogMessage("", "", "", 0)

	_, err := setupClient(sm)
	if err == nil {
		t.Errorf("setupClient function didn't error with invalid args.")
	}
}

type TestWriter struct {
}

func (tw TestWriter) Write(p []byte) (n int, err error) {
	n = 19
	err = nil
	return n, err
}

func Test_send(t *testing.T) {

	tw := TestWriter{}

	sm := message.NewSyslogMessage("udp", "192.168.48.10:514", "Error - Testing", 3)

	err := send(sm, tw)
	if err != nil {
		t.Errorf("send function did error with valid args.")
	}
}
