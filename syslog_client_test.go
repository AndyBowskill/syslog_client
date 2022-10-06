package main

import (
	"testing"
)

func TestAddressPortValidArgs(t *testing.T) {

	validAddressPort := "172.16.20.3:514"
	addressPort := AddressPort("172.16.20.3")
	if addressPort != validAddressPort {
		t.Errorf("AddressPort function with a valid adddress didn't return with valid addess and port")
	}
}

func TestCalculatePriotyValidArgs(t *testing.T) {

	priority := CalculatePriority(3)
	if priority != 11 {
		t.Errorf("CaulatePriority function with severity of 3 didn't return priority of 11.")
	}
}

func TestCalculatePriotyInvalidArgs(t *testing.T) {

	priority := CalculatePriority(20)
	if priority != 15 {
		t.Errorf("CaulatePriority function with severity of 20 didn't return priority of 15 (user debug).")
	}
}

func TestSetupClientValidArgs(t *testing.T) {

	sm := newSyslogMessage("udp", "192.168.48.10:514", 0, "")

	conn, err := sm.SetupClient()
	defer CloseClient(conn)
	if err != nil {
		t.Errorf("SetupClient function didn't error with valid args.")
	}
}

func TestSetupClientUnknownAddress(t *testing.T) {

	sm := newSyslogMessage("udp", ":514", 0, "")

	conn, err := sm.SetupClient()
	defer CloseClient(conn)
	if err != nil {
		t.Errorf("SetupClient function didn't error with unknown address.")
	}
}

func TestSendValidArgs(t *testing.T) {

	sm := newSyslogMessage("udp", "192.168.48.10:514", 3, "Testing")

	conn, _ := sm.SetupClient()
	defer CloseClient(conn)
	err := sm.Send(conn)
	if err != nil {
		t.Errorf("Send function didn't error with valid args.")
	}
}
