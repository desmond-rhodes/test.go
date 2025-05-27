package example_test

import (
	"testing"

	"github.com/desmond-rhodes/test.go/example"
)

func Test(t *testing.T) {
	t.Errorf(example.Greeting)
}
