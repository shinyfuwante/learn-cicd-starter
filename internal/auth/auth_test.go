package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	want := "152673 1"
	header := http.Header{}
	header.Add("Authorization", "ApiKey "+want)
	got, err := GetAPIKey(header)

	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
