// S19: An atomic file implementation.
//
// BSD License
// For atomicfile software
// Copyright (c) 2015, Facebook, Inc. All rights reserved.
//
// Package atomicfile provides the ability to write a file with an eventual
// rename on Close (using os.Rename). This allows for a file to always be in a
// consistent state and never represent an in-progress write.
//
// NOTE: `os.Rename` may not be atomic on your operating system.
//
// OUTPUT:
//
//     $ go run main.go
//     2017/02/26 16:27:03 tempfile at: hello.txt257898699
//
package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

// File behaves like os.File, but does an atomic rename operation at Close.
type File struct {
	*os.File
	path string
}

// New creates a new temporary file that will replace the file at the given
// path when Closed.
func New(path string, mode os.FileMode) (*File, error) {
	f, err := ioutil.TempFile(filepath.Dir(path), filepath.Base(path))
	if err != nil {
		return nil, err
	}
	log.Printf("tempfile at: %s", f.Name())
	if err := os.Chmod(f.Name(), mode); err != nil {
		f.Close()
		os.Remove(f.Name())
		return nil, err
	}
	return &File{File: f, path: path}, nil
}

// Close the file replacing the configured file.
func (f *File) Close() error {
	if err := f.File.Close(); err != nil {
		os.Remove(f.File.Name())
		return err
	}
	if err := os.Rename(f.Name(), f.path); err != nil {
		return err
	}
	return nil
}

// Abort closes the file and removes it instead of replacing the configured
// file. This is useful if after starting to write to the file you decide you
// don't want it anymore.
func (f *File) Abort() error {
	if err := f.File.Close(); err != nil {
		os.Remove(f.Name())
		return err
	}
	if err := os.Remove(f.Name()); err != nil {
		return err
	}
	return nil
}

func main() {
	file, err := New("hello.txt", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if _, err := io.WriteString(file, "Atomic gopher.\n"); err != nil {
		log.Fatal(err)
	}
	// If you run "ls" in the directory, you should see the temporary file,
	// e.g. hello.txt257898699.
	// When the program finishes (after 15 seconds), the temporary file will be gone.
	time.Sleep(15 * time.Second)
}
