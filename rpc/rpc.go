package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Content-length: %d\r\n\r\n%s", len(content), content)
}

func DecodeMessage(msg []byte) error {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return errors.New("Header not found")
	}
	return nil
}
