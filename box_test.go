package golang_united_school_homework

import (
	"testing"

	"gotest.tools/assert"
)

func TestNewBox(t *testing.T) {
	box := NewBox(0)
	assert.Equal(t, box == nil, false, `Created box is not nil`)
}
