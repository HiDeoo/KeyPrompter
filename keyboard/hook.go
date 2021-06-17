package keyboard

import (
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
				if !isModifier(event) &&
					len(modifiers) > 0 &&
					!isShiftOnlyModifier(modifiers) &&
					!isSpecialKey(event) &&
					!isSpecialKeyWithModifiers(event) {
					sendEvent(onEvent, event, modifiers)
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
				} else {
					if isSpecialKey(event) {
						sendEvent(onEvent, event, modifiers)
					} else if isSpecialKeyWithModifiers(event) && len(modifiers) > 0 {
						sendEvent(onEvent, event, modifiers)
					}
				}
				break
			}
		}
	}
}

func sendEvent(fn func(keyboardEvent KeyboardEvent), event hook.Event, modifiers map[KeyboardModifier]hook.Event) {
	keyboardEvent := newKeyboardEvent(event, modifiers)

	fn(keyboardEvent)
}
