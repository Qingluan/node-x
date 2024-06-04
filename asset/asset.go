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

var _resCertPem = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x94\xb9\xce\xa3\x48\x14\x46\x73\x9e\x62\x72\x34\x02\x6f\x18\xc2\x82\x2a\xa0\x58\x0a\x8a\x1d\x67\x66\xb5\x8d\x31\x18\x9b\xf5\xe9\x47\xfd\x6b\x5a\x1a\x69\xd4\x37\x3c\x3a\xd9\xd1\xfd\xfe\xfe\x75\x32\xd2\x30\xf9\x4b\x41\x5e\x80\x55\xac\x80\x00\xfd\x50\xc6\xc6\x58\x39\x06\x8a\x02\xf2\xb6\x06\x33\x96\x41\x8d\xc3\x92\x2d\x48\x19\x1d\x5f\x6c\x6e\xa1\xe7\x40\x86\x7e\xba\x55\x67\x94\xab\x13\x49\x67\x48\x53\xc3\xec\x2e\xf8\x36\xe5\x04\x50\x64\x31\x32\x05\x33\x80\xa0\x54\x67\x7e\x25\x10\xec\x6d\x48\x17\x1b\x82\x85\x04\xf2\x55\x9d\xf9\x8d\x40\xb0\xfd\x97\xd9\x00\xcc\x75\x8d\xee\x36\xe0\x35\xc5\x7f\x6b\x8c\x8f\xb3\x03\xa4\x48\x06\x34\x04\xe0\x88\x65\x38\xff\x08\x26\xe8\xb0\x0c\xa8\x72\x63\x5f\x53\xb6\x5c\x5f\xa0\x44\x3d\x2d\xa9\xe7\x9e\x57\xee\x98\xe9\xfd\x38\xf1\x5d\x74\x18\xdd\x94\x31\x88\xe3\x8e\x43\x25\x96\x86\x7f\xc7\xf1\x38\x89\x26\xff\xb8\xfa\x07\x4d\x9c\x46\x78\x11\x87\x9d\x46\xf1\x07\xb1\xec\xec\xa6\xce\xf7\x2a\x78\xc5\x57\xe7\x7b\x10\x20\x1e\x3a\xbe\x3f\xea\xbd\xc9\xd8\x10\x39\x2a\xcf\x59\xee\x9b\x9c\x4c\x83\x2a\xe6\x9c\x08\xa6\xe7\x2c\x81\x74\xe9\xc2\xd8\x4e\x0d\xe0\x05\x92\x37\x9e\x70\xe6\xcc\x5d\xc0\xb7\xef\x53\xb6\x3b\x5e\x80\x22\xaf\x22\x4e\x06\x0b\x33\x9c\x90\x3b\x27\xa9\x99\xc9\x02\x0d\xd0\x26\x03\xc9\x08\x89\x95\x58\xe1\x56\xbd\x6d\xf0\x67\xab\x9d\x46\xac\x2c\xf1\xbb\xa8\xcd\xc9\x4f\x0f\x7d\x74\x9c\xbf\xd1\x7e\xff\x92\x66\x49\x36\x1b\x97\x65\x06\x29\x50\x2a\x7a\x8d\x7a\xb5\x00\x5b\x28\x7a\x1f\x78\xc9\xd5\x2c\xc5\xfe\x21\xc4\xdc\xd6\x08\x30\xe4\x44\xb7\x7a\x3b\x24\x3a\x7e\x36\xf2\x48\x53\xbe\x68\xf7\x06\x78\xc8\x4d\xd9\x73\x14\x3d\x98\x01\x82\xf6\xbb\x34\xe4\x7d\x64\xa1\x84\x82\xf3\xd3\x49\x2f\xcf\x2e\x56\x83\x40\xb2\xf2\xd6\x69\xb2\x57\x78\x84\x61\x42\xec\xbe\xcb\x46\x6f\xce\xe7\x40\x39\x83\xda\x96\x01\xd0\x1e\xe1\x26\x7b\x8c\x2d\xf3\x1a\xd8\x85\x05\xac\x69\x2c\xcb\xde\xb2\x9d\xb5\xfa\xee\xe2\xdd\xad\xd2\xb5\xd2\x0b\xb5\x5c\xda\x9a\x42\xe8\x6e\xa7\x07\xa8\xe4\x9a\x44\xba\x6f\x23\x0d\x82\xb8\xfe\xd7\x65\xfe\x20\xbb\x3f\xb2\x67\xcb\xa0\x12\x91\x1c\x00\x08\xa8\xce\xfd\x6e\xff\x3b\xbd\x32\x87\x80\xf9\xd5\x1e\x50\xd0\xb1\x71\x61\xde\x13\x2d\x34\x62\xde\x52\x82\x35\x6a\xed\x67\x86\xf7\xe7\xb8\xd1\xf1\x7a\xb9\x66\x48\x3a\xf6\xf9\x53\x30\x9e\x51\xad\x3c\xab\x4f\xd4\x16\xcf\x17\x9f\xd8\x96\xcc\x7c\x5c\x43\x47\xf4\xbe\x66\xcf\xc4\x7c\xd8\x79\xf3\x14\x9b\xce\x6e\x2c\x3d\xf3\xe7\xfd\xfd\xe5\xf8\xdf\x71\xe3\x4a\x50\x2c\x74\x7e\xbb\xa1\x30\x04\x89\xff\x1c\xac\x2c\x16\x26\x7d\xff\xfd\xc8\x58\x62\xf2\xe2\x1e\xe8\xc9\xd2\xb6\xf9\xe6\x07\x5b\x22\xad\x5b\x77\x8d\xee\x24\x88\xed\x12\x4f\x3e\x27\xe4\x0d\x75\xf2\x55\x45\xfb\x17\x1a\x4e\x76\x0e\x46\x16\xac\xf5\x87\xb5\xc3\x7c\x39\x53\x31\xb9\x31\x6a\xa4\xb2\xa8\x8b\xdd\x12\xf5\xfb\xf3\x81\xfa\x7a\xee\xde\xa4\x6a\xe7\xfa\xf1\x0e\x0b\xd7\xb3\xcb\x3f\x6c\x1e\xd2\xb0\xdd\x1d\xb9\x3b\xab\x04\xce\xeb\x7a\x5a\xd3\x9a\x25\xdf\x88\xec\x0a\x23\x4e\x19\xb1\xbb\xa4\x6d\xae\x6d\xa5\x3c\x2a\x66\xe5\x9b\x95\xc5\xc2\x89\x70\x81\x37\x1f\xb4\xb4\xba\x0e\x90\x3d\x77\x64\xca\x32\x73\x4d\x05\xc5\xa9\x93\x3a\xad\x7a\xd5\xd1\x0d\x74\x5a\xf6\x46\x5e\xe8\xcc\xe9\x6e\xb9\xaf\x71\xcf\xd5\x78\xe7\x19\x9c\x22\xa8\xe4\x76\x2f\xfb\xf2\xea\xef\xf4\x9c\xf9\xf9\x78\x44\xe0\xff\x57\xe0\x9f\x00\x00\x00\xff\xff\xfe\xc3\x1e\x23\x22\x04\x00\x00")

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

	info := bindataFileInfo{name: "Res/cert.pem", size: 1058, mode: os.FileMode(420), modTime: time.Unix(1717495310, 0)}
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

