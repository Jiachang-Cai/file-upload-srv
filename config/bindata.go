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

	info := bindataFileInfo{name: "config/bindata.go", size: 0, mode: os.FileMode(420), modTime: time.Unix(1622446990, 0)}
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

	info := bindataFileInfo{name: "config/config.go", size: 362, mode: os.FileMode(420), modTime: time.Unix(1620877853, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configDevToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\x41\x6e\xc3\x20\x10\x45\xf7\x9c\x62\x44\xb7\x16\x02\x6c\x82\x1c\xc9\x27\x89\xba\x00\xcf\x24\xb5\x52\xc7\x16\xd8\x8d\xda\x55\x6f\x10\xf5\x08\xdd\xf5\x10\xb9\x8e\x7b\x8e\x0a\xb7\x75\xbc\x0b\x2b\x66\x1e\x7c\x1e\x13\xc6\x53\xdb\x21\x41\x05\x1c\xc9\x8f\x07\x0e\x00\x0f\x30\x5d\xdf\xa7\xcb\xc7\xf7\xd7\xe7\x74\xbd\x64\x30\x83\x0c\x02\x3d\x93\x8b\x94\xc1\x40\x71\x60\x0e\x31\xa4\x5b\xdb\xb2\xb4\x9a\x33\xd7\xf7\x0d\xa6\x3a\x37\x52\x15\x65\x6a\x34\x47\x7a\xad\x80\x1b\xb4\xd6\x29\x53\x6b\x59\x63\x91\x1b\xe9\x37\xb8\xd7\x92\x6a\xb2\xa6\xde\x2b\x43\x9c\xb1\x1d\xfa\x47\x06\x00\xb0\x43\x2f\x5a\x17\x07\x0a\xbf\x75\x5a\x4f\x5d\x1c\x52\xb0\xd2\x56\x48\x21\x85\xe2\x0b\xea\xbb\x90\x90\x29\x72\xbd\xf4\xd0\x9f\x5c\x3b\xff\xe7\xe0\x5a\x8a\x78\xbc\x1d\x1f\x23\xcd\xca\x6f\x67\xa4\x97\x55\x8a\x8b\xf1\xdc\x05\x5c\x90\x96\x5a\x26\xab\x40\xd8\xc4\x3f\xb1\x79\x2f\x52\xe4\xcd\xec\x7f\x04\x8b\xd9\x76\x93\xdb\x92\xaf\x54\xa0\x02\x75\xef\x9d\x9f\x00\x00\x00\xff\xff\xc1\xd3\x6f\x0d\x82\x01\x00\x00")

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

	info := bindataFileInfo{name: "config/dev.toml", size: 386, mode: os.FileMode(420), modTime: time.Unix(1622446250, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configReleaseToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x4f\x6e\xf2\x30\x10\xc5\xf7\x39\xc5\x7c\xfe\xb6\x10\xc5\x49\x1c\x37\x48\x48\xf4\xdf\xa6\x12\x5d\x56\xa8\xa8\x0b\xdb\x33\x50\x4a\x83\xc1\x0e\x02\x76\xbd\x01\xea\x11\xba\xeb\x21\xb8\x0e\x3d\x47\xe5\x94\x02\x8b\x7a\xe5\x79\xef\xf9\xf9\x67\xbb\xe5\xac\xb2\x48\xd0\x05\xe6\xe8\x95\x94\x27\x06\x00\xff\x61\xbf\x7b\xdb\x6f\xdf\xbf\x3e\x3f\xf6\xbb\x6d\x0b\x90\xf4\x72\xdc\x82\x43\xa2\x05\x35\xf9\x3a\x52\x88\x2e\x9c\xeb\x94\xa5\x4c\x59\xa4\xe6\xf3\x09\x86\x39\x13\x09\xcf\xcb\x20\x4c\xa6\xb4\xe9\x02\x13\x28\xa5\xe2\xc2\xa4\x89\xc1\x3c\x13\x89\x2e\x70\x94\x26\x64\x48\x0a\x33\xe2\x82\x58\x14\x0d\x51\x3f\x45\x00\x00\x43\xd4\x71\xa5\x7c\x4d\xee\x67\x0e\xeb\xd9\xfa\x3a\x14\x73\x99\xc6\x5c\xc6\x49\x9c\xb3\xa3\x37\xb7\x2e\x78\x22\xcf\xd2\xa3\x86\x7a\xa6\xaa\xe6\x49\x63\x55\x91\xc7\xe9\x29\xbe\xf4\xd4\x30\xaf\x57\xc1\x3a\xab\x51\xde\xaf\xac\x6b\xf8\x07\xfa\x5e\xa5\xbd\xbb\xb6\x1d\x3c\xf4\x69\x73\x5b\xdc\xa8\x00\xe8\x08\x27\xfe\xc0\xd8\xec\xe3\xd0\x70\x82\xfc\xfd\x8d\x23\x24\xcf\x3b\x45\x26\x4b\x76\x86\x05\x5d\xe0\x7f\x5e\x69\x9c\x6f\xab\x05\x1f\xaf\x84\x5d\x74\xae\x8a\xde\xc5\xa3\xb9\xbc\x7e\x59\xf7\x17\xd3\xd9\x3f\xf6\x1d\x00\x00\xff\xff\x58\xa8\x1c\xe4\xa3\x01\x00\x00")

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

	info := bindataFileInfo{name: "config/release.toml", size: 419, mode: os.FileMode(420), modTime: time.Unix(1622446250, 0)}
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
