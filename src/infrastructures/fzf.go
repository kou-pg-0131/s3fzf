package infrastructures

import (
	fzf "github.com/ktr0731/go-fuzzyfinder"
	"github.com/nsf/termbox-go"
)

// FZF .
type FZF struct {
	fzfFind      func(list interface{}, itemFunc func(int) string, opts ...fzf.Option) (int, error)
	termboxClose func()
	termboxSync  func() error
}

// NewFZF .
func NewFZF() *FZF {
	return &FZF{
		fzfFind:      fzf.Find,
		termboxClose: termbox.Close,
		termboxSync:  termbox.Sync,
	}
}

// Find .
func (f *FZF) Find(list interface{}, itemFunc func(int) string, previewFunc func(int, int, int) string) (int, error) {
	i, err := f.fzfFind(list, itemFunc, fzf.WithHotReload(), fzf.WithPreviewWindow(previewFunc))
	if err != nil {
		return -1, err
	}
	return i, nil
}

// Close .
func (f *FZF) Close() {
	f.termboxClose()
}

// Sync .
func (f *FZF) Sync() error {
	return f.termboxSync()
}