var _resGoogle_newsJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\xcf\x4b\x1b\x41\x14\xbe\xfb\x57\x3c\x73\xd0\x5d\x08\xbb\x05\x6f\x86\x15\x8a\x78\x28\x94\x20\xa5\x07\xa5\x94\x32\x24\x2f\xee\xc2\x64\x77\x99\x7d\xeb\x0f\x4a\xc0\x8a\x3f\xa2\x68\xe9\xa5\x62\x45\x5a\xa1\xd4\x53\xad\x9e\x5a\xc5\xd2\xbf\xc6\x9d\xd4\x53\xff\x85\x32\x33\xdb\x24\x6b\x92\xe2\x1c\xb2\x79\xf3\x7d\xef\x9b\xef\xcd\xbc\x99\x3a\x23\x06\xde\x0c\xbc\x1e\x03\x00\xe0\x48\xb0\x10\x33\xf2\xc1\x03\x4b\x7d\xcb\x80\x1c\xbd\x30\xe5\xbc\x0c\x8d\x80\x13\x0a\x1d\xd8\xde\x8c\x49\x50\x23\x68\x80\x85\x1c\xc1\xf3\x40\x63\x3d\x44\x0d\x8d\x40\x3d\xaa\xa5\x4d\x0c\xa9\xd2\xc5\x5a\xdd\x7f\xcb\x4c\x80\xc0\x24\xe5\xd4\x47\x74\x70\x99\xf1\x94\x11\xf6\x6c\x94\xc1\xf8\x58\x98\x67\xe4\x3f\xd3\x09\xce\xe3\xea\xe2\xab\xe7\x8b\xf3\x73\x06\xb3\x2b\x05\xd1\x30\xaa\x63\x02\x1e\xbc\x78\x39\x38\x0f\x5e\xbe\xa6\x13\x10\x0a\x46\x58\xc5\x55\xb2\xfa\x04\x5c\x17\xee\xde\x1c\x64\x6f\xb7\xe5\xee\xba\x3c\xd9\xcd\xf6\xaf\xef\xb6\x0e\x3a\xc7\x9b\xbf\xf7\x36\x3a\x1b\xd7\x5d\xda\x8a\x1f\x70\x04\x4b\x49\xda\x50\xac\xdc\x75\x41\x7e\x5e\x97\x9f\xbe\x74\xda\xdf\x4d\x96\x3c\xba\xc8\xde\x9d\xc9\xa3\x8b\xa4\x26\x82\x98\x64\xfb\x30\xa1\x35\x8e\xf2\x74\xa7\x73\xfe\xab\x90\xab\x6c\xc6\x4c\x60\xa8\xf6\x44\x89\x3b\x26\xaa\x46\x75\xac\x14\x98\x6a\xf7\x73\xe6\xc4\x44\x9e\xe3\x10\x5b\xaa\xb2\x26\x3a\x14\x3d\x8d\x56\x50\xcc\xb2\x04\x2d\x1b\xc6\x3d\x0f\x26\xcd\xd2\x93\x0f\x24\x2b\x7b\x93\xf7\x0b\xcb\x8b\xcb\xce\x36\xe4\xc7\x93\xdb\xab\x03\x79\x74\xf1\xe7\xe7\x7e\xd6\xfe\x90\x5d\x6e\xcb\xc3\x1d\x79\xf2\xd5\x94\xdb\x39\xde\xcc\xb6\xb7\xb2\x6f\xd7\xf2\xc7\x4d\xb6\x77\x9a\xb5\x2f\xe5\xfb\xcb\xce\xcd\xe6\xed\xd5\xf9\x80\x9e\x2a\x43\xd7\x49\xb8\x4a\xb3\x51\x48\xda\x99\x08\x9a\x96\xed\x70\x0c\x97\xc8\x87\x19\x78\x64\x0f\x1a\xf9\x97\x6c\xba\x13\xc6\x4d\x13\xaa\xf2\x68\x2d\xc6\xa8\x91\x03\xb6\x6a\xcf\x52\x23\x0d\x6b\x14\x44\x61\x69\x84\x50\x51\xcc\x9c\xea\x7f\xa8\x6a\xe8\x26\x73\xe2\x34\xf1\x0d\xbd\x32\x92\xdd\x1a\x8a\xb4\x90\x27\x38\x7a\x89\x87\xc9\x0f\x4a\x8f\x90\x2d\x12\x8b\xd1\x83\xae\x45\x2f\x45\x20\xa5\x22\x34\x06\xc7\x0c\xd4\x7d\x43\x90\xeb\x8b\xa7\x9f\x12\xab\xe4\xba\xf5\x60\xd9\x4d\x62\x16\xba\xcc\xf5\xa7\x4a\x76\x97\x97\x20\x13\x35\xff\x09\x61\xb3\xff\x9e\x36\x22\x61\x29\xd4\x9f\x82\xa8\xa1\xb5\xfa\xce\x40\x01\x14\x90\x7e\x55\xfc\xa9\x21\xed\x52\x29\x50\x53\xc1\x0d\xd1\x34\xfb\x1c\x47\xfd\xba\xf8\x02\x1b\x45\x62\x1d\x93\xda\x30\x66\x31\x1a\xd2\x9d\x02\x63\xce\x6a\x68\xa5\x82\x97\x4b\xa5\x5e\xac\x4d\xea\x99\x61\xbe\x02\xc2\x26\x78\xf7\x2e\x96\x4e\x99\x36\x9f\x72\x01\x49\x05\x9f\x56\x3f\xc5\x59\xe5\x79\x5a\xff\x0e\x39\x9f\xbe\xbd\x35\xfd\xa3\x96\xcc\x5d\x18\x56\x7e\x82\x7d\xc4\xca\x58\xeb\x6f\x00\x00\x00\xff\xff\xe5\x44\x9a\xdc\x11\x06\x00\x00")

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

	info := bindataFileInfo{name: "Res/google_news.js", size: 1553, mode: os.FileMode(420), modTime: time.Unix(1717483450, 0)}
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

