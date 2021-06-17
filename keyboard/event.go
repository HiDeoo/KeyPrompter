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
	Code      uint16            `json:"code"`
	Modifiers KeyboardModifiers `json:"modifiers"`
}

type KeyboardModifiers struct {
	Command bool `json:"command"`
	Control bool `json:"control"`
	Fn      bool `json:"fn"`
	Option  bool `json:"option"`
	Shift   bool `json:"shift"`
}

func newKeyboardEvent(event hook.Event, modifiers map[KeyboardModifier]hook.Event) KeyboardEvent {
	return KeyboardEvent{
		Character: hook.RawcodetoKeychar(event.Rawcode),
		Code:      event.Rawcode,
		Modifiers: newKeyboardModifiers(modifiers),
	}
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
