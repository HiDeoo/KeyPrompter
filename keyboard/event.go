package keyboard

import (
	hook "github.com/robotn/gohook"
)

type KeyboardModifier = uint16

const (
	CommandLeft  KeyboardModifier = 54
	CommandRight                  = 55
	ShiftLeft                     = 56
	OptionLeft                    = 58
	ControlLeft                   = 59
	ShiftRight                    = 60
	OptionRight                   = 61
	ControlRight                  = 62
	Fn                            = 63
)

type KeyboardEvent struct {
	Character string            `json:"character"`
	Code      KeyboardModifier  `json:"code"`
	Modifiers KeyboardModifiers `json:"modifiers"`
}

type KeyboardModifiers struct {
	Command bool `json:"command"`
	Control bool `json:"control"`
	Fn      bool `json:"fn"`
	Option  bool `json:"option"`
	Shift   bool `json:"shift"`
}

var SpecialKeyMap = map[KeyboardModifier]string{
	36:  "↩",   // Return
	51:  "⌫",   // Delete
	53:  "⎋",   // Escape
	64:  "F17", // F17
	71:  "⌧",   // Numpad Clear
	76:  "↩",   // Enter
	79:  "F18", // F18
	80:  "F19", // F19
	90:  "F20", // F20
	96:  "F5",  // F5
	97:  "F6",  // F6
	98:  "F7",  // F7
	99:  "F3",  // F3
	100: "F8",  // F8
	101: "F9",  // F9
	103: "F11", // F11
	105: "F13", // F13
	106: "F16", // F16
	107: "F14", // F14
	109: "F10", // F10
	111: "F12", // F12
	113: "F15", // F15
	115: "↖",   // Home
	116: "⇞",   // Page Up
	117: "⌦",   // Delete Forward
	118: "F4",  // F4
	119: "↘",   // End
	120: "F2",  // F2
	121: "⇟",   // Page Down
	122: "F1",  // F1
	123: "←",   // Arrow Left
	124: "→",   // Arrow Right
	125: "↓",   // Arrow Down
	126: "↑",   // Arrow Up
}

var SpecialKeyWithModifiersMap = map[KeyboardModifier]string{
	30: "]",  // Bracket Right
	33: "[",  // Bracket Left
	42: "\\", // Backslash
}

func newKeyboardEvent(event hook.Event, modifiers map[KeyboardModifier]hook.Event) KeyboardEvent {
	keyboardEvent := KeyboardEvent{
		Character: hook.RawcodetoKeychar(event.Rawcode),
		Code:      event.Rawcode,
		Modifiers: newKeyboardModifiers(modifiers),
	}

	if isSpecialKey(event) {
		keyboardEvent.Character = SpecialKeyMap[event.Rawcode]
	} else if isSpecialKeyWithModifiers(event) {
		keyboardEvent.Character = SpecialKeyWithModifiersMap[event.Rawcode]
	}

	return keyboardEvent
}

func newKeyboardModifiers(modifiers map[KeyboardModifier]hook.Event) KeyboardModifiers {
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

	return KeyboardModifiers{
		Command: containsCommand,
		Control: containsControl,
		Fn:      containsFn,
		Option:  containsOption,
		Shift:   containsShift,
	}
}

func isModifier(event hook.Event) bool {
	return isCommandModifier(event) ||
		isControlModifier(event) ||
		isFnModifier(event) ||
		isOptionModifier(event) ||
		isShiftModifier(event)
}

func isShiftOnlyModifier(modifiers map[KeyboardModifier]hook.Event) bool {
	if len(modifiers) > 1 {
		return false
	}

	for _, modifier := range modifiers {
		return isShiftModifier(modifier)
	}

	return false
}

func isSpecialKey(event hook.Event) bool {
	if _, ok := SpecialKeyMap[event.Rawcode]; ok {
		return true
	}

	return false
}

func isSpecialKeyWithModifiers(event hook.Event) bool {
	if _, ok := SpecialKeyWithModifiersMap[event.Rawcode]; ok {
		return true
	}

	return false
}

func isCommandModifier(event hook.Event) bool {
	return event.Rawcode == CommandLeft || event.Rawcode == CommandRight
}

func isControlModifier(event hook.Event) bool {
	return event.Rawcode == ControlLeft || event.Rawcode == ControlRight
}

func isFnModifier(event hook.Event) bool {
	return event.Rawcode == Fn
}

func isOptionModifier(event hook.Event) bool {
	return event.Rawcode == OptionLeft || event.Rawcode == OptionRight
}

func isShiftModifier(event hook.Event) bool {
	return event.Rawcode == ShiftLeft || event.Rawcode == ShiftRight
}
