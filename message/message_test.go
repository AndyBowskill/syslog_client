package message

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_NewSyslogMessage(t *testing.T) {

	expected := SyslogMessage{
		Protocol:    "tcp",
		AddressPort: "192.168.100.64",
		Message:     "Testing, testing.",
		Priority:    5,
	}

	actual := NewSyslogMessage("tcp", "192.168.100.64", "Testing, testing.", 5)

	if diff := cmp.Diff(&expected, actual); diff != "" {
		t.Error(diff)
	}
}
