// Code generated for package generator by go-bindata DO NOT EDIT. (@generated)
// sources:
// templates/client.gohtml
package generator

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesClientGohtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x52\xc1\xaa\xdb\x30\x10\xbc\xeb\x2b\x06\x53\x8a\x5d\x8c\x72\x2f\xbc\x43\x09\x14\xde\xa1\x21\xa4\xfd\x01\x55\xd9\x38\xa6\xb6\x6c\xd6\x72\x42\x30\xfa\xf7\xb2\x92\x1b\xdb\x25\xbd\xf6\x62\x24\xef\xcc\xec\xce\xac\x7a\x63\x7f\x99\x8a\x30\x4d\xd0\xc7\xf9\x1c\x82\x52\x75\xdb\x77\xec\x91\x2b\x20\xab\x6a\x7f\x1d\x7f\x6a\xdb\xb5\xbb\x81\xaa\x96\x9c\xaf\xbb\x5d\xd5\x8c\xb4\xb3\x4d\x4d\xce\x67\x0a\xc2\x67\xe3\x2a\x82\x7e\x8f\xcc\x41\x54\x80\x58\xd0\x07\xd3\x8a\x2a\xb2\xd4\xc5\x5f\x11\xc2\x4c\x22\x77\x16\x64\xa1\xd4\x65\x74\x16\x07\xba\x0b\xe6\x3b\xf1\xad\xb6\xc2\xd9\xc7\x0e\x39\xf7\x36\x9d\x90\x5a\xea\x74\x2b\xf0\x69\x0b\xc7\xa4\x00\x8b\xcf\x6f\x70\x74\xcf\xb7\xb5\x42\x4a\xfa\x74\xdc\xe3\x0d\x4f\x3d\x05\x30\xf9\x91\x1d\xac\x0a\x4a\xf9\x47\x9f\xb2\x58\x68\xef\x5f\x8d\x25\xd4\xce\x13\x5f\xe4\x34\x6d\xdc\x7e\x23\x7f\xed\xce\xaf\xdc\x4a\xf7\xfa\x02\xfd\x85\xab\x1f\xa2\x1a\x82\xe1\x6a\x88\x88\xe5\xd7\x33\x81\x02\x7f\xf0\x27\xea\x9b\xc7\x52\xde\xdc\xcb\x25\x32\x62\xee\xb8\xd8\x84\xf8\x7a\x7e\x0c\x9e\x47\xeb\xe3\xdc\x62\x7e\x13\xa0\x70\xfe\x61\x26\xee\x23\xb7\x31\xe1\x0f\xeb\x18\xff\xbb\xc9\x38\x3a\x70\x33\x0c\x16\x5c\x92\xb7\x7e\x34\xcd\x9a\x17\x41\xc4\x2c\xdb\x8f\x7b\xd6\x7b\xd3\x34\x79\xfc\x8d\xf8\xf6\xd6\x36\xf4\xca\x45\x56\xce\xa0\x97\x6e\x4a\x19\xa7\x19\xe4\x9a\xa2\x9c\xc2\x14\xca\x25\xf6\x44\xfd\x18\x47\x4b\x42\x45\xfc\xce\xef\xea\x6f\xc7\x21\x24\xe4\xc6\xa5\x02\x82\x5a\x24\x7f\x07\x00\x00\xff\xff\x64\xc9\x6f\xed\x96\x03\x00\x00"

func templatesClientGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesClientGohtml,
		"templates/client.gohtml",
	)
}

func templatesClientGohtml() (*asset, error) {
	bytes, err := templatesClientGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/client.gohtml", size: 918, mode: os.FileMode(420), modTime: time.Unix(1757395562, 0)}
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
	"templates/client.gohtml": templatesClientGohtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
	"templates": &bintree{nil, map[string]*bintree{
		"client.gohtml": &bintree{templatesClientGohtml, map[string]*bintree{}},
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
