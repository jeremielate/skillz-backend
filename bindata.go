// Code generated by go-bindata. DO NOT EDIT.
// sources:
// static/index.html
package main

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

var _staticIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x94\x5f\x6f\xd3\x30\x14\xc5\x9f\xd3\x4f\x71\xf1\xc3\x96\x4a\x5d\x5d\x21\x1e\x50\x49\xfa\xc0\x36\x51\xd0\xc6\x26\x1a\x24\x78\xf4\xe2\xbb\xc4\xc3\xb5\x83\x7d\x13\x1a\xaa\x7e\x77\xe4\xfc\x81\x0e\xc1\xe0\xa9\xd5\xbd\xbf\x73\x6c\x5f\x1f\x27\x79\x76\x71\x73\x9e\x7d\xbe\xbd\x84\x92\xb6\x7a\x35\x49\xc6\x1f\x14\x72\x35\x89\x12\x52\xa4\x71\x75\xa1\x8a\x42\x19\xd8\x7c\x51\x5a\x7f\x87\x33\xb8\xb2\xb9\xd0\x40\xe8\x49\x99\x02\x3c\xba\x06\x5d\xc2\x7b\x76\x92\xf0\x5e\x9c\xdc\x59\xd9\x06\x0f\x65\xaa\x9a\x80\xda\x0a\x53\x46\xb8\x23\x06\x95\x16\x39\x96\x56\x4b\x74\x29\x2b\x89\xaa\x25\xe7\x3a\x78\x96\xd6\xd3\xf2\xe5\x62\xb1\xe0\x0c\x94\x4c\x59\xed\x34\x03\xbe\x4a\xee\x1c\xf0\xff\xb0\x7a\x73\x99\xf5\xba\x2d\x52\x69\xe5\x23\x69\xc0\x85\x43\xf1\x48\x71\xba\x67\xe7\xd6\x10\x1a\x3a\xcb\xda\x0a\xd9\x12\x98\xa8\xb4\xca\x05\x29\x6b\xf8\x83\xb7\x86\x1d\x4e\x3b\xc7\x70\x26\x74\x9e\xad\x12\x3e\x3a\xfd\xdb\x7b\x8b\xde\x8b\xa2\xb3\x5d\xa3\xd6\x76\x34\x93\x82\xc4\x9f\x9d\xa4\x6a\x3a\xc2\xa1\xaf\xac\xf1\x18\x28\xa9\x9a\xd0\xf2\xb9\x53\xd5\xf1\xe1\xf9\x83\x68\x44\x5f\x65\xab\x49\x14\x35\xc2\xc1\xae\x74\x90\x82\xc1\x6f\xf0\xe9\xfa\x6a\x4d\x54\x7d\xc0\xaf\x35\x7a\x8a\xa7\xaf\x06\xa2\x76\x1a\x52\xd8\xef\x61\xfe\xd1\x69\x38\x1c\x42\x7d\x57\xba\xb9\xad\xd0\xc4\xec\xf6\x66\x93\xb1\x59\x80\x66\x40\xae\xc6\xe9\xd8\xf6\x48\x83\xd5\xba\x9b\x44\xfc\x78\x70\xb3\x30\xb8\xdf\x26\xf7\x53\x6b\x8d\x43\x21\x5b\x4f\x82\x30\x2f\x85\x29\x10\x52\xb8\xaf\x4d\x1e\x58\x88\xa7\xb0\x9f\x44\x51\xa4\xee\x21\x0e\x74\xc7\x6e\x02\x0b\x69\x9a\xc2\x0b\x38\x39\x81\x6e\x07\x24\xa8\xf6\x5d\xed\xf9\x62\x31\x88\x22\x69\xf3\x7a\x8b\x86\xe6\x05\xd2\xa5\xc6\xf0\xf7\x75\xfb\x56\xc6\xbf\x26\x38\x9d\x2b\x63\xd0\xad\xb3\xeb\x2b\x48\xa1\x5f\xa1\x6f\x65\xb8\xa3\xb0\xc7\xe8\x30\x89\xa2\xc3\x38\xa0\x70\x39\x90\xc2\xbb\xcd\xcd\xfb\xb9\x27\xa7\x4c\xa1\xee\xdb\xb8\x5b\xad\x76\x7a\x09\x7f\x5d\x31\x84\x75\x3a\x6f\x84\xae\x71\x16\xe8\x3e\x84\x4f\x08\x86\x94\x1e\x6b\x86\x98\x3d\x21\x1a\x83\x78\xac\x0a\x5b\x7e\x42\xd2\xc5\x6d\xe0\xc3\x49\x8f\x2e\xd5\xc8\x38\x74\x43\x25\xe1\x7d\x96\xc2\x0b\xee\x9f\x6e\xc2\xfb\xaf\xc1\x8f\x00\x00\x00\xff\xff\x9a\x84\x4d\x12\x25\x04\x00\x00")

func staticIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_staticIndexHtml,
		"static/index.html",
	)
}

func staticIndexHtml() (*asset, error) {
	bytes, err := staticIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/index.html", size: 1061, mode: os.FileMode(420), modTime: time.Unix(1523752955, 0)}
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
	"static/index.html": staticIndexHtml,
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
	"static": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{staticIndexHtml, map[string]*bintree{}},
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

