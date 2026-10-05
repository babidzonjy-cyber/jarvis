package tools

import "testing"

func TestInspectYandexMusic(t *testing.T) {
	if err := InspectYandexMusic(); err != nil {
		t.Fatal(err)
	}
}
