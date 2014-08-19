package uuid

import (
	"testing"
)

var (
	testStringUUID = "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"
	testByteUUID   = []byte("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
)

func TestZero(t *testing.T) {
	u := UUID{}

	if !u.IsZero() {
		t.Error("Default value is not zero")
	}
}

func TestZero2(t *testing.T) {
	u := UUID{}

	if u.String() != "00000000-0000-0000-0000-000000000000" {
		t.Error("Default zero value is not zero uuid")
	}
}

func TestFromString(t *testing.T) {
	u, err := FromString("00000000-0000-0000-0000-000000000000")
	if err != nil {
		t.Error("Failed to read zero-uuid")
	}

	if !u.IsZero() {
		t.Error("Zero-uuid from string does not yield zero-uuid")
	}

	if u.String() != "00000000-0000-0000-0000-000000000000" {
		t.Error("zero uuid is not zero uuid")
	}
}

func TestFromStringMultiple(t *testing.T) {
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
		u, err := FromString(i)
		if err != nil {
			t.Error("Failed to read uuid '" + i + "':" + err.Error())

			continue
		}

		if u.String() != i {
			t.Error("String representation of UUID '" + i + "' does not match, got '" + u.String() + "'.")
		}
	}
}

func TestAlternativeValid(t *testing.T) {
	uuid := "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"
	list := []string{
		"A0EEBC99-9C0B-4EF8-BB6D-6BB9BD380A11",
		"{a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11}",
		"a0eebc999c0b4ef8bb6d6bb9bd380a11",
		"a0ee-bc99-9c0b-4ef8-bb6d-6bb9-bd38-0a11",
		"{a0eebc99-9c0b4ef8-bb6d6bb9-bd380a11}",
	}

	for _, i := range list {
		u, err := FromString(i)
		if err != nil {
			t.Error("Failed to read uuid '" + i + "':" + err.Error())

			continue
		}

		if u.String() != uuid {
			t.Error("String representation of UUID '" + i + "' does not match, got '" + u.String() + "'.")
		}
	}
}

func TestInvalid(t *testing.T) {
	list := []string{
		"",
		"0",
		"abcdefg",
		"fffffff-ffff-ffff-ffff-fffffffffffff",   /* all - shifted one to the left */
		"fffffffff-ffff-ffff-ffff-fffffffffff",   /* all - shifted one to the right */
		"ffffffff-ffff-ffff-ffff-fffffffffff",    /* one char too short */
		"ffffffff-ffff-ffff-ffff-fffffffffffff",  /* one char too long */
		"ffffffff-ffff-ffff-ffff-ffffffffffffff", /* two char too long */
		"ffffffff-ffff-ffff-ffff--fffffffffff",   /* one dash instead of char */
	}

	for _, i := range list {
		_, err := FromString(i)
		if err == nil {
			t.Error("Expected error but got nothing when reading uuid '" + i + "'.")
		}
	}
}

func TestV4(t *testing.T) {
	u, err := V4()
	if err != nil {
		panic(err)
	}

	if u[6]&0xf0 != 0x40 {
		t.Error("UUID generated from V4() does not have the version byte set to 4: '" + u.String() + "'.")
	}

	if u[8]&0xf0 < 0x80 || u[8]&0xf0 > 0xB0 {
		t.Error("UUID generated from V4() does not have the 9th byte beginning with 8, 9, A or B: '" + u.String() + "'.")
	}
}

func BenchmarkFromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FromString("12345678-9abc-deff-edcb-a98765432100")
	}
}

func BenchmarkFromString2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FromString("123456789abcdeffedcba98765432100")
	}
}

func BenchmarkSetString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u := UUID{}

		u.SetString(testStringUUID)
	}
}

func BenchmarkReadBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u := UUID{}

		u.ReadBytes(testByteUUID)
	}
}

func BenchmarkString(b *testing.B) {
	u, err := FromString("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		u.String()
	}
}

func BenchmarkV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		V4()
	}
}
