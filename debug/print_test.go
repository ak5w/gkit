package debug_test

import (
	"testing"

	"github.com/ak5w/gkit/debug"
)

func TestSlice(t *testing.T) {

	s := []int{1, 2, 3}
	debug.Slice(s)
}
