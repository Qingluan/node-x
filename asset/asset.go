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

var _resCertPem = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x94\xc9\xae\xab\x3c\x16\x46\xe7\x3c\x45\xcd\x51\x89\x26\x0d\x64\x68\x6c\x70\x0c\x31\xc1\x34\xc9\x21\x33\x42\x80\xd0\x85\x1e\x02\x4f\x5f\xba\x47\x75\xa5\x92\x4a\xff\x1e\x2e\xad\xd9\xd2\xfe\xfe\xfd\xe7\x34\x1d\x13\xfb\x5f\x50\x77\x7d\x62\x10\x08\x7c\xfd\x97\x72\x94\x10\xb8\xf7\x21\x04\x71\x9d\x81\x85\x68\x20\x23\x81\x69\xca\x1d\xba\x0f\x89\xee\x3d\xe1\xfb\x73\x9a\xf8\x76\x7f\x97\xd7\x58\xf9\x19\xc5\x05\xb1\xd0\xb4\x9a\x07\x79\xcf\xb1\x0d\x98\x7e\xe1\x34\x06\x16\x80\x40\x62\x2c\xe2\x6a\x23\x7d\xa1\x9b\xfe\xa5\x3e\x15\xe9\x56\x45\xc6\x22\x6e\x36\x02\xbb\xff\x65\x14\x80\x25\xcb\xf4\x9c\x02\x11\x43\xaf\xc3\x9c\x47\x9e\x3b\xc4\x74\x0d\xb0\x00\x80\x3d\xd1\xd0\xf2\x2b\x58\xa0\x21\x1a\x60\x28\x2e\x3e\x13\x14\xd5\xd0\xc3\xc7\x18\x8e\x2f\xe8\x47\x3a\x98\x97\xb8\xbb\x4d\xe2\x43\x15\x0e\xa1\xcf\x45\x78\x39\xc7\x1f\x90\x12\x5f\x6d\x17\x93\x37\x9d\xad\xe8\xc2\x82\x54\xf8\x46\x9a\xfc\xa1\xef\x15\x41\xd6\x9e\x87\x46\x38\x23\x9c\x37\x37\xc7\x43\x12\x31\x8e\xc2\x7c\x08\x61\x3f\x20\x67\x54\x38\xa7\xf3\x5c\x53\xb4\x55\x2b\x28\x11\x6b\x49\x33\x7f\xfb\x4b\xf9\x58\x58\xfb\x65\x47\x45\xbb\x38\xc8\xfc\x42\xd6\x25\x17\xe0\x88\xb2\x68\xd5\xf0\x76\xbc\xbc\xb3\xe1\x8c\x19\x7b\x81\xca\x7c\x63\xee\x91\xfb\x72\xa0\xde\x2e\x9e\x73\xdb\x12\x55\x1b\xe8\x9a\xcd\x87\x8f\xd5\xca\xe4\x43\xf8\xc9\x8d\xa7\x1e\x9c\x0b\xa6\x0c\x78\x8d\x14\x29\x4c\xa2\x8f\x26\x84\xc2\xde\xa2\xa9\x73\x4e\xaf\x92\x7e\xe6\xa6\xc6\xed\xfa\xe3\x52\x19\xd8\xa8\x88\xeb\x94\xee\x59\x5c\x27\xbb\x82\xdf\x61\x48\xf9\xf7\xbe\xd6\xf7\xda\xd5\x05\xd7\x7a\x9a\xe1\x24\x0f\x01\xbd\x9a\xf5\xb7\x95\x8a\x71\x7e\xbe\x06\x44\x2b\xc2\xb1\x00\x0a\x9d\x9c\xbd\xc3\x30\xc8\xc4\xa7\x96\x96\x99\x49\x56\x7e\x94\xbf\xb5\xca\x1f\xa5\x26\xdb\x62\x46\x1f\xb5\xf0\xf1\xbd\xce\x4e\x2a\xe6\x08\xde\x33\x05\x19\xd5\x00\xc0\x45\xb0\x69\x2e\x47\x35\x11\x03\x29\x78\xa1\x8c\xdd\x35\xcd\xfd\x18\x4f\x35\xf7\x76\xc4\x98\x9a\x1c\x1a\xcb\x38\x99\x2e\x2c\x3e\xdb\x86\x27\x04\x52\x2d\xb3\x6f\x67\x8f\xea\x18\x81\x7b\xf6\x5f\x97\xfb\x07\xd9\xf9\x95\x5d\xaa\x81\x54\xd5\x35\x1f\x20\xc0\xce\xc2\xdf\xf6\x7f\xd3\xc3\x25\x00\xdc\x9f\xf6\x80\xa1\xe8\xbb\x31\xe3\xed\x3f\x57\x45\x55\x75\xe5\x80\xe4\xd5\x2c\x27\x60\x8a\x62\x77\xfd\xa6\x8f\xa4\x1c\x24\xe9\x94\x3a\xd1\x98\x4b\xd4\x16\xc6\x9e\x5a\x73\xe6\x9b\x59\x1d\x9b\x5c\x79\x71\x69\x5e\xc6\x5b\x7c\xf1\x9c\x47\x8b\x83\x79\xff\xc1\x50\x59\xbb\x98\xd1\x4b\xe4\x5f\xf7\xe5\x35\x54\xe4\x1d\x8a\x6e\xea\xeb\x25\x59\xba\xd9\xa5\x90\xc0\x11\x46\xf9\xcf\x82\x75\x72\xdb\x73\x92\xd0\x83\xbe\x7e\x24\xf7\x86\x9f\x73\xe3\x7e\x5c\x93\x49\xc8\x91\x5a\x7d\x03\x9c\x9c\x4e\x3f\x6c\x2a\x06\xd1\xe3\x1d\x2b\xca\x0c\x24\x3d\x7a\xea\xd7\xcf\xb7\xaa\x7b\x4d\xf7\xbc\x91\x66\x0f\x8f\x9c\x3b\x3a\xe3\x15\x2f\x91\x94\xf8\x1b\xbf\x8c\xfe\xc9\xf2\xeb\xc3\xae\x80\xe1\x9b\x9c\x8e\x69\xdd\x38\x45\x75\x3c\x61\xbc\x51\x8b\xc7\x96\x9d\xb4\x31\x0c\x9b\x64\xed\x0f\x08\xd7\x64\x34\xde\x06\x57\xf2\x7d\xbf\x96\x88\x54\x87\x8d\xe0\x30\x9b\x95\xf6\x3d\x87\x8e\x77\x52\xfd\xfc\xd2\xdf\xbb\xdd\xfd\xc8\x4f\xfc\x1d\xb6\x45\x0e\xd5\xe5\xf1\x99\x9a\xe2\xda\x14\xf5\x46\xda\xdd\x94\x56\xbb\xa6\xe7\xd4\x5e\x9e\xe7\xce\xae\x9b\x9f\xa2\x91\xab\xf2\xce\x23\x68\x25\x16\xb5\xcf\x2f\xa8\x70\xbf\x1f\xaf\xdb\xe8\xff\x57\xe0\x3f\x01\x00\x00\xff\xff\x9d\xd8\xc3\x72\x22\x04\x00\x00")

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

	info := bindataFileInfo{name: "Res/cert.pem", size: 1058, mode: os.FileMode(420), modTime: time.Unix(1730374479, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resGoogleJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\xdf\x4b\x1b\x4b\x14\x7e\xf7\xaf\x38\xe6\x41\x77\x21\xec\x5e\xf0\xcd\xb0\xc2\x45\x7c\xb8\x97\x4b\xae\xb4\x7d\x50\x4a\x29\x43\x72\xd6\x5d\x98\xec\x86\xd9\xb3\xfe\xa0\x04\xac\xf8\x23\x8a\x96\xbe\x54\xac\x48\x2b\x94\xfa\x54\xab\x4f\xad\x62\xe9\x5f\xe3\x4e\xea\x53\xff\x85\x32\x33\xdb\x24\x6b\x92\xe2\x3c\x64\x73\xe6\xfb\xce\x37\xdf\x99\x39\x33\x75\x46\x0c\xbc\x19\x78\x31\x06\x00\xc0\x91\x60\xa1\xc9\x28\x00\x0f\x2c\xf5\x2d\x03\x72\xf4\xa2\x94\xf3\x32\xf8\x21\x27\x14\x3a\xb0\xbd\x19\x93\xa0\x46\xe8\x83\x85\x1c\xc1\xf3\x40\x63\x3d\x44\x0d\x8d\x40\x3d\xae\xa5\x0d\x8c\xa8\xd2\xc5\x5a\xdd\x7f\xcb\x4c\x80\xc0\x24\xe5\xd4\x47\x74\x70\x99\xf1\x94\x11\xf6\x6c\x94\xc1\xf8\x58\x98\x67\x14\x3c\xd2\x09\xce\xdf\xd5\xc5\xe7\x4f\x16\xe7\xe7\x0c\x66\x57\x0a\xa2\x51\x5c\xc7\x04\x3c\x78\xfa\x6c\x70\x1e\xbc\x7c\x4d\x27\x24\x14\x8c\xb0\x8a\xab\x64\xf5\x09\xb8\x2e\xdc\xbd\x3c\xc8\x5e\x6d\xcb\xdd\x75\x79\xb2\x9b\xed\x5f\xdf\x6d\x1d\x74\x8e\x37\x7f\xec\x6d\x74\x36\xae\xbb\xb4\x95\x20\xe4\x08\x96\x92\xb4\xa1\x58\xb9\xeb\x82\xfc\xb0\x2e\xdf\x7f\xec\xb4\xbf\x98\x2c\x79\x74\x91\xbd\x3e\x93\x47\x17\x49\x4d\x84\x4d\x92\xed\xc3\x84\xd6\x38\xca\xd3\x9d\xce\xf9\xf7\x42\xae\xb2\xd9\x64\x02\x23\xb5\x27\x4a\xdc\x31\x51\x35\xae\x63\xa5\xc0\x54\xbb\x9f\x33\x27\x26\xf2\x1c\x87\xd8\x52\x95\x35\xd0\xa1\xf8\xbf\x78\x05\xc5\x2c\x4b\xd0\xb2\x61\xdc\xf3\x60\xd2\x2c\x3d\xf9\x40\xb2\xb2\x37\x79\xbf\xb0\xbc\xb8\xec\x6c\x43\xbe\x3b\xb9\xbd\x3a\x90\x47\x17\x3f\xbf\xed\x67\xed\xb7\xd9\xe5\xb6\x3c\xdc\x91\x27\x9f\x4c\xb9\x9d\xe3\xcd\x6c\x7b\x2b\xfb\x7c\x2d\xbf\xde\x64\x7b\xa7\x59\xfb\x52\xbe\xb9\xec\xdc\x6c\xde\x5e\x9d\x0f\xe8\xa9\x32\x74\x9d\x84\xab\x34\x1b\x47\xa4\x9d\x89\xb0\x61\xd9\x0e\xc7\x68\x89\x02\x98\x81\xbf\xec\x41\x23\xbf\x93\x4d\x77\xc2\xb8\x69\x42\x55\x1e\xad\x35\x31\xf6\x73\xc0\x56\xed\x59\xf2\xd3\xa8\x46\x61\x1c\x95\x46\x08\x15\xc5\xcc\xa9\xfe\x81\xaa\x86\x6e\x32\xa7\x99\x26\x81\xa1\x57\x46\xb2\x5b\x43\x91\x16\xf2\x04\x47\x2f\xf1\x30\xf9\x41\xe9\x11\xb2\x45\x62\x31\x7a\xd0\xb5\xe8\xa5\x08\xa4\x54\x44\xc6\xe0\x98\x81\xba\x6f\x08\x72\x7d\xf1\xf4\x53\x62\x95\x5c\xb7\x1e\x2e\xbb\x49\x93\x45\x2e\x73\x83\xa9\x92\xdd\xe5\x25\xc8\x44\x2d\xf8\x87\xb0\xd1\x7f\x4f\xfd\x58\x58\x0a\x0d\xa6\x20\xf6\xb5\x56\xdf\x19\x28\x80\x42\xd2\xaf\x4a\x30\x35\xa4\x5d\x2a\x05\x6a\x2a\xb8\x21\x9a\x66\x9f\xe3\xa8\x5f\x97\x40\xa0\x5f\x24\xd6\x31\xa9\x0d\x63\x16\xa3\x21\xdd\x29\xb0\xc9\x59\x0d\xad\x54\xf0\x72\xa9\xd4\x8b\xb5\x49\x3d\x33\xcc\x57\x48\xd8\x00\xef\xde\xc5\xd2\x29\xd3\xe6\x53\x2e\x20\xa9\xe0\xd3\xea\xa7\x38\xab\x3c\x4f\xeb\xdf\x21\xe7\xd3\xb7\xb7\xa6\x7f\xd4\x92\xb9\x0b\xc3\xca\x4f\xf0\xdf\xc7\xff\x57\x9d\x84\x44\x18\x2d\x85\xfe\x9a\xd5\x97\x67\x57\xc6\x5a\xbf\x02\x00\x00\xff\xff\xb9\x16\xae\x86\x21\x06\x00\x00")

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

	info := bindataFileInfo{name: "Res/google.js", size: 1569, mode: os.FileMode(420), modTime: time.Unix(1717568138, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resGoogle_newsJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x54\xcf\x4f\xdb\x48\x14\xbe\xf3\x57\x3c\x72\x00\x5b\xca\x3a\x2c\x62\x2f\xc9\x1a\xed\x0a\x71\xd8\xd5\x2a\x8b\x76\x7b\x00\x21\x54\x59\xf1\x33\x1e\x69\x62\x47\xe3\x71\x00\x55\x91\x28\xe2\x47\x40\x50\xf5\x52\x44\x11\x6a\x91\xaa\x72\x2a\x85\x53\x0b\xa2\xea\x5f\x83\x9d\x72\xea\xbf\x50\xcd\x4c\x9a\xc4\x38\x89\x5c\xb5\x52\x2d\x65\xe2\xf1\xf7\xbe\x6f\xde\x9b\xf9\xe6\xd9\x16\xb7\xc0\x9c\x86\x47\x23\x00\x00\x14\x39\xcc\xd7\x2c\xee\x82\x09\x9a\xf8\xcf\x03\x52\x34\xbd\x90\xd2\x3c\x38\x84\x72\x64\x72\xa2\x9b\xd3\x8a\x20\x1e\xe2\x80\x86\x14\xc1\x34\x41\x62\x5d\x44\x3c\x12\x01\xdb\xaf\x84\x55\xf4\x78\xa9\x83\x35\x3a\x6f\x75\x8b\x01\xc3\x20\xa4\xbc\x27\xd0\xc0\xba\x45\x43\x8b\x63\x37\x8d\x3c\xa8\x3c\xe6\xe7\x2c\xee\xfe\x27\x09\xc6\x9f\xe5\x85\x87\x0f\x16\xe6\x66\x15\xa6\x97\x12\xa2\x9e\x6f\x63\x00\x26\x2c\x2e\xa5\xbf\x83\xd9\x5e\xd3\x20\x1c\x99\xc5\xb1\x8c\xab\x5c\xeb\x11\x28\x14\xe0\xee\xf1\x41\xf4\x64\x3b\xde\x5d\x8f\x4f\x76\xa3\xfd\xeb\xbb\xad\x83\xd6\xf1\xe6\xa7\xbd\x8d\xd6\xc6\x75\x27\x6c\xc5\x25\x14\x41\x13\x92\x3a\x24\x2b\x2f\x14\x20\x7e\xb5\x1e\xbf\x7c\xdd\x6a\xbe\x53\xac\xf8\xe8\x22\x7a\x7a\x16\x1f\x5d\x04\x15\x46\x6a\x3c\x6e\x1e\x06\x7c\x8d\x62\x7c\xba\xd3\x3a\xff\x98\xe0\x8a\x34\x6b\x16\x43\x4f\xec\x89\x10\x37\xd4\xac\xec\xdb\x58\x4a\x44\x8a\xdd\x6f\x47\x8e\x8d\xb5\x39\x06\xb7\x96\xcb\x56\x15\x0d\xee\xff\xe3\xaf\x20\x9b\xb1\x02\xd4\x74\x18\x35\x4d\x18\x57\x4b\x8f\x67\x0c\x16\xe9\x8d\xdf\x2f\xac\x5d\x5c\x74\xb6\x11\xbf\x38\xb9\xbd\x3a\x88\x8f\x2e\x3e\x7f\xd8\x8f\x9a\xcf\xa3\xcb\xed\xf8\x70\x27\x3e\x79\xa3\xca\x6d\x1d\x6f\x46\xdb\x5b\xd1\xdb\xeb\xf8\xfd\x4d\xb4\x77\x1a\x35\x2f\xe3\x67\x97\xad\x9b\xcd\xdb\xab\xf3\x94\x9e\x28\x43\xd6\xc9\x71\x95\xcf\xf8\x1e\x97\x99\x31\x52\xd5\x74\x83\xa2\xb7\xcc\x5d\x98\x86\x09\x3d\x9d\xc8\x57\xb2\x72\x27\x8c\x2a\x13\x8a\xf2\xf8\x5a\x0d\x7d\xa7\x0d\xe8\xc2\x9e\x39\x27\xf4\x2a\x9c\xf8\x5e\x6e\x80\x50\x52\x4c\x9d\xea\x90\x50\xf1\x48\x93\x19\xb5\x30\x70\x55\x78\x69\x60\x74\xa3\x2f\xd2\x40\x1a\xe0\xe0\x25\xb2\xc9\xa7\xa5\x07\xc8\x26\x03\x93\xb3\x4c\xd7\xa2\x4b\x61\xc8\x43\xe6\xa9\x04\x47\x14\xd4\xe9\x21\x48\xe5\xc5\x93\xad\x44\xcb\x15\x0a\xd6\xe2\x1f\xa2\xd1\xfc\x52\x47\x7b\xa9\x60\x93\xba\xf8\xe5\xf4\x4e\x78\x80\x16\xab\xb8\x7f\x71\xac\xf6\x5e\x57\xc7\x67\x9a\x40\x6d\x52\x07\xdf\x91\x9a\x7a\xb2\xeb\xd8\xa4\x6e\x54\x5c\x42\x6d\x86\x5e\xd7\x25\x93\xe2\xf0\xfb\x41\xbf\xc3\x6f\xf7\x0e\x53\xc8\xd7\xfc\x40\x36\x9e\x1e\xc2\xe2\xc4\x52\x1f\x1f\x96\x52\x54\x4e\xb8\xea\x6e\xbd\xdc\x5f\xb3\x71\x43\x46\xdb\x4c\x75\x0d\x67\x29\xca\xbe\x97\x9c\xb9\x0c\x9d\x34\xd7\xb6\x78\x6a\xd9\xc9\xbe\xcb\x0a\x46\x8a\x4e\x38\x56\xc1\xec\x73\xa9\x65\x3d\x45\xf5\x97\x4f\xa1\x21\xa3\x45\x31\xa4\x11\xb1\x85\x45\x39\xa6\x31\x91\x6b\x51\x8e\xf9\x21\xde\xeb\x71\x80\x32\xbb\xc8\xf1\xde\xae\x75\x3d\x28\xbc\x3d\xc4\x00\x53\x7d\x4e\xf9\x7b\xb6\xfb\x27\x39\xc4\xc6\xa0\x92\xed\x94\x33\x19\x64\x2a\x1b\xf5\x87\x9b\x43\x94\x51\x94\x63\x76\x73\xc0\x10\x53\x7d\x9b\x71\xba\xd1\xc9\x17\xf5\xbd\xdd\xc3\xfe\xfe\xff\xdf\xb2\x11\x70\x46\xbc\x65\xe2\xac\x69\x3d\x9a\x7a\x69\xa4\xf1\x25\x00\x00\xff\xff\x34\x0c\x26\x14\x23\x09\x00\x00")

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

	info := bindataFileInfo{name: "Res/google_news.js", size: 2339, mode: os.FileMode(420), modTime: time.Unix(1717552630, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resJs_textJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x56\x4f\x6f\x1b\x45\x14\xbf\xfb\x53\xbc\xec\x21\x59\xd7\x61\xed\x04\x02\x04\x7b\x2d\x45\xa6\x87\x4a\x10\x2a\x28\x12\x55\x53\xac\x91\xfd\xd6\x9e\x30\xde\x59\xcd\xce\x36\xae\xa8\x0f\x15\x12\x14\x01\x15\x17\x0a\x82\x03\x1c\xa2\x72\x21\xe2\x8f\x38\x44\x51\x3f\x4e\x9c\x96\x53\xbf\x02\x9a\xd9\xf5\xee\x78\x67\x9d\xf6\x82\x84\x4f\xd9\x7d\xbf\xf7\x7e\xbf\xf7\x77\x33\x24\x92\x80\xdf\x85\x4f\x6b\x00\x00\x77\x88\x80\x69\x44\xe4\x18\x7c\x70\x9a\x4d\x89\x53\xe9\xd6\x9d\x76\x6e\x13\x18\x27\x4c\x82\x0f\x43\x3e\x48\x26\x18\x4a\x0f\xef\x10\x96\x10\x89\xae\x76\xdb\xcc\x0d\x9b\x10\x26\x8c\x6d\xc2\x47\xd7\x89\x1c\xbf\xaf\xdd\xbc\xbd\xfd\x9b\xfd\x1b\x37\xaf\x5f\x4d\x6d\xf5\x22\xac\xe2\x89\xc1\x87\x5b\xb7\x8b\x77\x21\x1f\x22\xf8\x19\xa3\x47\x25\x0a\x22\x71\x5f\x0b\x4a\x41\x0c\x25\x48\x2a\x19\x9a\x72\xf4\x0b\x4f\x0a\x3a\x31\x61\x8c\x86\x9f\x98\x28\xc6\x07\x44\x52\x1e\x7a\x63\x81\x41\x01\xa3\xc1\x61\x32\x89\xc0\x87\x80\xb0\x18\xdb\xb5\xc2\x10\x5f\x0d\x47\x8c\xc6\xe3\x77\xc9\x21\x17\x54\xde\x05\x1f\x62\x29\x8a\xba\x2d\x80\x98\xc2\x7a\x3c\x09\x55\x95\x5a\xed\x25\xeb\x60\x4c\x43\x8c\xb1\x64\xcd\x21\xcd\x26\xfc\x73\xff\x9b\xf9\xc3\xcf\xe7\x27\xdf\x3f\xfd\xed\xf1\xf9\xe9\x5f\xb9\x29\xe0\x02\x5c\xad\x44\xfb\x01\x85\x8e\x12\xe0\x31\x0c\x47\x72\xdc\x06\xda\x68\xd4\x0d\x29\x03\x1e\xc6\x8a\x8e\x88\x5e\x5a\x45\x85\x5d\x3c\xee\x49\x97\xd6\x6d\xea\xf9\x83\xe3\x8b\x47\x27\x29\xf5\xc5\x0f\xbf\xcf\xbf\x7d\x7c\x7e\x7a\xf6\xec\xab\x3f\x2f\x1e\x7d\x91\xbe\x7c\xfe\xe4\xc1\xfc\xe7\xb3\xf3\xb3\x87\x7b\x1f\xf4\xae\x5d\x7b\xfa\xcb\xfd\xe7\x4f\xbe\xcc\x03\xd0\x00\xdc\x9c\xaf\xeb\xc3\xeb\x3b\xb0\xbe\x5e\x28\xe8\xf8\xb0\xdb\x82\x7b\xf7\xc0\xc4\xec\xbe\x51\xc6\x6c\x6d\x6f\x9b\x69\xa8\x9f\x59\xd1\x46\xa3\xa8\xe7\xec\x85\xda\xcf\x4f\x4f\x6c\xed\x1f\x86\x74\xc0\x87\xf8\xec\xeb\xcf\xe6\x3f\xfd\x6d\x26\x80\x2c\x46\x2b\x8b\xd6\xf4\x35\x6c\xb5\xca\x2a\x5b\xd3\xdd\x80\xec\x94\x85\x9a\xcd\xad\x16\x3a\x5b\x51\xf3\xb4\xc8\xb9\xee\xf9\xf1\xaf\xf3\xe3\x1f\x2f\xbe\xfb\x23\x47\x0a\x94\x89\x08\x97\x87\xab\xbb\x44\x98\xd2\xa5\x04\x47\x63\xca\x10\x5c\xb5\x3f\xa6\x46\xb5\x53\x11\x11\xa8\x47\x4f\x19\xbd\xf4\x69\x9f\x0f\xb1\xbd\xd4\xc7\x0c\xb5\xbe\x9e\xc2\x24\x19\xed\x93\x09\xc2\x9a\x0f\x4e\x2c\xef\x32\x74\xaa\x4d\x03\x41\x23\xa9\x6d\x69\x80\x85\xd5\x93\xfc\x1d\x7e\x84\xa2\x47\x62\x74\xeb\xb0\xe6\xfb\xb0\x91\x82\x37\x5e\x12\xac\x48\x5f\x06\x0b\x0e\x0d\x04\x99\xa0\x53\xee\x8d\xca\x2a\x15\x8c\x53\xd9\xe3\xa1\xd4\x51\xf4\x91\xc8\x56\x08\xba\xb0\x53\x5f\x76\xca\x6f\xcc\x34\xaf\x98\xed\xde\xb6\x5c\xd4\xee\x71\x86\x1e\xe3\x23\xd7\x79\x1b\x09\xa3\xe1\xe8\x2d\x67\x53\x4e\x65\x05\x58\x1f\x3e\x2f\x4a\xe2\xf1\x2a\x7d\x15\x4e\xe5\x33\x65\xda\x66\x6a\x8e\xed\x3c\x54\x01\xd6\x32\xc7\x8a\x34\x0d\x19\xce\x41\xe8\x5c\xc6\x29\x45\x82\xb6\x79\x56\xab\x7e\x2a\xfe\x7a\xe1\x39\x9f\xe5\x97\x9f\x30\x76\x03\x75\xd1\x53\x5d\x87\x9c\x86\xee\xc6\x41\xb8\x91\x21\x23\x22\xc8\x48\x90\x68\xac\x3e\x19\x19\xd8\x8b\x23\x46\xa5\x52\xaf\xf4\x6b\x58\x12\x31\x4e\x86\xfd\xb4\x7d\x8e\xa3\xdf\x05\x5c\xb8\xd9\x07\x01\x81\x07\x46\x28\xa3\x2a\xaa\x58\x85\xf3\x62\x40\xd6\x7c\x68\x95\x4a\x67\x30\x34\x7c\x50\x95\xab\x48\x5c\xd1\x29\x1a\xf0\x35\xab\x27\x30\x62\x64\x80\x6e\xb3\x73\xeb\xe3\xee\xed\x2b\x5d\xef\x4a\xe7\xa0\xe9\x5d\xe9\x36\x47\x9b\xe0\x64\xd2\xa1\xd8\x44\xe2\xd1\x70\x88\xd3\xf7\x02\xd7\xe9\xe8\x45\x00\xa7\x0e\x5d\x78\x65\x0b\x4a\x62\x06\x3c\x94\x34\x4c\xb0\x42\x82\x1d\x33\xcb\xa9\x03\xdb\xad\x6c\xaf\x0c\x1e\x70\xea\x1d\x2b\xd7\x4b\xc2\xe7\x51\xb3\x8d\xca\x03\xa9\x56\x68\xad\xa5\x58\x55\x5d\x28\x1a\xe8\xd4\x2b\x26\x54\x71\x58\x5f\x61\x57\x05\xa8\x42\x2f\xea\x7e\xc4\xc5\x30\x5e\x14\x3e\x53\x97\xf1\x80\x51\xea\x32\x91\x76\x2b\x6a\xf4\xea\x0a\x86\xca\xb2\x98\xbf\x15\xcb\xb8\xf8\xd9\xe3\xd3\x30\x84\x56\x47\xb4\x37\x6f\x35\x87\x4a\x45\x07\x34\xbb\xfd\xbf\xce\xc5\x7a\xb3\xea\xac\xd8\x62\xac\x7b\x6f\xce\x79\xf5\x71\xaf\xd0\xac\x9c\x96\x39\xad\xf8\x33\xe3\x5a\x09\x24\xac\x1f\x11\xa1\x66\xcc\xb8\x18\xc6\x2d\xca\xaf\x11\x4f\x64\x94\xc8\xbe\xba\x68\xfd\x88\x8c\xd0\xbe\x49\xfa\x48\xf0\xa0\x08\x5a\x3a\x49\xf6\xf8\x2b\x8f\xf2\xf8\xe7\xa9\x17\x73\x5e\xf4\x7f\xab\xbc\xd5\xb0\xaa\xed\x2b\xda\x6d\x65\x61\x16\xae\xaa\xd9\xb3\xda\x65\x4d\xb3\xef\xd1\xd6\x9b\x55\x13\xfa\x1f\x4b\xac\x6e\xad\xfe\x9f\xab\x1c\xad\x36\xab\xfd\x1b\x00\x00\xff\xff\x4f\x6e\x4d\x1c\x32\x0d\x00\x00")

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

	info := bindataFileInfo{name: "Res/js_text.js", size: 3378, mode: os.FileMode(420), modTime: time.Unix(1717569217, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resKeyPem = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\xd5\x37\x12\xab\xd8\x16\x85\xe1\x9c\x51\x74\x4e\xbd\x12\x46\x12\x10\xe2\xbd\x37\x12\x64\xd8\x03\x1c\xac\xf0\x8c\xfe\x55\xdf\x1b\xf6\x0e\x77\xb4\xea\x4b\xfe\xff\xfd\x7b\x9c\x28\xab\xd6\x3f\x8e\xa7\x46\x6c\x20\xfe\xa3\x8b\xf1\x9f\x2f\x62\xaa\xaa\xb8\x03\x95\x63\x05\xd6\xe2\x00\x9c\x6b\xd8\xc8\xcc\x81\x71\xac\x2b\x4a\x2c\xeb\xf3\x9c\x0e\x0e\x00\x7c\xc8\x02\x91\x65\x47\x95\x63\x5d\x21\x6f\x87\x8d\xc7\xe8\xd8\x97\xdf\x39\x8f\xac\x05\x1f\xa4\x22\xbb\x1f\xf9\x1c\x6d\x58\x42\x3f\x5e\x71\x90\xca\x87\x92\x0f\x6c\xa5\x06\xf4\x74\x68\xa8\xe6\xdc\xed\x1c\xb7\x6a\x27\x47\xea\xd8\x24\xe2\x93\x7a\x10\x5c\xf6\x1a\x1f\x8a\x20\x37\xc8\x18\x39\xbe\x80\xab\xd2\xfb\xb1\xbf\x62\xfe\xb7\x08\xce\x4a\x39\xb3\xef\x69\x98\x45\xeb\x21\x14\xdc\x49\x1d\xf7\xf3\x67\xc0\xe4\x70\xa7\xd3\x7d\x53\x9c\xe1\x08\xda\xc9\xbb\x73\x69\xb0\x0e\x46\x20\x98\xde\xf3\xd1\xdb\xa8\xc1\xa2\xc8\xae\x5b\xb0\x9d\x56\xcb\x49\x13\x10\x21\x1d\x19\xbe\x13\xdd\x25\xcd\x2d\xe6\x05\xf6\xd7\xa0\x4f\x84\x3a\xa8\xe8\xe6\xe5\xdb\x8f\x55\x5a\x97\x5a\xe4\x2b\xa5\x10\x3c\x2e\xd3\x81\x7b\xc4\x8f\xa7\x6e\x56\x8e\x52\xd9\xb8\xa8\x6c\xa3\x37\xff\xde\x47\x27\xc9\x52\xa7\x7a\x0e\xf4\x14\xec\xda\xac\x8e\x3f\x97\xa5\x42\xeb\x67\x2f\x3e\x39\xdb\x63\xed\x7e\xdb\xf9\x0d\x21\x96\xd0\xb4\xb5\xfe\x9c\xf0\x76\xdd\xb3\x62\x11\xcc\x4e\x75\x43\xfe\x31\x13\xa0\x8e\xe3\x10\x60\x19\x57\x41\xa0\xa9\x17\xba\x12\x67\x4f\xa3\x6f\x7c\x04\x77\xee\x9a\x49\xff\x18\x02\x7f\xb6\x90\xb2\x73\x9d\x87\x9f\x55\x2c\x30\x39\x96\x15\x79\x00\x44\x56\x23\xe7\x75\x3e\x5f\xad\x73\x5c\xba\x9c\xe4\x5f\x2a\xfb\xc1\xb2\xaf\x2e\xaa\xab\xfb\x64\xde\xac\xef\x48\x52\x40\x22\xd9\xb9\x51\x3e\xc8\xf7\xf9\xe8\x9d\x5f\xc5\x24\x1e\x6d\x8e\x26\xde\x83\x41\x8c\x4a\xed\x0d\xf7\x83\xb0\xf9\xf7\xbe\xfa\x93\x9a\xab\x65\x7a\x69\x8e\x49\x12\xe2\x95\x00\x18\x85\xe3\xb4\x63\xf4\xad\x7f\xb9\x75\xf3\x10\xd6\x67\x9a\x58\xa0\x4c\xb4\x96\x86\xd8\x2f\x49\x45\x12\x6e\xe8\x68\xd6\xbb\xd1\x66\xd7\x3b\x19\x6b\xed\x21\x55\xfd\x86\x67\xb5\x3f\x89\x24\xa3\x7d\x09\xb8\xaf\x42\x0c\xc9\x99\x6a\xf8\x3d\x65\x11\xe7\xc9\x78\x01\x87\x7a\x54\x1b\xa3\x01\x29\x89\x52\x30\x09\xcf\x37\xec\xb0\x73\x24\xdf\xed\xc8\x04\xd5\x2a\xa8\x56\xaa\xd4\x90\x4e\x69\x3b\x7c\x95\x19\x46\x3b\x5e\xe6\x7f\x4f\x99\x2d\x54\xf4\x44\x46\x16\x1f\x20\xd7\x90\xbd\xc8\xc8\x93\x60\xe0\x5c\xf8\xdc\x69\xb3\x49\xdb\x9d\x14\x5a\x7e\x94\xa7\x29\x30\xc2\x6e\xab\x98\xec\xd7\xf7\xcc\x83\xf4\xbf\xec\x07\x16\x57\xf6\x02\xbc\x16\x6c\xb5\x8d\xb4\xe1\x8a\xa1\xf6\x96\xe9\xa9\x91\xf0\xf1\xc9\x7b\x59\x4a\xa9\xbf\xd2\xfd\x2c\x2f\x9c\x24\x8d\x12\xcb\xa4\x0d\x5a\x13\x47\xbb\x3a\x07\x5c\x61\x1b\x30\x5f\x7e\x5c\x5d\x68\x3f\x0d\xae\x11\x7d\x0d\x81\xac\xfe\x10\xa9\xb0\x87\x92\x10\x02\x55\x3c\xdf\xf1\x4c\x94\x5e\xa7\xc9\xb2\x35\x5a\x86\x3e\xca\x65\xa7\x5e\x56\xe1\xbf\x14\xf9\x2d\xf1\xc6\x56\x0c\x02\x1f\x86\xd3\xa5\x75\xfa\xef\xee\x17\x01\xb1\x58\x2f\x79\x8a\xbe\xdd\x4b\x69\x4a\x2c\xb5\xbe\x53\xd0\x88\xb7\xc8\xcc\xd4\x3d\x98\xe9\x2c\x32\x1b\x38\x66\xd8\x3a\x5a\xcd\xaf\x26\x78\x7b\x33\x51\xff\xbc\x52\xc7\x27\xdb\x4b\x59\x34\x3f\x42\x98\x8d\x6f\x66\xaf\x7b\xe9\xce\x56\x9a\xad\x47\x86\xe3\x5e\x97\xf6\xce\x49\xc7\x9f\xc9\xcb\xd0\xc3\x95\xc4\x00\x63\xdf\x26\x90\x8f\xc0\xc1\xcd\xbd\xa8\x52\x5d\x97\x18\x98\x6b\xea\x42\x23\xc0\xbe\x77\xc5\xf6\x7a\x77\xf1\x2e\x10\x1b\xa0\x30\x55\xd3\x2a\xbc\x2d\xdb\xbf\x1a\xa3\x6c\x14\x63\x8a\xe4\x96\xd2\xc6\x7e\x04\xfb\xa3\x44\xa9\xfe\xbd\xd5\x51\x29\xde\x87\x6f\xa6\xe1\xc1\xe0\x88\xc3\x5d\xb0\xd1\x9f\x89\x99\x78\x7e\x77\x27\x8d\xdd\x63\x17\x78\xf1\xc5\xad\x3b\xb8\x3a\x3a\x1f\x81\xe2\x62\xee\x2a\x61\x1c\xe8\x27\x09\xf1\xe4\x43\x31\xc5\xab\x44\x89\x55\x93\xd0\x97\x96\xb2\x88\x6f\xe0\xdf\xef\x01\xa2\xf2\xaf\x72\x4c\x6e\x72\x69\xd3\xdb\xa4\xc5\x30\x35\xd3\xe1\xb8\x94\xa0\x6f\xa6\x4a\xff\xd0\x92\x8e\x79\x47\x9a\xa2\x9a\xa6\xe2\xcd\x7c\x85\xfa\x17\xd2\x1f\x05\x79\x0b\xb7\xdb\x5a\x3c\xaa\xce\xa5\x44\x83\x70\x4a\xc9\x6f\xcd\x24\x5f\xbb\xbb\xf6\x75\x47\xe3\xbc\xe2\x17\xcc\xe4\x74\x94\x1e\x19\x63\x94\x47\x45\x49\xd3\xce\x4b\x3e\x1f\x7c\xf5\x9a\x24\x25\x34\xe4\x23\x79\xf8\xdd\xb7\x55\xb2\xda\xb0\xf2\xf2\xe7\xe0\xaf\xd6\xda\x72\x67\xb4\x3c\x88\xee\x67\x8a\x85\x6d\xca\x70\x65\x08\x3a\x59\x71\xdc\x88\xd8\x32\xd4\x36\xb0\x18\xd3\xf2\x17\x39\x42\x88\x72\xf8\x82\xb9\x51\x93\x69\xbb\xa3\x04\xa8\x5b\xbd\x5c\xa9\x74\xd7\x2d\xaa\x5e\x92\x00\x2b\x7e\x9e\x13\x2e\x08\x0c\x8c\x88\x09\xec\xf2\x42\xa6\x93\x88\x35\x75\x6e\xfc\xca\x2b\xbf\x7f\x7e\x10\x69\x74\x3f\x25\xef\xb5\x69\xf7\x5d\x85\xf0\xc3\xa2\xd3\x67\xfc\x51\x1e\x84\x5e\xa6\x2c\x0e\x1b\x76\xe9\xec\x65\x85\xd6\xd6\x86\x4e\x92\x3d\xe7\xe5\xd6\x4f\x76\x28\xd6\xfb\x36\x1f\xf1\xa5\x55\x00\xa9\x8c\x02\xd7\xde\x93\xf4\x25\x98\xe3\xc9\xd7\xb3\x11\x37\xa6\x95\x17\x0b\xdf\x71\x9c\x51\x77\x37\xc8\xa7\x4f\xc0\xc4\x8f\xe6\x5f\x64\x7e\xac\xa5\x02\xb5\x88\xa1\x08\x4e\x54\xd4\x02\x8e\xfd\x20\x91\x33\x7a\x6b\x93\xf9\x57\x0e\xc5\x81\x1f\x86\x61\x58\xfc\xad\x8c\xbf\x4e\x60\xb6\xfe\xbc\x9e\x9d\xc3\xe5\x98\x2a\xa8\x2e\x29\xb2\xc4\x28\xd3\x43\xb8\xc7\x3c\x07\xa5\x41\xfc\xea\x6d\x08\x6b\xc4\xe8\x85\x8d\x8c\x2d\xcb\x09\x37\xec\x0e\x6e\x45\xe3\x49\x7e\x19\x55\x46\x9c\xa7\x79\xb9\xe2\x9a\x1e\x6e\x12\xe6\x5d\x8a\xa1\x43\xf2\x5c\x2d\x56\x1f\x4c\x6c\x74\x1e\xd7\x96\x8b\x69\x40\xf1\x3c\x72\x6e\x4d\x00\x35\x40\xcb\x14\x61\x1f\x19\x30\x50\x1c\xcd\xe4\xbb\xe4\x90\x3f\xdd\x11\x2d\xe1\xbf\x2d\xfa\x7f\x00\x00\x00\xff\xff\xee\x3e\x80\x9b\xa8\x06\x00\x00")

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

	info := bindataFileInfo{name: "Res/key.pem", size: 1704, mode: os.FileMode(384), modTime: time.Unix(1730374479, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resNodeXIni = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xc1\x6a\x84\x30\x10\x86\xef\xf3\x14\x1e\x8a\x47\x85\x1e\x17\xc2\x9e\x7a\xea\x13\x94\x65\x09\x31\x99\x46\xad\x26\xae\x33\x69\xb6\x94\xbe\x7b\x89\xb1\x28\xd2\xc3\x1e\x02\xc9\x3f\xff\xf7\x31\xe4\x62\xf0\x5d\x85\x81\xaf\x00\x8c\x77\x96\x3d\x15\xa2\xe8\x49\xa6\x47\xd5\x13\x00\x5c\x08\xd5\xac\x5b\x89\xce\x76\x0e\xaf\x00\xd6\x7b\x3b\x60\x21\x8a\x96\x79\xa2\x53\x5d\xc7\x18\xab\x1c\x56\xda\x8f\x75\xee\x9f\x6f\xe2\xe9\xfb\xf5\xe5\xed\xa7\x24\xad\x64\x98\x3e\xc5\x73\x49\x3e\xcc\x1a\xc5\xe0\xb8\xe4\x86\xc4\xcd\xcc\xa7\x71\xd5\x49\x87\x91\x1e\x77\x72\x33\x0a\x17\xd3\x7a\x8d\xea\x4c\x38\x80\x4b\x96\xb9\x73\x34\x2b\x03\x4d\xe7\xec\xb1\xd8\x39\xfb\x9f\x1f\xbe\x54\xeb\xfd\xae\x9c\xe7\xd5\x12\xef\x81\x69\x03\x9c\xc1\xfb\x41\x9f\xc3\xa5\x9f\xae\x2b\x93\xbe\xf6\x0f\x33\x41\x7f\xa4\x63\xfd\x01\xdd\x06\x0b\xbe\x6d\xf6\x1b\x00\x00\xff\xff\xa1\xfe\xf8\xac\xb1\x01\x00\x00")

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

	info := bindataFileInfo{name: "Res/node-x.ini", size: 433, mode: os.FileMode(420), modTime: time.Unix(1717941212, 0)}
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
