package item

const (
	eofKeyCode      = 4
	nullKeyCode     = 0
	spacebarKeyCode = 32
	tabKeyCode      = 9
	lfKeyCode       = 10
	delKeyCode      = 127
)

// todo how to make these vars non reassignable by other packages
var (
	EofKey          = Key{eofKeyCode}
	NullKey         = Key{nullKeyCode}
	BashKey         = Key{tabKeyCode}
	RepeatKey       = Key{lfKeyCode}
	PreviousMenuKey = Key{spacebarKeyCode}
	ReloadKey       = Key{delKeyCode}
)

var labels = map[byte]string{
	spacebarKeyCode: "spacebar",
	tabKeyCode:      "tabulation",
	lfKeyCode:       "return",
	delKeyCode:      "backspace",
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
	if label, found := labels[k.Byte]; found {
		return label
	}
	return string(k.Byte)
}
