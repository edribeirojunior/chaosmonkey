// Code generated by go-bindata.
// sources:
// migration/mysql/1.0.0_initial_schema.sql
// DO NOT EDIT!

package migration

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

var _migrationMysql100_initial_schemaSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x54\xd1\x6e\xda\x30\x14\x7d\xcf\x57\xdc\xb7\x82\x96\x4c\x50\x8d\x6d\x12\xda\x43\x20\x66\x8b\x16\x02\x0b\xce\xd4\x3e\x45\xae\x73\x01\x8b\xe0\x58\xb1\x51\xbb\x7d\xfd\xe4\xa4\xa1\x29\x6c\x15\x5a\xcb\x13\xbe\x3e\xf7\x9c\x63\xfb\xe4\x7a\x1e\xbc\xdb\x8b\x4d\xc5\x0c\x42\xaa\x1c\xcf\x83\xd5\x8f\x08\x84\x04\x8d\xdc\x88\x52\xc2\x55\xaa\xae\x40\x68\xc0\x07\xe4\x07\x83\x39\xdc\x6f\x51\x82\xd9\x0a\x0d\x4d\x9f\x05\x09\x0d\x4c\xa9\x42\x60\xee\x4c\x13\xe2\x53\x02\xd4\x9f\x44\x04\xc2\x19\xc4\x0b\x0a\xe4\x26\x5c\xd1\x15\x68\xbe\xc5\xfc\x50\xa0\x86\x9e\x03\x00\x20\x72\x08\x63\x5a\x23\xe2\x34\x8a\xc0\x4f\xe9\x22\x0b\xe3\x69\x42\xe6\x24\xa6\xb0\x4c\xc2\xb9\x9f\xdc\xc2\x77\x72\xeb\xd6\xf8\xdc\x9a\x6c\x7f\x81\x55\x69\x5b\xdd\xb6\xea\x79\x0d\xaa\x5c\x83\xc1\x6a\x2f\x64\xe3\xaf\x55\x76\xed\xc9\x8a\x92\xb3\x02\x8c\xd8\x23\xfc\x2e\x25\xd6\xd4\xf5\xaa\x4b\x4d\xc3\xf9\x09\xbd\xe7\x35\x28\x21\x21\xa5\xd3\xf7\x30\x41\xce\x0e\xba\x91\xb2\xf5\x5c\xac\xd7\x58\xa1\xe4\xe8\xc2\x9e\xfd\x7a\x5c\xc3\xba\x2a\xf7\xb5\xa7\x5a\x87\x29\x75\x94\x81\x9f\x7e\x32\xfd\xe6\x27\xbd\xd1\xf0\xba\xff\xa4\xd5\xe0\x38\x2f\x0f\xd2\x3c\xc7\x0d\x07\x83\x53\x5c\x85\x1b\x7b\xbe\x13\xbe\x41\x1f\x3a\xde\x3d\x0f\xac\xcf\xbb\x82\xc9\x1d\x68\x53\x09\xb9\x01\x53\x82\x90\xb9\xe0\xf6\xae\x64\x69\x40\x55\xa8\x51\x9a\x9a\x53\x1b\xc6\x77\xa7\x1e\xaf\x47\xa3\xfe\x2b\x38\x79\x71\xd0\x06\xab\xe7\x9c\x9f\x3e\x7e\x7e\x0d\x67\x18\x07\xe4\xa6\xbe\xda\x4c\xc8\x1c\x1f\xa0\x67\xff\xf7\xeb\xbd\xbe\x43\xe2\xaf\x61\x4c\xbe\x84\x52\x96\xc1\x64\xec\xbc\x94\xcb\x4e\x52\xfe\x37\x9a\x6f\xfd\xae\x17\xbc\xc1\xa5\xf7\xfa\x72\x4e\x4e\xec\xe9\xcd\xf9\x31\x86\x83\x73\x7f\x42\x6a\xc3\x24\xc7\x4c\xe4\x4f\xc0\x0f\x67\xb2\x3b\x51\x14\x98\x67\xcc\xfc\xf5\xb3\x82\x80\xcc\xfc\x34\xa2\x30\x4d\x93\x84\xc4\x34\xb3\xbb\x2b\xea\xcf\x97\x4d\x77\x81\x4c\x6f\x31\x6f\xdc\x4c\x16\x8b\x88\xf8\xf1\x79\xf3\xcc\x8f\x56\xc4\xed\x24\x82\x29\x95\x1d\x85\xdb\x68\x30\xa5\xdc\x63\xf1\x5f\x19\x71\xba\xd3\x30\x28\xef\x65\x3b\x0f\x8f\xc3\xd0\x16\x2f\x1a\x87\x55\x69\xb5\xe0\x8e\xf1\x9d\x13\x24\x8b\xe5\x63\xf0\x8e\x23\x70\xdc\xad\x76\x03\x38\x76\xfe\x04\x00\x00\xff\xff\x9f\x18\xee\xd3\x93\x05\x00\x00")

func migrationMysql100_initial_schemaSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationMysql100_initial_schemaSql,
		"migration/mysql/1.0.0_initial_schema.sql",
	)
}

func migrationMysql100_initial_schemaSql() (*asset, error) {
	bytes, err := migrationMysql100_initial_schemaSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migration/mysql/1.0.0_initial_schema.sql", size: 1427, mode: os.FileMode(420), modTime: time.Unix(1477837063, 0)}
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
	"migration/mysql/1.0.0_initial_schema.sql": migrationMysql100_initial_schemaSql,
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
	"migration": {nil, map[string]*bintree{
		"mysql": {nil, map[string]*bintree{
			"1.0.0_initial_schema.sql": {migrationMysql100_initial_schemaSql, map[string]*bintree{}},
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
