package mouse

import "testing"

func TestMoveMouse(t *testing.T) {
	MoveMouse(100, 100)
}

func TestClickMouse(t *testing.T) {
	if err := ClickMouse(); err != nil {
		t.Fatal(err)
	}
}
