package backup

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"
)

const (
	volumeContainerFilename = "volume-container.json"
)

type containerVolume struct {
	path     string
	hostPath string
	tw       *tar.Writer
}

func newContainerVolume(path, hostPath string, tw *tar.Writer) *containerVolume {
	return &containerVolume{
		path:     path,
		hostPath: hostPath,
		tw:       tw,
	}
}

func (v *containerVolume) Store() error {
	return filepath.Walk(v.hostPath, v.addFile)
}

func (v *containerVolume) addFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.Mode().IsDir() {
		return nil
	}

	relPath := path[len(filepath.Dir(v.hostPath))+1:] // relative to volume parent directory (<docker>/vfs/dir/)
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	th, err := tar.FileInfoHeader(info, relPath)
	if err != nil {
		return err
	}
	th.Name = relPath

	if err := v.tw.WriteHeader(th); err != nil {
		return err
	}

	if _, err := io.Copy(v.tw, file); err != nil {
		return err
	}
	return nil
}
