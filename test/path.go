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
