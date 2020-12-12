package gateways

import "io"

// IFileSystem .
type IFileSystem interface {
	Copy(dest io.Writer, src io.Reader) error
	Create(path string) (io.WriteCloser, error)
}
