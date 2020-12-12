package infrastructures

import (
	"errors"
	"reflect"
	"testing"

	fzf "github.com/ktr0731/go-fuzzyfinder"
	"github.com/nsf/termbox-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/*
 * NewFZF()
 */

func Test_NewFZF_ReturnFZF(t *testing.T) {
	f := NewFZF()

	assert.Equal(t, reflect.ValueOf(fzf.Find).Pointer(), reflect.ValueOf(f.fzfFind).Pointer())
	assert.Equal(t, reflect.ValueOf(termbox.Close).Pointer(), reflect.ValueOf(f.termboxClose).Pointer())
	assert.Equal(t, reflect.ValueOf(termbox.Sync).Pointer(), reflect.ValueOf(f.termboxSync).Pointer())
}

/*
 * FZF.Find()
 */

func TestFZF_Find_ReturnIndexWhenSucceeded(t *testing.T) {
	mfzf := new(mockFZF)
	mfzf.On("Find", []string{"item1", "item2"}, mock.Anything, mock.Anything).Return(1, nil)

	f := &FZF{fzfFind: mfzf.Find}

	idx, err := f.Find([]string{"item1", "item2"}, func(i int) string { return "" }, func(i, w, h int) string { return "" })

	assert.Equal(t, 1, idx)
	assert.Nil(t, err)
	mfzf.AssertNumberOfCalls(t, "Find", 1)
}

func TestFZF_Find_ReturnErrorWhenFindFailed(t *testing.T) {
	mfzf := new(mockFZF)
	mfzf.On("Find", []string{"item1", "item2"}, mock.Anything, mock.Anything).Return(0, errors.New("SOMETHING_WRONG"))

	f := &FZF{fzfFind: mfzf.Find}

	idx, err := f.Find([]string{"item1", "item2"}, func(i int) string { return "" }, func(i, w, h int) string { return "" })

	assert.Equal(t, -1, idx)
	assert.EqualError(t, err, "SOMETHING_WRONG")
	mfzf.AssertNumberOfCalls(t, "Find", 1)
}

/*
 * FZF.Close()
 */

func TestFZF_Close(t *testing.T) {
	mbox := new(mockTermbox)
	mbox.On("Close")

	f := &FZF{termboxClose: mbox.Close}

	f.Close()

	mbox.AssertNumberOfCalls(t, "Close", 1)
}

/*
 * FZF.Sync()
 */

func TestFZF_Sync_ReturnNilWhenSucceeded(t *testing.T) {
	mbox := new(mockTermbox)
	mbox.On("Sync").Return(nil)

	f := &FZF{termboxSync: mbox.Sync}

	err := f.Sync()

	assert.Nil(t, err)
	mbox.AssertNumberOfCalls(t, "Sync", 1)
}

func TestFZF_Sync_ReturnErrorWhenSyncFailed(t *testing.T) {
	mbox := new(mockTermbox)
	mbox.On("Sync").Return(errors.New("SOMETHING_WRONG"))

	f := &FZF{termboxSync: mbox.Sync}

	err := f.Sync()

	assert.EqualError(t, err, "SOMETHING_WRONG")
	mbox.AssertNumberOfCalls(t, "Sync", 1)
}
