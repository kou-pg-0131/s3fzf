package cmd

import (
	"bytes"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/kou-pg-0131/s3fzf/src/infrastructures"
	"github.com/stretchr/testify/assert"
)

/*
 * NewFileWriter()
 */

func Test_NewFileWriter_ReturnFileWriter(t *testing.T) {
	w := NewFileWriter()

	assert.Equal(t, os.Stdout, w.defaultWriter)
	assert.IsType(t, &infrastructures.FileSystem{}, w.fileSystem)
	assert.NotNil(t, w.fileSystem)
}

/*
 * FileWriter.Write()
 */

func TestFileWriter_Write_ReturnNilWhenOutputSpecifiedAndSucceeded(t *testing.T) {
	r := strings.NewReader("FILE")

	mf := new(mockIOWriteCloser)
	mf.On("Close").Return(nil)

	mfs := new(mockFileSystem)
	mfs.On("Create", "OUTPUT").Return(mf, nil)
	mfs.On("Copy", mf, r).Return(nil)

	fw := &FileWriter{fileSystem: mfs}

	err := fw.Write("OUTPUT", r)

	assert.Nil(t, err)
	mf.AssertNumberOfCalls(t, "Close", 1)
	mfs.AssertNumberOfCalls(t, "Create", 1)
	mfs.AssertNumberOfCalls(t, "Copy", 1)
}

func TestFileWriter_Write_ReturnNilWhenOutputNotSpecifiedAndSucceeded(t *testing.T) {
	r := strings.NewReader("FILE")
	w := new(bytes.Buffer)

	mfs := new(mockFileSystem)
	mfs.On("Copy", w, r).Return(nil)

	fw := &FileWriter{defaultWriter: w, fileSystem: mfs}

	err := fw.Write("", r)

	assert.Nil(t, err)
	mfs.AssertNumberOfCalls(t, "Copy", 1)
}

func TestFileWriter_Write_ReturnErrorWhenCreateFailed(t *testing.T) {
	r := strings.NewReader("FILE")

	mf := new(mockIOWriteCloser)

	mfs := new(mockFileSystem)
	mfs.On("Create", "OUTPUT").Return(mf, errors.New("SOMETHING_WRONG"))

	fw := &FileWriter{fileSystem: mfs}

	err := fw.Write("OUTPUT", r)

	assert.EqualError(t, err, "SOMETHING_WRONG")
	mfs.AssertNumberOfCalls(t, "Create", 1)
}

func TestFileWriter_Write_ReturnErrorWhenCopyFailed(t *testing.T) {
	r := strings.NewReader("FILE")
	w := new(bytes.Buffer)

	mfs := new(mockFileSystem)
	mfs.On("Copy", w, r).Return(errors.New("SOMETHING_WRONG"))

	fw := &FileWriter{defaultWriter: w, fileSystem: mfs}

	err := fw.Write("", r)

	assert.EqualError(t, err, "SOMETHING_WRONG")
	mfs.AssertNumberOfCalls(t, "Copy", 1)
}
