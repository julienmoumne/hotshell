// Code generated by go-bindata.
// sources:
// dslrunner.js
// submodule_test.js
// DO NOT EDIT!

package dslrunner

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

var _dslrunnerJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x91\x41\x6e\xf2\x30\x10\x85\xf7\x3e\xc5\xec\x62\x0b\xe4\x03\xfc\x12\xcb\xbf\x37\xe8\x0a\x21\x44\x9d\x09\x58\x4d\x9c\x74\x6c\x23\x2a\xc4\xdd\x3b\x76\x5c\x13\x28\xde\x24\x79\x6f\xf2\xbd\xb1\xde\xf9\x40\xb0\x87\x0d\x10\x7e\x45\x4b\x28\x9b\xe8\x5a\x24\x6f\x46\xc2\x46\x09\x91\x6c\x1b\x70\xf0\x3c\xb2\xdd\xe5\x4f\x13\x89\xd0\x05\x16\xae\xd9\xf9\x37\x0f\xdc\x84\xe8\xa2\x33\xc1\x8e\x2e\x0b\xd2\x8c\xae\xb3\xc7\x35\x98\x43\xdf\x7f\x1c\xcc\xa7\x82\xab\x10\xc0\xa7\x8e\x11\x32\xcb\xa3\x4c\x0e\x94\x93\x12\x26\xc2\x33\xe3\x4b\x50\xb5\xee\xc1\x33\x5a\x54\x27\xd0\xf7\x02\x91\x67\x4b\xa8\x54\x55\xbe\xb1\x18\xcc\x09\x24\x5e\x0c\x4e\x69\x01\xf5\xf4\x53\x8b\x3d\x06\x2c\x74\x9d\x6f\xf5\x08\x9d\x8d\x16\xbd\x81\xd5\x06\x1a\xd8\xfe\xff\x45\x31\x3b\x1e\x4f\x61\xcd\xe2\x0a\x6a\x00\xbf\x37\xbb\x46\x07\xb2\xc3\x72\x11\xf1\xe2\x4a\xe9\xce\x62\xe1\xda\x0e\xe4\x5e\x5b\xff\xce\x75\x74\xd6\x61\x2b\xcb\xec\xbc\x97\x52\xcf\x0c\xbd\xa8\x49\xfc\xd1\xf5\x14\xfd\xa9\x54\xa2\x1e\x13\xde\x4a\x1d\xb2\x16\x75\x67\xd7\x86\x04\xaf\x35\x8c\x6d\xec\x51\xe3\x65\x1a\x29\xf8\xcc\xe5\xb8\xf4\x78\x65\xf9\xe2\xf9\x9f\x00\x00\x00\xff\xff\x5e\x00\x8a\xaa\x62\x02\x00\x00")

func dslrunnerJsBytes() ([]byte, error) {
	return bindataRead(
		_dslrunnerJs,
		"dslrunner.js",
	)
}

func dslrunnerJs() (*asset, error) {
	bytes, err := dslrunnerJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dslrunner.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _submodule_testJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x24\xcc\x4d\x0a\x02\x31\x14\x03\xe0\x7d\x4f\x91\xdd\x6b\x41\xe6\x00\x82\x87\xd1\x4e\xa4\xc5\x76\x9e\xf6\x47\x94\x61\xee\x6e\x8b\x59\x05\xf2\x91\xf7\xb5\x20\x36\x66\x5c\x50\xf8\xea\xb1\xd0\x4a\xd0\x56\x03\x53\x12\xb7\xcc\xc9\x98\xac\x6b\x4f\x5c\xf8\x79\x6a\x69\x75\xd0\x7b\xdf\x7c\x8b\xba\xc1\x3a\xec\x06\x23\x13\xda\xfd\xc1\xef\x19\x42\x39\xc1\xe7\x75\x36\x1f\x14\xb5\xdf\xfe\x07\x72\x38\x73\xfc\x02\x00\x00\xff\xff\x1a\xf2\xa5\xae\x71\x00\x00\x00")

func submodule_testJsBytes() ([]byte, error) {
	return bindataRead(
		_submodule_testJs,
		"submodule_test.js",
	)
}

func submodule_testJs() (*asset, error) {
	bytes, err := submodule_testJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "submodule_test.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"dslrunner.js": dslrunnerJs,
	"submodule_test.js": submodule_testJs,
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
	"dslrunner.js": &bintree{dslrunnerJs, map[string]*bintree{}},
	"submodule_test.js": &bintree{submodule_testJs, map[string]*bintree{}},
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

