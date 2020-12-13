package cmd

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * NewConfirmer()
 */

func Test_NewConfirmer(t *testing.T) {
	c := NewConfirmer()

	assert.NotNil(t, c)
	assert.NotNil(t, c.scanner)
}

/*
 * confirmer.Confirm()
 */

func TestConfirmer_Confirm_ReturnTrueWhenYes(t *testing.T) {
	mscn := new(mockScanner)
	mscn.On("Scan").Return("y", nil)

	c := &Confirmer{scanner: mscn}

	y, err := c.Confirm("MESSAGE")

	assert.Equal(t, true, y)
	assert.Nil(t, err)
	mscn.AssertNumberOfCalls(t, "Scan", 1)
}

func TestConfirmer_Confirm_ReturnFalseWhenNotYes(t *testing.T) {
	mscn := new(mockScanner)
	mscn.On("Scan").Return("no", nil)

	c := &Confirmer{scanner: mscn}

	y, err := c.Confirm("MESSAGE")

	assert.Equal(t, false, y)
	assert.Nil(t, err)
	mscn.AssertNumberOfCalls(t, "Scan", 1)
}

func TestConfirmer_Confirm_ReturnErrorWhenScanFailed(t *testing.T) {
	mscn := new(mockScanner)
	mscn.On("Scan").Return("", errors.New("SOMETHING_WRONG"))

	c := &Confirmer{scanner: mscn}

	y, err := c.Confirm("MESSAGE")

	assert.Equal(t, false, y)
	assert.EqualError(t, err, "SOMETHING_WRONG")
	mscn.AssertNumberOfCalls(t, "Scan", 1)
}
