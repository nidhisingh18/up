// Code generated by go-bindata.
// sources:
// bindata.go
// byline.js
// index.js
// shim.go
// DO NOT EDIT!

package shim

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

var _bindataGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func bindataGoBytes() ([]byte, error) {
	return bindataRead(
		_bindataGo,
		"bindata.go",
	)
}

func bindataGo() (*asset, error) {
	bytes, err := bindataGoBytes()
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD
	info := bindataFileInfo{name: "bindata.go", size: 0, mode: os.FileMode(420), modTime: time.Unix(1507926480, 0)}
=======
<<<<<<< HEAD
	info := bindataFileInfo{name: "bindata.go", size: 0, mode: os.FileMode(420), modTime: time.Unix(1507633021, 0)}
=======
<<<<<<< HEAD
	info := bindataFileInfo{name: "bindata.go", size: 0, mode: os.FileMode(420), modTime: time.Unix(1506687428, 0)}
=======
	info := bindataFileInfo{name: "bindata.go", size: 0, mode: os.FileMode(420), modTime: time.Unix(1506522157, 0)}
>>>>>>> misc
>>>>>>> misc
>>>>>>> misc
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bylineJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x58\xff\x4f\xe3\x46\x16\xff\x3d\x7f\xc5\x2b\x27\x35\x89\xce\x18\xb6\xd2\x49\xd5\x22\x4e\x32\xc1\x2c\xbe\x0b\x09\x72\x4c\xb9\xd5\xb2\xaa\x1c\x7b\x8c\xa7\x38\x1e\xdf\xcc\x78\xd3\xa8\xcb\xff\x7e\xef\xcd\xd8\xb1\x0d\x81\x6b\xab\x2d\xd8\x33\xef\x7d\xde\xf7\x2f\xe6\xe4\x04\x66\xa2\xda\x49\xfe\x98\x6b\x98\xcc\xa6\xf0\xd3\xe9\x87\x0f\xc7\xf8\xe3\x1f\xf0\x2f\x91\x97\x70\xcd\xb6\x4a\x94\xa3\x93\x13\xfc\x07\xb7\x4c\x6e\xb8\x52\x5c\x94\xc0\x15\xe4\x4c\xb2\xf5\x0e\x1e\x65\x5c\x6a\x96\x3a\x90\x49\xc6\x40\x64\x90\xe4\xb1\x7c\x64\x0e\x68\x01\x71\xb9\x83\x8a\x49\x44\x00\xb1\xd6\x31\x2f\x79\xf9\x08\x31\x24\x28\x91\xf0\x90\x58\xe7\x88\xa4\x44\xa6\xb7\xb1\x64\x48\x9f\x42\xac\x94\x48\x78\x8c\x90\x90\x8a\xa4\xde\xb0\x52\xc7\x9a\x44\x66\xbc\x60\x0a\x26\x3a\x67\x70\xb4\x6a\x38\x8e\xa6\x24\x87\xb0\x52\x16\x17\xc0\x4b\xa0\xeb\xf6\x16\xb6\x5c\xe7\xa2\xd6\x20\x99\xd2\x92\x27\x04\xe3\x20\x51\x52\xd4\x29\x69\xd2\x5e\x17\x7c\xc3\x1b\x21\xc8\x4e\x68\xc6\x1f\x8a\x4c\xa8\x15\x9a\x42\x0a\x3b\xb0\x11\x29\xcf\xe8\x37\x33\xf6\x55\xf5\xba\xe0\x2a\x77\x20\xe5\x84\xbe\xae\x35\x1e\x2a\x3a\x4c\x58\x49\x5c\x68\xcd\x89\x90\x04\xa7\x58\x51\x10\x08\x47\x03\x8c\xd1\x9d\x8e\x86\x8c\x04\x55\xe4\x5c\xdd\xb8\xcb\x88\xde\xe6\x62\x33\xb4\x87\x2b\x42\xcb\x6a\x59\xa2\x60\x66\xd8\x52\x81\xee\x33\x72\x7f\x63\x89\xa6\x13\xe2\xc8\x44\x51\x88\x2d\xd9\x98\x88\x32\xe5\x64\x9a\xfa\xd8\x44\x31\xc2\xfb\x78\x2d\xbe\x31\x63\x96\x8d\x7c\x29\x34\x6a\x6d\x55\xa1\x88\x54\x5d\xa4\x9b\x2b\x95\xc7\x68\xc3\x9a\x35\xee\x43\xe1\x9c\xd2\x02\xe2\x9e\x65\x92\xd4\x50\x1a\xf3\x81\x63\x30\x2a\x21\x8d\xdc\x97\x16\xbb\xad\x1e\xd7\x3e\xac\x96\x57\xd1\xbd\x17\xfa\x10\xac\xe0\x36\x5c\xfe\x12\x5c\xfa\x97\x70\xe4\xad\xf0\xfd\xc8\x81\xfb\x20\xba\x5e\xde\x45\x80\x14\xa1\xb7\x88\x3e\xc3\xf2\x0a\xbc\xc5\x67\xf8\x77\xb0\xb8\x74\xc0\xff\xcf\x6d\xe8\xaf\x56\xb0\x0c\x09\x2d\xb8\xb9\x9d\x07\x3e\x1e\x07\x8b\xd9\xfc\xee\x32\x58\x7c\x82\x0b\x64\x5d\x2c\x23\x98\x07\x37\x41\x84\xb8\xd1\xd2\xc8\x6c\xd0\x02\x7f\x45\x78\x37\x7e\x38\xbb\xc6\x57\xef\x22\x98\x07\xd1\x67\x87\xb0\xae\x82\x68\x41\xc8\x57\xcb\x10\x3c\xb8\xf5\xc2\x28\x98\xdd\xcd\xbd\x10\x6e\xef\xc2\xdb\xe5\xca\x47\x25\x2e\x11\x79\x11\x2c\xae\x42\x14\xe4\xdf\xf8\x8b\xc8\x45\xc1\x78\x06\xfe\x2f\xf8\x02\xab\x6b\x6f\x3e\x27\x69\x04\xe7\xdd\xa1\x19\x21\x29\x0a\xb3\xe5\xed\xe7\x30\xf8\x74\x1d\xc1\xf5\x72\x7e\xe9\xe3\xe1\x85\x8f\xfa\x79\x17\x73\xdf\x4a\x43\xeb\x66\x73\x2f\xb8\x71\xe0\xd2\xbb\xf1\x3e\xf9\x86\x6b\x89\x40\xc6\x48\xa2\xb4\x6a\xc2\xfd\xb5\x4f\xa7\x24\xd5\xc3\x7f\xb3\x28\x58\x2e\xc8\x9e\xd9\x72\x11\x85\xf8\xea\xa0\xb9\x61\xb4\xe7\xbe\x0f\x56\xbe\x03\x5e\x18\xac\x50\x61\x63\x63\xb8\x44\x21\xe4\x5d\x64\x5a\x1a\x1c\x64\x5d\xf8\x16\x88\x3c\x3f\x0c\x10\x92\xd0\xfb\xdd\xca\xdf\x63\xc2\xa5\xef\xcd\x11\x6e\x65\xfc\xbf\x18\xd0\xbb\xa3\xd1\xb7\x18\xf3\x41\x4b\x16\x6f\xe0\x1c\x2b\xf0\xbf\x35\x97\x6c\x32\xb6\x27\xe3\xa9\x33\x02\xfc\xaf\xd6\xbc\xe8\xdf\xd2\xfb\x78\x7a\x36\x22\x44\xcc\xdc\x6f\xd8\x31\x58\x89\xd9\xe7\xdd\x06\x23\x2c\xbf\xba\x60\x2e\xfb\x9d\x12\x4b\x21\x57\x56\x97\xa6\xa4\x27\x88\x98\xae\x0c\xae\x03\xa2\x32\x49\x37\x85\x3f\x50\x80\x64\x1a\x6b\x05\x86\x9c\x6e\x82\x94\x9a\x59\x86\x83\xbc\x67\xa3\x67\xab\xc2\x3a\x56\x3c\x39\x20\x7c\x00\xf1\xa7\x34\xe1\x19\xf4\x2e\xed\xd9\x5e\x3f\x8b\x36\xe7\xe5\xfb\x4a\x01\x3c\x03\x2b\x14\x1b\x32\x97\x6c\x0b\x3d\xd6\x01\x79\x6b\x47\xca\x2a\xc9\x12\xd3\x53\xdf\x34\xa6\xc3\x38\x6c\x90\xd5\x19\x83\xa2\x04\xf2\x16\xe2\x71\x32\xc6\x40\x63\x0d\x7c\xfa\x08\xeb\x5d\x81\xdc\x7f\x7b\x05\x84\x9d\xa4\x27\x9a\xba\xcb\x96\xdb\x2e\x22\xd9\x06\x3b\x50\x8a\xbd\x4b\x94\x63\xa3\xec\xff\xf7\x45\x13\x97\x56\xb7\x3f\xe7\xb6\xbd\xf7\x7f\x78\xed\x7e\x9d\x4b\xb1\x35\x0e\xf4\xa5\x14\x72\x32\x46\x8f\x60\x13\x45\xb5\x3a\xda\x71\xe3\xc9\x57\x20\x2e\x3d\xc6\xeb\x82\xbd\x89\xd6\xd1\xc2\xa6\x56\xda\x9a\x6d\x79\x3a\x54\xaa\x92\x82\xb2\xf9\x9d\x38\xf6\x84\x56\xbc\x62\x93\x42\xf5\x3d\x56\x28\xf4\xcb\xa8\x69\xaa\xb5\xa2\xa6\x4f\xed\x96\xf0\x4a\x91\x32\xf8\x76\xea\x7e\x38\x85\x23\x5b\x78\xea\xa7\x23\x93\x02\x48\xfd\x32\x0d\x06\x09\xd0\xbd\xf4\x3d\x7e\x40\x3f\x63\xbc\xc5\x76\x23\x5c\x05\x54\x26\xe4\xc6\x4d\x70\x2c\x4c\x68\x92\x0c\xf3\xb7\x79\x46\xfc\xf6\xe9\xfb\x77\xf8\x83\x82\x0a\x60\x94\xc7\x05\xc2\xcc\xb1\x1b\x52\x1c\x67\x99\xd2\xa2\x32\xd6\xe0\x98\xae\x70\x52\x67\x12\x47\xe2\x9a\x91\x8d\xeb\x3a\xcb\x70\x03\x49\x2d\xeb\x36\xe7\x49\x8e\x1e\x39\xc6\x04\xc5\x6c\x8b\x4b\xcc\x0c\x65\x38\x29\x33\x51\x8d\xdf\x28\x02\xed\xc0\x47\xdf\x98\x63\x77\x04\x66\xde\xb9\xbf\xb6\x81\x59\xe1\x1e\xc0\xdc\x9e\x12\xe7\xa0\x65\xcd\xce\xf6\x84\xc4\x77\x61\x44\xe3\xd5\x97\xaf\xdd\xc5\x13\x63\x95\xbf\xa9\xf4\x8e\x9c\xd4\x33\xd1\x7d\x71\x81\x16\x67\x31\x56\x71\x0f\x32\x56\x7a\x96\xd7\xe5\x93\x5f\xe2\x5c\xbd\x47\x1d\x67\x21\xd5\xa0\xa5\xb2\xf6\xe9\xf8\x89\x19\x73\x94\xa8\x65\xc2\xc6\x0a\xb0\x2f\x0a\xb3\xc5\x60\x5e\x6e\x19\xee\x01\xe5\x58\x43\x1e\xe3\x58\x17\x25\x6b\xb1\xb1\x82\xc7\x94\x32\x63\xa7\x2b\x69\x25\x93\x36\x65\x4d\x4a\x1b\xc2\x16\xad\xbd\x31\x42\x71\xa5\x21\xe8\x24\x26\xe8\x54\xd8\xd5\x00\xe3\x0b\xa2\x48\x8f\x95\xde\x15\xac\x89\xbc\x6a\x78\x08\x0f\xe1\x71\x35\xa0\x2d\x20\x61\x38\xf9\x9b\xd4\x08\x5f\x14\x8b\x2d\x98\x9e\x60\x34\x18\x39\x5f\x06\xa2\xbd\x3d\x6b\x98\x9e\x47\xed\xcf\x67\xea\x06\x23\x9a\x17\x2e\x2f\x71\x15\xe5\x5a\x4d\xba\xf4\x74\x5e\xa5\x24\x0d\x95\xee\xde\xad\xa4\xd0\x42\xef\x2a\xe6\xfe\xaa\x5b\x92\x7e\xdf\x4b\x28\x1e\xce\xde\xc9\x0e\xf9\xb7\x51\xde\xf4\xd3\x84\x72\x63\xcd\xcb\x58\xee\xc0\xd0\x2a\x5c\x5c\xe1\x2e\xba\x3a\xfe\x19\x49\x7a\x46\xed\x1f\x31\xee\x38\xde\xb2\x9f\xc7\x26\xa4\xe4\x29\x9b\x46\x2e\x57\xf6\xc1\xca\x9c\xf6\x63\xd3\xe1\x9c\xc3\xd8\x26\xfc\xb8\xf3\xa0\xa1\x47\x11\xe6\xb7\xab\x05\x9a\x86\xb4\x93\xe9\x99\xa9\x25\x14\xd5\xd0\xf5\xb4\x69\x35\xe8\x7c\xd9\x1b\x28\x6f\x01\xee\x73\xa3\xe3\x7b\xde\x27\x6f\x62\x13\xf7\x95\xbd\xc6\x4a\xd3\xdc\x9a\x72\xb0\xa0\xaa\x2a\xb8\x9e\x9c\x3c\xc8\x87\xf2\xfb\x83\xfc\xfe\x50\x9e\x10\x6c\xe3\x55\x93\xc2\x86\x02\x66\xe1\xfc\xaa\x29\x6a\x55\x61\x7c\x1a\x27\x37\x9e\x7b\xa7\x6e\x7e\xfc\xd1\x92\x7e\x39\xfd\x6a\xbc\xf6\x50\xee\x3d\x66\x4b\x5e\xe5\x3c\xd3\x93\xa6\xff\x0e\xf1\xf6\xa5\xed\x16\xac\x7c\xd4\x39\xfc\x13\x4e\xbb\xee\x3e\x24\xf9\xf2\x16\xcf\x31\x7c\xf8\x0a\x7f\x3f\xb7\xd2\x50\x8b\xb3\x77\x64\xbf\x5b\xff\xd6\x0c\xeb\xb6\x3e\xb6\xb1\x4a\x8e\xdf\x68\x49\xaf\xd4\xb2\x0d\x71\x62\x14\x98\x76\x4c\x55\xad\xf2\x26\xef\xba\x24\xff\xd0\xe4\xb9\x1d\xb4\x87\xcb\xa5\x63\xec\xd7\x4b\x87\x41\xed\xee\x45\xb9\xc4\xc5\x36\xde\xa9\xa6\x63\xdb\xa6\x8c\x26\xc3\xa4\x12\xf8\x89\xb1\x2e\xf0\x1b\x31\x96\xf4\xc9\x30\x35\x6e\x42\x1e\x0c\x3c\xf6\x96\x77\xc2\x42\x42\xda\xc8\xb4\x49\x76\xc8\xf8\x9e\xc3\x8d\x26\xea\x89\x57\xc0\xa8\x19\xdb\x88\xec\x2b\xed\x60\x0f\xc7\x92\x25\xaa\x5e\x36\x40\x57\x7d\x5d\xef\x24\x8f\x4c\xda\x31\x62\x1c\xc1\x8c\xbb\xbb\xf6\x31\x9d\xf6\xfb\x9e\x19\x57\xcc\x7e\xaa\xe6\xf8\xf9\x75\xbc\xc5\x56\x27\x61\x13\xcb\x27\xda\x98\xd0\xe1\x49\x4e\x1f\xd4\x29\x23\x77\x11\x3a\xea\x52\x97\xb4\x2a\xdb\xc1\xfe\x3b\x7e\xe7\xf1\xe4\x69\x0f\x68\x36\x6d\x56\x64\x8d\x07\xce\xf6\x17\x8a\xe9\x60\xb3\x61\x29\x7d\x4e\x4f\xf6\xb1\xea\xeb\x02\x86\xf1\x8d\x74\xe8\x85\xb2\xc3\x7c\xee\x3d\xdb\x0d\xe4\x40\x73\xc6\xff\x89\x6d\xf2\x6e\x26\x65\x05\x0a\xed\x27\x51\x97\x34\xaf\x72\xf4\x40\xb3\x71\xe0\x74\x90\xae\xe6\x3b\x9b\x41\x3b\x69\x3e\x7e\x24\xf6\x37\x44\xb7\x71\xea\x4b\xb7\x21\x1b\x88\xe8\xb6\xc7\xe1\xac\xc2\x26\x33\x3c\xf8\xe1\xfc\x10\xe3\x60\x47\x6f\x0c\x39\x24\xa5\x6b\xb5\xc3\x59\xdc\xee\x88\xa6\x47\xbf\xd2\xa2\x95\x41\xdb\x81\xf9\x2b\x0a\xae\x36\x45\x4a\x5b\x26\x65\xc9\x46\x60\x85\x25\x62\xb3\xa1\x45\x39\xa6\x3f\x47\x70\x97\xb9\x38\xce\x71\x2d\x6d\xb6\xc4\xb8\xb4\x09\x6a\xd6\x70\xda\x2a\x9a\xa1\xd9\xd7\x9c\xd4\x1d\xa8\xf1\x57\xec\xda\x7f\x83\xfc\x2f\x00\x00\xff\xff\xbb\x06\x82\x6d\x66\x12\x00\x00")

func bylineJsBytes() ([]byte, error) {
	return bindataRead(
		_bylineJs,
		"byline.js",
	)
}

func bylineJs() (*asset, error) {
	bytes, err := bylineJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "byline.js", size: 4710, mode: os.FileMode(420), modTime: time.Unix(1507548061, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _indexJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x54\x51\x73\xe3\x34\x10\x7e\xcf\xaf\xd8\xbb\x99\x4e\xec\x5e\xea\xb4\x65\xfa\x92\x4e\x1e\x80\x16\xae\x0c\xdc\x31\x53\x18\x1e\x4a\xb9\x93\xed\x4d\x22\x70\x24\x23\xc9\x4d\x33\xa5\xff\x9d\x6f\x25\x3b\xc9\x31\xe5\xe0\xa1\xcd\x78\xf5\xed\xee\xb7\x9f\x3e\xed\xe8\x41\x39\xaa\x56\xba\xa9\x69\x4e\x8e\xff\xec\xb4\xe3\x6c\x1c\x03\x1f\x5a\x67\x2b\xf6\x7e\x9c\x5f\x46\x54\xb9\x6d\xb4\xe1\x43\x58\x31\x4d\x31\x41\x8c\xa6\xc7\xc7\x23\x3a\xa6\x2b\x2e\xbb\x25\xb1\x79\x20\xe4\x14\x88\x4c\x47\xa3\xca\x1a\x1f\xa8\x8e\x27\x73\xea\xcb\x16\xc0\x14\x57\xd7\x5f\xfd\xfc\xed\x87\xdb\xb7\x37\x3f\xec\x2b\x7c\x49\x6b\xd5\x92\x5d\x90\x0f\x4e\x9b\x65\xa6\xeb\x9c\x82\xa5\x4a\x35\x4d\xa9\xaa\x3f\x68\xd1\x99\x2a\x68\x6b\x26\xd4\x79\xae\x69\x61\x1d\x6d\x56\x6c\x24\x75\xad\xcc\x96\xd0\xad\xea\x9c\x63\x13\x22\x53\xf6\xc1\x93\x72\x4c\xb6\x0b\x3e\x28\x53\xa3\x66\xcf\x2b\xce\xde\x97\xf5\x60\xf6\xf4\xbc\x67\xf1\xd3\x8a\xa9\x51\xa0\xad\x6b\x52\x21\xa8\x6a\x85\x5e\xa0\xa1\x86\xa2\x34\xdd\x53\x6a\x95\x76\xfb\x92\x92\x76\x23\x7a\x66\x57\x2a\x70\x61\xec\x26\xcb\x81\x3e\x3b\x3d\x3d\xcd\xe9\x2f\x3a\xdd\x37\x31\xfc\x28\xc8\x25\x1b\x76\x80\x7a\x34\xf3\x18\x46\x57\x2b\xda\xe8\xa6\x21\x6b\x9a\x2d\x95\x8c\x96\x2d\xe3\xbc\x26\x7e\x60\xb7\xa5\xf3\xdf\x2e\xce\x29\xe8\x35\x12\x4a\xc6\x38\xbb\x02\x75\x22\x31\x28\xd4\xd7\xcf\xf2\xa7\x11\xd1\x74\x4a\x3f\x3a\xe4\x43\x16\xdc\x83\xb4\x61\x88\xa2\x4d\xe0\x25\x3b\x5c\x0a\x57\xda\x4b\xce\xc6\xa9\x56\x04\xb3\x9d\xa9\x41\x80\x16\x8d\x55\x41\x9a\xb4\x16\x60\x32\xdd\xba\x64\xe7\x53\xc1\xac\xf3\x1d\x34\xd8\x0e\xf0\x8b\xf3\x93\x8b\x2f\xa8\xd4\xc1\xe7\x00\x88\x16\x3a\xea\xd0\x2b\xf2\x86\xce\x7a\x01\x88\xf4\x82\x32\x39\x9c\xcf\x7b\xbd\x72\x12\x96\x94\x32\xce\x04\xf2\x8c\xbf\x9d\x96\xba\x96\x90\xe3\xd0\x39\x43\xb7\x3b\x67\x5c\x8e\x9e\x77\x6a\x7e\x1d\x6d\xdc\xdb\x2b\xfa\xa2\xd4\x46\x41\xaf\x9b\xe9\xfb\x83\x1b\x17\x00\x0a\x46\x8f\x17\xbe\x55\x1b\x23\x4e\x5e\x2b\x6d\xc6\x13\x7a\x82\xeb\x6a\x6d\x67\x74\x37\x6e\x75\xcb\x88\x0c\xbf\x83\x6d\x71\xce\xce\xdd\xd3\xb3\x98\x5e\x82\x85\x45\x01\x84\xac\x03\x6a\x90\x3e\x43\x20\xca\x2e\xde\xb7\x0d\x17\x11\x90\x8d\xef\xfc\x4a\xaf\xef\x29\x7e\xcd\xe8\xc8\x23\x45\x90\x32\xdb\xee\x5d\x3c\xea\x90\x9d\xc9\x64\xf9\x61\x03\x44\x0f\xeb\x57\xb6\xe6\x09\x79\xbd\x34\xaa\xf9\x5c\x23\xa4\xcd\x48\xc0\xf3\x23\xdf\xc3\xe7\xb1\xed\x27\x05\xfe\xb5\x7f\x2f\xed\x3b\xde\xc8\x43\x3f\xa9\xb9\xd1\x6b\x2d\x4e\xfc\xee\xf6\xfd\x3b\xd1\x0a\xaf\xea\x40\x5b\x7c\x41\xda\xb4\x15\xb2\xc8\x3d\x41\x50\x49\x80\x32\x48\xad\x82\x3a\x1c\x44\xa0\x71\x00\x71\x44\xdc\x11\xf9\x6e\x96\xc6\x2e\x77\x93\xb4\xca\x79\x5c\xfa\x8c\x3e\x1e\xf9\x8f\x28\x10\xf3\x46\xbd\xcd\xd6\x7e\x29\x23\x04\x5c\x76\x72\x11\x02\x20\x22\x24\x0b\x49\xe4\xd4\x26\xba\x0a\xcf\x36\xe0\x81\xc5\x1b\xea\xd1\x2f\xf5\xeb\xf0\x78\x5a\xae\x64\x56\x63\xcd\xc9\xef\x1e\x6f\x41\x8a\x7c\x4a\xe0\x32\xe6\x27\x5b\x46\xcb\xf6\x83\x84\x6d\xcb\x58\x61\xa0\x51\xc0\xd0\xaf\xe0\xf2\x71\x5a\x67\xe3\xff\xd9\x33\xee\xda\x13\xaa\x2d\xba\x07\xd9\x74\xbd\xd8\xff\xdd\x3e\x6d\xdb\x68\xf1\x61\xb7\xdd\x25\x1e\xf7\x02\xc7\x15\x72\xe0\x97\xce\x7a\xea\xaf\xaa\xdd\x53\xfc\xfc\x8d\x1c\xb0\xad\xbb\xb6\xd1\xd0\x55\x56\x95\x6f\x81\xfd\x87\x4c\x2f\xd1\xcc\xa4\x71\x74\xeb\x24\xea\xf4\xa0\x9a\x8e\xa3\xed\x40\x65\x30\xde\x5b\x2c\xec\x86\x29\xee\x2d\x9f\x8c\x86\xa6\xd6\xe1\x63\x95\x8e\xe6\x07\x6f\x4e\x50\x70\x76\x78\xc4\xbf\x32\x4d\x81\x8f\x62\x98\xf5\x17\x85\xcd\xf4\x8d\x75\xd7\xeb\x36\x6c\xaf\x05\xfc\xbd\xb5\xad\x54\x50\x8d\xe7\xcb\xbd\x78\x71\x07\x0d\xcb\x53\x44\xdb\xab\x05\xa5\x44\xd9\x32\xa2\x07\x87\x6b\x53\x6c\x1c\x9e\x45\x16\x0d\x97\x2e\x5a\x2f\xb6\x59\x92\xf1\xb5\xae\x5f\xcf\x50\x73\x92\xbe\x22\x4b\x04\x12\xdb\x14\x43\xdb\x80\x76\x88\x82\xaf\x28\x94\xbf\x19\xff\x6a\xc6\x71\xbd\xfd\x1d\x00\x00\xff\xff\x8f\x23\x4e\xfc\xa3\x07\x00\x00")

func indexJsBytes() ([]byte, error) {
	return bindataRead(
		_indexJs,
		"index.js",
	)
}

func indexJs() (*asset, error) {
	bytes, err := indexJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.js", size: 1955, mode: os.FileMode(420), modTime: time.Unix(1507922653, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _shimGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x8c\xc1\xad\x02\x31\x0c\x05\xef\xa9\xe2\x35\x90\xcd\x3f\xff\x1a\x38\xd0\x82\x43\x8c\xb1\x96\xd8\x91\x93\x45\xa2\x7b\x90\xf6\xc2\x71\x46\x9a\x29\x45\xfc\x5f\xd8\x38\x68\x31\xc4\x73\x55\x6b\xb4\x08\xb9\x7b\x5b\xda\x19\x7f\xc8\x63\x17\xcc\x87\x76\x6c\x29\x95\x82\x2b\xdd\x76\x12\x3e\xd5\x08\x7f\x69\xe3\x09\x3a\xf9\xee\x81\x38\xcc\xd4\x04\x14\x55\x57\x50\xbc\xf1\x24\x93\xe3\xdb\x4c\xb8\xe1\x42\xbd\x36\xda\xd2\xf8\xf9\xa4\x4f\x00\x00\x00\xff\xff\x42\x73\x98\x6c\x89\x00\x00\x00")

func shimGoBytes() ([]byte, error) {
	return bindataRead(
		_shimGo,
		"shim.go",
	)
}

func shimGo() (*asset, error) {
	bytes, err := shimGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "shim.go", size: 137, mode: os.FileMode(420), modTime: time.Unix(1507548061, 0)}
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
	"bindata.go": bindataGo,
	"byline.js": bylineJs,
	"index.js": indexJs,
	"shim.go": shimGo,
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
	"bindata.go": &bintree{bindataGo, map[string]*bintree{}},
	"byline.js": &bintree{bylineJs, map[string]*bintree{}},
	"index.js": &bintree{indexJs, map[string]*bintree{}},
	"shim.go": &bintree{shimGo, map[string]*bintree{}},
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

