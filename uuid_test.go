package uuid

import (
	"fmt"
	"testing"
)

var (
	testStringUUID = "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"
	testZeroString = "00000000-0000-0000-0000-000000000000"
	testByteUUID   = []byte("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	testMixed      = map[string]validator{
		"00000000-0000-0000-0000-00000000000f": valid{"00000000-0000-0000-0000-00000000000f"},
		"00000000-0000-0000-0000-0000000000f0": valid{"00000000-0000-0000-0000-0000000000f0"},
		"00000000-0000-0000-0000-100000000000": valid{"00000000-0000-0000-0000-100000000000"},
		"00000000-0000-0000-f000-000000000000": valid{"00000000-0000-0000-f000-000000000000"},
		"00000000-0000-f000-0000-000000000000": valid{"00000000-0000-f000-0000-000000000000"},
		"00000000-f000-0000-0000-000000000000": valid{"00000000-f000-0000-0000-000000000000"},
		"f0000000-0000-0000-0000-000000000000": valid{"f0000000-0000-0000-0000-000000000000"},
		"0f000000-0000-0000-0000-000000000000": valid{"0f000000-0000-0000-0000-000000000000"},
		"000f0000-0000-0000-0000-000000000000": valid{"000f0000-0000-0000-0000-000000000000"},
		"00000000-00f0-0000-0000-000000000000": valid{"00000000-00f0-0000-0000-000000000000"},
		"f0000000-f000-f000-f000-f00000000000": valid{"f0000000-f000-f000-f000-f00000000000"},
		"12345678-9abc-deff-edcb-a98765432100": valid{"12345678-9abc-deff-edcb-a98765432100"},
		"ffffffff-ffff-ffff-ffff-ffffffffffff": valid{"ffffffff-ffff-ffff-ffff-ffffffffffff"},
		"12345678-1234-5678-1234-567812345678": valid{"12345678-1234-5678-1234-567812345678"},
		"c56a4180-65aa-42ec-a945-5fd21dec0538": valid{"c56a4180-65aa-42ec-a945-5fd21dec0538"},
		"00000000-0000-0000-0000-000000000000": valid{"00000000-0000-0000-0000-000000000000"},
		"00000000-0000-0000-0000-00000000000":  invalid{},
		"":        invalid{},
		"0":       invalid{},
		"abcdefg": invalid{},
		"fffffff-ffff-ffff-ffff-fffffffffffff":   invalid{}, /* all - shifted one to the left */
		"fffffffff-ffff-ffff-ffff-fffffffffff":   invalid{}, /* all - shifted one to the right */
		"ffffffff-ffff-ffff-ffff-fffffffffff":    invalid{}, /* one char too short */
		"ffffffff-ffff-ffff-ffff-fffffffffffff":  invalid{}, /* one char too long */
		"ffffffff-ffff-ffff-ffff-ffffffffffffff": invalid{}, /* two char too long */
		"ffffffff-ffff-ffff-ffff--fffffffffff":   invalid{}, /* one dash instead of char */
		"C56A4180-65AA-42EC-A945-5FD21DEC":       invalid{},
		"x56a4180-h5aa-42ec-a945-5fd21dec0538":   invalid{},
		"null": invalid{},
		"something invalid here":                  invalid{},
		"A0EEBC99-9C0B-4EF8-BB6D-6BB9BD380A11":    valid{"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"},
		"{a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11}":  valid{"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"},
		"a0eebc999c0b4ef8bb6d6bb9bd380a11":        valid{"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"},
		"a0ee-bc99-9c0b-4ef8-bb6d-6bb9-bd38-0a11": valid{"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"},
		"{a0eebc99-9c0b4ef8-bb6d6bb9-bd380a11}":   valid{"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"},
	}
)

// validator validates errors and UUIDs for tests
type validator interface {
	ReadError(error) error
	Validate(UUID) error
	ValidateEmpty(UUID) error
}

// valid is an expected valid UUID read.
type valid struct {
	Str string
}

// ReadError validates any errors which might arise when reading the UUID.
func (v valid) ReadError(err error) error {
	if err != nil {
		return fmt.Errorf("expected '%s', got error: %s", v.Str, err.Error())
	}

	return nil
}

// Validate validates a UUID with the expectation that it should not validate
// upon read failure.
func (v valid) Validate(u UUID) error {
	if v.Str != u.String() {
		return fmt.Errorf("string representation '%s' does not match '%s'.", u.String(), v.Str)
	}

	return nil
}

// ValidateEmpty validates a UUID where an error condition must set the UUID
// to zero.
func (v valid) ValidateEmpty(u UUID) error {
	return v.Validate(u)
}

// invalid is an expected failure to read a UUID.
type invalid struct{}

// ReadError validates any errors which might arise when reading the UUID.
func (v invalid) ReadError(err error) error {
	if err == nil {
		return fmt.Errorf("expected error")
	}

	return nil
}

