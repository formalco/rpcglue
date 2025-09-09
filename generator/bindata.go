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

var _templatesClientGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\xc1\x8a\xdc\x30\x0c\x86\xef\x7e\x8a\x9f\x50\x4a\x52\x06\xcf\xbd\xb0\x87\x32\x50\xd8\x43\x97\x65\xdb\x17\x70\xbd\x9a\x4c\x68\xe2\x04\xc5\xd9\x65\x08\x7e\xf7\x22\x3b\x99\x24\x9d\xf4\xd2\xde\x6c\x4b\xff\x2f\x7d\x92\x3b\x63\x7f\x99\x92\x30\x8e\xd0\xcf\xd3\x39\x04\xa5\xaa\xa6\x6b\xd9\x23\x57\x40\x56\x56\xfe\x32\xfc\xd4\xb6\x6d\x8e\x3d\x95\x0d\x39\x5f\xb5\xc7\xb2\x1e\xe8\x68\xeb\x8a\x9c\xcf\x14\x44\xcf\xc6\x95\x04\xfd\x18\x95\xbd\xb8\x00\x31\xa0\x9f\x4c\x23\xae\xc8\x52\x15\x7f\x41\x08\x93\x88\xdc\xab\x64\x16\x4a\x9d\x07\x67\xf1\x44\xef\x92\xf3\x9d\xf8\xad\xb2\xa2\x39\xc5\x0a\x39\x77\x36\x9d\x90\x4a\xea\x74\x2b\xf0\x69\x9b\x8e\x51\x01\x16\x9f\x1f\xe0\xe8\x3d\xdf\xc6\x0a\x09\xe9\x97\xe7\x13\x1e\x70\xf3\x53\x00\x93\x1f\xd8\xc1\xaa\xa0\x94\xbf\x76\x69\x16\x8b\xec\xf1\xab\xb1\x84\xca\x79\xe2\xb3\x9c\xc6\x0d\xed\x37\xf2\x97\xf6\x75\x8f\x56\xaa\x57\x67\xe8\x2f\x5c\xfe\x10\xd7\x10\x0c\x97\x7d\xcc\x58\x9e\x6e\x13\x28\x30\xe7\xbf\x50\x57\x5f\x97\xf0\xe6\x7e\x58\x46\x46\xcc\x2d\x17\x9b\x21\xee\xf7\x8f\xde\xf3\x60\x7d\xec\x5b\xe0\x37\x03\x14\xcd\x5f\x60\xe2\x3e\x72\x1b\x27\xfc\x61\x3d\xc6\xff\x81\xdc\x61\xcc\xef\x21\x13\x9b\xc8\xea\x9e\x66\xd6\x9b\x4b\x24\x01\xde\x0c\x83\x45\x96\xaa\x59\x3f\x98\x7a\x6d\x33\xef\x63\xdb\x5e\x7c\x25\x66\xf9\x22\xf1\x33\xe8\x93\xa9\xeb\x3c\xfb\x03\x52\xaf\x18\xb3\x03\x04\xea\x80\x8f\xb1\x5c\x31\x1b\x4f\xcd\xfd\x9b\x63\x5a\xca\x18\xc6\x70\xef\x9b\xb6\xb9\x6a\xff\x0e\x6b\xfa\xb2\x51\x17\xc7\xb5\xd7\xd3\x94\xb4\x8e\xce\xce\x41\x2d\xb7\xdf\x01\x00\x00\xff\xff\x43\x49\x49\x7e\x02\x04\x00\x00")

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

	info := bindataFileInfo{name: "templates/client.gohtml", size: 1026, mode: os.FileMode(420), modTime: time.Unix(1757395199, 0)}
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
