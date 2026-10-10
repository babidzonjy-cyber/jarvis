package mouse

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func MoveMouse(x, y int) {
	robotgo.MouseSleep = 100
	robotgo.Move(x, y)
}

func ClickMouse() error {
	if err := robotgo.Click(); err != nil {
		return fmt.Errorf("click mouse: %w", err)
	}

	return nil
}

func ClickAt(x, y int) error {
	robotgo.Move(x, y)
	if err := robotgo.Click(); err != nil {
		return fmt.Errorf("cannot click: %w", err)
	}
	return nil
}
