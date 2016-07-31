// Code generated by go-bindata.
// sources:
// assets/css/body.css
// assets/html/index.html
// assets/js/app.js
// DO NOT EDIT!

package handler

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _assetsCssBodyCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4a\xca\x4f\xa9\x54\xa8\xe6\x52\x50\x48\xcb\xcf\x2b\xd1\x4d\x4b\xcc\xcd\xcc\xa9\xb4\x52\xf0\x48\x2c\xca\x4d\x2c\x29\x49\xcc\xb3\x86\xc9\x14\x67\x56\xa5\x5a\x29\x18\xea\x19\xa7\xe6\x5a\x73\xd5\x72\x01\x02\x00\x00\xff\xff\x30\xd8\xd3\xf8\x37\x00\x00\x00")

func assetsCssBodyCssBytes() ([]byte, error) {
	return bindataRead(
		_assetsCssBodyCss,
		"assets/css/body.css",
	)
}

func assetsCssBodyCss() (*asset, error) {
	bytes, err := assetsCssBodyCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/css/body.css", size: 55, mode: os.FileMode(420), modTime: time.Unix(1469988667, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _assetsHtmlIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x90\xb1\x5a\xc4\x20\x10\x84\x7b\x9f\x62\x3f\x2a\x6d\x42\xef\x07\x67\x63\x71\xb5\x6f\xc0\x91\x4d\xc2\xdd\x86\x45\x76\xcf\xcf\xbc\xbd\x41\xa2\xf6\x56\x0c\xcc\xec\xf0\x83\x5b\x74\xa5\xd3\x03\x80\x5b\x30\x8c\x4d\xec\x92\x52\xbe\x41\x45\xf2\x46\x74\x23\x94\x05\x51\x0d\xe8\x56\xd0\x1b\xc5\x4f\xb5\x51\xc4\xc0\x52\x71\xf2\x66\x51\x2d\xf2\x6c\xed\xc4\x59\x65\x98\x99\x67\xc2\x50\x92\x0c\x91\xd7\x96\x7b\x99\xc2\x9a\x68\xf3\xe7\x50\xd7\xa0\x1a\xb2\xf9\xc7\x25\x36\x88\xa0\x4a\x3b\xb2\x17\x1e\xb7\xa1\x79\x47\x8f\x26\x25\xec\x1a\xe0\xcc\x19\x45\xe1\x35\x61\x37\xed\xaf\xeb\xec\xcf\x03\x5d\x6b\x38\x86\x2f\x77\x55\xce\xc0\x39\x52\x8a\x37\x6f\x2a\xbe\xdf\xf7\x82\x37\x26\x7a\x7c\x32\xa7\xb6\x3a\xdb\x43\xc7\xc4\x98\x3e\x20\x8d\xde\x10\xcf\xe6\xe4\xec\xbe\x3d\x0c\x89\x35\x15\x05\xa9\xf1\x8f\xf7\x2a\x36\x94\x32\x5c\xa5\x45\x7b\xa0\xb3\x74\x84\x9d\xe9\xfb\xf7\xbf\x02\x00\x00\xff\xff\xfe\x8d\x1c\x0d\x85\x01\x00\x00")

func assetsHtmlIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsHtmlIndexHtml,
		"assets/html/index.html",
	)
}

