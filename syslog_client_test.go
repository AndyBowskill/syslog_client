package main

import (
	"testing"
)

func TestSetupClientValidArgs(t *testing.T) {

	conn, err := SetupClient("udp", "192.168.48.10:514")
	defer CloseClient(conn)
	if err != nil {
		t.Errorf("SetupClient function didn't error with valid args.")
	}
}

func TestSetupClientUnknownAddress(t *testing.T) {

	conn, err := SetupClient("udp", ":514")
	defer CloseClient(conn)
	if err != nil {
		t.Errorf("SetupClient function didn't error with unknown address.")
	}
}

func TestSendValidArgs(t *testing.T) {

	conn, _ := SetupClient("udp", "192.168.48.10:514")
	defer CloseClient(conn)
	err := Send(conn, 5, "Testing")
	if err != nil {
		t.Errorf("Send function didn't error with valid args.")
	}
}
