package keyboard

import hook "github.com/robotn/gohook"

type KeyboardModifier = uint16
type KeyboardModifiers = map[KeyboardModifier]hook.Event

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

// We are not interested in localized keys that could come up with various modifiers and the table used in
// github.com/robotn/gohook is broken on macOS.
// https://github.com/robotn/gohook/issues/14
// https://eastmanreference.com/complete-list-of-applescript-key-codes
var rawCodeMap = map[KeyboardModifier]string{
	10: "§",
	50: "`",
	18: "1",
	19: "2",
	20: "3",
	21: "4",
	23: "5",
	22: "6",
	26: "7",
	28: "8",
	25: "9",
	29: "0",
	27: "-",
	24: "=",
	48: "⇥", // Tab
	12: "q",
	13: "w",
	14: "e",
	15: "r",
	17: "t",
	16: "y",
	32: "u",
	34: "i",
	31: "o",
	35: "p",
	33: "[",
	30: "]",
	42: "\\",
	0:  "a",
	1:  "s",
	2:  "d",
	3:  "f",
	5:  "g",
	4:  "h",
	38: "j",
	40: "k",
	37: "l",
	41: ";",
	39: "'",
	6:  "z",
	7:  "x",
	8:  "c",
	9:  "v",
	11: "b",
	45: "n",
	46: "m",
	43: ",",
	47: ".",
	44: "/",
	49: "␣",
	75: "/", // Numpad /
	67: "*", // Numpad *
	78: "-", // Numpad -
	89: "7", // Numpad 7
	91: "8", // Numpad 8
	92: "9", // Numpad 9
	86: "4", // Numpad 4
	87: "5", // Numpad 5
	88: "6", // Numpad 6
	69: "+", // Numpad +
	83: "1", // Numpad 1
	84: "2", // Numpad 2
	85: "3", // Numpad 3
	82: "0", // Numpad 0
	65: ".", // Numpad .
}

// Some keys may not trigger a key down event (and can only be detected on a key up event) or may not require a
// modifier.
// https://github.com/robotn/gohook/issues/26
// http://macbiblioblog.blogspot.com/2014/12/key-codes-for-function-and-special-keys.html
var SpecialKeyMap = map[KeyboardModifier]string{
	53:  "⎋", // Escape
	122: "F1",
	120: "F2",
	99:  "F3",
	118: "F4",
	96:  "F5",
	97:  "F6",
	98:  "F7",
	100: "F8",
	101: "F9",
	109: "F10",
	103: "F11",
	111: "F12",
	105: "F13",
	107: "F14",
	113: "F15",
	106: "F16",
	64:  "F17",
	79:  "F18",
	80:  "F19",
	90:  "F20",
	51:  "⌫",  // Delete
	36:  "↩",  // Return
	76:  "↩",  // Enter
	114: "?⃝", // Help or Insert
	117: "⌦",  // Delete Forward
	115: "↖",  // Home
	119: "↘",  // End
	116: "⇞",  // Page Up
	121: "⇟",  // Page Down
	123: "←",  // Arrow Left
	124: "→",  // Arrow Right
	125: "↓",  // Arrow Down
	126: "↑",  // Arrow Up
	71:  "⌧",  // Numpad Clear
}

func isModifier(event hook.Event) bool {
	return isCommandModifier(event) ||
		isControlModifier(event) ||
		isFnModifier(event) ||
		isOptionModifier(event) ||
		isShiftModifier(event)
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

func isShiftOnlyModifier(modifiers KeyboardModifiers) bool {
	_, isShiftLeftPressed := modifiers[ShiftLeft]
	_, isShiftRightPressed := modifiers[ShiftRight]

	if ((isShiftLeftPressed || isShiftRightPressed) && len(modifiers) == 1) ||
		(isShiftLeftPressed && isShiftRightPressed && len(modifiers) == 2) {
		return true
	}

	return false
}
