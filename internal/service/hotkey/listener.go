package hotkey

import (
	"fmt"

	"golang.design/x/hotkey"
)

func ListenPushToTalk(onDown func() error, onUp func() error) error {
	hk := hotkey.New([]hotkey.Modifier{hotkey.ModCmd}, hotkey.KeyZ)

	if err := hk.Register(); err != nil {
		return fmt.Errorf("cannot register hotkey %w", err)
	}

	for {
		<-hk.Keydown()
		onDown()

		<-hk.Keyup()
		onUp()
	}
}

func ListenExit(onExit func()) error {
	hk := hotkey.New([]hotkey.Modifier{hotkey.ModOption}, hotkey.KeyQ)
	if err := hk.Register(); err != nil {
		return err
	}

	<-hk.Keydown()
	onExit()

	return nil
}
