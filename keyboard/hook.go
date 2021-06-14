package keyboard

import (
	"fmt"

	hook "github.com/robotn/gohook"
)

type KeyboardEvent struct {
	KeyChar string `json:"key_char"`
}

func HandleEvents(onEvent func(keyboardEvent KeyboardEvent)) {
	eventChannel := hook.Start()
	defer hook.End()

	for event := range eventChannel {
		if event.Kind == hook.KeyDown || event.Kind == hook.KeyHold {
			fmt.Println("hook: ", event)

			keychar := hook.RawcodetoKeychar(event.Rawcode)

			keyboardEvent := KeyboardEvent{KeyChar: keychar}

			onEvent(keyboardEvent)
		}
	}
}
