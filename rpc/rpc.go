package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Content-length: %d\r\n\r\n%s", len(content), content)
}

func DecodeMessage(msg []byte) (int, error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return 0, errors.New("Header not found")
	}
	contentLenghtBytes := header[len("Content-length: "):]
	contentLength, err := strconv.Atoi(string(contentLenghtBytes))
	if err != nil {
		return 0, errors.New("There is no content")
	}
	_ = content
	return contentLength, nil
}
