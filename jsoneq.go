package jsoneq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/fatih/color"
)

var red = color.New(color.FgRed)
var green = color.New(color.FgGreen)
var yellow = color.New(color.FgYellow)
var blue = color.New(color.FgBlue)

func FormatJSON(j map[string]interface{}) string {
	b, _ := json.Marshal(j)

	var s bytes.Buffer
	json.Indent(&s, b, "", "  ")

	return s.String()
}

func AssertEqual(t *testing.T, a, b []byte) {
	var ja map[string]interface{}
	var jb map[string]interface{}

	err := json.Unmarshal(a, &ja)
	err = json.Unmarshal(b, &jb)
	if err != nil {
		t.Error(err)
	}

	errors := checkEqual(t, ja, jb, "")

	msg := fmt.Sprintf("Expected:\n%s\n\nActual:\n%s\n\nDifference:\n%s", FormatJSON(ja), FormatJSON(jb), strings.Join(errors, "\n"))
	t.Errorf("\n%s", msg)
}

func formatDiff(key string, lhs, rhs interface{}) []string {
	return []string{
		red.Sprintf("- %v: %v", key, lhs),
		green.Sprintf("+ %v: %v", key, rhs),
	}
}

func checkEqual(t *testing.T, a, b map[string]interface{}, root string) []string {
	var errors []string

	for key, val := range a {
		var newRoot string
		if root == "" {
			newRoot = key
		} else {
			newRoot = fmt.Sprintf("%s.%s", root, key)
		}

		if b[key] != nil {
			lhs := val
			rhs := b[key]

			lhsV := reflect.ValueOf(lhs)
			rhsV := reflect.ValueOf(rhs)

			if lhsV.Kind() != rhsV.Kind() {
				errors = append(errors, yellow.Sprintf("%s: type mismatch %s != %s", newRoot, lhsV.Kind(), rhsV.Kind()))
			}

			switch lhsV.Kind() {
			case reflect.Float64, reflect.Bool:
				if lhs != rhs {
					errors = append(errors, formatDiff(newRoot, lhs, rhs)...)
				}
			case reflect.String:
				if strings.Compare(lhs.(string), rhs.(string)) != 0 {
					errors = append(errors, formatDiff(newRoot, lhs, rhs)...)
				}
			case reflect.Slice:
				if !reflect.DeepEqual(lhs, rhs) {
					errors = append(errors, formatDiff(newRoot, lhs, rhs)...)
				}
			case reflect.Map:
				errors = append(errors, checkEqual(t, lhs.(map[string]interface{}), rhs.(map[string]interface{}), newRoot)...)
			default:
				errors = append(errors, yellow.Sprintf("%s: has an invalid type: %s", lhsV.Kind()))
			}
		} else {
			errors = append(errors, red.Sprintf("- %s", newRoot))
		}
	}

	return errors
}
