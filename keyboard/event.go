package keyboard

import (
	hook "github.com/robotn/gohook"
)

type KeyboardEvent struct {
	Character string                 `json:"character"`
	Code      uint16                 `json:"code"`
	Modifiers KeyboardEventModifiers `json:"modifiers"`
}

type KeyboardEventModifiers struct {
	Command bool `json:"command"`
	Control bool `json:"control"`
	Fn      bool `json:"fn"`
	Option  bool `json:"option"`
	Shift   bool `json:"shift"`
}

func newKeyboardEvent(event hook.Event, modifiers KeyboardModifiers) KeyboardEvent {
	var character string

	if char, ok := rawCodeMap[event.Rawcode]; ok {
		character = char
	} else {
		character = hook.RawcodetoKeychar(event.Rawcode)
	}

	keyboardEvent := KeyboardEvent{
		Character: character,
		Code:      event.Rawcode,
		Modifiers: newKeyboardModifiers(modifiers),
	}

	if isSpecialKeyEvent(event) {
		keyboardEvent.Character = SpecialKeyMap[event.Rawcode]
	}

	return keyboardEvent
}

func newKeyboardModifiers(modifiers KeyboardModifiers) KeyboardEventModifiers {
	containsCommand := false
	containsControl := false
	containsFn := false
	containsOption := false
	containsShift := false

	for _, modifier := range modifiers {
		if isCommandModifier(modifier) {
			containsCommand = true
		} else if isControlModifier(modifier) {
			containsControl = true
		} else if isFnModifier(modifier) {
			containsFn = true
		} else if isOptionModifier(modifier) {
			containsOption = true
		} else if isShiftModifier(modifier) {
			containsShift = true
		}
	}

	return KeyboardEventModifiers{
		Command: containsCommand,
		Control: containsControl,
		Fn:      containsFn,
		Option:  containsOption,
		Shift:   containsShift,
	}
}

func isShortcutEvent(event hook.Event, modifiers KeyboardModifiers) bool {
	return !isModifier(event) && len(modifiers) > 0 && !isShiftOnlyModifier(modifiers) && !isSpecialKeyEvent(event)
}

func isSpecialKeyEvent(event hook.Event) bool {
	if _, ok := SpecialKeyMap[event.Rawcode]; ok {
		return true
	}

	return false
}
