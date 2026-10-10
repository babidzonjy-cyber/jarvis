package tools

import (
	"testing"
)

func TestOpenProjectZed(t *testing.T) {
	_, err := OpenProject("study", "zed")
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpenProjectVsCode(t *testing.T) {
	_, err := OpenProject("jarvis", "vscode")
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpenProjectCursor(t *testing.T) {
	_, err := OpenProject("jarvis", "cursor")
	if err != nil {
		t.Fatal(err)
	}
}
