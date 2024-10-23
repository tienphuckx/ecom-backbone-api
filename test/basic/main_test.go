package basic

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestAddOne(t *testing.T) {
	var (
		input  = 1
		output = 2
	)
	actual := AddOne(input)
	if actual != output {
		t.Errorf("Error: AddOne(%d) Expected %d, but got %d", input, output, actual)
	}
}

func TestAddOneAdvance(t *testing.T) {
	assert.Equal(t, AddOne(2), 4, "AddOne(2) should return 3")
}
