// Code generated by "esc "; DO NOT EDIT.

package js

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// _escFS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func _escFS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// _escDir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func _escDir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// _escFSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func _escFSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// _escFSMustByte is the same as _escFSByte, but panics if name is not present.
func _escFSMustByte(useLocal bool, name string) []byte {
	b, err := _escFSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// _escFSString is the string version of _escFSByte.
func _escFSString(useLocal bool, name string) (string, error) {
	b, err := _escFSByte(useLocal, name)
	return string(b), err
}

// _escFSMustString is the string version of _escFSMustByte.
func _escFSMustString(useLocal bool, name string) string {
	return string(_escFSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/helpers.js": {
		local:   "pkg/js/helpers.js",
		size:    14682,
		modtime: 0,
		compressed: `
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
`,
	},

	"/": {
		isDir: true,
		local: "pkg/js",
	},
}
