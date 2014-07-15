package backup

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"
	"syscall"
)

const (
	volumeContainerFilename = "volume-container.json"
)

type containerVolume struct {
	path     string
	hostPath string
	tw       *tar.Writer
	size     uint
}

func newContainerVolume(path, hostPath string, tw *tar.Writer) *containerVolume {
	return &containerVolume{
		path:     path,
		hostPath: hostPath,
		tw:       tw,
	}
}

func (v *containerVolume) Store() (uint, error) {
	return v.size, filepath.Walk(v.hostPath, v.addFile)
}

func (v *containerVolume) addFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	relPath := path[len(filepath.Dir(v.hostPath))+1:] // relative to volume parent directory (<docker>/vfs/dir/)
	th, err := tar.FileInfoHeader(info, relPath)
	if err != nil {
		return err
	}
	th.Name = relPath
	if si, ok := info.Sys().(*syscall.Stat_t); ok {
		th.Uid = int(si.Uid)
		th.Gid = int(si.Gid)
	}

	if err := v.tw.WriteHeader(th); err != nil {
		return err
	}

	if !info.Mode().IsDir() {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		n, err := io.Copy(v.tw, file)
		if err != nil {
			return err
		}
		v.size = v.size + uint(n)
	}
	return nil
}
