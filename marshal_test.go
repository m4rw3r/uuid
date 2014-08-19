package uuid

import (
	"fmt"
	"bytes"
	"testing"
	"reflect"
)

func TestUUIDMarshalText(t *testing.T) {
	u, err := FromString("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	if err != nil {
		panic(err)
	}

	a, b := u.MarshalText()

	if b != nil {
		t.Error("expected UUID.MarshalText() to have err == nil, got '"+reflect.TypeOf(err).String()+"'.")
	}

	if bytes.Compare(a, []byte("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")) != 0 {
		t.Error(fmt.Sprintf("expected UUID.MarshalText() to return '%x', got '%x'.", []byte("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"), a))
	}
}

func TestUUIDMarshalJSON(t *testing.T) {
	u, err := FromString("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	if err != nil {
		panic(err)
	}

	a, b := u.MarshalJSON()

	if b != nil {
		t.Error("expected UUID.MarshalJSON() to have err == nil, got '"+reflect.TypeOf(err).String()+"'.")
	}

	if bytes.Compare(a, []byte("\"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11\"")) != 0 {
		t.Error(fmt.Sprintf("expected UUID.MarshalText() to return '%x', got '%x'.", []byte("\"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11\""), a))
	}
}

func TestUUIDUnmarshalText(t *testing.T) {
	list := []string{
		"00000000-0000-0000-0000-00000000000f",
		"00000000-0000-0000-0000-0000000000f0",
		"00000000-0000-0000-0000-100000000000",
		"00000000-0000-0000-f000-000000000000",
		"00000000-0000-f000-0000-000000000000",
		"00000000-f000-0000-0000-000000000000",
		"f0000000-0000-0000-0000-000000000000",
		"0f000000-0000-0000-0000-000000000000",
		"000f0000-0000-0000-0000-000000000000",
		"00000000-00f0-0000-0000-000000000000",
		"f0000000-f000-f000-f000-f00000000000",
		"12345678-9abc-deff-edcb-a98765432100",
		"ffffffff-ffff-ffff-ffff-ffffffffffff",
	}

	for _, i := range list {
		u := UUID{}

		err := u.UnmarshalText([]byte(i))
		if err != nil {
			t.Error(fmt.Sprintf("Failed to read '%s': %s", i, err.Error()))
		}

		if u.String() != i {
			t.Error("String representation of UUID '" + i + "' does not match, got '" + u.String() + "'.")
		}
	}
}

func TestUUIDUnmarshalJSON(t *testing.T) {
	list := []string{
		"00000000-0000-0000-0000-00000000000f",
		"00000000-0000-0000-0000-0000000000f0",
		"00000000-0000-0000-0000-100000000000",
		"00000000-0000-0000-f000-000000000000",
		"00000000-0000-f000-0000-000000000000",
		"00000000-f000-0000-0000-000000000000",
		"f0000000-0000-0000-0000-000000000000",
		"0f000000-0000-0000-0000-000000000000",
		"000f0000-0000-0000-0000-000000000000",
		"00000000-00f0-0000-0000-000000000000",
		"f0000000-f000-f000-f000-f00000000000",
		"12345678-9abc-deff-edcb-a98765432100",
		"ffffffff-ffff-ffff-ffff-ffffffffffff",
	}

	for _, i := range list {
		u := UUID{}

		err := u.UnmarshalJSON([]byte("\""+i+"\""))
		if err != nil {
			t.Error(fmt.Sprintf("Failed to read '%s': %s", i, err.Error()))
		}

		if u.String() != i {
			t.Error("String representation of UUID '" + i + "' does not match, got '" + u.String() + "'.")
		}
	}
}
