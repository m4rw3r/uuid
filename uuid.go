/*
Package uuid implements a fast representation of UUIDs (Universally Unique Identifiers) and integrates with JSON and SQL drivers.

This package supports reading of multiple formats of UUIDs, including but not limited to:

	a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11
	A0EEBC99-9C0B-4EF8-BB6D-6BB9BD380A11
	{a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11}
	a0eebc999c0b4ef8bb6d6bb9bd380a11
	a0ee-bc99-9c0b-4ef8-bb6d-6bb9-bd38-0a11
	{a0eebc99-9c0b4ef8-bb6d6bb9-bd380a11}

The parsing-speed of UUIDs in this package is achieved in several ways:

A lookup table converts hexadecimal digits to bytes.

Scanning and parsing is done in place without allocating anything.

Resulting bytes are written to the UUID as it is parsed. On parse errors
this will leave the UUID only partially populated with data from the
input string, leaving the rest of the UUID unmodified.

This package just ignores non-hexadecimal digits when scanning. This can cause
some odd representations of hexadecimal data to be parsed as valid UUIDs, and
longer strings like these will parse successfully:

	a0eebc99,9c0b,4ef8,bb6d,6bb9bd380a11
	a0eebc99This9cIs0b4eOKf8bb6d6bb9bdLOL380a11
	a0-ee-bc-99-9c-0b-4e-f8-bb-6d-6b-b9-bd-38-0a-11

However, the hexadecimal digits MUST come in pairs, and the total number of bytes
represented by them MUST equal 16, or it will generate a parse error.
For example, UUIDs like these will not parse:

	a0eebc999-c0b-4ef8-bb6d-6bb9bd380a11
	a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a
	a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a111
	a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a1111

All string-creating functions will generate UUIDs in the canonical format of:

	a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11

*/
package uuid

import (
	"crypto/rand"
	"fmt"
)

// UUID represents a Universally-Unique-Identifier.
type UUID [16]byte

// zero is the zero-UUID, every single byte set to 0.
var zero = [16]byte{}

// ScanError contains the scanner-state for when the error occurred.
type ScanError struct {
	// Scanned is the number of bytes of the source string which has been considered.
	Scanned int
	// Written is the number of decoded hexadecimal bytes which has
	// been written to the UUID instance.
	Written int
	// Length is the length of the source string.
	Length  int
}

// ErrTooShort occurs when the supplied string does not contain enough
// hexadecimal characters to represent a UUID.
type ErrTooShort ScanError

func (e *ErrTooShort) Error() string {
	return fmt.Sprintf("invalid UUID: too few bytes (scanned characters: %d, written bytes: %d, string length: %d)", e.Scanned, e.Written, e.Length)
}

// ErrTooLong occurs when the supplied string contains more than the
// required number of hexadecimal characters to represent a UUID.
type ErrTooLong ScanError

func (e *ErrTooLong) Error() string {
	return fmt.Sprintf("invalid UUID: too many bytes (scanned characters: %d, written bytes: %d, string length: %d)", e.Scanned, e.Written, e.Length)
}

// ErrUneven occurs when a hexadecimal digit is not part of a pair, making it
// impossible to decode it to a byte.
type ErrUneven ScanError

func (e *ErrUneven) Error() string {
	return fmt.Sprintf("invalid UUID: uneven hexadecimal bytes (scanned characters: %d, written bytes: %d, string length: %d)", e.Scanned, e.Written, e.Length)
}

// hexchar2byte contains the integer byte-value represented by a hexadecimal character,
// 255 if it is an invalid character.
var hexchar2byte = []byte{
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 255, 255, 255, 255, 255, 255,
	255, 10, 11, 12, 13, 14, 15, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 10, 11, 12, 13, 14, 15, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
}

// V4 creates a new random UUID with data from crypto/rand.Read().
func V4() (UUID, error) {
	u := UUID{}

	_, err := rand.Read(u[:])
	if err != nil {
		return u, err
	}

	u[8] = (u[8] | 0x80) & 0xBF
	u[6] = (u[6] | 0x40) & 0x4F

	return u, nil
}

// FromString reads a UUID into a new UUID instance.
func FromString(str string) (UUID, error) {
	u := UUID{}

	err := u.SetString(str)

	return u, err
}

// SetString reads the supplied string-representation of the UUID into the instance.
// On invalid UUID an error is returned and the UUID state will be undetermined.
// This function will ignore all non-hexadecimal digits.
func (u *UUID) SetString(str string) error {
	i := 0
	x := 0
	c := len(str)

	for x < c {
		a := hexchar2byte[str[x]]
		if a == 255 {
			// Invalid char, skip
			x++

			continue
		}

		// We need to perform this check after the attempted hex-read in case
		// we have trailing "}" characters
		if i >= 16 {
			return &ErrTooLong{x, i, c}
		}
		if x+1 >= c {
			// Not enough to scan
			return &ErrTooShort{x, i, c}
		}

		b := hexchar2byte[str[x+1]]
		if b == 255 {
			// Uneven hexadecimal byte
			return &ErrUneven{x, i, c}
		}

		u[i] = (a << 4) | b

		x += 2
		i++
	}

	if i != 16 {
		// Can only be too short here
		return &ErrTooShort{x, i, c}
	}

	return nil
}

// IsZero returns true if the UUID is zero.
func (u UUID) IsZero() bool {
	return u == zero
}

// String returns the string representation of the UUID.
// This method returns the canonical representation of
// ``xxxxxxxx-xxxx-Mxxx-Nxxx-xxxxxxxxxxxx``.
func (u UUID) String() string {
	return fmt.Sprintf("%8.x-%4.x-%4.x-%4.x-%12.x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:16])
}