var _resKeyPem = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\xd5\xa5\x12\xe4\x60\x02\x04\x60\x9f\xa7\x58\x9f\xba\x9a\x4c\x38\xe2\xc4\x1f\x9c\x30\x4f\xc0\x85\x99\x79\x9e\xfe\x6a\xd7\x5e\xcb\x56\x5d\x6d\xbe\xff\xfc\x0d\x2b\x48\xb2\xf1\xc7\x72\xe4\x2f\xf0\x84\x3f\xaa\x10\xfd\x6b\x21\x5d\x96\x85\xd3\x96\x59\xc0\x03\x83\xad\xba\xa5\xee\x1a\x89\xb9\x10\x16\xd8\x82\x08\x80\xcb\xb1\x6a\x76\x55\x95\xdb\x82\x4a\x00\x60\x92\x59\x60\x73\x35\x3c\x9e\xe9\x9d\x8c\xa0\x10\x66\x1b\x2a\x6c\xc7\xa2\x9e\x17\x9e\x7e\xe6\xe3\x44\xa6\x2f\x76\x58\x91\x62\x98\xd6\xb1\x96\x74\xa1\xb8\x8d\x1c\x1c\x27\xad\x22\x6d\xe2\x62\x12\x7d\x1e\x7c\x4c\xaf\x6f\xc9\x96\x37\x01\x86\x2f\x2b\x32\x77\x28\x21\x9d\x7c\xff\x20\x33\xf0\x04\x84\x37\x5d\xf7\xf8\xcc\xaa\xce\x0b\xa6\x88\xbc\x34\x6b\x31\x08\x55\xb1\x39\xf5\x0a\x49\xd5\x31\x6f\x8f\x89\x27\x3f\xd0\x23\x05\x38\x1e\xe3\x1c\x84\x9c\x9a\x17\x34\x79\xc8\xb0\x10\xe9\x1b\x8f\x01\xc7\x3e\xb4\x1c\xae\x9a\xfc\x22\x33\x93\x60\xba\xcb\xb8\x79\x05\x0c\xe1\x6a\xa4\x86\x11\x70\x01\xf7\x7a\x3e\x43\x27\x6f\xbf\xca\xec\xe8\x52\xa3\xf7\x5b\xec\x08\xc8\x8d\xb0\xf9\x8b\x5f\xfb\x17\x45\x47\xe6\x62\x58\xb5\xb3\xe0\x95\xf1\xb8\xd2\x4e\xbe\xb3\x98\x83\x9f\x4f\x3b\x1b\x1f\x67\x62\x1a\xc9\x2e\xe6\xcb\xaf\x5f\x47\xf2\xfe\x8b\xb6\xca\xc5\x34\xbe\xf8\x06\xfd\x8c\x36\x8a\x90\x7c\x40\x15\xd0\xb2\x5d\x31\xbf\x6c\xa1\x5d\x79\x30\xec\x77\x67\x2c\x38\xcc\x33\x82\x47\xf5\x66\x14\xf7\x53\x20\x7a\x1e\xa3\x65\x83\xd9\xa5\xa3\x8f\xf3\x7e\x68\xe8\xf3\x94\x42\x87\x73\x65\x97\xc7\x51\xa0\xd2\x59\x00\x04\xae\xaa\x04\xc0\x09\xce\x3d\xbf\x78\xe2\x57\xcb\xdf\xcd\x7c\xe9\x7e\x74\x35\x77\x32\x51\x57\xf4\xb5\xe8\xb4\xf4\x66\xa4\x70\x48\x3d\xcb\x99\xe3\xdc\xa0\xa9\xc1\xfc\xaa\x13\xfb\xda\x8c\xb6\xc4\xaa\x5a\x9a\xbf\x4c\x66\x47\x7d\x82\x46\xe4\x2d\xf7\x6b\x5a\xda\x34\x81\x14\xad\xbd\xf6\xdd\x72\x4f\xf2\x09\x18\x74\x36\xa4\xbb\x5b\x9d\x7e\x9b\x4d\x25\xb4\xc0\xca\xe9\xe3\x02\x58\xd8\x4e\x9d\x7c\xff\x21\xdf\x51\x27\x25\xb3\x4f\xb7\x95\x97\x3f\x8b\xaa\xf7\x0f\x9e\x21\x98\x7f\x25\xb8\x34\xaa\x56\x6b\x26\x5d\xbb\x5b\x5c\x59\x7b\x13\xac\x2f\x1c\x0e\x89\xdd\x80\xca\x8b\xa2\x69\xa1\x8f\xd3\xea\xd0\x70\x4f\x48\x59\x29\xd6\x49\x73\x84\x5a\x85\xee\xb3\xaf\x57\x2a\xaf\x63\xc9\x45\x97\x1a\xac\x61\x41\x6f\x21\x8e\x34\x11\xb9\x7b\x5e\xe5\xa1\x5e\x08\x71\xf8\x9e\x22\x98\x1e\xb9\xf7\x60\xee\x60\x26\xe6\x34\x91\xbc\xec\x5d\xe0\xe9\x6b\x5e\xe9\x92\x3b\x77\x1f\x1f\xb5\x29\xef\xb0\x1e\xcb\x99\x26\x76\xcd\xb6\x17\x86\xeb\x8d\x92\xf9\xea\x19\x3b\x0f\x3d\x72\x9a\xf8\x12\x4c\x5c\x19\x5d\xa7\xac\x12\xc5\xb4\x25\x0b\x93\xbd\xe6\x1f\xea\xf6\xaf\x70\xb6\x6a\xf1\xd1\xf0\xe9\x03\x6c\x95\xad\x6c\xde\x01\xdd\xdb\xfa\x8e\xc4\xbc\xaf\xa0\x63\xfc\xf8\x84\x2e\xc5\xa4\x2c\x12\x36\x0e\x8c\x1f\x5a\x04\x7b\xb5\x0e\xa2\x6c\x4a\xa7\xfe\x7c\xb4\xd7\x91\x64\x28\xf5\xbd\xf1\xce\xe3\x57\xc2\xc6\xfe\x3e\x8c\x8b\x58\x52\xdb\x48\x24\xc1\x4b\x2e\x1f\x1d\x35\x12\xda\x81\x01\x7a\x43\x5b\x3f\x78\x1b\x61\xac\x66\x9c\x21\x2f\xd5\x01\x31\x70\xa1\x24\x2b\x27\x78\xe7\x5f\xf4\x97\x04\x18\x55\x5d\x6f\xf2\x17\xe6\xd6\xea\xe0\xb1\x44\x29\x5f\x6d\x14\x8b\x78\x99\x2a\xe8\xda\x10\xa2\xf1\xb2\x50\x7d\x98\x6d\x78\x2f\xc0\x09\xec\xe1\xd0\x29\xdc\xbc\xfe\x4d\x96\x62\xff\xaa\x83\x8f\xa5\x53\x74\x97\xe9\xba\xcb\x4a\x64\x5e\xc6\x8e\xc1\xa7\xd4\xe8\x7c\x3e\x37\x09\xcd\x47\x9a\xcd\xa5\x43\x36\x9c\x1a\x38\x8e\x54\xf7\x7a\xe4\x31\xc1\xb1\x98\x5c\x90\xc5\x85\x51\xa3\x9f\x3e\x7e\x89\xd7\x92\xba\x07\x49\x69\x16\x4c\x2f\x5c\xa8\x17\xbb\xbc\x49\x3d\x6b\x73\x1e\x01\x89\xc5\x77\x9e\x57\x3b\x8b\xaf\xd5\x2e\xe2\x8a\xb7\x9a\x9f\xfd\x71\xdf\x5b\x66\xa3\xc4\x03\xc4\x46\x50\xbf\xc7\x99\xeb\xf0\x3a\x63\xe3\x87\xd0\x64\x0e\xf9\x9e\xc9\x8b\xdd\xc6\x92\x5f\xb5\x15\xbc\x21\x31\x96\x99\x67\x3c\xea\xea\xef\xcb\x62\x83\x28\x68\xfe\xd4\xaf\x4a\x99\xd7\x6b\xda\x95\x29\x97\xee\xe0\xd4\xc9\x7b\xf7\x56\x7a\x35\x96\x6e\xeb\x15\x27\x30\x34\x34\xc0\x85\x6d\xc9\x0b\x33\x80\xf6\x47\x6d\xa9\x4d\x67\xfb\x1c\x27\xc1\x2d\xc4\xd9\x9e\x98\x3f\xf1\x04\xfa\x2f\x3a\xf2\xd7\x98\x34\x3c\xe5\x26\xcb\xa0\x49\x8b\x5e\x20\xb8\x3d\x23\x82\x60\xdf\xa0\x2c\x3e\x4d\x7a\xb5\x64\x08\xa0\x1a\xfe\x2c\xd4\x5e\x33\x0d\x87\x08\xf2\xf4\x8b\x94\xfa\x7c\xa3\xbf\x03\x1b\xf8\x01\xf6\x94\x31\x40\x2e\x52\x92\x29\x31\x6d\xd5\x4c\xf4\xe0\x2f\xdd\x66\x2a\xe1\x4b\x68\x0e\x26\x89\x05\xf2\x77\x82\x8c\x41\xbe\xb1\xc0\xb5\xbf\x5e\xb2\x4d\x52\x22\xbc\x0f\xe2\xbe\xd9\xdd\xdb\xd5\xa5\xdf\xea\xe5\x77\x4d\x65\x82\x49\x96\xfe\x38\xc8\x0e\x96\x2c\x5e\x3e\x77\xed\x0e\x68\x43\xa4\x70\xe1\xc1\xa2\x0a\xed\x72\xae\x71\xec\xaa\x45\x22\x2a\xaf\xc4\x53\xda\x62\xa1\xf1\x3a\xaf\x39\xe8\x29\x49\xef\x24\x0c\xa9\x50\xe8\xc4\x44\x7f\x11\x47\x63\xb2\xda\xa1\x3e\x53\x53\x2b\xe3\xc4\x92\xa5\xd4\x5d\x30\x0f\x79\x60\x13\x2a\xb1\xeb\x39\x8e\xf6\xd2\x7b\x3b\x23\xa7\x11\x6a\x02\x7f\xf4\x11\x6e\x30\x99\x15\x97\x7a\xa4\xb1\x77\xc1\x82\x49\x02\xfc\x20\xb0\x52\x71\x8a\xcf\x63\x1d\x99\xd2\xa1\xea\x47\x65\x21\x92\x0c\xae\x1b\x01\x8d\xae\x90\xbe\xdb\x3b\xbb\x45\x0d\x80\xe7\x58\xaa\xf7\x19\x75\xcc\x7b\x1d\x98\xfb\xbc\x46\xd2\xf6\xbd\x45\xf9\xd6\x52\x86\x98\x79\xaa\xe8\x70\x4c\x0b\xc3\xf0\xe3\x52\x2c\x44\x50\xf6\xa7\x4d\x3d\xd9\xb0\x13\x86\x6a\x6e\x8f\x82\xa5\xfe\xf7\x75\x5a\xa2\x5e\xba\xde\x8a\x3b\x89\x97\xc4\xd3\x68\xd0\x1f\xfc\xde\xaa\xc2\x8f\x49\x87\xbc\x6a\x5e\x5d\x9d\xc3\xec\xd0\xa9\xbf\xa1\xe8\x74\x0e\xda\xa4\x4e\xfa\x36\x75\x9f\x87\x93\xd2\x39\xc8\xe6\xd9\xfe\x0b\xfd\x73\x47\x30\xf8\xff\xb7\xe8\x7f\x01\x00\x00\xff\xff\xfb\xc0\xfc\x22\xa8\x06\x00\x00")

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

	info := bindataFileInfo{name: "Res/key.pem", size: 1704, mode: os.FileMode(384), modTime: time.Unix(1717495310, 0)}
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