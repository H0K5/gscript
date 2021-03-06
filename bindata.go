// Code generated by go-bindata.
// sources:
// templates/compile_template.go.tmpl
// DO NOT EDIT!

package gscript

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

var _templatesCompile_templateGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x52\x51\x6b\xdb\x30\x10\x7e\xd7\xaf\x38\x44\x06\x32\xa4\x4a\xdb\xc7\xc1\x1e\xb2\xb6\xdb\x0a\x6b\x19\x1b\x74\x0f\x5d\x29\x8a\x75\x56\xc5\x2c\xc9\x48\x72\x92\x62\xf4\xdf\x87\x64\x2f\x89\x4b\x5f\x8c\xee\x7c\xf7\xdd\x77\xdf\x77\x9d\xa8\xff\x0a\x85\x60\x84\xb6\x84\x68\xd3\x39\x1f\x81\x11\x00\xba\x79\x8d\x18\x68\x7e\xd5\xce\x74\x1e\x43\x58\x35\xad\x88\x58\x52\xda\xad\xb4\xeb\xa3\x6e\x4b\x14\x5e\x6d\x4d\x49\x7e\x29\x1d\x5f\xfa\x0d\xaf\x9d\x59\x29\xb4\xe7\xb5\x96\xb8\x52\xa1\xf6\xba\x8b\x94\x54\x84\x34\xbd\xad\xcb\x2c\x56\xc1\x40\x00\xb6\xc2\xc3\x4e\x41\x06\xe0\xbf\x85\x8e\x5f\xbd\xeb\x3b\x02\x30\x0c\x67\xe0\x85\x55\x08\x8b\xe7\x25\x2c\xb6\x06\x3e\x7e\x82\x05\x7f\xb8\x0b\x90\x12\x01\xd8\x29\xbe\x96\x92\x5d\x54\x04\x40\x39\xc8\xb0\x13\x22\x80\xc4\x06\x33\x2a\xbf\x76\x16\x59\x55\x72\x2a\x60\x46\x98\xa8\xf0\x7b\xdc\x31\x3a\x0c\x19\x98\xdf\x5e\x43\x4a\xf4\x50\xc6\xaf\x3c\x8a\x88\x0f\x77\x53\xe7\x30\x9c\x12\x41\xb3\x41\x59\xb8\x6c\x0d\xbf\xc9\x41\x80\xb3\xc2\x68\x6c\x5e\x4b\x79\x5b\x34\x1c\xe1\x4b\x39\xff\xa2\x5b\xb4\xc2\x60\x9e\xb3\x84\x63\xfe\x5e\x18\xfc\x26\xc2\x0b\xa4\x74\x98\x85\x56\xce\x00\xbf\x3b\x21\x7f\x15\xd2\x2c\x87\x23\x78\x78\x9c\x91\xe7\x2a\xd0\x27\x56\x1d\x57\xb8\xd9\x63\xdd\x47\xfc\xd1\x0a\x5b\xb6\x48\xe5\x7b\x82\x3e\x2a\x98\x05\x67\x15\x49\x84\xcc\x97\x7c\xa3\xf6\x1b\x2f\xde\x93\x20\xa5\xd1\xd9\x77\x77\x63\x15\x3c\x3e\xe5\x6b\x2a\xfe\x84\x25\x3c\xe7\xee\xf1\x7c\xf8\x4f\x14\x72\xdd\xb6\xac\x9c\x56\xf6\x25\x27\xd0\xb3\x72\x7d\x39\xfe\xdc\x37\x0d\x7a\x36\x22\x0c\x93\x4c\x07\x42\x5a\xee\x97\xb0\xd8\x14\x3a\xe3\xe0\xab\xe9\x58\xf1\xa8\xe3\xd8\xa1\x1b\x30\x4e\x96\x16\xb8\xb8\x9c\xfd\x84\xce\x6b\x1b\x1b\xa0\x7f\x6c\x4e\x50\x98\x35\x9e\x5a\x32\xab\x3e\xdf\x7f\x38\xbf\xdc\x2f\x81\x66\x06\xff\x2b\x26\x95\x4b\x94\xaa\x62\x8a\xc7\xd8\x7b\x0b\x81\x64\x2d\x0f\x70\xc7\xe7\xbf\x00\x00\x00\xff\xff\xbf\x15\xf7\x33\x84\x03\x00\x00")

func templatesCompile_templateGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesCompile_templateGoTmpl,
		"templates/compile_template.go.tmpl",
	)
}

func templatesCompile_templateGoTmpl() (*asset, error) {
	bytes, err := templatesCompile_templateGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/compile_template.go.tmpl", size: 900, mode: os.FileMode(420), modTime: time.Unix(1511057853, 0)}
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
	"templates/compile_template.go.tmpl": templatesCompile_templateGoTmpl,
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
		"compile_template.go.tmpl": &bintree{templatesCompile_templateGoTmpl, map[string]*bintree{}},
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

