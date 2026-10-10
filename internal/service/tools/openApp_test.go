package tools

import (
	"testing"
)

func TestOpenAppSuccess(t *testing.T) {
	if err := OpenApp("Yandex Music"); err != nil {
		t.Fatal(err)
	}
}

func TestOpenAppNotSuccess(t *testing.T) {
	appName := "afaof"
	err := OpenApp(appName)
	if err == nil {
		t.Fatal("there is no such app, there must be an error")
	}
}
