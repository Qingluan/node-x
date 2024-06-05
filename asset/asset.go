// Code generated for package asset by go-bindata DO NOT EDIT. (@generated)
// sources:
// Res/cert.pem
// Res/google.js
// Res/google_news.js
// Res/js_text.js
// Res/key.pem
// Res/node-x.ini
package asset

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

var _resCertPem = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x94\x4b\xaf\xaa\x48\x14\x46\xe7\xfc\x8a\x9e\x93\x8e\x78\x44\x28\x86\xbb\x28\x84\xe2\x25\xc8\xdb\x99\x88\x14\xa2\x87\x87\x28\xaf\x5f\xdf\xb9\x27\x7d\x93\x4e\x3a\x77\x0f\x57\xd6\x6c\x65\x7f\x7f\xff\x3a\xac\xe9\xd4\xfd\x4b\xd5\x4e\x21\x3d\x50\x15\x42\xed\x87\x72\x0e\xa5\xaa\x18\xaa\x2a\x5c\xbf\x19\x4c\x14\x03\xa3\x51\x11\x89\x49\x9d\xad\xd9\xb5\x90\xdf\x52\x22\xa6\xf2\xa7\xb8\xe5\x77\x87\xea\x36\x4c\xc4\xcf\x4c\xab\x3d\xd3\x6a\xbc\xba\xe0\x6b\x36\x87\x7d\x98\x80\xc0\xed\x30\x09\x8b\x4b\xe0\xcb\x21\xfe\xec\x10\x10\x5d\x82\x2f\x87\x49\x58\x5d\x02\xeb\x7f\x99\x03\x30\x31\xa6\xdd\x1d\x10\x74\x35\xe8\x75\x2e\xa0\xf9\x8e\xf8\x1a\x06\x3f\x02\x10\x29\x26\xd3\x8f\x60\x41\x4b\x31\xf8\xe4\x1c\x1f\x25\x47\xf3\xee\xdb\xa4\x5b\xf7\x4a\x65\x6d\xbe\xec\x49\x8c\xc4\xa6\xf3\x9a\x00\x95\x2f\xe7\xc4\xc9\x26\x3b\x34\x62\x33\xda\x59\x22\xec\x6f\x49\x53\x07\xf9\x06\x87\xcf\x3b\xae\xb7\xa0\x8b\x4d\x51\xce\x63\x37\x8f\xd0\xe4\x6c\xcd\xb0\xf4\xd4\xfa\xe8\xf8\x94\xc7\x40\x56\x14\x24\xdd\x7a\x1d\x73\xf1\x75\x9f\x59\xd5\x7e\x78\x5e\xeb\x2a\x55\x54\x43\xba\xda\xd1\xd0\x2d\x3d\x51\xd6\xe2\x51\x51\x59\xb1\x98\xb8\x57\xba\xe3\x59\x54\xda\x21\x79\x9e\x54\x66\xd9\x96\x96\x97\xe5\xb5\x1f\x6e\xb5\x75\xe4\xce\x69\x58\x24\x9f\x00\xd5\xad\xaf\xd4\x58\x4c\x7d\x33\x3c\x26\x5e\x60\x1b\xde\x12\xb9\x08\x1f\x7d\x03\x9b\x37\x97\xbe\x52\xa9\x12\x24\xf9\xb5\xbb\x2c\x3b\x68\x26\xe7\x8c\x77\x52\xfc\x9a\x9b\x9c\x13\x28\x2f\xf2\x2c\x1e\x92\x7b\xd9\xa0\x6d\x6b\xc6\x36\xbf\xa7\x37\x21\x95\x8a\x60\x22\xdf\xb2\xbf\x45\x9e\x61\xae\xb1\x42\xd3\x65\x5e\x95\x77\x7e\x09\xb7\x49\xa0\x0f\xb0\x78\xe3\x68\xb3\x30\xe3\x8a\x4f\x24\xc1\xe3\x6c\xd4\x75\x12\x96\xfa\xb5\xf6\xfb\x4d\xd3\x8c\x8a\xba\x1f\xbb\xcf\xdb\x25\x88\x36\xf8\xa2\x4b\xa8\xd0\x2c\x1d\x94\xa5\xdb\x99\xd5\x10\x02\x73\x30\x80\x5e\x47\x2b\x3e\x71\x0e\x16\x74\xd8\x46\x05\x61\x7e\x82\xf1\xe9\x91\x47\xeb\x5e\x9c\x04\x90\xb4\x43\x06\xbb\x43\x2b\xd6\x73\x99\x5f\x0c\xe7\xb3\x42\x89\x99\x1b\x1b\x81\xa3\xe9\x04\x12\xf6\xaf\xcb\xfd\x41\xf6\x7e\xe4\x93\x83\xa1\x44\x1a\x0e\x81\x80\x6f\x6c\x7e\xb7\xff\x9d\x5e\x9d\x22\xe0\x7e\xb5\x07\x9f\x68\x76\x5a\x9b\x13\xfe\x90\xf1\x30\x32\xe9\x3e\xeb\xfc\x60\xd7\x5f\x57\xb9\x9f\x04\x75\x2a\xca\xf8\xec\xb5\x47\x35\xdd\x9e\x0e\x7c\x84\x18\x19\xca\xab\x64\xf0\x1d\x62\xda\x86\x2b\xc6\x35\x56\xbf\xdd\x2f\x78\xed\x09\x72\xcb\x4c\xd1\xe3\x23\xe0\xe2\xb3\xb9\x21\x33\xb9\xbd\xcf\x3c\xad\x8c\xbb\x2b\xa4\x81\x7c\x8c\x8d\x69\x2b\x5a\x6d\x29\x7f\x03\x3e\xa0\xe6\x70\xfd\xd2\xce\x0a\x97\x8d\x8f\x93\x32\x08\x89\xc3\x3c\xdf\xf9\x44\xed\x16\x5d\x82\x0b\x5a\x2e\x3a\x72\x2d\xe7\x11\xac\xe8\xf9\x16\x08\x62\xc8\x8d\x84\xdc\xcb\xc2\x74\x1d\xb2\x74\x69\xcd\x46\x0c\x4f\x6a\x9f\x25\x03\x77\x6e\xd9\x93\x69\xbc\xb1\x9b\x73\xbf\xff\x54\x3d\x15\xe1\xae\x47\xc9\x6c\xf9\x06\x9f\xdf\x07\x22\xb1\xe4\xdb\x6a\x8f\x1a\x56\x28\xd9\x7b\x37\x12\x47\x59\x9d\xc6\xaf\xf6\x5d\xcd\xb3\x6f\x57\x1b\xae\x2b\xa1\x7f\x48\xe6\xab\xee\x78\xfe\x91\xa1\x3d\x56\xaa\xc8\xa2\xdb\xad\x11\x33\x7c\xe4\x83\x95\x05\x55\x67\x75\xaf\x57\x28\x4b\x07\x11\xad\xae\x58\xde\x61\x70\xbc\x75\x57\x6a\xfc\x72\x2f\x6a\x6e\x79\x18\xe7\x9d\xac\xc8\x46\xe2\x7f\xbd\x97\x61\xe9\x26\x31\xbb\x58\x63\x30\x9b\x64\xc3\xfd\x7c\xbc\xe6\x92\xff\xaf\xc0\x3f\x01\x00\x00\xff\xff\x04\x4b\x56\xa9\x22\x04\x00\x00")

