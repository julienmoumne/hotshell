// Code generated by go-bindata.
// sources:
// ../../../examples/default/default.hs.js
// DO NOT EDIT!

package definitionloader

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _ExamplesDefaultDefaultHsJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x58\xed\x6e\x1b\xb9\x15\xfd\xef\xa7\x38\x10\x16\x19\xc9\xf6\xcc\xd8\x41\xda\x6e\x9d\x66\xb7\x2d\xd2\xd6\x2d\xb6\xdd\x05\x9c\xa2\x05\x92\x34\xa0\x38\x57\x1a\xd6\x1c\x92\x21\x39\xfa\x68\xec\x3e\x4b\x9f\xa5\x4f\x56\x70\x66\x34\x1f\x92\x46\xb1\x81\xea\x87\x05\xf1\xe3\x9c\x7b\x0f\x2f\x0f\x49\xaf\x98\x85\xf0\x54\xe0\x0d\x2c\x7d\x2e\x85\xa5\x69\x94\x6b\xef\x72\x92\x32\x9a\x25\xa1\xeb\x2c\x8c\xa1\x0d\xf1\x91\x31\xa1\xab\x1a\xf3\xa9\x3f\xa0\x54\x19\x59\xc7\xb5\xa5\x68\x76\x76\x16\x70\xa6\x5f\x32\x72\xfc\x06\xd1\xed\x6e\xee\xe3\x25\x16\xa5\xe2\x5e\x68\x85\xe9\x0c\x5f\xce\xce\x00\x20\x40\x49\xa1\xca\x0d\xde\x54\xb4\x01\x8b\x15\x14\xa2\x51\x19\x6d\x7e\x5c\x4c\xa3\x1f\x42\x77\x34\xc3\x77\x88\xaf\xdb\x39\x73\xab\xd7\x8e\x2c\xde\x34\xb3\xbf\x47\xe4\x48\x39\x31\x97\x14\x37\x7d\x11\x6e\x10\x69\x43\x2a\xaa\xa9\xfa\x61\x4d\xde\xe5\xc2\x41\x38\xb4\xe1\x39\x64\xb4\x60\xa5\xf4\x28\x48\x95\x97\xc8\x84\x33\x92\x6d\x29\xc3\x3a\x27\x05\xa5\x43\xbf\x50\xa2\x8a\x7f\x21\x24\x85\xd9\xc6\xea\x95\xc8\x28\x4b\x3e\xa8\xc9\xe3\xec\xf5\x11\x1e\xe0\xad\x56\x91\xc7\xbd\xd2\x6b\xac\x73\xe6\x5b\xc6\x30\x9f\xcd\x75\xe9\xf1\x7d\x33\xbb\x9b\x7c\x4f\xdb\x1b\x44\x2c\xba\x44\x03\x73\x2b\x7c\xf8\x0d\xaf\xb1\x24\x1f\xbe\x2a\x40\xe1\x27\xc7\x65\x3d\x88\x63\xc0\x0a\xae\x8b\x82\xa9\x2c\x96\x42\x11\x98\x31\x52\x70\x56\x21\x78\x0d\x5a\x2c\x04\x17\xa4\xbc\xdc\xc2\x12\x67\x52\x82\xa9\x0c\x2e\x67\x96\x76\x13\xdd\x20\xe1\xdd\x9a\x08\xa5\xc8\xbe\xa5\x42\xe3\x4d\xdb\x1e\x3e\x13\x00\x69\xba\x9b\x5b\xc5\xe5\x3e\xa8\x09\x2e\x0e\x46\xf5\xb3\xb7\x6d\xf6\x91\x25\xe7\x99\xf5\x60\x86\xf1\x9c\xa2\x4b\xf0\x22\xbb\x41\xe4\xca\x4c\xc3\x91\x5d\x09\x4e\x4d\xdf\x4b\x34\x63\xa3\xc7\xd9\xd7\x29\x5c\x47\xe1\xb6\x8a\xe7\x56\x2b\xf1\x2f\x82\x17\xc5\x1e\x49\x5a\x3a\x9b\xba\xb9\x50\xa9\xf2\x26\x63\x9e\x60\xb4\x96\x89\xf2\x26\xd1\x76\x39\xca\x75\xbc\x35\x4d\xe1\xca\x79\xa8\xb2\x6a\x39\xad\x2e\x0d\xa4\x5e\xc2\x92\x64\x9e\xb2\x56\xe3\xaf\xc7\x2f\xbb\xf8\xeb\xf4\x03\x8e\x3b\xd8\x69\xc7\x81\xc6\x6a\x2d\x62\x9c\x93\x73\x89\xd4\xcb\x56\x04\x49\xce\xe1\xe2\xf7\x48\x57\xcc\xa6\x52\x2f\xd3\x46\xed\xb4\x37\x76\x4c\x84\x21\x0f\x75\x3c\x64\xad\xb6\x4f\xa3\xe9\x86\x8e\xb1\x54\xed\xc3\x72\xcc\xaa\x4a\xc4\x64\x60\x47\xa1\xf1\xb8\x40\xbd\xf2\xbd\xc0\x64\x1f\x70\xb0\x99\xfe\x4c\xaa\x74\x08\xfb\xa1\xf2\x04\xca\x50\x3a\xa1\x96\x60\xf8\x13\x5b\xb1\x3b\x6e\x85\xf1\xe0\x5a\x2d\xc4\xb2\xb4\xf5\xc6\x7a\x7b\xf7\x43\x82\xdf\x6d\x58\x61\x24\xe1\xa6\xb7\xe1\xf7\xc1\x43\x80\x8f\xb3\xe3\xbc\x95\x69\x71\x4b\xcc\x93\x83\xcf\x09\x0b\x2d\xa5\x5e\x07\xea\xaa\x9c\xf6\x71\x69\xc5\xe4\xb4\x4d\x6b\x36\x92\xcd\x07\x85\xbf\x31\x55\x99\x8a\x24\x66\x15\x0a\x6d\xa9\x31\xa5\xbd\x19\xf5\x1a\x16\x7b\xbe\x54\x54\xbe\x14\x9c\xb6\x6f\xa7\x86\x2d\x69\xd2\xac\xec\xce\xad\x2f\x10\x21\xf7\xde\xb8\x9b\x34\x5d\x0a\x9f\x97\xf3\x84\xeb\x22\xfd\x67\x29\x05\xa9\x42\x97\x85\xa2\xb4\x3d\x6c\x1e\x4f\x04\x1c\x78\x83\x00\xce\x30\x4e\x73\x66\xab\x9d\xa4\xc3\xdf\xd0\x6a\x2c\xad\x84\x2e\x5d\xad\x8a\xb6\xe0\xde\xca\x0b\x1e\xba\x3f\x97\xa2\xb3\xe0\x5d\x82\x3b\xa6\x7d\x16\xe0\x96\xad\x82\xba\xde\xea\x72\x2e\x09\x6b\xe1\xf3\x30\xbb\x49\x71\xe7\x68\x95\x8d\x6a\x13\x16\xda\x8d\x99\x79\xbe\x27\x5a\x5e\x89\xd6\x9c\x32\x55\xd4\x39\x49\xb3\x13\x2c\xca\x1d\xe2\x38\xb4\x44\xe3\xc1\xed\x56\xcd\x11\x81\xea\xd2\x72\xd0\x8b\xd6\x40\xc6\x42\xa1\xbd\x50\x28\xfa\xea\x29\x52\x4f\x5c\x77\x9b\x77\x4d\xcc\xe7\x64\xf1\xa2\x76\xcb\x83\xf9\xfd\xfd\xd9\x87\xe0\x1d\x04\xcf\x89\xdf\x57\x99\x37\x60\xad\x0f\xf0\xd2\x4a\xac\xbd\xb7\x89\x50\x51\xaf\x08\xf7\xc1\xfe\x7f\xfe\x3d\x4a\xe1\x7b\x29\x33\xcf\xf3\x2a\xde\x40\x10\xca\x6d\xbe\x6d\x69\xea\xce\x58\xe1\x1a\x01\xbf\x0f\x79\x58\xc7\x07\xd1\xdf\xdd\xdd\x62\xab\x4b\x8b\x42\x3b\x8f\xda\x54\x29\xab\xce\x36\xb2\x87\x86\x3e\x1a\x6d\xef\x40\x90\x9a\x33\x99\x6b\xe7\x3b\x25\x5c\x8e\xae\xf5\x44\xce\xcb\x0e\xe5\x0f\x96\x36\xf8\xa9\x9c\x4b\xc1\xf1\x9b\x2a\x2c\xfc\xf5\x2f\x7f\xfc\xfb\x00\x53\xd1\xba\x74\x64\x7f\xbd\xb4\xb4\xd9\x97\x73\x2c\xf7\x45\xef\x18\xc8\x84\xaf\x93\x1f\x3a\x66\xb8\x63\x3d\x23\xf5\x79\x87\xb8\x12\x05\xfe\x9d\x26\x73\xe6\xf2\x4f\xc6\xea\x00\xd4\x06\x7c\xac\xef\x84\x14\xbf\xed\xc5\x59\x30\xee\x46\x71\x8f\xf7\x9e\x40\xce\x87\xf1\xa6\xe4\x79\x1a\x16\xc6\x0d\x2b\x77\xaf\xef\x04\xe0\xed\x7e\xa8\x63\x90\x07\xbd\x4f\x58\x2f\xd5\x81\x2b\xf2\x6b\x6d\xef\xf1\x02\x6e\xeb\xc2\x63\xa2\xf4\x42\x3e\x63\xa1\x7a\x4b\xbf\x10\x2a\x83\xa7\x8d\x87\xd8\xad\x77\xab\x27\xcf\x75\xd8\x4e\x93\xf7\xa1\x64\x03\xea\x47\xbc\x37\xcc\x7b\xb2\xea\x23\x26\xaf\x61\x89\x65\x90\x30\xaf\xb1\xb4\x64\x10\x5b\xb5\x76\xf8\x46\x22\x26\x7c\x63\x4e\xe9\xa4\x77\xfc\xff\xfd\x4f\x63\x41\xd5\xa6\x80\xd1\xd6\x1f\xa1\x0f\xcd\x1d\xa1\x79\x0d\xce\x3c\x7e\x85\x34\xa3\x55\xea\xb9\x49\xaf\x5f\xfe\x22\xb9\x4a\xae\x92\xeb\xf4\x34\x6d\xd9\xf7\xaa\x5a\x37\x33\x30\xaa\xe6\xe7\x51\x88\x41\x53\xf8\xb4\x66\x7a\xd0\xd3\x70\x54\x46\x52\xba\xde\x75\xf2\xc8\xd8\x5e\xae\x93\x5c\x38\xaf\xed\x76\x82\x07\x84\x12\x46\x2c\xf0\xf2\xbb\x2a\x4b\x55\x4a\x89\x07\x04\xac\x89\x4b\xff\x81\xf3\xf7\x57\xf1\x2f\x3f\x9e\xe3\x3c\x4d\xc3\x68\xa7\xad\xc7\x03\x4a\x25\x3e\x23\xe6\xbb\x86\x58\x59\x3c\x20\x27\x96\x45\x03\xda\x13\x0a\x99\x4e\x21\x53\x3b\x8e\x30\x60\x59\x66\x83\xeb\xbc\xc0\x92\xf4\xae\x14\x86\x07\x45\xb8\x50\xdc\xa4\xa9\x30\x42\x2d\x74\x22\xf4\xb3\x24\x14\xe3\x12\x0a\x15\x8a\x8d\xaa\x9b\x9c\xa2\xba\xb2\x9d\x21\xca\x46\x95\x5c\x87\xb7\x59\xfc\x23\x3a\xd9\x9a\xd8\x2c\x49\x62\x8e\x5c\x52\xce\x4b\xe5\xcb\xea\xd2\x73\xfd\x2a\xb9\x7a\x95\xbc\x4a\xeb\xa6\xb8\xf9\x19\x67\xe4\xee\xbd\x36\x31\x2b\xb2\x9f\xbf\x4a\x84\xd3\x4f\xd6\xaf\x67\xdb\x26\x5c\x5a\xaa\x87\x71\x30\x64\x4c\xbb\xfb\x8f\xf3\xda\xcc\x5a\xfd\x86\xe3\x9e\x7a\xd4\x86\x53\x09\xbc\xb4\x96\x94\x47\x26\x2c\xf1\x50\x39\xd0\xaa\xda\x43\xf8\xf6\xea\xdb\xeb\x8e\x61\xeb\x73\xad\x10\x17\xb8\x13\xe1\x82\x72\xfb\xee\xdd\x4f\x77\xd5\xb1\x56\x8f\x7b\x82\xf7\x0c\x0e\xb5\xe5\xd8\xb9\xf0\xa4\x33\x91\x85\x4d\x51\x5d\x00\xb2\x9d\x7b\x75\x90\xed\x2b\xc4\x21\x96\xbe\x7d\x84\xec\xca\xf8\x84\x38\xbd\xc7\x53\xf5\x84\x19\x3c\x5b\x9e\xf7\xba\x19\x70\x7c\x9a\xd6\xff\x12\x09\x11\x65\x5d\x44\xe9\x79\x72\xde\x44\x15\x3c\xea\x67\xd1\x2c\x71\x46\x0a\x3f\x8d\x3e\xa8\x68\x36\x4b\x88\xf1\x7c\xda\xa9\x43\xf2\x12\x62\xb3\x6f\xc9\x7b\x39\x88\xcd\x30\x85\x08\x17\x08\x13\x87\xa1\xd7\xad\x7b\x51\x0e\x97\xb0\xf9\x7a\x9c\xfd\x2f\x00\x00\xff\xff\x8e\x0e\xa8\xdd\x6d\x12\x00\x00")

func ExamplesDefaultDefaultHsJsBytes() ([]byte, error) {
	return bindataRead(
		_ExamplesDefaultDefaultHsJs,
		"../../../examples/default/default.hs.js",
	)
}

func ExamplesDefaultDefaultHsJs() (*asset, error) {
	bytes, err := ExamplesDefaultDefaultHsJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../../examples/default/default.hs.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"../../../examples/default/default.hs.js": ExamplesDefaultDefaultHsJs,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"..": &bintree{nil, map[string]*bintree{
		"..": &bintree{nil, map[string]*bintree{
			"..": &bintree{nil, map[string]*bintree{
				"examples": &bintree{nil, map[string]*bintree{
					"default": &bintree{nil, map[string]*bintree{
						"default.hs.js": &bintree{ExamplesDefaultDefaultHsJs, map[string]*bintree{}},
					}},
				}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
