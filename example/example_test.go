package example_test

import (
	"testing"

	"github.com/desmond-rhodes/test.go/example"
)

func Test(t *testing.T) {
	t.Errorf(example.Greeting)
}

func Fuzz(f *testing.F) {
	f.Fuzz(func(t *testing.T, _ byte) {
		t.Skip()
	})
}
