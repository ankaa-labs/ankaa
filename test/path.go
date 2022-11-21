// Copyright (c) 2022 The ankaa-labs Authors
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package test

import (
	"os"
	"path/filepath"
	"runtime"

	xerrors "github.com/pkg/errors"
)

// CurrentProjectPath get the project root path
func CurrentProjectPath() string {
	path := currentFilePath()

	ppath, err := filepath.Abs(filepath.Join(filepath.Dir(path), "../"))
	if err != nil {
		panic(xerrors.Wrapf(err, "Get current project path with %s failed", path))
	}

	f, err := os.Stat(ppath)
	if err != nil {
		panic(xerrors.Wrapf(err, "Stat project path %v failed", ppath))
	}

	if f.Mode()&os.ModeSymlink != 0 {
		fpath, err := os.Readlink(ppath)
		if err != nil {
			panic(xerrors.Wrapf(err, "Readlink from path %v failed", fpath))
		}
		ppath = fpath
	}

	return ppath
}

func currentFilePath() string {
	_, file, _, _ := runtime.Caller(1)
	return file
}
