package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveAccent(t *testing.T) {
	assert.Equal(t, 'e', RemoveAccent('é'))
	assert.Equal(t, 'e', RemoveAccent('è'))
	assert.Equal(t, 'e', RemoveAccent('ê'))
	assert.Equal(t, 'e', RemoveAccent('ë'))
	assert.Equal(t, 'a', RemoveAccent('à'))
	assert.Equal(t, 'u', RemoveAccent('ü'))
}
