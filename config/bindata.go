// Code generated for package config by go-bindata DO NOT EDIT. (@generated)
// sources:
// config/bindata.go
// config/config.go
// config/dev.toml
// config/release.toml
package config

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

var _configBindataGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func configBindataGoBytes() ([]byte, error) {
	return bindataRead(
		_configBindataGo,
		"config/bindata.go",
	)
}

func configBindataGo() (*asset, error) {
	bytes, err := configBindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/bindata.go", size: 0, mode: os.FileMode(420), modTime: time.Unix(1623724776, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configConfigGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8f\xb1\x4a\x03\x41\x10\x40\xeb\x9d\xaf\x18\x17\x84\x5d\x08\x77\x88\x5d\x24\x45\x4c\x25\xa8\x85\xb1\xb1\x4c\xce\xd9\x73\x31\xb7\x77\xec\xce\x29\x21\xa4\x14\x2b\xb1\xb3\xf6\x0b\x52\x8a\x68\xe2\xd7\xc4\x44\xff\x42\xb2\x27\xa2\x58\xce\x83\x79\xf3\xa6\x1a\x64\x97\x83\x9c\x30\x2b\x9d\xb1\x39\x80\x2d\xaa\xd2\x33\x2a\x10\x72\x38\x66\x0a\x12\x84\x34\x05\x4b\x00\x21\x73\xcb\x17\xf5\x30\xc9\xca\x22\x0d\x95\xd9\xd9\x4d\xaf\x6c\x45\x5e\x82\x06\x30\xb5\xcb\xf0\xc0\x59\x56\x19\x06\xf6\xd6\xe5\x1a\xc9\xfb\xd2\xe3\x04\x44\x14\xf5\x4a\xc7\xe4\xb8\xb5\xc1\xd8\xee\x60\x37\x04\x62\x65\x0a\x4e\xfa\x95\xb7\x8e\x8d\x92\x4d\x43\xba\x1d\x12\x2e\x8b\x91\x6c\x61\xa6\x35\x08\x6b\xe2\xca\x56\x07\x9d\x1d\xe1\x04\x40\x08\x4f\x5c\x7b\xb7\xc1\x20\xa6\x20\x62\x46\xd2\x27\xee\x45\xc1\xe9\xb8\x22\x25\xa3\x42\x63\x9a\xe2\xc7\xec\x6d\xbd\x98\x7d\xde\xdc\xad\x17\xb3\xd5\xc3\xed\xf2\xf5\x69\xf5\x38\x7f\x9f\xdf\x2f\x9f\x5f\xce\xba\x47\x87\x3f\x17\xda\x1d\x6c\x4c\x27\x34\x38\x6f\x54\x2a\xa6\x27\xc7\x74\xbd\x5f\x1b\x43\x5e\xfd\x7e\x45\xeb\xbd\xbf\x65\xff\xc2\xbe\x47\x67\x47\x30\x85\xaf\x00\x00\x00\xff\xff\x99\x58\x64\x37\x6a\x01\x00\x00")

func configConfigGoBytes() ([]byte, error) {
	return bindataRead(
		_configConfigGo,
		"config/config.go",
	)
}

