package parser_test

import (
	"reflect"
	"testing"

	"github.com/errybase/go-dotenv/parser"
)

var s = []byte(`
# comment line
FOO=foo
BAR = bar
BAZ = "baz"
`)

func TestParse(t *testing.T) {
	want := map[string]string{
		"FOO": "foo",
		"BAR": "bar",
		"BAZ": "baz",
	}
	got := parser.Parse(s)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %+v, got: %+v", want, got)
	}
}
