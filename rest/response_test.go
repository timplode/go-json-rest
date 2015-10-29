package rest

import (
	"testing"
)

func TestResponseNotIndent(t *testing.T) {

	writer := responseWriter{
		nil,
		false,
		nil,
	}

	got, err := writer.EncodeJson(map[string]bool{"test": true})
	if err != nil {
		t.Error(err.Error())
	}
	gotStr := string(got)
	expected := "{\"test\":true}"
	if gotStr != expected {
		t.Error(expected + " was the expected, but instead got " + gotStr)
	}
}
