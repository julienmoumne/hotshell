package item

// todo these keys are working on most systems and are not likely used for menu items, are there any other?
// todo Escape, Home, End, Delete all produces code 27, how to name this key?
var KeyCodes = struct {
	Backspace byte
	Tab       byte
	Return    byte
	Space     byte
}{
	127, 9, 10, 32,
}

var labels = map[byte]string{
	KeyCodes.Backspace: "backspace",
	KeyCodes.Tab:       "tabulation",
	KeyCodes.Return:    "return",
	KeyCodes.Space:     "spacebar",
}

func KeyName(key byte) string {
	if label, found := labels[key]; found {
		return label
	}
	return string(key)
}
