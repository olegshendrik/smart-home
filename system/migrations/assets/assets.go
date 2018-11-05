// Code generated by go-bindata.
// sources:
// migrations/20170121_004649_nodes.sql
// DO NOT EDIT!

package database

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

var _migrations20170121_004649_nodesSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x52\x4d\x6f\x9b\x40\x14\xbc\xf3\x2b\xe6\x86\xa3\xda\x97\x48\x39\xe5\x44\x0d\x91\x2c\x21\xdc\xc4\x50\xe5\x86\xd6\xec\x53\xfc\x64\x58\xb6\xbb\x0f\x39\xee\xaf\xaf\x16\xa8\xeb\xb8\xad\xd4\xee\x89\x8f\xf9\x58\xcd\xcc\x6a\x85\x4f\x1d\xbf\x39\x25\x84\xca\x46\xab\x15\x76\xcf\x39\xd8\xc0\x53\x23\xdc\x1b\xc4\x95\x8d\xc1\x1e\xf4\x4e\xcd\x20\xa4\x71\x3a\x90\x81\x1c\xd8\x63\xe2\x05\x10\x7b\x28\x6b\x5b\x26\x1d\x35\x8e\x82\x96\x9c\x2d\xc1\xf4\x9a\x7c\xed\x45\xc9\xe0\xa1\x3c\xc8\x0c\x1d\x16\x31\x19\xb5\x6f\x49\xc7\x4b\xc4\x9a\xfd\xf4\x7c\xf7\x18\x45\xeb\x97\x2c\x29\x33\x94\xc9\xe7\x3c\x9b\xb8\x58\x44\x00\x6b\x5c\xce\x9e\xdf\x3c\x39\x56\x2d\x6e\x8e\xe9\x05\x66\x68\x5b\x34\xbd\xf1\xe2\x14\x1b\x99\xed\xed\x91\xce\xb0\x8e\x3b\xe5\xce\x38\xd2\x79\x19\x24\xed\x2f\xe6\x57\xe5\xd6\x07\xe5\x16\xf7\x0f\x0f\x77\x1f\x24\x8b\x6d\x89\xa2\xca\xf3\x40\xb0\xbd\x93\x8b\xd5\xd0\x91\xe3\xe6\xf6\x06\x1f\x09\x46\x75\xf4\x5f\x0e\x9a\x7c\xe3\xd8\x8e\x71\x96\xf4\x2e\xbf\xcb\xdf\x10\xa6\xa0\x75\xad\x04\x10\xee\xc8\x8b\xea\x2c\x4e\x2c\x87\xf1\x15\xdf\x7b\x43\x01\x37\x58\xad\x84\x46\xd8\x5f\x71\x63\x74\x01\x3c\x77\x35\x47\x7a\xd5\xde\x9f\x6e\x81\x34\x7b\x4a\xaa\xbc\xc4\xa5\xd1\xe8\xaa\xc6\xaa\xd8\x3c\x57\x19\x36\x45\x9a\xbd\x8e\xf9\xd5\x6c\xeb\xfb\x7a\x52\x1d\xcc\x37\x6c\x8b\xc9\x62\x11\x7e\x2e\xc1\x36\x90\xaf\xf7\x98\xf6\x27\xf3\x73\x91\x97\x39\x86\x8f\xff\x34\x48\xd7\xb7\x2d\x69\xec\x55\x73\x8c\xd2\x97\xed\x97\x79\x57\x9b\x27\x64\xaf\x9b\x5d\xb9\x9b\x17\xb6\x4e\x76\xeb\x24\xcd\x1e\xa3\x1f\x01\x00\x00\xff\xff\x15\x60\x3b\xf6\x0a\x03\x00\x00")

func migrations20170121_004649_nodesSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20170121_004649_nodesSql,
		"migrations/20170121_004649_nodes.sql",
	)
}

func migrations20170121_004649_nodesSql() (*asset, error) {
	bytes, err := migrations20170121_004649_nodesSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20170121_004649_nodes.sql", size: 778, mode: os.FileMode(420), modTime: time.Unix(1541428124, 0)}
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
	"migrations/20170121_004649_nodes.sql": migrations20170121_004649_nodesSql,
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
	"migrations": &bintree{nil, map[string]*bintree{
		"20170121_004649_nodes.sql": &bintree{migrations20170121_004649_nodesSql, map[string]*bintree{}},
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

