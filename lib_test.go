package stripmargin_test

import (
	"testing"

	"github.com/cstanze/stripmargin"
)

func TestStripMargin(t *testing.T) {
	input := `
    |This is a
    |multiline string
    |with a margin
  `
	expected := `
This is a
multiline string
with a margin
`
	actual := stripmargin.StripMargin(input)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
