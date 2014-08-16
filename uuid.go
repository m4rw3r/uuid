// Package uuid implements a fast UUID representation and integrates with JSON and SQL drivers
package uuid

import (
	"crypto/rand"
	"fmt"
)

// UUID represents a Universally-Unique-Identifier
type UUID [16]byte

var zero = [16]byte{}

type scanError struct {
	scanned int
	bytes   int
	length  int
}

type ErrTooShort scanError

func (e *ErrTooShort) Error() string {
	return fmt.Sprintf("invalid UUID: too few bytes (scanned: %d, length: %d, bytes: %d)", e.scanned, e.length, e.bytes)
}

type ErrTooLong scanError

func (e *ErrTooLong) Error() string {
	return fmt.Sprintf("invalid UUID: too many bytes (scanned: %d, length: %d, bytes: %d)", e.scanned, e.length, e.bytes)
}

type ErrUneven scanError

func (e *ErrUneven) Error() string {
	return fmt.Sprintf("invalid UUID: uneven hexadecimal bytes (scanned: %d, length: %d, bytes: %d)", e.scanned, e.length, e.bytes)
}

// hexchar2byte contains the integer byte-value represented by a hexadecimal character,
// 255 if it is an invalid character
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

// hex2byte reads the first two bytes of the input string and returns the byte matched
// by their hexadecimal value
func hex2byte(x string) (byte, bool) {
	a := hexchar2byte[x[0]]
	b := hexchar2byte[x[1]]

	return (a << 4) | b, a != 255 && b != 255
}

// V4 creates a new random UUID with data from crypto/rand.Read
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

// FromString reads a UUID into a new UUID instance
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

// IsZero returns true if the UUID is zero
func (u UUID) IsZero() bool {
	return u == zero
}

// String returns the string representation of the UUID
func (u UUID) String() string {
	return fmt.Sprintf("%8.x-%4.x-%4.x-%4.x-%12.x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:16])
}
