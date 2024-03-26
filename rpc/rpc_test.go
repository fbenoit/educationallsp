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
		t.Fatalf("Expected %s, but got %s", expected, actuale)
	}
}
func TestDecode(t *testing.T) {
	incoming := "Content-length: 16\r\n\r\n{\"Testing\":true}"
	contentLength, err := rpc.DecodeMessage([]byte(incoming))
	if err != nil {
		t.Fatal(err)
	}

	if contentLength != 16 {
		t.Fatalf("Expected 16 got %d", contentLength)
	}
}
