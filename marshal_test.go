package uuid

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

var (
	testByteJSONUUID = []byte("\"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11\"")
	testByteNull     = []byte("null")
)

func TestUUIDMarshalText(t *testing.T) {
	u, err := FromString("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	if err != nil {
		panic(err)
	}

	a, b := u.MarshalText()

	if b != nil {
		t.Error("expected UUID.MarshalText() to have err == nil, got '" + reflect.TypeOf(b).String() + "'.")
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
		t.Error("expected UUID.MarshalJSON() to have err == nil, got '" + reflect.TypeOf(b).String() + "'.")
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

		err := u.UnmarshalJSON([]byte("\"" + i + "\""))
		if err != nil {
			t.Error(fmt.Sprintf("Failed to read '%s': %s", i, err.Error()))
		}

		if u.String() != i {
			t.Error("String representation of UUID '" + i + "' does not match, got '" + u.String() + "'.")
		}
	}
}

func BenchmarkUnmarshalText(b *testing.B) {
	u := UUID{}

	for i := 0; i < b.N; i++ {
		_ = u.UnmarshalText(testByteUUID)
	}
}

func BenchmarkUnmarshalJSON(b *testing.B) {
	u := UUID{}

	for i := 0; i < b.N; i++ {
		_ = u.UnmarshalText(testByteJSONUUID)
	}
}

func BenchmarkMarshalText(b *testing.B) {
	u, err := FromString("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		u.MarshalText()
	}
}

func BenchmarkMarshalJSON(b *testing.B) {
	u, err := FromString("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		u.MarshalJSON()
	}
}

func TestNullUUIDMarshalText1(t *testing.T) {
	u, err := FromString("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	if err != nil {
		panic(err)
	}

	n := NullUUID{Valid: true, UUID: u}

	a, b := n.MarshalText()

	if b != nil {
		t.Error("expected NullUUID.MarshalText() to have err == nil, got '" + reflect.TypeOf(b).String() + "'.")
	}

	if bytes.Compare(a, []byte("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")) != 0 {
		t.Error(fmt.Sprintf("expected NullUUID.MarshalText() to return '%x', got '%x'.", []byte("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"), a))
	}
}

func TestNullUUIDMarshalText2(t *testing.T) {
	n := NullUUID{Valid: false}

	a, b := n.MarshalText()

	if b != nil {
		t.Error("expected NullUUID.MarshalText() to have err == nil, got '" + reflect.TypeOf(b).String() + "'.")
	}

	if bytes.Compare(a, []byte("null")) != 0 {
		t.Error(fmt.Sprintf("expected NullUUID.MarshalText() to return '%x', got '%x'.", []byte("null"), a))
	}
}

func TestNullUUIDMarshalJSON1(t *testing.T) {
	u, err := FromString("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	if err != nil {
		panic(err)
	}

	n := NullUUID{Valid: true, UUID: u}

	a, b := n.MarshalJSON()

	if b != nil {
		t.Error("expected NullUUID.MarshalJSON() to have err == nil, got '" + reflect.TypeOf(b).String() + "'.")
	}

	if bytes.Compare(a, []byte("\"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11\"")) != 0 {
		t.Error(fmt.Sprintf("expected NullUUID.MarshalText() to return '%x', got '%x'.", []byte("\"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11\""), a))
	}
}

func TestNullUUIDMarshalJSON2(t *testing.T) {
	n := NullUUID{Valid: false}

	a, b := n.MarshalJSON()

	if b != nil {
		t.Error("expected NullUUID.MarshalJSON() to have err == nil, got '" + reflect.TypeOf(b).String() + "'.")
	}

	if bytes.Compare(a, []byte("null")) != 0 {
		t.Error(fmt.Sprintf("expected NullUUID.MarshalText() to return '%x', got '%x'.", []byte("null"), a))
	}
}

func TestNullUUIDUnmarshalText1(t *testing.T) {
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
		u := NullUUID{}

		err := u.UnmarshalText([]byte(i))
		if err != nil {
			t.Error(fmt.Sprintf("Failed to read '%s': %s", i, err.Error()))
		}

		if u.Valid != true {
			t.Error(fmt.Sprintf("Parsed UUID was not marked as valid despite not raising error, on '%s'.", i))
		}

		if u.UUID.String() != i {
			t.Error("String representation of UUID '" + i + "' does not match, got '" + u.UUID.String() + "'.")
		}
	}
}

