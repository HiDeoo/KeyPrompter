package keyboard

import (
	"fmt"

	hook "github.com/robotn/gohook"
)

func HandleEvents(onEvent func(keyboardEvent KeyboardEvent)) {
	eventChannel := hook.Start()
	defer hook.End()

	modifiers := make(map[KeyboardModifier]hook.Event)

	for event := range eventChannel {
		if event.Kind == hook.KeyDown || event.Kind == hook.KeyHold || event.Kind == hook.KeyUp {
			switch event.Kind {
			case hook.KeyDown:
				if !isModifier(event) {
					keyboardEvent := newKeyboardEvent(event, modifiers)

					// TODO(HiDeoo)
					fmt.Printf("%+v\n", keyboardEvent)

					onEvent(keyboardEvent)
				}
				break
			case hook.KeyHold:
				if isModifier(event) {
					modifiers[event.Rawcode] = event
				}
				break
			case hook.KeyUp:
				if isModifier(event) {
					if _, ok := modifiers[event.Rawcode]; ok {
						delete(modifiers, event.Rawcode)
					}
				}
				break
			}
		}
	}
}