// Validate validates a UUID with the expectation that it should not validate
// upon read failure.
func (v invalid) Validate(u UUID) error {
	return nil
}

// ValidateEmpty validates a UUID where an error condition must set the UUID
// to zero.
func (v invalid) ValidateEmpty(u UUID) error {
	if u != [16]byte{} {
		return fmt.Errorf("expected empty UUID, got '%s'", u.String())
	}

	return nil
}

func TestZero(t *testing.T) {
	u := UUID{}

	if !u.IsZero() {
		t.Error("Default value is not zero")
	}
}

func TestZero2(t *testing.T) {
	u := UUID{}

	if u.String() != testZeroString {
		t.Error("Default zero value is not zero uuid")
	}
}

func TestFromString(t *testing.T) {
	for i, v := range testMixed {
		u, err := FromString(i)
		if e := v.ReadError(err); e != nil {
			t.Errorf("FromString(%s): %s", i, e.Error())
		} else if e := v.Validate(u); e != nil {
			t.Errorf("FromString(%s): %s", i, e.Error())
		}
	}
}

func TestFromStringZero(t *testing.T) {
	u, err := FromString(testZeroString)
	if err != nil {
		t.Error("Failed to read zero-uuid")
	}

	if !u.IsZero() {
		t.Error("Zero-uuid from string does not yield zero-uuid")
	}

	if u.String() != testZeroString {
		t.Error("zero uuid is not zero uuid")
	}
}

func TestReadBytes(t *testing.T) {
	for i, v := range testMixed {
		u := UUID{}

		err := u.ReadBytes([]byte(i))
		if e := v.ReadError(err); e != nil {
			t.Errorf("FromString(%s): %s", i, e.Error())
		} else if e := v.Validate(u); e != nil {
			t.Errorf("FromString(%s): %s", i, e.Error())
		}
	}
}

func TestSetString(t *testing.T) {
	for i, v := range testMixed {
		u := UUID{}

		err := u.SetString(i)
		if e := v.ReadError(err); e != nil {
			t.Errorf("SetString(%s): %s", i, e.Error())
		} else if e := v.Validate(u); e != nil {
			t.Errorf("SetString(%s): %s", i, e.Error())
		}
	}
}

func TestMaybeFromString(t *testing.T) {
	for i, v := range testMixed {
		u := MaybeFromString(i)
		if e := v.ValidateEmpty(u); e != nil {
			t.Errorf("MaybeFromString(%s): %s", i, e.Error())
		}
	}
}

func TestMustFromString(t *testing.T) {
	for i, v := range testMixed {
		testMustFromString(t, i, v)
	}
}

func testMustFromString(t *testing.T, i string, v validator) {
	var u UUID

	defer func() {
		r := recover()

		err, ok := r.(error)
		if !ok && r != nil {
			panic(fmt.Sprintf("in test for MustFromString: got non-error: %+v", r))
		}

		if e := v.ReadError(err); e != nil {
			t.Errorf("MustFromString(%s): %s", i, e.Error())
		}
	}()

	u = MustFromString(i)
	if e := v.Validate(u); e != nil {
		t.Errorf("MustFromString(%s): %s", i, e.Error())
	}
}

func TestV4(t *testing.T) {
	u, err := V4()
	if err != nil {
		panic(err)
	}

	if u[6]&0xf0 != 0x40 {
		t.Errorf("UUID generated from V4() does not have the version byte set to 4: '%s'.", u.String())
	}

	if u[8]&0xf0 < 0x80 || u[8]&0xf0 > 0xB0 {
		t.Errorf("UUID generated from V4() does not have the 9th byte beginning with 8, 9, A or B: '%s'.", u.String())
	}
}

func TestSetZero(t *testing.T) {
	u, err := FromString("12345678-9abc-deff-edcb-a98765432100")
	if err != nil {
		panic(err)
	}

	if u == [16]byte{} {
		panic("zero UUID returned from FromString")
	}

	u.SetZero()

	if u != [16]byte{} {
		t.Errorf("SetZero() did not zero UUID")
	}
}

func TestVersion(t *testing.T) {
	list := map[string]int{
		"10a7f7c0-1011-11e5-ad77-0002a5d5c51b": 1,
		"22220da4-d863-3451-98ce-02cc7288bf9a": 3,
		"ebd435d3-63eb-43c6-8e92-342238da6b58": 4,
		"92ce3c5b-5c3e-51e5-8490-2a2334346357": 5,
	}

	for i, v := range list {
		u := MustFromString(i)

		if u.Version() != v {
			t.Errorf("Version(%s) returned %d, expected %d", i, u.Version(), v)
		}
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

func BenchmarkMaybeFromStringOk(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaybeFromString("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	}
}

func BenchmarkMaybeFromStringFail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaybeFromString("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11abcdef")
	}
}