func resCertPemBytes() ([]byte, error) {
	return bindataRead(
		_resCertPem,
		"Res/cert.pem",
	)
}

func resCertPem() (*asset, error) {
	bytes, err := resCertPemBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "Res/cert.pem", size: 1058, mode: os.FileMode(420), modTime: time.Unix(1717495720, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resGoogleJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\xcf\x4b\x1b\x41\x14\xbe\xfb\x57\x3c\x73\xd0\x5d\x08\xbb\x05\x6f\x86\x15\x8a\x78\x28\x94\x20\xa5\x07\xa5\x94\x32\x24\x2f\xee\xc2\x64\x77\x99\x7d\xeb\x0f\x4a\xc0\x8a\x3f\xa2\x68\xe9\xa5\x62\x45\x5a\xa1\xd4\x53\xad\x9e\x5a\xc5\xd2\xbf\xc6\x9d\xd4\x53\xff\x85\x32\x33\xdb\x24\x6b\x92\xe2\x1c\xb2\x79\xf3\x7d\xef\x9b\xef\xcd\xbc\x99\x3a\x23\x06\xde\x0c\xbc\x1e\x03\x00\xe0\x48\xb0\x10\x33\xf2\xc1\x03\x4b\x7d\xcb\x80\x1c\xbd\x30\xe5\xbc\x0c\x8d\x80\x13\x0a\x1d\xd8\xde\x8c\x49\x50\x23\x68\x80\x85\x1c\xc1\xf3\x40\x63\x3d\x44\x0d\x8d\x40\x3d\xaa\xa5\x4d\x0c\xa9\xd2\xc5\x5a\xdd\x7f\xcb\x4c\x80\xc0\x24\xe5\xd4\x47\x74\x70\x99\xf1\x94\x11\xf6\x6c\x94\xc1\xf8\x58\x98\x67\xe4\x3f\xd3\x09\xce\xe3\xea\xe2\xab\xe7\x8b\xf3\x73\x06\xb3\x2b\x05\xd1\x30\xaa\x63\x02\x1e\xbc\x78\x39\x38\x0f\x5e\xbe\xa6\x13\x10\x0a\x46\x58\xc5\x55\xb2\xfa\x04\x5c\x17\xee\xde\x1c\x64\x6f\xb7\xe5\xee\xba\x3c\xd9\xcd\xf6\xaf\xef\xb6\x0e\x3a\xc7\x9b\xbf\xf7\x36\x3a\x1b\xd7\x5d\xda\x8a\x1f\x70\x04\x4b\x49\xda\x50\xac\xdc\x75\x41\x7e\x5e\x97\x9f\xbe\x74\xda\xdf\x4d\x96\x3c\xba\xc8\xde\x9d\xc9\xa3\x8b\xa4\x26\x82\x98\x64\xfb\x30\xa1\x35\x8e\xf2\x74\xa7\x73\xfe\xab\x90\xab\x6c\xc6\x4c\x60\xa8\xf6\x44\x89\x3b\x26\xaa\x46\x75\xac\x14\x98\x6a\xf7\x73\xe6\xc4\x44\x9e\xe3\x10\x5b\xaa\xb2\x26\x3a\x14\x3d\x8d\x56\x50\xcc\xb2\x04\x2d\x1b\xc6\x3d\x0f\x26\xcd\xd2\x93\x0f\x24\x2b\x7b\x93\xf7\x0b\xcb\x8b\xcb\xce\x36\xe4\xc7\x93\xdb\xab\x03\x79\x74\xf1\xe7\xe7\x7e\xd6\xfe\x90\x5d\x6e\xcb\xc3\x1d\x79\xf2\xd5\x94\xdb\x39\xde\xcc\xb6\xb7\xb2\x6f\xd7\xf2\xc7\x4d\xb6\x77\x9a\xb5\x2f\xe5\xfb\xcb\xce\xcd\xe6\xed\xd5\xf9\x80\x9e\x2a\x43\xd7\x49\xb8\x4a\xb3\x51\x48\xda\x99\x08\x9a\x96\xed\x70\x0c\x97\xc8\x87\x19\x78\x64\x0f\x1a\xf9\x97\x6c\xba\x13\xc6\x4d\x13\xaa\xf2\x68\x2d\xc6\xa8\x91\x03\xb6\x6a\xcf\x52\x23\x0d\x6b\x14\x44\x61\x69\x84\x50\x51\xcc\x9c\xea\x7f\xa8\x6a\xe8\x26\x73\xe2\x34\xf1\x0d\xbd\x32\x92\xdd\x1a\x8a\xb4\x90\x27\x38\x7a\x89\x87\xc9\x0f\x4a\x8f\x90\x2d\x12\x8b\xd1\x83\xae\x45\x2f\x45\x20\xa5\x22\x34\x06\xc7\x0c\xd4\x7d\x43\x90\xeb\x8b\xa7\x9f\x12\xab\xe4\xba\xf5\x60\xd9\x4d\x62\x16\xba\xcc\xf5\xa7\x4a\x76\x97\x97\x20\x13\x35\xff\x09\x61\xb3\xff\x9e\x36\x22\x61\x29\xd4\x9f\x82\xa8\xa1\xb5\xfa\xce\x40\x01\x14\x90\x7e\x55\xfc\xa9\x21\xed\x52\x29\x50\x53\xc1\x0d\xd1\x34\xfb\x1c\x47\xfd\xba\xf8\x02\x1b\x45\x62\x1d\x93\xda\x30\x66\x31\x1a\xd2\x9d\x02\x63\xce\x6a\x68\xa5\x82\x97\x4b\xa5\x5e\xac\x4d\xea\x99\x61\xbe\x02\xc2\x26\x78\xf7\x2e\x96\x4e\x99\x36\x9f\x72\x01\x49\x05\x9f\x56\x3f\xc5\x59\xe5\x79\x5a\xff\x0e\x39\x9f\xbe\xbd\x35\xfd\xa3\x96\xcc\x5d\x18\x56\x7e\x82\x7d\xc4\xca\x58\xeb\x6f\x00\x00\x00\xff\xff\xe5\x44\x9a\xdc\x11\x06\x00\x00")

func resGoogleJsBytes() ([]byte, error) {
	return bindataRead(
		_resGoogleJs,
		"Res/google.js",
	)
}

func resGoogleJs() (*asset, error) {
	bytes, err := resGoogleJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "Res/google.js", size: 1553, mode: os.FileMode(420), modTime: time.Unix(1717471063, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resGoogle_newsJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\xdf\x4b\x1b\x4b\x14\x7e\xf7\xaf\x38\xe6\x41\x77\x21\xec\x5e\xf0\xcd\xb0\xc2\x45\x7c\xb8\x97\x4b\xae\xb4\x7d\x50\x4a\x29\x43\x72\xd6\x5d\x98\xec\x86\xd9\xb3\xfe\xa0\x04\xac\xf8\x23\x8a\x96\xbe\x54\xac\x48\x2b\x94\xfa\x54\xab\x4f\xad\x62\xe9\x5f\xe3\x4e\xea\x53\xff\x85\x32\x33\xdb\x24\x6b\x92\xe2\x3c\x64\x73\xe6\xfb\xce\x37\xdf\x99\x39\x33\x75\x46\x0c\xbc\x19\x78\x31\x06\x00\xc0\x91\x60\xa1\xc9\x28\x00\x0f\x2c\xf5\x2d\x03\x72\xf4\xa2\x94\xf3\x32\xf8\x21\x27\x14\x3a\xb0\xbd\x19\x93\xa0\x46\xe8\x83\x85\x1c\xc1\xf3\x40\x63\x3d\x44\x0d\x8d\x40\x3d\xae\xa5\x0d\x8c\xa8\xd2\xc5\x5a\xdd\x7f\xcb\x4c\x80\xc0\x24\xe5\xd4\x47\x74\x70\x99\xf1\x94\x11\xf6\x6c\x94\xc1\xf8\x58\x98\x67\x14\x3c\xd2\x09\xce\xdf\xd5\xc5\xe7\x4f\x16\xe7\xe7\x0c\x66\x57\x0a\xa2\x51\x5c\xc7\x04\x3c\x78\xfa\x6c\x70\x1e\xbc\x7c\x4d\x27\x24\x14\x8c\xb0\x8a\xab\x64\xf5\x09\xb8\x2e\xdc\xbd\x3c\xc8\x5e\x6d\xcb\xdd\x75\x79\xb2\x9b\xed\x5f\xdf\x6d\x1d\x74\x8e\x37\x7f\xec\x6d\x74\x36\xae\xbb\xb4\x95\x20\xe4\x08\x96\x92\xb4\xa1\x58\xb9\xeb\x82\xfc\xb0\x2e\xdf\x7f\xec\xb4\xbf\x98\x2c\x79\x74\x91\xbd\x3e\x93\x47\x17\x49\x4d\x84\x4d\x92\xed\xc3\x84\xd6\x38\xca\xd3\x9d\xce\xf9\xf7\x42\xae\xb2\xd9\x64\x02\x23\xb5\x27\x4a\xdc\x31\x51\x35\xae\x63\xa5\xc0\x54\xbb\x9f\x33\x27\x26\xf2\x1c\x87\xd8\x52\x95\x35\xd0\xa1\xf8\xbf\x78\x05\xc5\x2c\x4b\xd0\xb2\x61\xdc\xf3\x60\xd2\x2c\x3d\xf9\x40\xb2\xb2\x37\x79\xbf\xb0\xbc\xb8\xec\x6c\x43\xbe\x3b\xb9\xbd\x3a\x90\x47\x17\x3f\xbf\xed\x67\xed\xb7\xd9\xe5\xb6\x3c\xdc\x91\x27\x9f\x4c\xb9\x9d\xe3\xcd\x6c\x7b\x2b\xfb\x7c\x2d\xbf\xde\x64\x7b\xa7\x59\xfb\x52\xbe\xb9\xec\xdc\x6c\xde\x5e\x9d\x0f\xe8\xa9\x32\x74\x9d\x84\xab\x34\x1b\x47\xa4\x9d\x89\xb0\x61\xd9\x0e\xc7\x68\x89\x02\x98\x81\xbf\xec\x41\x23\xbf\x93\x4d\x77\xc2\xb8\x69\x42\x55\x1e\xad\x35\x31\xf6\x73\xc0\x56\xed\x59\xf2\xd3\xa8\x46\x61\x1c\x95\x46\x08\x15\xc5\xcc\xa9\xfe\x81\xaa\x86\x6e\x32\xa7\x99\x26\x81\xa1\x57\x46\xb2\x5b\x43\x91\x16\xf2\x04\x47\x2f\xf1\x30\xf9\x41\xe9\x11\xb2\x45\x62\x31\x7a\xd0\xb5\xe8\xa5\x08\xa4\x54\x44\xc6\xe0\x98\x81\xba\x6f\x08\x72\x7d\xf1\xf4\x53\x62\x95\x5c\xb7\x1e\x2e\xbb\x49\x93\x45\x2e\x73\x83\xa9\x92\xdd\xe5\x25\xc8\x44\x2d\xf8\x87\xb0\xd1\x7f\x4f\xfd\x58\x58\x0a\x0d\xa6\x20\xf6\xb5\x56\xdf\x19\x28\x80\x42\xd2\xaf\x4a\x30\x35\xa4\x5d\x2a\x05\x6a\x2a\xb8\x21\x9a\x66\x9f\xe3\xa8\x5f\x97\x40\xa0\x5f\x24\xd6\x31\xa9\x0d\x63\x16\xa3\x21\xdd\x29\xb0\xc9\x59\x0d\xad\x54\xf0\x72\xa9\xd4\x8b\xb5\x49\x3d\x33\xcc\x57\x48\xd8\x00\xef\xde\xc5\xd2\x29\xd3\xe6\x53\x2e\x20\xa9\xe0\xd3\xea\xa7\x38\xab\x3c\x4f\xeb\xdf\x21\xe7\xd3\xb7\xb7\xa6\x7f\xd4\x92\xb9\x0b\xc3\xca\x4f\xf0\xdf\xc7\xff\x57\x9d\x84\x44\x18\x2d\x85\xfe\x9a\xd5\x97\x67\x57\xc6\x5a\xbf\x02\x00\x00\xff\xff\xb9\x16\xae\x86\x21\x06\x00\x00")

func resGoogle_newsJsBytes() ([]byte, error) {
	return bindataRead(
		_resGoogle_newsJs,
		"Res/google_news.js",
	)
}

func resGoogle_newsJs() (*asset, error) {
	bytes, err := resGoogle_newsJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "Res/google_news.js", size: 1569, mode: os.FileMode(420), modTime: time.Unix(1717495686, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resJs_textJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x55\xcf\x8f\xdb\x44\x14\xbe\xe7\xaf\x78\xeb\xc3\xae\xd3\x04\x27\x5d\x58\xd0\x92\x38\xd2\x2a\xf4\x50\x09\x96\x0a\x8a\x44\xd5\x2d\xd5\x28\x79\x8e\x67\x99\xcc\x58\xe3\x71\x37\x15\xcd\xa1\x42\x82\x22\xa0\xe2\x42\x41\x70\x80\xc3\xaa\x5c\x58\xf1\x43\x1c\x56\xab\xfe\x39\x9b\x6d\x39\xf5\x5f\x40\x33\x76\xec\x89\xed\xb4\x3d\x92\x53\xec\xf7\xbd\xf7\x7d\xef\x9b\x79\xcf\x63\xa2\x08\xf8\x03\xf8\xac\x01\x00\x70\x87\x48\x98\x45\x44\x85\xe0\x83\xd3\xe9\x28\x9c\x29\xb7\xe9\xf4\xf2\x98\xc4\x38\x61\x0a\x7c\x18\x8b\x51\x32\x45\xae\x3c\xbc\x43\x58\x42\x14\xba\x26\xad\x9d\x07\xda\xc0\x13\xc6\xda\xf0\xf1\x35\xa2\xc2\x0f\x4c\x9a\xb7\xb7\x7f\xe3\xf6\xf5\x1b\xd7\xae\xa4\xb1\x66\x51\x56\xf3\xc4\xe0\xc3\xcd\x5b\xc5\x3b\x2e\xc6\x08\x7e\xc6\xe8\x51\x85\x92\x28\xdc\x37\x82\x52\x10\x43\x05\x8a\x2a\x86\xb6\x1c\xf3\xc2\x53\x92\x4e\x6d\x18\xa3\xfc\x53\x1b\xc5\xc4\x88\x28\x2a\xb8\x17\x4a\x0c\x0a\x18\x0d\x0e\x93\x69\x04\x3e\x04\x84\xc5\xd8\x6b\x14\x81\xf8\x0a\x9f\x30\x1a\x87\xef\x91\x43\x21\xa9\xba\x0b\x3e\xc4\x4a\x16\xbe\x2d\x81\x98\xc2\x86\x22\xe1\xda\xa5\x6e\x6f\x25\x3a\x0a\x29\xc7\x18\x4b\xd1\x1c\xd2\xe9\xc0\xbf\xf7\xbf\x5d\x3c\xfc\x62\x71\xf2\xc3\xd3\xdf\x1f\x9f\x9f\xfe\x9d\x87\x02\x21\xc1\x35\x4a\x4c\x1e\x50\xe8\x6b\x01\x1e\x43\x3e\x51\x61\x0f\x68\xab\xd5\xb4\xa4\x8c\x04\x8f\x35\x1d\x91\xc3\xd4\x45\x8d\x5d\x3e\xee\x29\x97\x36\xab\xd4\x8b\x07\xc7\x17\x8f\x4e\x52\xea\x8b\x1f\xff\x58\x7c\xf7\xf8\xfc\xf4\xec\xd9\xd7\x7f\x5d\x3c\xfa\x32\x7d\xf9\xfc\xc9\x83\xc5\x2f\x67\xe7\x67\x0f\xf7\x3e\x1c\x5e\xbd\xfa\xf4\xd7\xfb\xcf\x9f\x7c\x95\x17\xa0\x01\xb8\x39\xdf\xc0\x87\x37\x77\x60\x73\xb3\x50\xd0\xf7\x61\xb7\x0b\xf7\xee\x81\x8d\xd9\x7d\xab\x8c\xb9\xbc\xbd\x6d\xb7\xa1\x7f\xb6\xa3\xad\x56\xe1\xe7\xfc\xa5\xda\xcf\x4f\x4f\xaa\xda\x3f\xe2\x74\x24\xc6\xf8\xec\x9b\xcf\x17\x3f\xff\x63\x37\x80\x2c\xc6\x4a\x17\xdd\xd9\x1b\xd8\xed\x96\x55\x76\x67\xbb\x01\xd9\x29\x0b\xb5\x0f\xb7\x5e\xe8\x7c\x8d\xe7\xa9\xc9\xb9\xee\xc5\xf1\x6f\x8b\xe3\x9f\x2e\xbe\xff\x33\x47\x4a\x54\x89\xe4\xab\x97\x6b\xb0\x42\x98\xd2\xa5\x04\x47\x21\x65\x08\xae\x9e\x1f\x5b\xa3\x9e\xa9\x88\x48\x34\x57\x4f\x07\xbd\xf4\x69\x5f\x8c\xb1\xb7\x72\x8e\x19\x6a\x73\x33\x85\x29\x32\xd9\x27\x53\x84\x0d\x1f\x9c\x58\xdd\x65\xe8\xd4\x87\x46\x92\x46\xca\xc4\xd2\x02\xcb\xa8\xa7\xc4\xbb\xe2\x08\xe5\x90\xc4\xe8\x36\x61\xc3\xf7\x61\x2b\x05\x6f\xbd\x22\x58\x93\xbe\x0a\x16\x1c\x1a\x48\x32\x45\xa7\x7c\x36\xba\xab\x54\x30\xce\xd4\x50\x70\x65\xaa\x98\x25\x91\x8d\x10\x0c\x60\xa7\xb9\x9a\x94\xef\x98\x59\xee\x58\x35\xbd\x57\x49\xd1\xb3\x27\x18\x7a\x4c\x4c\x5c\xe7\x1d\x24\x8c\xf2\xc9\xdb\x4e\x5b\xcd\x54\x0d\xd8\x2c\x3e\x2f\x4a\xe2\x70\x9d\xbe\x9a\xa4\xf2\x9a\xb2\x63\x73\x7d\x8f\xab\x7d\x68\x03\x36\xb2\xc4\x9a\x36\x2d\x19\xce\x01\x77\x5e\xc4\xa9\x64\x82\xd5\xf0\xbc\x51\xff\x54\xfc\x7b\xe9\x3a\x9f\xe7\x9b\x9f\x30\x76\x1d\x8d\xe9\xa9\xae\x43\x41\xb9\xbb\x75\xc0\xb7\x32\x64\x44\x24\x99\x48\x12\x85\xfa\x93\x91\x81\xbd\x38\x62\x54\x69\xf5\x5a\xbf\x81\x25\x11\x13\x64\x7c\x3b\x3d\x3e\xc7\x31\xef\x02\x21\xdd\xec\x83\x80\x20\x02\xab\x94\xe5\x8a\x36\xab\x48\x5e\x5e\x90\x0d\x1f\xba\x25\xeb\x2c\x86\x96\x0f\xda\xb9\x9a\xc6\x35\x9d\xa6\x01\xdf\xb0\x7a\x12\x23\x46\x46\xe8\x76\xfa\x37\x3f\x19\xdc\xba\x34\xf0\x2e\xf5\x0f\x3a\xde\xa5\x41\x67\xd2\x06\x27\x93\x0e\xc5\x24\x12\x8f\xf2\x31\xce\xde\x0f\x5c\xa7\x6f\x06\x01\x9c\x26\x0c\xe0\xb5\xcb\x50\x12\x33\x12\x5c\x51\x9e\x60\x8d\x84\x6a\xcd\xac\xa7\x3e\x6c\x77\xb3\xb9\xb2\x78\xc0\x69\xf6\x2b\xbd\xbe\xa0\x7c\x5e\x35\x9b\xa8\xbc\x90\x3e\x0a\xa3\xb5\x54\xab\xee\x14\x8a\x03\x74\x9a\x35\x37\x54\x73\x54\xbe\xc2\xae\x2e\x50\x87\x5e\xfa\x7e\x24\xe4\x38\x5e\x1a\x9f\xa9\xcb\x78\xc0\xb2\xba\x4c\x64\xd2\x0a\x8f\x5e\x5f\xc3\x50\x6b\x8b\xfd\x5b\x33\x8c\xcb\x5f\xf5\xfa\xb4\x2c\xa1\xf5\x15\xab\x93\xb7\x9e\x43\xb7\x62\x0a\xda\xa7\xfd\xbf\xee\xa5\xf2\x66\xdd\x5a\xa9\x8a\xa9\xec\x7b\xfb\x9e\xd7\x2f\xf7\x1a\xcd\x3a\x69\x95\xb3\x52\x7f\x6e\x6d\xab\xec\xc3\x5c\x14\x6a\xcc\x1b\xff\x05\x00\x00\xff\xff\x90\xea\xcf\x01\x51\x0b\x00\x00")

func resJs_textJsBytes() ([]byte, error) {
	return bindataRead(
		_resJs_textJs,
		"Res/js_text.js",
	)
}

func resJs_textJs() (*asset, error) {
	bytes, err := resJs_textJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "Res/js_text.js", size: 2897, mode: os.FileMode(420), modTime: time.Unix(1717468748, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resKeyPem = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\xd5\x37\x12\xab\xc8\x02\x85\xe1\x9c\x55\xdc\x5c\xf5\x0a\x10\x3e\x78\x41\x37\xb6\x01\x09\x6f\xa4\x0c\x6f\x84\xb0\xc2\xae\x7e\x6a\x6e\x3a\x27\x3c\xd1\x9f\x7d\xff\xfb\x77\x50\x56\xd1\xf3\x8f\xed\xa2\x10\xf8\xf2\x1f\x43\x7e\xfd\x7d\xb1\x07\x42\xf2\xe6\x20\x08\x24\xf0\x84\xd5\x67\xaa\x3f\x8d\x2a\xec\x04\x04\x8e\xac\x00\xe0\x89\xd0\xc8\xf6\xaa\xf2\x5a\x50\xc9\x00\x0c\x08\x02\x47\x7a\x87\x16\xfb\x90\xed\x86\x8c\xc6\x8b\xc1\x84\xda\xc0\xef\xe6\x4e\x07\x74\x3f\xda\xbd\xc7\x97\xf3\xc3\xe5\xf4\x4a\xe9\xe9\x7e\x33\x5f\x11\xc1\x14\x51\xdf\x7a\x29\x0e\xfd\xae\x81\x2d\x09\x54\xba\xcf\xcb\x63\x1b\x8f\x0d\xf4\x69\x75\xbd\x30\xc8\x76\xf2\x14\x58\x1d\xb7\x79\x9c\x20\xf0\x6c\x31\xa9\x30\xcc\x98\x97\x51\x33\x4b\x97\xb5\x75\x2c\x88\x1a\x9b\x99\xc1\x32\x9e\x93\x24\x5c\xf9\xa7\x46\x9c\x60\x54\x34\x23\x8c\xd6\x9b\x16\x06\x6c\x89\x3a\x57\xac\x0c\xd3\x90\xd3\xb2\xcc\xa6\xa5\x68\x0d\xeb\x1d\xfb\x79\xb4\x7a\x7c\x3b\x38\x42\x0b\xe9\xd8\xd1\x7d\x2b\xb2\x3d\x53\xb3\xcf\xe0\xc9\x43\xcb\xd1\xa0\x5e\x3c\xd1\x1c\xb3\x35\x81\xb1\xdc\x4c\x25\x27\x05\xfa\xfd\xf1\x86\x14\x1b\xce\x47\x9f\x12\xe8\x46\xdf\xaa\x70\x89\x9a\xb2\xe7\xc9\x41\x0f\xcd\x1b\x83\x0a\x22\x66\x73\x6f\x97\xbe\x9c\x43\xf2\xb6\xa6\x5f\xa1\x80\xe2\xf3\xc0\x2e\xe1\x97\x26\x3e\x19\x79\xea\x02\x4e\x7b\xdb\xcc\xca\x7f\xe5\x6b\xc0\x82\xcf\x5b\x6b\xdb\xc8\x2f\xd5\xac\x75\x26\xbc\xef\x37\x41\x64\xb6\x71\xfd\x3d\x25\x1e\xf5\x30\x51\x59\x3e\x97\x0d\x15\x60\xc2\x39\x52\x7a\xbd\xf8\xa0\x7a\x40\x00\x64\xb1\xaa\x64\x08\xf4\x58\x7b\x98\x45\x8a\xa6\x63\x6c\x35\x0d\xa0\x1f\xee\x7c\xcc\x0c\xb1\xec\x79\x94\xea\x5b\x5a\x88\xaf\x2b\x8f\x7b\x5e\x06\xce\x1d\xbb\x82\x40\xae\xd8\x3c\x7d\xd5\xd3\x25\xdb\x1c\x97\x64\x1c\x3b\xb1\xe6\x8b\xb3\x63\xc3\x59\x76\xfd\xaa\x1d\x1e\xdc\xa9\x25\x0b\x45\xc2\x95\x86\x82\x17\x96\xb9\x43\x2c\x0f\x4a\x41\x26\x39\xae\xc6\x68\x26\x2c\x91\xd7\x8f\x7c\xdf\xe3\xbf\xde\x53\xea\xc6\x22\xcb\x77\xe1\xc4\x57\x86\x86\x89\xf0\x9e\x55\xd8\x70\x04\x40\x39\x70\xed\x9c\x6b\x9a\x73\x91\xca\xf8\xeb\x82\x86\xc5\xdf\x6c\x1f\xe3\x98\xeb\xd4\xac\x9d\x6e\x2e\x63\x9d\x28\x6f\x6e\xef\x5f\xb7\xcc\xea\x15\x56\x29\x52\x8b\xf8\x30\x59\x7d\x24\x05\x7b\x0b\xf4\x7c\x9a\xf9\xb5\xbb\x3f\x6e\x8d\x37\xc7\x83\x72\x0f\xaf\xed\x98\x73\x08\xb1\x0f\x4a\xac\x31\x5d\x94\x98\xf4\xaa\x4f\x76\x4d\x64\x41\x28\x7c\x20\xdf\x40\x5b\x7d\x89\x60\xd4\xe4\x3e\x94\x15\x72\x34\xbf\xa2\x1e\x2f\x71\x73\x92\xe2\x6a\xc6\x46\xbd\xc2\x5a\x8a\xea\x63\x6f\xb0\x33\xb9\xcc\x3b\xf5\xea\x5b\x42\x51\x40\x3d\x3a\x3c\x15\x5a\x44\xbc\x73\x8a\x3b\x9b\xda\xcb\xeb\x84\xbe\x55\x51\xbe\x9c\x6b\x2a\x8b\xd5\x4b\x06\x9c\x05\x88\xeb\xae\xd0\x5c\xd7\x5e\x99\x9c\x00\xec\x65\xe8\x1d\x84\xc4\xca\xf7\x20\x5c\x19\x8a\x60\x91\x7b\x3b\x6b\x12\x25\xbf\x0d\xca\x5e\xe4\x9c\xd3\x83\xc9\x73\xab\x7f\xeb\xfb\xe9\x7c\x07\x3b\x08\x5d\xe0\x7c\xa9\x8c\x73\xc5\x96\x0c\xfc\x11\xc3\xc3\x85\x16\xcf\x51\x77\x2b\xd0\x9a\x2c\x2e\x2d\xcd\xd8\xcc\xb2\xfa\xac\x34\x3c\x7c\x0c\x17\xca\xc8\x98\xab\xf4\x82\x9c\xd6\xb4\x8e\xa6\xc5\xe6\x7d\xa3\x8e\xe1\x5e\x99\xe6\xf7\xb1\xa7\x86\x8f\xa5\x42\x7b\xa8\xbe\xa0\x7e\xe7\xc5\xcb\xe5\x53\x1a\xcc\xa5\xe5\xbb\x7c\xa4\x3e\xc4\xdf\x64\x76\x95\x74\x6d\x58\xdc\x74\xf8\xac\x89\x6e\xce\x84\x62\xdc\x85\x74\xbe\x34\x7a\xd5\xb7\xe8\xb1\x61\x00\xb0\x47\x49\x77\xaa\xae\xad\x01\xfe\xf6\xe4\x1c\xbd\xf0\xcb\x7c\x26\xa2\xf1\x96\xa9\x22\x3e\x88\xd5\x2d\x84\x7b\xfe\x5e\x68\x6f\xbd\xdc\x9b\x1a\x4b\xef\x54\x69\xa4\xf9\x12\xf8\xcc\x5c\xba\x15\x73\xed\xb1\x4f\x57\xb3\x69\xa2\xc5\x10\xb6\xee\xac\x76\x8d\x6a\x9e\x86\xb3\xa5\xfd\x1a\xe5\x49\x73\xe5\x4a\x2d\x2a\x44\x97\xb1\x95\x32\x50\x6a\xe3\x49\xdc\x7b\x38\x24\xf6\xc4\x13\x12\x7e\x1e\x07\xb6\x95\xf8\xce\x85\x14\xeb\xff\x78\xb1\x7a\x41\x4e\x06\xb0\xdc\x02\x27\xfa\x40\x7c\x7c\x9a\xb9\x50\x5a\xd2\xcd\x69\x96\xfa\x74\xbf\x30\xa6\xb9\xdc\x2e\x9b\x5a\xd4\xa3\x88\x1b\xdd\x32\x14\xfd\x0e\x43\x45\x41\xfb\xd7\x7d\x80\x14\x1d\x82\xdd\x3f\x2d\xe1\xf1\x9d\x6f\x44\x6c\x23\xaa\xb6\xd7\xc3\xfd\x41\x7d\x5c\x07\x26\xba\xb8\x1f\x2d\xfe\x9a\xea\x77\xf7\x1e\xea\x4b\x9f\x6c\xfc\x26\x39\xac\x54\x61\xde\xf2\x2b\x01\xa5\x17\x21\xe3\xf1\x75\x63\x8a\x26\x5a\x07\x7f\x73\x07\x17\xa7\x5e\xc3\x8d\x22\xcd\x0c\x04\x71\xf6\x73\x3a\x01\xbf\x35\xbc\x16\x6d\x77\xf3\x57\x6b\x07\x0f\x1c\x03\x56\x50\x9f\x30\xe0\x4e\x0a\xab\x53\x17\x45\xaa\x76\x75\x29\x31\xbb\x38\x9d\x56\x32\x07\x12\x95\x84\x23\xc6\x92\x42\x49\x68\x0c\x67\x3e\x71\x49\xbe\x8a\xe7\x59\xa5\x3d\x3d\x27\x87\xcf\x99\x0c\x20\x16\xb6\xc7\xa8\x27\xfa\x8c\x8d\xba\xde\x09\x92\xec\x02\xd9\x2d\x72\x76\x70\x6a\xcf\x16\x48\x6b\x5e\x8e\xdf\xe1\xb7\x6c\xed\xee\x61\x0b\x6c\xde\xbe\xdf\x8f\x1f\xef\xec\x3d\x2c\x63\x86\xe1\xca\x32\x2e\xc1\x80\x31\xa2\x17\xae\x96\x7d\xc6\x1f\xee\x21\x09\xdc\x4d\xfb\x52\x4b\xfc\x24\x88\x5d\xae\xa9\x44\xf4\xb7\x67\xa1\xab\x7a\x61\xe5\x60\x50\x81\xca\xe0\xa2\x7e\xa9\x89\x03\x8c\x8f\xce\xa9\xdf\x69\x60\x44\xec\xde\x51\x09\x0d\x2e\x39\x39\xbe\xe8\xfb\xed\xea\x76\xce\xdf\x76\x05\x1e\x99\x4a\xbb\xf7\x64\x4d\x8c\xca\xba\xba\x62\xe8\x32\x4d\x5f\x94\x48\x85\x68\xd4\x07\x98\xe0\x13\x7b\xe3\x2e\xd4\x1b\x14\xb6\xa1\xf0\x08\x3f\x6e\xf0\xc2\xf9\xc2\x99\xac\x4d\x53\x98\xa9\xa8\x4d\x9e\x72\xe7\x8b\x20\x66\xa6\xe8\x6b\x69\x77\xa2\x08\xe7\xa8\xa4\xc9\xc7\x5f\x22\x03\x4a\xa7\xa4\xe7\xf7\x8e\x3b\xc9\x83\xf3\xb1\xa5\x4e\xf2\x43\xa4\xea\xa1\xf4\x89\x09\x8e\x6f\xad\xcc\xf6\x05\x82\xff\x63\x7f\xdd\x91\x9f\xd2\x7f\x2d\xfa\x27\x00\x00\xff\xff\x42\x8a\x2c\x15\xa8\x06\x00\x00")

func resKeyPemBytes() ([]byte, error) {
	return bindataRead(
		_resKeyPem,
		"Res/key.pem",
	)
}

func resKeyPem() (*asset, error) {
	bytes, err := resKeyPemBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "Res/key.pem", size: 1704, mode: os.FileMode(384), modTime: time.Unix(1717495720, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resNodeXIni = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xc1\x8e\x83\x20\x10\x86\xef\x3c\x05\x87\xcd\x1e\xf1\xbe\x09\xf1\xb4\xa7\x7d\x82\x8d\x31\x04\x65\x8a\x58\x05\x5b\x30\xd8\x34\x7d\xf7\x06\xd0\x6a\x48\x0f\x3d\x90\xc0\x3f\xdf\xc7\x4c\xa6\x12\x70\xe2\xf3\xe0\x6a\x84\x1c\x2c\x8e\xf5\x16\x53\xdc\x5b\x16\x1e\xa4\xb7\x08\xa1\xca\x02\xbf\xb6\x1d\x03\x2d\x95\x86\x1a\x21\x69\x8c\x1c\x00\x53\xdc\x39\x37\xd9\x9f\xa2\xf0\xde\x93\x14\x92\xd6\x8c\x45\xe2\xcb\x0b\xfd\xba\xff\xfd\xfe\x3f\x56\x9e\x69\xf0\xf6\x63\xe9\xdb\x35\x23\xd5\x3e\xf4\x6f\xb8\x12\x73\x26\xc6\x2c\x79\xa5\x17\x5b\xa3\x46\x69\x99\x83\x4a\xcb\xb7\x43\xdd\x78\x67\xcc\x01\x4e\x75\x12\xe3\xa3\x30\xed\x82\x16\xb0\x64\xdf\xa7\x30\xf2\xe1\xba\x3a\x61\x77\x9b\x26\xe6\xf6\x1c\x8e\x34\x99\xba\x17\xa2\xfe\x9a\xec\x19\x00\x00\xff\xff\x4d\xbe\x05\xee\x91\x01\x00\x00")

func resNodeXIniBytes() ([]byte, error) {
	return bindataRead(
		_resNodeXIni,
		"Res/node-x.ini",
	)
}

func resNodeXIni() (*asset, error) {
	bytes, err := resNodeXIniBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "Res/node-x.ini", size: 401, mode: os.FileMode(420), modTime: time.Unix(1717483614, 0)}
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
	"Res/cert.pem":       resCertPem,
	"Res/google.js":      resGoogleJs,
	"Res/google_news.js": resGoogle_newsJs,
	"Res/js_text.js":     resJs_textJs,
	"Res/key.pem":        resKeyPem,
	"Res/node-x.ini":     resNodeXIni,
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
	"Res": &bintree{nil, map[string]*bintree{
		"cert.pem":       &bintree{resCertPem, map[string]*bintree{}},
		"google.js":      &bintree{resGoogleJs, map[string]*bintree{}},
		"google_news.js": &bintree{resGoogle_newsJs, map[string]*bintree{}},
		"js_text.js":     &bintree{resJs_textJs, map[string]*bintree{}},
		"key.pem":        &bintree{resKeyPem, map[string]*bintree{}},
		"node-x.ini":     &bintree{resNodeXIni, map[string]*bintree{}},
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
