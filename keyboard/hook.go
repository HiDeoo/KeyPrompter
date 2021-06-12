package keyboard

import (
	"fmt"

	hook "github.com/robotn/gohook"
)

func HandleEvents(onEvent func()) {
	eventChannel := hook.Start()
	defer hook.End()

	for event := range eventChannel {
		if event.Kind == hook.KeyDown || event.Kind == hook.KeyHold {
			keychar := hook.RawcodetoKeychar(event.Rawcode)

			fmt.Println("hook: ", event)
			fmt.Println("keychar: ", keychar)
			fmt.Println("------------------------------------------------------------------------")

			onEvent()
		}
	}
}
