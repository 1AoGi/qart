// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"qart/appfs/proto"
)

// Root is the root of the local file system.  It has no effect on App Engine.
var Root = "."

type Context struct{}

type CacheKey struct{}

func NewContext(req *http.Request) *Context {
	return &Context{}
}

func (*Context) CacheRead(ckey CacheKey, path string) (CacheKey, []byte, bool) {
	return ckey, nil, false
}

func (*Context) CacheWrite(ckey CacheKey, data []byte) {
}

func (*Context) Read(path string) ([]byte, *proto.FileInfo, error) {
	p, err := filepath.Abs(filepath.Join(Root, path))
	if err != nil {
		panic(err)
	}
	log.Printf("Read file from: %s\n", p)
	dir, err := os.Stat(p)
	if err != nil {
		return nil, nil, err
	}
	fi := &proto.FileInfo{
		Name:    dir.Name(),
		ModTime: dir.ModTime(),
		Size:    dir.Size(),
		IsDir:   dir.IsDir(),
	}
	data, err := ioutil.ReadFile(p)
	return data, fi, err
}

func (*Context) Write(path string, data []byte) error {
	p, err := filepath.Abs(filepath.Join(Root, path))
	if err != nil {
		panic(err)
	}
	log.Printf("Write file to: %s\n", p)
	return ioutil.WriteFile(p, data, 0666)
}

func (*Context) Remove(path string) error {
	p, err := filepath.Abs(filepath.Join(Root, path))
	if err != nil {
		panic(err)
	}
	log.Printf("Remove file to: %s\n", p)
	return os.Remove(p)
}

func (*Context) Mkdir(path string) error {
	p, err := filepath.Abs(filepath.Join(Root, path))
	if err != nil {
		panic(err)
	}
	log.Printf("Mkdir for: %s\n", p)
	fi, err := os.Stat(p)
	if err == nil && fi.IsDir() {
		return nil
	}
	return os.Mkdir(p, 0777)
}

func (*Context) Readdir(path string) ([]proto.FileInfo, error) {
	p, err := filepath.Abs(filepath.Join(Root, path))
	if err != nil {
		panic(err)
	}
	log.Printf("Readdir from: %s\n", p)
	dirs, err := ioutil.ReadDir(p)
	if err != nil {
		return nil, err
	}
	var out []proto.FileInfo
	for _, dir := range dirs {
		out = append(out, proto.FileInfo{
			Name:    dir.Name(),
			ModTime: dir.ModTime(),
			Size:    dir.Size(),
			IsDir:   dir.IsDir(),
		})
	}
	return out, nil
}