func assetsHtmlIndexHtml() (*asset, error) {
	bytes, err := assetsHtmlIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/html/index.html", size: 389, mode: os.FileMode(420), modTime: time.Unix(1469991594, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _assetsJsAppJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x53\xc1\x8a\xdb\x30\x10\xbd\xe7\x2b\x06\x5d\xd6\xa1\xbb\x0a\xa5\xf4\x92\xdd\x06\x5a\x28\xb4\x65\xdb\xc2\x6e\xa1\x87\xd2\x83\x2c\x8d\x13\x25\xf6\x28\x2b\xc9\x31\xa1\xe4\xdf\x3b\x92\x02\xf1\xc5\x90\xf5\x45\x96\xe6\xcd\x7b\x4f\x33\xa3\x83\xf2\x10\x9c\xde\x61\x84\x0f\x40\x38\xc0\x6f\xac\x9f\xf3\xbe\x12\x43\x58\x2e\x16\x1b\x47\x18\xa2\xb1\x28\xdf\x6f\xba\x6e\x78\xd9\xeb\x77\x47\xd9\x87\x3b\x54\x21\xde\xbd\x95\xd8\xf2\x6a\x75\x8d\x8a\x42\x54\xed\x4e\x6a\xd7\x2d\x06\xac\x0b\xa7\x98\xdf\xcf\x92\x42\xeb\xd6\x4c\x6f\x9c\xee\x3b\xa4\x28\xd7\x18\x3f\xb7\x98\x7e\x3f\x1d\xbf\x9a\x4a\x70\x38\x21\x67\x25\x49\x3a\xea\x30\x04\xb5\x46\xce\x69\x7a\xd2\xd1\x3a\x82\x0a\x0f\x8c\x9f\xc3\xbf\x19\x40\xa2\x6c\xbc\xea\x12\xe0\xdb\xf3\xcf\x1f\x72\xaf\x7c\xc0\x82\x90\x46\x45\xc5\x64\x00\x61\xb0\x51\x6f\xaa\x0c\x94\xf1\xb8\xc7\x92\x0c\xa0\x55\x40\x10\x6c\x37\xf6\x41\x2c\xf3\x11\x24\x8b\xd2\x52\x40\x1f\x3f\x9a\xad\xd2\xcc\xf4\xe5\xd7\xf7\xc7\xea\xa6\xc6\xc6\x79\x44\x32\x37\xb7\x45\x53\x96\x44\x78\x03\xe2\xa1\xf6\xb0\x58\x89\xac\x96\xbe\xda\xa3\xda\xdd\x8f\x34\xb4\xeb\x29\x1a\x37\xd0\xab\x64\x84\x77\x6d\x6b\x69\x0d\x96\x40\xb0\xce\x59\x16\xb5\x23\x13\x1e\xb1\x89\x57\x69\x27\x92\xd7\xcb\x02\xef\x47\xa2\x5a\x91\xb1\x5c\x51\x0c\x72\xeb\x2c\x55\x0c\xbb\x99\x27\xf9\x25\x3c\xd4\xab\x0b\x6e\xb0\x44\xe8\xb3\xaf\x45\xbd\x9a\xf4\x76\x9a\x9d\x46\x5d\x76\x7b\xa4\xa9\x16\x5f\x65\x38\x11\xa0\xb9\xa8\x9d\xc6\x33\xa4\x5b\x17\x26\x27\xe8\x2a\xfa\xcc\x30\x45\x8f\xde\x73\xa5\x26\xe8\xb9\x53\xc1\xb5\x28\x59\xe6\x7c\x5c\xb2\xf3\x63\x48\x65\x7e\xc2\x97\x9e\x9f\x15\xe7\x27\xb8\x48\xd3\xc9\x25\x15\xa3\x90\xb8\x4d\x81\x4b\xfd\x39\xfc\x47\xd4\x48\x82\x8d\x6d\xb1\x69\xd2\xba\xc3\x83\xcd\x07\xbc\xa4\x3b\x88\xbf\x59\x25\x8b\x14\x96\xa7\xd4\xd2\xb1\xc9\xe2\xef\x7c\x8b\xc0\x17\xad\xf2\x03\x0a\xd1\xf3\xc0\xd9\xe6\x58\x8d\x2c\xcc\x8b\xeb\xff\x01\x00\x00\xff\xff\x56\x62\x1c\xae\x23\x04\x00\x00")

func assetsJsAppJsBytes() ([]byte, error) {
	return bindataRead(
		_assetsJsAppJs,
		"assets/js/app.js",
	)
}

func assetsJsAppJs() (*asset, error) {
	bytes, err := assetsJsAppJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/js/app.js", size: 1059, mode: os.FileMode(420), modTime: time.Unix(1470001812, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"assets/css/body.css": assetsCssBodyCss,
	"assets/html/index.html": assetsHtmlIndexHtml,
	"assets/js/app.js": assetsJsAppJs,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"assets": &bintree{nil, map[string]*bintree{
		"css": &bintree{nil, map[string]*bintree{
			"body.css": &bintree{assetsCssBodyCss, map[string]*bintree{
			}},
		}},
		"html": &bintree{nil, map[string]*bintree{
			"index.html": &bintree{assetsHtmlIndexHtml, map[string]*bintree{
			}},
		}},
		"js": &bintree{nil, map[string]*bintree{
			"app.js": &bintree{assetsJsAppJs, map[string]*bintree{
			}},
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

