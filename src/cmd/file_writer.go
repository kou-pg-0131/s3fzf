package cmd

import (
	"io"
	"os"

	"github.com/kou-pg-0131/s3fzf/src/infrastructures"
	"github.com/kou-pg-0131/s3fzf/src/interfaces/gateways"
)

// IFileWriter .
type IFileWriter interface {
	Write(out string, r io.Reader) error
}

// FileWriter .
type FileWriter struct {
	fileSystem    gateways.IFileSystem
	defaultWriter io.Writer
}

// NewFileWriter .
func NewFileWriter() *FileWriter {
	return &FileWriter{
		fileSystem:    infrastructures.NewFileSystem(),
		defaultWriter: os.Stdout,
	}
}

// Write .
func (fw *FileWriter) Write(out string, r io.Reader) error {
	var w io.Writer
	if out != "-" {
		f, err := fw.fileSystem.Create(out)
		if err != nil {
			return err
		}
		defer f.Close()
		w = f
	} else {
		w = fw.defaultWriter
	}

	if err := fw.fileSystem.Copy(w, r); err != nil {
		return err
	}

	return nil
}
