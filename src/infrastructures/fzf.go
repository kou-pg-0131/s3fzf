package infrastructures

import (
	fzf "github.com/ktr0731/go-fuzzyfinder"
	"github.com/nsf/termbox-go"
)

// FZF .
type FZF struct{}

// NewFZF .
func NewFZF() *FZF {
	return new(FZF)
}

// Find .
func (f *FZF) Find(list interface{}, itemFunc func(int) string, previewFunc func(int, int, int) string) (int, error) {
	i, err := fzf.Find(list, itemFunc, fzf.WithHotReload(), fzf.WithPreviewWindow(previewFunc))
	if err != nil {
		return -1, err
	}
	return i, nil
}

// Close .
func (f *FZF) Close() {
	termbox.Close()
}
