package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * New()
 */

func Test_New_ReturnCommand(t *testing.T) {
	mcfm := new(mockConfirmer)
	mfw := new(mockFileWriter)
	ms3c := new(mockS3Controller)

	ic := New(&CommandConfig{
		FileWriter:   mfw,
		Confirmer:    mcfm,
		S3Controller: ms3c,
	})

	c, ok := ic.(*Command)
	if !ok {
		t.Fatal()
	}

	assert.NotNil(t, c)
	assert.Equal(t, mcfm, c.confirmer)
	assert.Equal(t, mfw, c.fileWriter)
	assert.Equal(t, ms3c, c.s3Controller)
}
