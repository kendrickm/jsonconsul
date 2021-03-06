package jsonconsul

import (
	"fmt"
	"testing"
)

func ExampleJsonSet_RunBadJson() {
	ji := &JsonSet{}
	ji.ParseFlags([]string{"blah/blah", "\"a"})
	err := ji.Run()
	fmt.Println(err)

	// Output:
	// Can't set the key blah/blah invalid value: "a
	// If it is a string please wrap the value with quotes. Ex. \"value\"
}

func ExampleJsonSet_RunBadExpectedType() {
	ji := &JsonSet{}
	ji.ParseFlags([]string{"blah/blah", "true"})
	ji.setExpectedType("int")
	err := ji.Run()
	fmt.Println(err)

	// Output:
	// Invalid type. Value is a bool. Expected number
}

func TestJsonSet_RunGoodExpectedTypeBool(t *testing.T) {
	ji := &JsonSet{}
	ji.ParseFlags([]string{"blah/blah", "true"})
	ji.setExpectedType("bool")
	err := ji.Run()

	if err != nil {
		t.Error(err)
	}
}

func TestJsonSet_RunGoodExpectedTypeString(t *testing.T) {
	ji := &JsonSet{}
	ji.ParseFlags([]string{"blah/str", "\"this is a string\""})
	ji.setExpectedType("string")
	err := ji.Run()

	if err != nil {
		t.Error(err)
	}
}
