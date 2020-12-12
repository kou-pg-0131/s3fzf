package infrastructures

import (
	"io"
	"os"
)

// FileSystem .
type FileSystem struct{}

// NewFileSystem .
func NewFileSystem() *FileSystem {
	return new(FileSystem)
}

// Create .
func (fs *FileSystem) Create(path string) (io.WriteCloser, error) {
	return os.Create(path)
}

// Copy .
func (fs *FileSystem) Copy(dist io.Writer, src io.Reader) error {
	_, err := io.Copy(dist, src)
	return err
}
