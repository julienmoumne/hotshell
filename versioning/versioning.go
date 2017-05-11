//go:generate go-bindata -nometadata -pkg versioning ../VERSION

package versioning

func GetVersion() ([]byte, error) {
	return Asset("../VERSION")
}
