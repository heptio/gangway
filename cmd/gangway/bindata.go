// Code generated by "esc -o cmd/gangway/bindata.go templates/"; DO NOT EDIT.

package main

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

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
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

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/templates/commandline.tmpl": {
		local:   "templates/commandline.tmpl",
		size:    4350,
		modtime: 1537502709,
		compressed: `
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
XogNWt3yGrT6b8d/AwAA///cMjut/hAAAA==
`,
	},

	"/templates/home.tmpl": {
		local:   "templates/home.tmpl",
		size:    2088,
		modtime: 1532555682,
		compressed: `
H4sIAAAAAAAC/4xWUW/jNgx+76/g6V42LLKbuxXYcnaGrd2AYnfrgHYH7JGWGZutLPkkOW5u2H8fJMdJ
mhuwBUgsUhT5kR9ppXh1c3f98OfvP0MbOr2+KOIDNJqmFGREVBDW6wuAoqOA0IbQS/o08LYU19YEMkE+
7HoSoCapFIGeQx7dvAPVovMUyj8efpHfifzoxmBHpdgyjb114eTwyHVoy5q2rEgmYQFsODBq6RVqKpcL
6PCZu6GbFdnl3nXgoGndoGlG3BX5JF7EnVdSwvX9PYCUyVKzeYLW0aYUMSO/yvONNcFnjbWNJuzZZ8p2
OStrfthgx3pXfsBAjlF/c6us8QIc6VL4sNPkW6Igjo7Pd84iqdo8+kxpO9QbjY5SJHzE51xz5fNuH4c/
U36ZLS8vsze58i/0WccmU95PMVOkuALoHWUV+hb+SmL8VKieGmcHU0tltXUrqDSqp3cHg7329bffb6q3
V0d9LIj0/JlW4DvU+mxnqsoKrq3xVqNffLAGlV28HxTXuFfT4j1X5DCwNfDBGru4oUf8OMA9Gj8pfuLg
gyPs4CM5PNm4toNjcvAbjQvorLG+R0VHFHZLbqPtuAIcgj3qR+tqOTrsV2Cs6/AE+NhyIJn8rGKtpp2/
L9Iji2TLSlv1tC9fj3XNplnBJSyv+ufZ+tw4m5mRUelPan9SQjYtOQ6zjyLfs1bk04AVla13iU6DW1Aa
vS+F5qYNstIDQfwR4KymUhjccpNqKibai5oPZwxuU/Y9uTRWyIacWBcIXJdC28bKo3rfma/FfLpyaGoZ
rcRxkHC9z6gY9GzoIjRouSZpjeyolvFgbceIaTbPB30UBp0QRHydrVjTIaiPXgxuz0/u13hEWWNAiSrw
FgP5f3VWDSFYE3tdY+9JrAuet17yJNYdmaHIeX3IsMhr3iYScoPT4qSynlTqY2Nlj7WsbBApIzY1PcsK
TSr0F4SckDBnVLl1/M5iu5xtYy9E3sgEclCjeyIj3x6ogF+HipyhQB5+HEJLJrBKjVDk7fLg7yS4s+Pe
2yF6DHh1HtBq8Ms3kBpOrB9a9jAE1hx2MLLW0JLuYWeHOGtzWIKRQxu17hSY0oOP6NMmGrjrydzexBeC
IRXgq7vbm+uvIU5uBvfcGGADwUJDAXxAF6jOiry9OmQzc/L/UpubJde2YSMmhmJbahs5S81x7JVgpEbX
EIy4JS9ps4kQJyHVYpq7dcJ5a05H4QWsmdDzLkqLqD3cQnCvHPfB7y8inyTwTp3cD7am7PHTQG6XroZp
Kd9ky2yZ3vyPXqyLfDr6hZfHlzfFf9nGy/XMKCGfXkdFPv0tuPgnAAD//yEDNLMoCAAA
`,
	},

	"/templates/kubeconfig.tmpl": {
		local:   "templates/kubeconfig.tmpl",
		size:    589,
		modtime: 1537503171,
		compressed: `
H4sIAAAAAAAC/3SRQW/DIAyF7/wKq3cm7cqtaneoNE1Tu+2OiNNaTSAyEG2K8t8nIGzp1N149mfeM+iB
PpA9OatgfBSmiz4geyUkLGclAAAMcqCWjA4odQwXxxS+FEwTPOwKt9vCPGfWI4/ICjapu309nLJ+Pz7D
PG8EgNU93oy+6B7TsHE24Gco7uW8uC9R7g2lfvS1+dRr6kr5f5vIjDbIanGPuZJtFOycbeksBsYWGa1B
r2CaRbLLIX8tVr45TI6VHkoO7EZqagnSYi2dq0q7UQpDTc2R5GFfV1shHg1jWGOnXFmj1MjgrmgLddi/
JXELDJK8j8gycrdgWZcP+gEZW0Z/WV93LKU/d5ZHcNSY7wAAAP//w1bUIk0CAAA=
`,
	},

	"/": {
		isDir: true,
		local: "",
	},

	"/templates": {
		isDir: true,
		local: "templates",
	},
}
