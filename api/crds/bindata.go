// Package crds Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// installer.stash.appscode.com_stashpostgreses.yaml
package crds

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// ModTime return file modify time
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

var _installerStashAppscodeCom_stashpostgresesYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\x5f\x8f\xdc\x34\x10\x7f\xdf\x4f\x31\x12\x48\x05\x44\xb2\x9c\x2a\x21\xc8\x0b\x42\x07\x48\x15\x05\xaa\x6e\xe9\xcb\xa9\x48\xb3\xf6\x6c\x76\x38\xc7\x36\x9e\x49\xe8\xdd\xa7\x47\x76\x92\xbd\xfd\x7b\xd0\x13\x08\x3f\xc5\xbf\xf9\xff\x9b\x19\x2b\x18\xf9\x2d\x25\xe1\xe0\x1b\xc0\xc8\xf4\x5e\xc9\xe7\x9b\xd4\xb7\x5f\x49\xcd\x61\x39\x5c\xad\x49\xf1\x6a\x71\xcb\xde\x36\x70\xdd\x8b\x86\xee\x35\x49\xe8\x93\xa1\xef\x68\xc3\x9e\x95\x83\x5f\x74\xa4\x68\x51\xb1\x59\x00\x98\x44\x98\xc1\x37\xdc\x91\x28\x76\xb1\x01\xdf\x3b\xb7\x00\x70\xb8\x26\x27\x59\x07\x00\x63\x6c\x40\x14\x65\xbb\x00\xf0\xd8\xd1\x74\x8b\x41\xb4\x4d\x24\x24\x35\x7b\x51\x74\x8e\x52\x5d\x24\x35\xc6\x28\x26\x58\xaa\x4d\xe8\x16\x12\xc9\x64\x4f\x6d\x0a\x7d\x6c\xe0\x51\xdd\x31\xc0\x14\xd8\xa0\x52\x1b\x12\xcf\xf7\x6a\x97\x45\xfe\x9e\xed\xca\x75\x2c\x7a\x95\xc5\xaf\xa6\xb4\x0a\xee\x58\xf4\xc7\x53\xd9\x4b\x16\x2d\xf2\xe8\xfa\x84\xee\xa4\xa0\x22\x13\xf6\x6d\xef\x30\x1d\x49\x17\x00\x31\x2b\xa5\x81\x7e\xf5\xb7\x3e\xfc\xe9\x7f\x60\x72\x56\x1a\xd8\xa0\x93\x9c\x8e\x98\x10\xa9\x81\x9f\x73\x25\x11\x0d\xd9\x05\xc0\x80\x8e\x6d\x21\x7b\xac\x25\x44\xf2\xdf\xbe\x7a\xf1\xf6\xf9\xca\x6c\xa9\xc3\x11\xcc\x9e\x43\xa4\xa4\xbb\x92\x47\xfe\x77\x9d\xdf\x61\x00\x96\xc4\x24\x8e\xc5\x23\x3c\xcb\xae\x46\x1d\xb0\xb9\xd7\x24\xa0\x5b\x82\x61\xc4\xc8\x82\x94\x30\x10\x36\xa0\x5b\x16\x48\x54\x6a\xf0\x5a\x52\xda\x73\x0b\x59\x05\x3d\x84\xf5\xef\x64\xb4\x86\x55\xae\x33\x09\xc8\x36\xf4\xce\x82\x09\x7e\xa0\xa4\x90\xc8\x84\xd6\xf3\xfd\xce\xb3\x80\x86\x12\xd2\xa1\xd2\x44\xee\x7c\xd8\x2b\x25\x8f\x2e\x93\xd0\xd3\xe7\x80\xde\x42\x87\x77\x90\x28\xc7\x80\xde\xef\x79\x2b\x2a\x52\xc3\x4f\x21\x11\xb0\xdf\x84\x06\xb6\xaa\x51\x9a\xe5\xb2\x65\x9d\x67\xdd\x84\xae\xeb\x3d\xeb\xdd\xd2\x04\xaf\x89\xd7\xbd\x86\x24\x4b\x4b\x03\xb9\xa5\x70\x5b\x61\x32\x5b\x56\x32\xda\x27\x5a\x62\xe4\xaa\x24\xee\xb5\x2c\x4c\x67\x3f\x4a\xd3\x62\xc8\xb3\xbd\x4c\xf5\x2e\x96\xd9\x4e\xec\xdb\x1d\x5c\x26\xeb\x22\xef\x79\xb6\x80\x05\x70\x32\x1b\xf3\x7f\xa0\x37\x43\x99\x95\xd7\xdf\xaf\xde\xc0\x1c\xb4\xb4\xe0\x90\xf3\xc2\xf6\x83\x99\x3c\x10\x9f\x89\x62\xbf\xa1\x34\x36\x6e\x93\x42\x57\x3c\x92\xb7\x31\xb0\xd7\x72\x31\x8e\xc9\x1f\x92\x2e\xfd\xba\x63\xcd\x9d\xfe\xa3\x27\xd1\xdc\x9f\x1a\xae\xd1\xfb\xa0\xb0\x26\xe8\xa3\x45\x25\x5b\xc3\x0b\x0f\xd7\xd8\x91\xbb\x46\xa1\xff\x9c\xf6\xcc\xb0\x54\x99\xd2\xbf\x27\x7e\xff\xa1\x9a\xcf\xb9\xf5\xc8\xa7\xbc\x4a\x07\x08\x40\x87\xef\x5f\x92\x6f\x75\xdb\xc0\x97\xcf\x8f\x64\x11\x35\x8f\x64\x03\xbf\xdd\x60\x75\xff\xee\x93\x9b\x0a\xab\xfb\x2f\xaa\xaf\xdf\x7d\x76\x33\x7d\x7c\xfa\xcd\xc7\x47\x36\x67\x93\x9c\xe1\xb1\x81\x3b\x78\x7e\xf2\xce\x0e\xcd\xc1\x4b\xb4\x8a\x64\xf2\x00\xe5\x2e\x4e\x3b\xba\x09\x69\xd4\x81\x59\x69\xda\x0a\xd8\xb0\xa3\x7f\xc0\xc6\x1a\xcd\x6d\x1f\x8f\xf9\xb8\xa4\x9d\x0f\xa6\xf6\x0c\x7a\xb1\xe6\x7c\xf2\x5c\x71\x22\x7b\x6c\x56\x15\x67\x67\xb9\x3b\x22\x29\x9f\x4d\xef\x5c\x6e\xde\x2f\x03\xa5\xc4\xf6\xa4\x89\x17\x13\xe0\x0e\xdb\x13\xed\xc7\x4a\x4c\xd4\xb2\x68\xba\xfb\xc0\x32\xb3\x61\x0c\xc2\x1a\x9e\x60\xaa\xd8\xfe\x6b\xac\xce\xf9\x9f\x11\xcc\xf9\x9d\x88\x14\x8f\xfd\x5f\x6c\xc4\x93\x9a\x90\x48\x34\xa4\x0f\x6a\xc3\xff\x34\x69\xe7\x7c\x54\xd3\xa2\x1c\x40\x65\xae\x0e\x90\xa9\xc8\xc7\x57\xfe\x08\x1a\xe6\x3f\xb5\xe1\x0a\x5d\xdc\xe2\xd5\x03\x56\xaa\xaf\xa6\xff\xa8\x3d\x31\x40\xf9\xa5\xb0\x0d\x68\xea\xc7\x68\x39\x6e\x9e\xf2\x11\xf9\x2b\x00\x00\xff\xff\x2b\xb3\x82\xb2\x01\x0a\x00\x00")

func installerStashAppscodeCom_stashpostgresesYamlBytes() ([]byte, error) {
	return bindataRead(
		_installerStashAppscodeCom_stashpostgresesYaml,
		"installer.stash.appscode.com_stashpostgreses.yaml",
	)
}

func installerStashAppscodeCom_stashpostgresesYaml() (*asset, error) {
	bytes, err := installerStashAppscodeCom_stashpostgresesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "installer.stash.appscode.com_stashpostgreses.yaml", size: 2561, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
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
	"installer.stash.appscode.com_stashpostgreses.yaml": installerStashAppscodeCom_stashpostgresesYaml,
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
	"installer.stash.appscode.com_stashpostgreses.yaml": &bintree{installerStashAppscodeCom_stashpostgresesYaml, map[string]*bintree{}},
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