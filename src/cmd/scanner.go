package cmd

import (
	"bufio"
	"os"
	"strings"
)

// IScannerAPI .
type IScannerAPI interface {
	Scan() bool
	Text() string
	Err() error
}

// IScanner .
type IScanner interface {
	Scan() (string, error)
}

// Scanner .
type Scanner struct {
	scannerAPI IScannerAPI
}

// NewScanner .
func NewScanner() *Scanner {
	return &Scanner{scannerAPI: bufio.NewScanner(os.Stdin)}
}

// Scan .
func (scn *Scanner) Scan() (string, error) {
	if !scn.scannerAPI.Scan() {
		return "", scn.scannerAPI.Err()
	}

	return strings.TrimSpace(scn.scannerAPI.Text()), nil
}