func configConfigGo() (*asset, error) {
	bytes, err := configConfigGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/config.go", size: 362, mode: os.FileMode(420), modTime: time.Unix(1623706559, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configDevToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\x41\x4e\xc4\x30\x0c\x45\xf7\x39\x85\x15\xb6\x55\x66\xda\x01\xaa\x56\x9a\x93\x54\x2c\x5c\x6c\xcd\x54\x34\x4d\x94\xa4\x42\xec\xb8\xc1\x88\x23\xb0\xe3\x10\x73\x9d\x72\x0e\xe4\x16\x5a\x24\x24\xb2\xfa\xf6\xb3\x95\xe7\x30\x0e\xd6\x11\xc3\x11\x34\x71\x3b\x9e\x34\x00\xdc\xc0\x74\x7d\x9d\x2e\x6f\x9f\x1f\xef\xd3\xf5\x92\xc1\x0c\x32\x08\xdc\x33\x46\xce\x20\x71\x4c\x0a\x89\x82\x6c\xd5\x55\x55\x16\x5a\xa1\xf7\x1d\x49\x2d\xb1\x7b\xe2\x97\x25\x93\xb3\xd8\x0d\x92\xcf\x29\xf9\x7a\xb7\xeb\xdd\x23\xf6\x67\x17\x93\x56\xaa\xa1\xf6\x41\x01\x00\x34\xd4\x1a\x8b\x31\x71\x58\x6a\x79\x32\x23\x7b\x79\x51\x9a\xbd\xd9\x9b\x5c\xaf\xc8\xbb\x20\xe8\xee\xf6\x50\xac\x3d\x6a\x07\xb4\xf3\x19\x62\xb7\xcd\x8e\x91\xc3\xdf\xae\xc7\x18\x9f\x5d\xa0\x8d\xa8\x26\x30\x75\xf1\xdb\x67\xce\xe6\x84\x96\x37\xa1\x9f\x83\x57\xa1\xfa\xfe\x50\x56\xfa\x97\x01\x1c\x21\xff\xe7\x8b\xaf\x00\x00\x00\xff\xff\x44\xf1\x10\x8f\x6b\x01\x00\x00")

func configDevTomlBytes() ([]byte, error) {
	return bindataRead(
		_configDevToml,
		"config/dev.toml",
	)
}

func configDevToml() (*asset, error) {
	bytes, err := configDevTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/dev.toml", size: 363, mode: os.FileMode(420), modTime: time.Unix(1623724665, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configReleaseToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\x41\x4a\xc5\x30\x10\x86\xf7\x39\xc5\x10\xb7\x25\xbc\xf6\xa9\xa5\x0f\xde\x49\x8a\x8b\x84\x19\x9e\x45\xd3\x84\x99\x14\x71\xe7\x0d\x8a\x47\x70\xe7\x21\x7a\x9d\x7a\x0e\x49\xad\xad\x20\x2f\xab\x7f\xe6\xff\x7f\xf2\x0d\x0f\xbd\x0f\x48\x70\x06\xcd\xf4\x4c\x56\x48\x03\xc0\x0d\xcc\xd3\xdb\x3c\xbe\x7f\x7d\x7e\xcc\xd3\x58\x00\x92\x1b\x2e\x05\xac\x89\x02\x12\x49\x52\x16\x91\x73\xef\xd4\x34\x75\xa5\x95\x8d\xb1\xc3\x3c\x67\xd9\x3d\xd1\xeb\x22\x55\x8b\xee\x41\x01\x00\xb4\xe8\x8c\xb7\x92\x88\x7f\xe6\xfc\x1e\x83\xa4\x5c\x29\xab\xda\x1c\xcc\xc1\x94\x7a\xb3\x62\xe0\x6c\xdd\xdd\x1e\xab\x6d\x87\xae\xb7\x7e\x61\xcd\x00\x7b\x76\x10\xe2\xff\xdb\x68\x45\x5e\x02\xe3\xee\xa8\x96\x09\x3b\x59\x79\x16\x6d\x2e\xd6\xd3\x0e\xf4\x7b\xd3\x06\x74\xba\x3f\xd6\x8d\xfe\x43\x00\x67\x28\xaf\x7f\xf1\x1d\x00\x00\xff\xff\xf9\xaf\xaa\x0d\x4f\x01\x00\x00")

func configReleaseTomlBytes() ([]byte, error) {
	return bindataRead(
		_configReleaseToml,
		"config/release.toml",
	)
}

func configReleaseToml() (*asset, error) {
	bytes, err := configReleaseTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/release.toml", size: 335, mode: os.FileMode(420), modTime: time.Unix(1623721333, 0)}
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
	"config/bindata.go":   configBindataGo,
	"config/config.go":    configConfigGo,
	"config/dev.toml":     configDevToml,
	"config/release.toml": configReleaseToml,
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
	"config": &bintree{nil, map[string]*bintree{
		"bindata.go":   &bintree{configBindataGo, map[string]*bintree{}},
		"config.go":    &bintree{configConfigGo, map[string]*bintree{}},
		"dev.toml":     &bintree{configDevToml, map[string]*bintree{}},
		"release.toml": &bintree{configReleaseToml, map[string]*bintree{}},
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
