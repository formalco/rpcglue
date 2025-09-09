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

var _templatesClientGohtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x52\xcb\x8a\xe3\x30\x10\xbc\xeb\x2b\x0a\xb3\x2c\xf6\x62\x94\xfb\x42\x0e\x4b\x60\x21\x87\x0d\x21\x3b\x3f\xa0\x51\x64\xc7\x8c\xfc\xa0\x2d\x27\x04\xa3\x7f\x1f\x5a\xf2\xc4\xf6\x90\xb9\xce\xc5\x48\xee\xaa\xea\xae\x6a\x75\x4a\xbf\xa9\xd2\x60\x1c\x21\x8f\xd3\xd9\x7b\x21\xaa\xba\x6b\xc9\x21\x15\x40\x52\x56\xee\x32\xbc\x4a\xdd\xd6\x9b\xa2\xa5\x5a\x59\xdd\x6e\xa8\xd3\xa5\x1d\xcc\x46\xdb\xca\x34\x2e\x11\x60\x05\x52\x4d\x69\x20\xf7\x81\xdb\xb3\x0e\x10\x0a\xf2\xa0\x6a\xd6\x45\x12\xfb\xb8\x0b\xbc\x9f\x48\xa6\x39\x33\x32\x13\xa2\x18\x1a\x8d\x83\xb9\x31\xe6\xbf\xa1\x6b\xa5\x99\xb3\x0b\x1d\x52\xea\x74\x3c\x21\xb6\x94\xf1\x96\xe1\xd7\x1a\x8e\x51\x00\x1a\xbf\xb7\x68\xcc\x2d\x5d\xd7\x32\x2e\xc9\xd3\x71\x87\x2d\x1e\x7a\x02\x20\xe3\x06\x6a\xa0\x85\x17\xc2\xdd\xbb\x98\xc6\x4c\xdb\xff\x55\xda\xa0\x6a\x9c\xa1\x82\x4f\xe3\xca\xed\x3f\xe3\x2e\xed\xf9\x99\x5b\xee\x5e\x15\x90\x7f\xa8\x7c\x61\x55\xef\x15\x95\x7d\x40\xcc\xbf\x1e\x09\x64\xf8\xc0\x9f\x4c\x67\xef\x73\x79\x75\xcf\xe7\xc8\x0c\x51\x4b\xd9\x2a\xc4\xe7\xf3\xa3\x77\x34\x68\x17\xe6\x66\xf3\xab\x00\x99\xf3\x85\x99\xb0\x8f\x54\x87\x84\x7f\x2c\x63\xfc\x76\x93\x61\x74\xe0\xaa\x08\xc4\xb8\x28\xaf\xdd\xa0\xec\x92\x17\x40\x86\x88\xb7\x1f\xf6\x2c\x77\xca\xda\x34\xfc\x46\x78\x7b\x4b\x1b\x72\xe1\x22\xc9\x27\xd0\x53\x37\x39\x8f\x63\x7b\xbe\xc6\x28\x47\x3f\xfa\x7c\x8e\x3d\x52\x7f\x86\xd1\xa2\x50\x16\xbe\xd3\xbb\xfa\xec\xd8\xfb\x88\x5c\xb9\x14\x80\x17\xb3\xe4\x7b\x00\x00\x00\xff\xff\xbf\x9f\x59\x69\x98\x03\x00\x00"

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

	info := bindataFileInfo{name: "templates/client.gohtml", size: 920, mode: os.FileMode(420), modTime: time.Unix(1757396953, 0)}
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
