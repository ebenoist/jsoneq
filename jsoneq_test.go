package jsoneq_test

import (
	"testing"

	"github.com/ebenoist/jsoneq"
)

func Test_ReturnsFalse(t *testing.T) {
	a := `{
		"bar": "foo",
		"baz": "zed",
		"foo": "zed",
		"a": [2,3,1],
		"l":2,
		"persona": {
			"id": 10,
			"name": "Erik"
		}
	}`

	b := `{
		"foo": "bar",
		"fooze":"baz",
		"id": 12,
		"a": [3,2,1],
		"l":2 ,
		"persona": {
			"id": 12
		}
	}`

	jsoneq.AssertEqual(t, []byte(a), []byte(b))
}

func Test_Arrays(t *testing.T) {
}
