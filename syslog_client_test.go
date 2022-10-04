package main

import (
	"testing"
)

func TestSetupClientValidArgs(t *testing.T) {

	conn, err := SetupClient("udp", "192.168.48.10:514")
	if err != nil {
		t.Errorf("SetupClient function didn't return with valid args.")
	}

	CloseClient(conn)

}

func TestSendRaw(t *testing.T) {

	conn, _ := SetupClient("udp", "192.168.48.10:514")
	defer CloseClient(conn)
	err := SendRaw(conn)
	if err != nil {
		t.Errorf("SendRaw function didn't return with valid args.")
	}
}