func TestNullUUIDUnmarshalText2(t *testing.T) {
	u, err := FromString("12345678-9abc-deff-edcb-a98765432100")
	if err != nil {
		panic(err)
	}

	n := NullUUID{Valid: true, UUID: u}

	err = n.UnmarshalText([]byte("null"))
	if err != nil {
		t.Error(fmt.Sprintf("Failed to read 'null': %s", err.Error()))
	}

	if n.Valid != false {
		t.Error("Parsing 'null' did not set Valid to false.")
	}

	if n.UUID.String() != "12345678-9abc-deff-edcb-a98765432100" {
		t.Error("UUID modified when parsing 'null'")
	}
}

func TestNullUUIDUnmarshalJSON1(t *testing.T) {
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
		u := NullUUID{}

		err := u.UnmarshalJSON([]byte("\"" + i + "\""))
		if err != nil {
			t.Error(fmt.Sprintf("Failed to read '%s': %s", i, err.Error()))
		}

		if u.Valid != true {
			t.Error(fmt.Sprintf("Parsed UUID was not marked as valid despite not raising error, on '%s'.", i))
		}

		if u.UUID.String() != i {
			t.Error("String representation of UUID '" + i + "' does not match, got '" + u.UUID.String() + "'.")
		}
	}
}

func TestInvalidStringForUUID(t *testing.T) {
	u := NullUUID{}
	err := u.UnmarshalJSON([]byte("0"))
	if err == nil {
		if u.Valid {
			t.Error("We should have returned an Invalid UUID")
		}
	}
}

func TestNullUUIDUnmarshalJSON2(t *testing.T) {
	u, err := FromString("12345678-9abc-deff-edcb-a98765432100")
	if err != nil {
		panic(err)
	}

	n := NullUUID{Valid: true, UUID: u}

	err = n.UnmarshalJSON([]byte("null"))
	if err != nil {
		t.Error(fmt.Sprintf("Failed to read 'null': %s", err.Error()))
	}

	if n.Valid != false {
		t.Error("Parsing 'null' did not set Valid to false.")
	}

	if n.UUID.String() != "12345678-9abc-deff-edcb-a98765432100" {
		t.Error("UUID modified when parsing 'null'")
	}
}

func BenchmarkNullUnmarshalText1(b *testing.B) {
	n := NullUUID{UUID: UUID{}}

	for i := 0; i < b.N; i++ {
		_ = n.UnmarshalText(testByteUUID)
	}
}

func BenchmarkNullUnmarshalText2(b *testing.B) {
	n := NullUUID{UUID: UUID{}}

	for i := 0; i < b.N; i++ {
		_ = n.UnmarshalText(testByteNull)
	}
}

func BenchmarkNullUnmarshalJSON1(b *testing.B) {
	n := NullUUID{UUID: UUID{}}

	for i := 0; i < b.N; i++ {
		_ = n.UnmarshalText(testByteJSONUUID)
	}
}

func BenchmarkNullUnmarshalJSON2(b *testing.B) {
	n := NullUUID{}

	for i := 0; i < b.N; i++ {
		_ = n.UnmarshalText(testByteNull)
	}
}

func BenchmarkNullMarshalText1(b *testing.B) {
	u, err := FromString("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	if err != nil {
		panic(err)
	}

	n := NullUUID{Valid: true, UUID: u}

	for i := 0; i < b.N; i++ {
		n.MarshalText()
	}
}

func BenchmarkNullMarshalText2(b *testing.B) {
	n := NullUUID{Valid: false}

	for i := 0; i < b.N; i++ {
		n.MarshalText()
	}
}

func BenchmarkNullMarshalJSON1(b *testing.B) {
	u, err := FromString("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	if err != nil {
		panic(err)
	}

	n := NullUUID{Valid: true, UUID: u}

	for i := 0; i < b.N; i++ {
		n.MarshalJSON()
	}
}

func BenchmarkNullMarshalJSON2(b *testing.B) {
	n := NullUUID{Valid: false}

	for i := 0; i < b.N; i++ {
		n.MarshalJSON()
	}
}
