package tools

import "testing"

func TestSearchWeb(t *testing.T) {
	_, err := SearchWeb("как работает интернет")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSearchYoutube(t *testing.T) {
	_, err := SearchWeb("как работает интернет")
	if err != nil {
		t.Fatal(err)
	}
}
