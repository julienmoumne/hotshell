// Code generated by go-bindata.
// sources:
// examples/default/default.hs.js
// VERSION
// DO NOT EDIT!

package hotshell

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

var _examplesDefaultDefaultHsJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x58\xeb\x6e\x14\xc7\x12\xfe\x0d\x4f\x51\x5a\x21\x66\xd7\x66\x66\x6c\xe4\x73\x0e\xc7\x09\x90\x44\x24\x71\x22\x12\x90\x4c\x94\x48\x40\x50\xef\x4c\xed\x4e\xc7\x33\xdd\x4d\x77\xcf\x5e\x02\xce\xb3\xe4\x59\xf2\x64\xa9\xee\xb9\x8f\x77\x37\x46\x8a\x41\x78\xb7\x2f\x5f\xdd\xbf\xaa\x86\x5b\x2c\xa6\x1f\x52\x34\xc9\x39\x04\x17\xd2\x9a\x0c\xf3\x3c\xb8\x7e\x00\x8b\x52\x24\x96\x4b\x31\x9d\xc1\x87\xbb\x77\x81\x7e\x72\x2e\xca\x0d\x3c\x06\xdc\x60\x32\x0d\x4a\xc1\x0a\x0c\x66\x11\x17\x29\x6e\x5e\x2c\xa6\xc1\x73\xb7\x1d\xcc\xe0\x09\x84\xa7\xfe\xfc\x5c\xcb\xb5\x41\x4d\x37\xaa\x9b\x4f\x21\x30\x28\x0c\x9f\xe7\x18\xd6\x7b\x01\x9c\xff\xf5\x67\x20\x15\x8a\xa0\x92\xd1\x57\x67\xf2\x2a\xe3\x06\xe8\x6f\xab\x96\x81\x14\x17\xac\xcc\x2d\x14\x28\xca\x07\x90\x72\xa3\x72\xb6\xc5\x14\xd6\x19\x0a\x10\xd2\xed\x73\xc1\x9d\xde\xb0\xe0\x39\xba\xdb\x4a\xcb\x15\x4f\x31\x8d\xde\x88\xc9\xf5\xec\xb3\x1d\x72\x00\x9e\x49\x11\x58\xb8\x12\x72\x4d\x48\xcc\xb6\x12\xdd\x7d\x36\x97\xa5\x85\xa7\xf5\xed\xee\xf2\x15\x6e\xc9\x65\x2c\x20\x35\x2a\x98\x0b\x6e\xdd\x77\xb0\x12\x96\x68\xdd\x2f\x0f\xc8\xed\x64\xa7\x3b\x6f\xa8\x31\x10\x0a\x89\x2c\x0a\x26\xd2\x90\x7c\x87\xc0\x94\xca\x79\xc2\xbc\x5d\x84\x8b\x8b\x05\x4f\x38\x0a\x9b\x6f\x41\x63\xc2\xe8\x0e\x1d\x05\x93\x31\x8d\xcd\x45\x33\xb0\xd7\x0b\x13\x02\xf5\x33\x2c\x24\x3c\x6e\xd7\xc8\x76\x88\xe3\xe6\x8e\xd7\xc7\xd0\x3d\x38\x1e\x9c\xe8\x1b\xac\x5b\x83\x03\x8d\xc6\x32\x6d\x49\x3b\x96\x64\x48\xeb\x49\x91\xd2\xb2\x29\x53\x09\x14\xdc\x15\x4f\xb0\xde\x7b\x08\xf5\xd9\xe0\x7a\x76\x18\xde\x74\xf0\x66\x2b\x92\x4c\x4b\xc1\x7f\x47\xb0\xbc\x18\x09\x88\x4b\xa3\x63\x33\xe7\x22\x16\x56\xa5\xcc\x22\x28\x29\xf3\x88\xbe\x44\x52\x2f\x77\xca\xb9\xb9\x42\xa6\x9b\x72\xee\x92\xc9\x47\x4d\xcb\x52\x41\x2e\x97\xa4\x6d\x4e\x88\x69\xeb\xcb\xc3\x3a\xe7\x9d\xce\x95\xb9\x0e\xc3\x8c\x8b\xe8\x26\xc6\xbe\x4c\x0a\x58\x92\xa0\x31\x11\xa1\xb4\x36\xe7\xb4\x00\xc7\xdf\x40\xbc\x62\x3a\xa6\x8d\xb8\x76\x6c\xdc\x3b\xbb\xcb\xe6\xa1\x0c\xec\x64\xa0\xd6\x52\xdf\x4e\x44\x77\x74\x97\x04\xbf\xd6\x65\x59\xea\x13\x0c\x26\x03\x62\x71\x8b\x3b\xfd\xd1\x4b\xca\x63\x42\x1b\x61\x0d\xca\xe3\x07\x8a\x12\x15\x06\x65\xb8\x2f\x72\x8a\x4e\x69\xb8\x58\x52\xa9\x7c\xcf\x56\xec\x32\xd1\x5c\x59\x0a\x98\x58\xf0\x65\xa9\xab\x52\x79\x76\xf9\x3c\x82\xaf\x37\xac\x50\x44\x05\xe7\xbd\x0a\x1e\x83\x3b\xfd\xae\x67\xbb\xe5\x7a\x16\x4a\x34\x52\x42\x18\xb0\x14\xdc\x85\xcc\x73\xb9\x76\xa2\x7d\xe2\x8c\x71\x71\xc5\xf2\x69\x6b\xd6\x6c\x8f\x35\x6f\x04\xfc\xcc\x84\x67\x89\x1c\x99\x16\x50\x48\xb2\xac\x62\x99\xd1\x8d\x2a\x74\xc5\x88\x68\x0a\x4f\x34\x8e\x3a\xfb\xfc\xa8\xd8\x12\x27\x75\x40\x1b\xfe\x3d\x86\x00\x32\x6b\x95\x39\x8f\xe3\x25\xb7\x59\x39\x8f\x28\xaf\xe3\xdf\xca\x9c\x28\xa4\x90\x65\x21\x30\xce\x5a\xe6\x3f\xa0\xb0\x93\xeb\x1c\x60\x28\x2f\x70\xce\xb4\xaf\x19\xe9\xfe\x75\xab\x4a\xe3\x8a\x4b\x0a\x91\xf7\x8a\xd4\x90\x58\x9d\x1f\x27\x6e\xfb\x7d\xc9\x3b\x4e\x6d\x0c\x6c\x24\x8d\xa5\x00\x5c\xb0\x95\xf3\xae\xa5\x72\xa4\x5e\x01\x6b\xd2\xd9\xdd\xae\x4d\x6c\xb8\xca\x13\xa3\x54\x2e\xd0\x66\x1f\x3b\x67\x23\xa7\x65\xde\x69\x75\xdb\xf0\x5a\x13\xa8\x6a\x1c\x16\x64\x06\xc2\xd0\xad\x04\xfb\x95\x6b\xa2\x66\x10\xa9\x13\xfa\xd4\x32\x20\x17\x2d\x55\xec\x53\x05\x47\xaa\x60\xf0\x4f\x6d\xa1\xba\xb7\xee\x4a\x76\x4d\x39\x98\x51\x40\xef\x57\x74\x38\xbe\xee\x6e\xde\x19\x5f\x4f\xba\xeb\x54\xcb\xc9\x95\x37\xba\x06\x6a\x2b\x3f\x29\x75\x0e\x6b\x6b\x35\x75\xf3\xa0\x0e\xcf\x0d\xa4\x7f\x8f\x99\x77\xe3\xdb\x9e\xa1\xcc\x26\x99\xd7\xd4\xa1\xbb\x1c\x9b\x6f\x5b\x19\xd5\x66\x28\xe0\x14\x1c\xb8\xc3\xbb\xd3\xcf\xa7\x03\x7a\x5f\x5e\x5e\xc0\x56\x96\x9a\x8a\xcd\x50\xdb\xf2\xcc\x49\x3c\xe2\x7a\x15\xea\x1b\x84\xdd\xa2\x8d\x11\x7b\x7c\x9f\x4b\xea\xbe\x19\xa1\x75\x2e\x30\x19\x74\xab\xbd\x62\x1e\xa3\x2c\x3b\x94\x6f\x35\x6e\xe0\x25\x25\x3b\x4f\xe0\x4b\xaf\x15\xfc\xf4\xe3\x77\xbf\x0c\x30\x05\xae\x4b\x52\xf4\x8b\x25\x9d\xed\xfb\xf1\x90\xe9\x8b\x1e\xdd\xa7\x94\x74\xde\xf6\x21\x45\xba\x29\xe9\xf6\x96\xcf\x3b\xc0\x15\x2f\xe0\x8f\x38\x9a\x33\x93\xbd\xa3\x11\xcb\xe1\xb4\xfa\xee\xda\x3b\xe0\x89\xaf\x7a\x6a\x16\x2c\x31\x7b\x71\x77\xef\x1e\x40\xce\x86\xfa\xc6\x68\x93\xd8\xc5\xc5\x0c\x33\x76\xb4\x77\x00\xf0\x62\xac\xea\x3e\xc8\x1b\xbb\xb7\x08\x97\xe8\xc0\x05\xda\xb5\xd4\x57\x54\xea\x66\x6b\xe8\x08\x94\x96\xe7\xb7\x8f\x53\x2f\xf0\xd4\x2a\x53\xb0\xb8\xb1\xd4\x6a\xeb\x68\xb7\xee\x4c\x32\xe9\x0a\x69\xf2\xda\x25\xac\x03\x7d\x0b\xaf\x15\xb3\x16\x35\x7d\x9a\xc0\xfd\xfb\x34\x08\x31\x62\x5a\x50\xee\x33\x65\x9e\x82\x50\x8b\xb5\x81\x7b\x39\x84\x08\xf7\xd4\x21\x57\xc9\x46\x07\x9a\xf0\x2b\xea\xf1\x65\x41\x64\xa0\xed\x0e\x15\xdc\x72\x5f\xa8\x17\x49\x5a\xc1\xe7\x10\xa7\xb8\x8a\x6d\xa2\xe2\xd3\x87\xff\x8b\x4e\xe8\xcf\x69\x7c\x58\x74\xd9\xa7\xaa\xca\x7d\x6a\xc0\x53\xf5\xd7\x9d\x10\x83\x25\xf7\xd3\x12\xe9\x8d\x9d\x5a\x86\x67\x93\xd2\xf4\xe6\xc5\x1d\x67\x7b\xf6\x4e\x68\xa2\xb0\x52\x6f\x27\xf0\x11\x5c\x26\x43\xc8\xe1\xe1\x13\x6f\xa5\x28\x69\x96\xff\x08\x0e\x6b\x62\xe2\x5f\xe1\xe8\xf5\x49\xf8\xff\xb7\x47\x70\x14\xc7\xee\xb4\x21\x27\xd1\xaf\x52\xf0\xf7\x10\x26\xcd\x42\x28\x34\x7d\xcc\xc8\x6b\xc1\x40\xec\x01\x0f\xa9\xce\x43\xaa\xe2\x1d\xae\x80\xa5\xa9\x76\xdc\x43\xa1\x46\xd9\xa4\xc4\xb0\x49\xb8\x39\x82\xc6\x08\xae\xb8\x58\xc8\x88\xcb\x4f\x72\x21\xdf\xef\x42\x2e\x5c\xd2\xa1\x1f\xe0\x04\xfa\x04\xa7\x31\x03\x31\xdd\xeb\xc9\xb5\x7b\x63\x85\x2f\xa0\x73\x5b\xad\x1b\xcd\xee\xc8\x0c\x9a\xa8\x9c\x97\xc2\x96\x7e\xd6\x39\x3d\x8b\x4e\xce\xa2\xb3\xb8\x5a\x0a\xeb\xaf\x21\x09\xbf\xb2\x52\x85\xac\x48\xff\x7b\x16\x71\x23\x6f\xed\xbf\x1e\x79\x2b\x37\xab\xf8\x17\xae\xa3\x65\x98\x76\x63\x0f\x05\x59\xcd\x5a\xff\x0d\xcf\x1d\x4a\xdf\x7e\xa7\x75\xad\x09\xc8\xf5\x9a\x5e\x7b\x34\xb8\xd0\x6b\xcf\x65\x0e\x90\x7b\x5c\xc1\xc0\xa3\x93\x47\xa7\x9d\x84\xad\xcd\x68\x23\x2c\xe0\x92\xbb\xb9\xe4\xe2\xd5\xab\x97\x97\xbe\xb7\x55\xe7\x6e\x41\x41\x83\xd6\xb6\xdc\xd3\x1d\x6e\xd5\x18\x99\xab\x09\xdf\xfe\xd3\x86\xc3\x3a\xc4\xf6\xc9\x41\x03\x17\xbd\xe6\x9b\x17\x47\x93\xc5\x07\x7c\xd3\x7b\x25\xf9\xf7\xca\xe0\x8d\xf2\x69\x4f\x99\x81\x8c\x77\x11\xd2\x89\x69\xf5\x7f\x1b\x4e\xad\xb4\x53\x2b\x3e\x8a\x8e\x6a\xd5\x1c\x57\xfd\x27\x98\x45\x34\x40\x72\x3b\x0d\xde\x88\x60\xd6\x73\x0e\xe6\x0f\x80\x6f\xc6\xc4\x3c\x32\x81\x6f\x86\x16\x04\x34\xa0\xbb\x8b\x43\xcd\xab\xd5\x91\x92\xc3\x00\xd6\xbf\xae\x67\x7f\x07\x00\x00\xff\xff\xd6\x43\xd2\x22\xc4\x11\x00\x00")

func examplesDefaultDefaultHsJsBytes() ([]byte, error) {
	return bindataRead(
		_examplesDefaultDefaultHsJs,
		"examples/default/default.hs.js",
	)
}

func examplesDefaultDefaultHsJs() (*asset, error) {
	bytes, err := examplesDefaultDefaultHsJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "examples/default/default.hs.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _version = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x32\xd0\x33\xd4\x33\x00\x04\x00\x00\xff\xff\xb1\x21\x99\x2f\x05\x00\x00\x00")

func versionBytes() ([]byte, error) {
	return bindataRead(
		_version,
		"VERSION",
	)
}

func version() (*asset, error) {
	bytes, err := versionBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "VERSION", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"examples/default/default.hs.js": examplesDefaultDefaultHsJs,
	"VERSION": version,
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
	"VERSION": &bintree{version, map[string]*bintree{}},
	"examples": &bintree{nil, map[string]*bintree{
		"default": &bintree{nil, map[string]*bintree{
			"default.hs.js": &bintree{examplesDefaultDefaultHsJs, map[string]*bintree{}},
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

