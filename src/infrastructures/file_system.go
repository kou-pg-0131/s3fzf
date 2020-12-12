package infrastructures

import (
	"io"
	"os"
)

// FileSystem .
type FileSystem struct {
	osCreate func(string) (*os.File, error)
	ioCopy   func(io.Writer, io.Reader) (int64, error)
}

// NewFileSystem .
func NewFileSystem() *FileSystem {
	return &FileSystem{
		osCreate: os.Create,
		ioCopy:   io.Copy,
	}
}

// Create .
func (fs *FileSystem) Create(path string) (io.WriteCloser, error) {
	w, err := fs.osCreate(path)
	if err != nil {
		return nil, err
	}
	return w, nil
}

// Copy .
func (fs *FileSystem) Copy(dist io.Writer, src io.Reader) error {
	_, err := fs.ioCopy(dist, src)
	return err
}
