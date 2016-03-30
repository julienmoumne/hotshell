package item

const EOF_KEY_CODE = 4
const NUL_KEY_CODE = 0
const SPACE_KEY_CODE = 32
const TAB_KEY_CODE = 9
const LF_KEY_CODE = 10
const DEL_KEY_CODE = 127

var EOF_KEY = Key{EOF_KEY_CODE}
var NUL_KEY = Key{NUL_KEY_CODE}
var BASH_KEY = Key{TAB_KEY_CODE}
var REPEAT_KEY = Key{LF_KEY_CODE}
var PREVIOUS_MENU_KEY = Key{SPACE_KEY_CODE}
var RELOAD_KEY = Key{DEL_KEY_CODE}

var keyRegistery = map[byte]string{
	SPACE_KEY_CODE: "spacebar",
	TAB_KEY_CODE:   "tabulation",
	LF_KEY_CODE:    "return",
	DEL_KEY_CODE:   "backspace",
}

type Key struct {
	Byte byte
}

func MakeKey(key string) Key {
	if key == "" {
		return Key{}
	}
	return Key{key[0]}
}

func (k *Key) String() string {
	if label, found := keyRegistery[k.Byte]; found {
		return label
	}
	return string(k.Byte)
}
