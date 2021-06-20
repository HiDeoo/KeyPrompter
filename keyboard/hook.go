package keyboard

import (
	"github.com/google/go-cmp/cmp"
	hook "github.com/robotn/gohook"
)

type EventHandler = func(keyboardEvent KeyboardEvent)

func HandleEvents(onEvent EventHandler) {
	eventChannel := hook.Start()
	defer hook.End()

	modifiers := make(KeyboardModifiers)
	var lastKeyDownShortcutEvent KeyboardEvent

	for event := range eventChannel {
		if event.Kind == hook.KeyDown || event.Kind == hook.KeyHold || event.Kind == hook.KeyUp {
			switch event.Kind {
			case hook.KeyDown:
				if isShortcutEvent(event, modifiers) {
					lastKeyDownShortcutEvent = sendEvent(onEvent, event, modifiers)
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
				} else if isSpecialKeyEvent(event) {
					sendEvent(onEvent, event, modifiers)
				} else {
					keyboardEvent := newKeyboardEvent(event, modifiers)

					// Some specific basic shortcuts are still missing a key down event. If the event for this specific combo was
					// not sent yet, send it.
					if !cmp.Equal(lastKeyDownShortcutEvent, keyboardEvent) && isShortcutEvent(event, modifiers) {
						sendEvent(onEvent, event, modifiers)
					}
				}
			}
		}
	}
}

func sendEvent(send EventHandler, event hook.Event, modifiers KeyboardModifiers) KeyboardEvent {
	keyboardEvent := newKeyboardEvent(event, modifiers)

	send(keyboardEvent)

	return keyboardEvent
}
