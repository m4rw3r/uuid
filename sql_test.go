package uuid

import (
	"reflect"
	"testing"
)

func TestUUIDScanString(t *testing.T) {
	u := UUID{}

	err := u.Scan("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	if err != nil {
		t.Error("UUID.Scan failed on string \"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11\".")
	}

	if u.String() != "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11" {
		t.Error("UUID.Scan failed to properly read string \"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11\".")
	}
}

func TestUUIDScanByte(t *testing.T) {
	u := UUID{}

	err := u.Scan([]byte("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"))
	if err != nil {
		t.Error("UUID.Scan failed on byte array \"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11\".")
	}

	if u.String() != "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11" {
		t.Error("UUID.Scan failed to properly read byte array \"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11\".")
	}
}

func TestUUIDScanInt(t *testing.T) {
	u := UUID{}

	err := u.Scan(12345)
	if err == nil {
		t.Error("Expected UUID.Scan to fail on integer, did not fail")
		return
	}

	if _, ok := err.(*ErrInvalidType); !ok {
		t.Error("Expected UUID.Scan to fail with ErrInvalidType, failed with "+reflect.TypeOf(err).String())
		return
	}
}

func TestUUIDScanNil(t *testing.T) {
	u := UUID{}

	err := u.Scan(nil)
	if err == nil {
		t.Error("Expected UUID.Scan to fail on nil, did not fail")
		return
	}

	if _, ok := err.(*ErrInvalidType); !ok {
		t.Error("Expected UUID.Scan to fail with ErrInvalidType, failed with "+reflect.TypeOf(err).String())
		return
	}
}

func TestUUIDValue(t *testing.T) {
	u, err := FromString("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	if err != nil {
		panic(err)
	}

	v, err := u.Value()
	if err != nil {
		t.Error("err is set")
	}

	if s, ok := v.(string); ok {
		if s != "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11" {
			t.Error("expected 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', got '"+s+"'.")
		}
	} else {
		t.Error("expected string, got "+reflect.TypeOf(v).String())
	}
}

func TestNullUUIDScanString(t *testing.T) {
	u := NullUUID{}

	err := u.Scan("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	if err != nil {
		t.Error("NullUUID.Scan failed on string \"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11\".")
	}

	if !u.Valid {
		t.Error("NullUUID.Scan failed to set Valid for string \"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11\".")
	}

	if u.UUID.String() != "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11" {
		t.Error("NullUUID.Scan failed to properly read string \"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11\".")
	}
}

func TestNullUUIDScanByte(t *testing.T) {
	u := NullUUID{}

	err := u.Scan([]byte("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"))
	if err != nil {
		t.Error("NullUUID.Scan failed on byte array \"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11\".")
	}

	if !u.Valid {
		t.Error("NullUUID.Scan failed to set Valid for byte array \"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11\".")
	}

	if u.UUID.String() != "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11" {
		t.Error("NullUUID.Scan failed to properly read byte array \"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11\".")
	}
}

func TestNullUUIDScanInt(t *testing.T) {
	u := NullUUID{}

	err := u.Scan(12345)
	if err == nil {
		t.Error("Expected NullUUID.Scan to fail on integer, did not fail")
		return
	}

	if _, ok := err.(*ErrInvalidType); !ok {
		t.Error("Expected NullUUID.Scan to fail with ErrInvalidType, failed with "+reflect.TypeOf(err).String())
		return
	}
}

func TestNullUUIDScanNil(t *testing.T) {
	u := NullUUID{}

	err := u.Scan(nil)
	if err != nil {
		t.Error("Expected UUID.Scan to not fail on nil, did fail")
		return
	}

	if u.Valid {
		t.Error("NullUUID.Scan failed to set Valid to false for nil.")
	}
}

func TestNullUUIDValue(t *testing.T) {
	u, err := FromString("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	if err != nil {
		panic(err)
	}

	nu := NullUUID{
		Valid: true,
		UUID: u,
	}

	v, err := nu.Value()
	if err != nil {
		t.Error("err is set")
	}

	if s, ok := v.(string); ok {
		if s != "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11" {
			t.Error("expected 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', got '"+s+"'.")
		}
	} else {
		t.Error("expected string, got "+reflect.TypeOf(v).String())
	}
}

func TestNullUUIDValue2(t *testing.T) {
	u, err := FromString("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	if err != nil {
		panic(err)
	}

	nu := NullUUID{
		Valid: false,
		UUID: u,
	}

	v, err := nu.Value()
	if err != nil {
		t.Error("err is set")
	}

	if v != nil {
		t.Error("expected nil, got "+reflect.TypeOf(v).String())
	}
}

func BenchmarkUUIDValue(b *testing.B) {
	u, err := FromString("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		u.Value()
	}
}

func BenchmarkUUIDScanString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u := UUID{}

		_ = u.Scan(testStringUUID)
	}
}

func BenchmarkUUIDScanByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u := UUID{}

		_ = u.Scan(testByteUUID)
	}
}

func BenchmarkNullUUIDValue(b *testing.B) {
	u, err := FromString("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	if err != nil {
		panic(err)
	}

	nu := NullUUID{
		Valid: true,
		UUID:  u,
	}

	for i := 0; i < b.N; i++ {
		nu.Value()
	}
}

func BenchmarkNullUUIDValueNil(b *testing.B) {
	u, err := FromString("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	if err != nil {
		panic(err)
	}

	nu := NullUUID{
		Valid: false,
		UUID:  u,
	}

	for i := 0; i < b.N; i++ {
		nu.Value()
	}
}

func BenchmarkNullUUIDScanString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nu := NullUUID{}

		_ = nu.Scan(testStringUUID)
	}
}

func BenchmarkNullUUIDScanByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nu := NullUUID{}

		_ = nu.Scan(testByteUUID)
	}
}

func BenchmarkNullUUIDScanNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nu := NullUUID{}

		_ = nu.Scan(nil)
	}
}
