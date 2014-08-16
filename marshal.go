package uuid

import (
	"fmt"
)

// MarshalText returns the string-representation of the UUID.
func (u UUID) MarshalText() ([]byte, error) {
	return []byte(u.String()), nil
}

// MarshalJSON returns the string-representation of the UUID as a JSON-string.
func (u UUID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%8.x-%4.x-%4.x-%4.x-%12.x\"", u[0:4], u[4:6], u[6:8], u[8:10], u[10:16])), nil
}

// UnmarshalText reads an UUID from a string into the UUID instance.
// If this fails the state of the UUID is undetermined.
func (u *UUID) UnmarshalText(data []byte) error {
	err := u.SetString(string(data))
	if err != nil {
		return err
	}

	return nil
}

// UnmarshalJSON reads an UUID from a JSON-string into the UUID instance.
// If this fails the state of the UUID is undetermined.
func (u *UUID) UnmarshalJSON(data []byte) error {
	err := u.SetString(string(data[1 : len(data)-1]))
	if err != nil {
		return err
	}

	return nil
}
