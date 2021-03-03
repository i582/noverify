// Package embeddedrules Code generated for package embeddedrules by go-bindata DO NOT EDIT. (@generated)
// sources:
// embeddedrules/rules.php
package embeddedrules

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

var _embeddedrulesRulesPhp = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x59\x6d\x6f\xdb\x38\x12\xfe\x1e\x20\xff\x61\x6c\xa8\xa9\x9d\xf5\x4b\x8b\x1e\x70\x40\x13\x27\x7b\xd7\x2f\x77\x40\xb1\x59\x6c\x6f\xbf\x5c\x91\x3a\x94\x34\xb2\xb8\x91\x49\x1d\x49\x25\xd6\xc5\xf9\xef\x07\x92\x92\x2c\x4b\xb2\x5e\xda\xbd\x02\x75\x6c\x6a\x38\xcf\x33\x33\xe4\xcc\x90\xba\xbe\x8d\xc3\xf8\xfc\xec\xfc\x6c\x79\x79\x79\x7e\x06\x97\xf0\x33\xe3\x94\xc9\x18\x3d\x45\x39\x83\xbf\x7d\xfe\x6c\x47\x23\xca\x14\x0a\xb0\xff\x7c\x2a\x89\x1b\xa1\x7e\xb2\x3c\x9a\xeb\xf1\xed\x16\x99\x82\xdf\x30\xe6\x42\x81\x42\xc1\x88\x48\x01\x77\xb1\x40\x29\x29\x67\x12\x54\x48\x14\x78\x84\x81\x8b\x20\xe9\x36\x8e\x68\x40\xd1\x5f\xd8\xf9\x2e\x06\x5c\x20\x80\xb3\x83\x5b\xfd\xf1\x11\x9c\xd4\x3e\x21\x81\x85\xd7\x4f\xf2\xd1\xe5\xf9\x59\x90\x30\x4b\x34\x83\xfa\x62\x35\xa6\x93\x29\xbc\x9c\x9f\x01\x58\x66\xa0\x35\x6c\x49\xea\x22\x78\x3c\x89\x7c\x10\x18\x47\xc4\x43\x50\x21\x16\x24\x9f\xa9\x0a\xe1\x8f\x44\x2a\x70\x3c\xce\xfc\x7c\x9a\x4a\x63\x04\x97\xf3\xa8\x3c\xbc\xd4\x7f\xcc\x6f\xb8\x05\x25\x12\x84\x8f\x10\x90\x48\xe2\x95\x76\x47\x0b\xec\xb3\xa0\x0a\x81\x48\x78\x98\x68\x9d\x53\xa3\xe3\xe1\x08\x0b\x46\xff\x07\xb4\xdc\x6d\x05\x54\x40\x77\x65\x5f\x9a\xb1\x38\x11\x08\xce\xae\x0c\x5a\x8a\xc3\x30\xb0\xdb\x26\xb0\xdb\x02\xcc\x68\xa7\x52\xa2\x9a\x38\xbb\xe9\x11\xc8\x6b\xdb\x8a\x8a\xb9\x42\xa6\x28\x89\x80\xc7\x28\x88\x09\x7d\x2c\xd0\x43\x1f\x99\x87\x5a\x63\x82\xb2\xbe\x98\x2e\xc0\xd9\x12\xf9\x08\xab\x15\xbc\xbb\x82\xe5\x52\x7f\x09\x89\x84\x90\x6e\x42\x14\x65\x0d\x2a\x24\x0c\x2e\x8e\xd7\xdc\xe4\xa0\x61\x6a\x54\x54\x17\xdf\x61\x7e\xb1\xee\x96\xf0\x0b\x57\xf8\x11\x9e\x51\x2f\x36\x4d\xfd\xa1\xc4\x63\xb4\xd2\xee\xd1\xce\x22\x35\xfa\x80\x4f\xc8\x32\x1d\x34\x00\x5a\xec\x16\x8f\x24\x9b\x50\xd9\x75\xfa\xa0\x97\x8a\x17\xa2\xf7\x78\x17\x3f\x64\x7b\x4a\xff\x92\xe0\x72\x15\x5a\xe7\xd8\xe5\xb3\x5c\x82\x96\x95\x60\x16\x1c\x50\x09\x8c\x2b\x20\xb0\xe1\xdc\xcf\xe5\x20\xe0\x02\x5c\xaa\x9e\xa9\xc4\x83\x63\xa7\x0b\x3b\x3f\xd3\x62\x23\x40\xd9\x06\x1e\x28\x7b\x22\x11\xf5\xad\xe2\x99\xde\xde\xe8\x29\xf4\x81\x25\x5b\x17\x05\x04\x3c\x61\xbe\xd9\x33\x0f\x40\x65\x36\x5d\xa3\x1a\x9e\x21\x46\x71\x90\x44\x33\x70\xd1\x23\x89\xb4\x9b\x50\x70\xae\x80\x07\xe6\x7b\x2c\xb8\x1b\xe1\x56\x53\x3d\xf8\x26\xe3\x02\xff\x2c\x43\x03\x11\x08\x04\x04\xca\x24\xca\xa6\x13\xb5\x28\xd6\x29\xfc\xfc\x4c\x04\xd3\x8c\x3b\xc2\x9d\xad\x47\xc2\xd2\x35\xfe\x67\xed\x52\x45\x98\xff\xd1\x06\x12\xc0\x59\xeb\xe9\xce\x5a\x07\x6f\x7d\x55\x8c\x5d\x14\x0f\xcc\xd8\x6b\x0d\x73\xd4\x1b\x93\x35\x82\x8e\x4e\x81\x8e\x5a\x40\x57\x43\x2c\xfd\xd0\x68\xea\x69\x5b\x5b\x8d\x1d\x62\x6d\x13\xf0\xe8\x24\xf0\xe8\x08\x78\x60\x6c\xf7\xb5\xd8\x72\x51\x0f\xed\xfe\x08\x75\xff\x63\xa1\xdd\xd7\x43\x7b\x8c\x39\x3a\x85\xf9\x03\x91\xdd\xd7\x23\x5b\x35\xf4\xb4\xa5\xdf\x1f\xd8\x7d\x43\x60\xab\xc6\x9e\xb6\xb6\x23\xae\xad\xc8\xb7\xb7\x79\x8d\x2a\x6c\xd3\xf5\x65\x7d\xd5\xa0\xaa\xc3\x88\xb2\xaa\x51\x87\xaa\x21\xa4\x3a\x38\xf5\xa7\xd4\xaa\xe8\xa6\xaf\x9e\x9b\x76\x35\xbd\xf9\xdc\xb4\xf3\xb9\xee\xab\xe7\xba\x5d\x4d\x6f\x3e\xd7\x47\x7c\x5a\xdb\x07\x22\x25\xdd\x30\x3d\xd2\xbf\x11\x5d\xe9\x8f\x9f\x6c\x6f\x52\xe9\x44\x7f\x5a\xe5\xc3\xe5\x6e\xc0\x82\xdc\xc5\x9d\x3d\xe8\x51\xc7\x64\x74\x55\x3b\x26\x33\xd8\xda\x9e\x95\xd9\xf5\x07\x9b\x37\x81\xcd\xfb\x81\xcd\x87\x82\x5d\x36\x81\x5d\xf6\x03\xbb\x1c\x0a\xb6\x6c\x02\x5b\xf6\x03\x5b\x0e\x05\x7b\xd3\x04\xf6\xa6\x1f\xd8\x9b\xa1\x60\x17\x4d\x60\x17\xfd\xc0\x2e\x86\x82\xed\x9b\xc0\xf6\xfd\xc0\xf6\x43\xc1\xbe\x35\x81\x7d\xeb\x07\xf6\x6d\x28\xd8\xf5\x75\x13\x9a\x1d\xed\x86\xbb\xbe\x1e\x8a\x77\x73\xd3\x84\x67\x47\xbb\xf1\x6e\x6e\x86\xe2\x2d\x9a\xe0\x16\xfd\xd0\x16\xc3\x8f\x79\x4d\x68\x76\xb4\x1b\xce\x9c\x07\x07\x9c\xf8\x82\x60\xee\xa6\x73\xce\x10\xb6\x54\x2a\xf2\x58\x3f\xe9\x91\xaf\x1e\x4f\x98\x9a\x38\x64\x7a\x5f\xc9\xd6\xa5\x47\xf3\xf7\xf7\xd5\x9c\xcd\x83\xe0\xef\xe9\xfb\x86\x8c\x9d\x97\x23\x7d\x28\x21\x6e\x94\x02\x65\x0a\x99\x8f\x3e\x28\x0e\xfa\xf4\x62\xb4\xce\xdf\x9b\xb3\x1c\x03\xca\x7c\xdc\x95\xdd\x01\x15\x6c\x98\x83\x41\x2f\x7c\x51\xa6\x5c\xf7\x7d\x27\xbc\xa4\xff\x45\x1e\x74\xe2\x5b\xb1\x13\x04\x0e\x0f\xef\xbb\xa2\x21\x13\x19\x53\x8f\xf2\x44\x9f\xbd\x36\x89\xad\xa4\x5c\xf8\x28\x2a\xa1\x90\x4a\xc4\x5c\x4e\xde\x2e\xdf\xce\xc0\x91\xd3\xe3\x60\x64\x0f\x1d\x39\x83\xb7\xcb\xb7\xd3\x5a\x01\x15\x1b\x79\xa7\x75\xb6\xc5\x23\x5f\x16\xc6\x27\x1e\x17\x02\x3d\x05\x21\x49\xa5\x22\xde\x23\xe8\x43\x2d\x43\xf4\x23\xac\xf2\x2c\x99\xae\x9b\xd6\x7c\xc6\xda\x4a\x17\x9d\x6b\x4e\xf1\x65\xec\x85\x44\x8c\x5f\x67\xe0\xbc\x8c\x2f\xc7\xaf\xd3\xab\xe2\x39\xed\x12\xe8\xd0\x90\xb8\x52\x89\x75\x16\xfd\x97\xb1\x54\x35\xa1\xd7\xd3\xeb\xa1\xd1\xfc\xec\x5a\xcb\x04\x4c\x7b\x40\x26\xee\x1f\x7a\xfc\xb4\x0b\x62\x81\x9b\x75\x36\x6d\xe2\xac\x67\x60\xfe\x37\x90\xf9\x1e\x1e\x96\x83\x12\x5a\xb0\x3d\x0a\xda\x11\xd9\xa4\x72\x04\x9a\x98\x9d\xf4\x76\x4d\x78\x3c\xfe\x31\x6f\xfa\x18\xd1\x2d\xd5\x2b\xb6\x97\x1d\xb8\x8b\x23\xee\x67\xf0\x4d\x3c\xfb\x6e\xac\x44\x92\x0d\x02\x0f\xea\x37\x32\xd5\x7c\x47\x03\x98\x38\x54\xfe\xfe\xdb\x67\x5d\xe1\x9f\x50\xd0\x20\x9d\xc2\x62\xb1\x38\xde\x6e\x65\xb1\xba\x5c\x79\xe7\x65\x88\x77\xb1\x6c\xd9\x7a\xbf\x4b\xf4\xe1\xa2\xc2\x8e\x0b\xe0\x4f\x28\xec\x4d\x69\x76\xad\x24\x67\x10\xa3\x08\x49\x2c\x35\x2e\x95\x45\xf2\xba\xad\x76\x31\x17\xa5\x7a\x51\xba\x70\xdd\x35\x8c\xa5\xc7\x85\xa4\xb9\xb3\x39\xa2\xba\x1f\x46\x75\xbf\x6f\xa3\xba\xdf\x7f\x37\xd5\x7d\x8f\x6a\xe7\x91\x28\x1a\x7e\x5d\x4e\xd9\x9a\x08\x41\xd2\x89\xf3\x38\x03\xf3\x6d\xfd\x88\xa9\x9c\x38\x2a\xa4\x72\x7e\xe3\x13\x45\xa6\x95\x14\x5c\x48\xad\x71\x47\xa5\x92\x66\x6a\x59\xbe\xba\x34\x34\xb3\xce\xdb\xf5\x4f\xa6\x49\xc8\x78\xa6\xba\x46\x35\xe0\x60\x3a\x03\xc7\x00\x94\xaa\x54\x97\x9c\xbd\x30\x3e\xd8\xa9\x1f\x96\x2d\x25\xd3\xce\x2d\xc6\xb8\x9a\xeb\x5d\xec\xa9\x39\x32\x9e\x6c\x42\xf0\xf8\x36\x26\x82\xca\x86\x9d\x95\x23\x8d\x9f\x43\xa2\xc6\x0d\x35\xac\x41\x62\x66\x2e\xe7\x6b\x9e\xb3\x98\x9f\xb6\xc5\x81\xf0\xe0\xb8\xa3\xc5\xca\x38\xcb\xf8\x95\x88\xc1\x44\x57\xf9\xd5\x6a\x35\xcd\x27\x2c\xed\x17\x7b\xd1\x93\x90\xa8\x48\x98\x87\x1b\x2d\x4d\xe3\xea\x30\x68\x5e\x19\x1c\x2e\x5b\xca\x92\xf9\x6b\x84\x7c\xd4\xfc\x3e\x21\xcb\x92\x28\x2a\x0d\xea\x9f\x47\x92\x59\x7e\xed\xb6\x2e\x4b\xa5\x6d\x46\x66\xfb\x28\x93\xcc\x76\x57\x7d\x3c\xad\x38\xc5\xb4\x96\xab\x43\x4a\x18\xee\xe9\x51\xb3\xa7\x19\x57\xcd\xde\x1e\x35\x7b\x7b\x54\xf7\xe0\xe8\x94\xb7\x9b\x65\x9b\xbc\x3d\xfa\x13\xbd\x3d\xfa\x71\x6f\x3b\x3b\xfb\x22\xa3\xdd\xd9\x1f\x84\x5f\x54\x4d\x5d\xd5\xf2\x8d\x03\xdb\x44\x2a\x9d\xd6\x8c\xc3\x9e\x43\x64\x39\x47\xb6\xc9\x20\xe5\x09\x22\x6e\x85\xc8\x21\x2f\xb8\xba\xf4\x4f\x87\x11\xb2\x69\x44\x22\x11\x5e\xf8\x27\x92\x2a\xab\x2d\x11\x6b\x4d\x51\x2a\x44\xd3\xd5\xf3\x00\xbc\x44\x44\x29\xb8\x82\x78\x28\xcd\xfb\x19\xd3\xda\x53\xb6\xa9\xdd\x50\xbd\xbc\x7b\xad\xde\x4c\x7d\x7d\x57\x3b\xe1\xe4\xd3\xbf\xa4\x4c\x91\x5d\x47\x2a\x22\x2f\xf4\xb5\x98\xa1\xcb\xa1\x8f\xb1\x40\x8f\x28\xd4\xd9\x9d\x79\x08\xbf\xfe\xe3\x57\xf8\xeb\xe2\x2f\x33\x43\x97\x7c\xa5\xf7\x40\x99\x54\x48\xfc\x83\x32\x5b\x31\xbf\x3a\xe9\x7d\x6d\xdd\xbc\x38\xe9\x6b\xe7\x5d\x9d\x2d\x86\x3e\x06\x94\x21\x14\x76\x98\xd7\x5f\x2a\xa4\xa5\x20\x4a\x1b\x2d\xc5\x73\x61\x8f\x48\xd4\x7c\x90\x49\xaa\xe8\x93\x3e\xa4\x31\xa9\x08\x53\xd5\x34\x6f\xe5\x27\xe3\x7f\xaf\x3f\xdd\xfd\xf2\xe5\x5f\xe3\x19\xbc\xcf\x92\x78\xe5\xb2\xaf\x41\x70\x5a\xbb\xf8\xb3\x42\x9f\x88\xc4\x0e\xf7\x96\x5d\xf8\x61\xf1\x6e\x66\xe2\x6e\x66\xd3\xec\x34\xda\x62\x82\xb9\x16\x75\x11\x59\x29\x28\x15\x07\x67\x6c\xf3\x46\x38\x37\xa8\x85\x51\xe6\x38\x1f\x03\x92\x44\xda\xa1\xdc\xf8\x5f\x1f\x30\x6b\x0c\x2c\xdd\xe3\x10\xe4\x2d\x8a\xc0\x2d\x7f\xc2\xca\x22\x38\xa2\x93\x77\xec\xa7\xc8\x9a\x9c\x98\x6d\x94\xff\x05\x00\x00\xff\xff\x39\xc2\xca\x83\x9d\x20\x00\x00")

func embeddedrulesRulesPhpBytes() ([]byte, error) {
	return bindataRead(
		_embeddedrulesRulesPhp,
		"embeddedrules/rules.php",
	)
}

func embeddedrulesRulesPhp() (*asset, error) {
	bytes, err := embeddedrulesRulesPhpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "embeddedrules/rules.php", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"embeddedrules/rules.php": embeddedrulesRulesPhp,
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
	"embeddedrules": {nil, map[string]*bintree{
		"rules.php": {embeddedrulesRulesPhp, map[string]*bintree{}},
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
