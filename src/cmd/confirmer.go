package cmd

import "fmt"

// IConfirmer .
type IConfirmer interface {
	Confirm(msg string) (bool, error)
}

// Confirmer .
type Confirmer struct {
	scanner IScanner
}

// NewConfirmer .
func NewConfirmer() *Confirmer {
	return &Confirmer{scanner: NewScanner()}
}

// Confirm .
func (c *Confirmer) Confirm(msg string) (bool, error) {
	fmt.Print(msg)
	i, err := c.scanner.Scan()
	if err != nil {
		return false, err
	}

	return i == "y", nil
}
