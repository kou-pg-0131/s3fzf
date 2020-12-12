package infrastructures

import (
	"bytes"
	"errors"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * NewFileSystem()
 */

func Test_NewFileSystem_ReturnFileSystem(t *testing.T) {
	fs := NewFileSystem()

	assert.Equal(t, reflect.ValueOf(os.Create).Pointer(), reflect.ValueOf(fs.osCreate).Pointer())
	assert.Equal(t, reflect.ValueOf(io.Copy).Pointer(), reflect.ValueOf(fs.ioCopy).Pointer())
}

/*
 * FileSystem.Create()
 */

func TestFileSystem_Create_ReturnIOWriteCloserWhenSucceeded(t *testing.T) {
	mos := new(mockOS)
	mos.On("Create", "PATH").Return(&os.File{}, nil)

	fs := &FileSystem{osCreate: mos.Create}

	w, err := fs.Create("PATH")

	assert.Equal(t, &os.File{}, w)
	assert.Nil(t, err)
	mos.AssertNumberOfCalls(t, "Create", 1)
}

func TestFileSystem_Create_ReturnErrorWhenCreateFailed(t *testing.T) {
	mos := new(mockOS)
	mos.On("Create", "PATH").Return((*os.File)(nil), errors.New("SOMETHING_WRONG"))

	fs := &FileSystem{osCreate: mos.Create}

	w, err := fs.Create("PATH")

	assert.Nil(t, w)
	assert.EqualError(t, err, "SOMETHING_WRONG")
	mos.AssertNumberOfCalls(t, "Create", 1)
}

/*
 * FileSystem.Copy()
 */

func TestFileSystem_Copy_ReturnNilWhenSucceeded(t *testing.T) {
	w := bytes.NewBufferString("DIST")
	r := strings.NewReader("SRC")

	mio := new(mockIO)
	mio.On("Copy", w, r).Return(int64(0), nil)

	fs := &FileSystem{ioCopy: mio.Copy}

	err := fs.Copy(w, r)

	assert.Nil(t, err)
	mio.AssertNumberOfCalls(t, "Copy", 1)
}

func TestFileSystem_Copy_ReturnErrorWhenCopyFailed(t *testing.T) {
	w := bytes.NewBufferString("DIST")
	r := strings.NewReader("SRC")

	mio := new(mockIO)
	mio.On("Copy", w, r).Return(int64(0), errors.New("SOMETHING_WRONG"))

	fs := &FileSystem{ioCopy: mio.Copy}

	err := fs.Copy(w, r)

	assert.EqualError(t, err, "SOMETHING_WRONG")
	mio.AssertNumberOfCalls(t, "Copy", 1)
}
