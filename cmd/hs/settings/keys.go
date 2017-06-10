package settings

var KeyCodes = struct {
	Backspace string
	Tab       string
	Return    string
	Space     string
	Escape    string
	Delete    string
}{
	string(127),
	string(9),
	string(10),
	string(32),
	string(27),
	string([]byte{27, 91, 51, 126}),
}

var labels = map[string]string{
	KeyCodes.Backspace: "backspace",
	KeyCodes.Tab:       "tabulation",
	KeyCodes.Return:    "return",
	KeyCodes.Space:     "spacebar",
	KeyCodes.Escape:    "escape",
	KeyCodes.Delete:    "delete",
}

func KeyName(key string) string {
	if label, found := labels[key]; found {
		return label
	}
	return key
}
