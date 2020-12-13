package cmd

import (
	"bufio"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * NewScanner()
 */

func Test_NewScanner(t *testing.T) {
	s := NewScanner()

	assert.IsType(t, &bufio.Scanner{}, s.scannerAPI)
	assert.NotNil(t, s.scannerAPI)
}

/*
 * Scanner.Scan()
 */

func TestScanner_Scan_ReturnStringWhenSuccess(t *testing.T) {
	msc := new(mockScannerAPI)
	msc.On("Scan").Return(true)
	msc.On("Text").Return(" STRING ")

	scn := &Scanner{scannerAPI: msc}

	s, err := scn.Scan()

	assert.Equal(t, "STRING", s)
	assert.Nil(t, err)
	msc.AssertNumberOfCalls(t, "Scan", 1)
	msc.AssertNumberOfCalls(t, "Text", 1)
}

func TestScanner_Scan_ReturnErrorWhenScanFailed(t *testing.T) {
	msc := new(mockScannerAPI)
	msc.On("Scan").Return(false)
	msc.On("Err").Return(errors.New("SOMETHING_WRONG"))

	scn := &Scanner{scannerAPI: msc}

	s, err := scn.Scan()

	assert.Equal(t, "", s)
	assert.EqualError(t, err, "SOMETHING_WRONG")
	msc.AssertNumberOfCalls(t, "Scan", 1)
	msc.AssertNumberOfCalls(t, "Err", 1)
}
