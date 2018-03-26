// Code generated by go-bindata.
// sources:
// dslrunner.js
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

var _dslrunnerJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x52\x4d\x6f\xd4\x30\x14\xbc\xfb\x57\x0c\x27\x3b\xca\x2a\xbd\xaf\x14\x21\x54\x81\xc4\x8d\x03\x88\x43\x55\x55\xc6\x7e\xdb\x35\x64\x9d\xc5\x1f\x74\xab\x6a\xff\x0b\xbf\x85\x5f\x86\x9c\xb8\xde\x24\x1b\x24\x7c\xf3\x7b\xf3\x66\xe6\x8d\xfd\x4b\x3a\x3c\xa0\x85\xa3\x9f\xd1\x38\x12\x3c\x5a\x4d\xce\xab\xde\x11\xaf\x58\xea\x9a\x40\x07\x8f\x16\x77\xf7\xec\xd0\xeb\xd8\x51\x43\xa7\x63\xef\x82\x6f\x54\x74\x8e\x6c\x40\x8b\x97\x01\xb4\x1d\xb1\x1b\x3c\xe9\x2d\x78\x73\xc3\xcf\xcb\x89\xd4\x47\x3b\xc0\xd6\x5a\x3e\xf7\x3c\x63\xbb\x68\x55\x30\xbd\x1d\xee\x42\xf5\x76\x67\x1e\x37\x50\xb2\xeb\xbe\x49\xf5\xa3\xc2\x0b\x03\x80\xe4\xef\xe2\x62\xdd\x1e\x1b\x90\x85\x8f\x4e\xa4\x62\xa0\xdb\xcc\x24\x5e\xa9\xd2\x31\x3b\x88\x37\x0f\x8d\xf1\x1f\x32\x5a\x14\xc1\xaa\x80\xd2\x71\x14\xa2\xb3\xac\xd4\xfe\x19\xcc\x68\xbc\xe0\x82\x7b\x9e\xc8\xa5\xa3\x8a\x8f\x52\x3e\x43\xc9\xa0\xf6\x10\x74\x52\x74\x4c\x36\xaa\xc5\x90\xa6\x8e\x02\x65\xf2\x31\xb9\x39\xe9\xd8\xd0\xe4\x15\xea\x16\x1c\x77\xef\x5f\xa9\xa0\x64\x7c\xdc\x87\x0d\x38\x6a\x14\x01\xd4\xe0\xf7\xbc\x09\xce\x1c\xa6\x46\xfe\x63\xbf\x9c\xf1\x08\x9f\x27\x2d\xf5\xf7\xe8\xc3\x57\x3d\x8b\x38\x5b\x7b\xd2\x97\xe1\x74\xa9\x21\x52\xee\x5f\xac\xa6\x9d\xb1\xa4\x45\xc1\x55\x78\x0b\xce\xff\xfc\xde\x4e\x46\x6b\xf0\x1b\x5e\xad\x8a\x1a\x6b\xc2\xed\xde\x74\xfa\x9d\x73\xf2\xf9\xea\x75\x17\x22\xd9\xc0\x90\xe0\xe2\x89\x67\xbd\xf1\xfb\xaf\x2f\xa9\x3f\x06\x3a\x7c\xee\x3f\xc9\x04\x9f\xef\x3a\xa5\x68\x8e\xd1\xef\xf3\x5a\x33\xeb\x97\x98\x86\xeb\x72\x81\x8c\x59\xa8\x0c\xd5\xab\xaf\xcc\xce\x7f\x03\x00\x00\xff\xff\xe3\x4a\x0d\x67\xcf\x03\x00\x00")

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

