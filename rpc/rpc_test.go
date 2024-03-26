package rpc_test

import (
	"educationlsp/rpc"
	"testing"
)

type EncodingExample struct {
	Testing bool
}

func TestEncode(t *testing.T) {
	expected := "Content-length: 16\r\n\r\n{\"Testing\":true}"
	actuale := rpc.EncodeMessage(EncodingExample{Testing: true})
	if expected != actuale {
		t.Errorf("Expected %s, but got %s", expected, actuale)
	}
}
