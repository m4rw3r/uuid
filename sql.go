package uuid

import (
	"database/sql/driver"
	"fmt"
	"reflect"
)

// ErrInvalidType occurs when *UUID.Scan() does not receive a string.
type ErrInvalidType struct {
	Type reflect.Type
}

func (e *ErrInvalidType) Error() string {
	return fmt.Sprintf("uuid Scan(): invalid type '%s', expected string.", e.Type.String())
}

// Scan scans a uuid from the given interface instance and stores it.
// If scanning fails the state of the UUID is undetermined.
func (u *UUID) Scan(val interface{}) error {
	if s, ok := val.(string); ok {
		return u.SetString(s)
	}

	return &ErrInvalidType{reflect.TypeOf(val)}
}

// Value gives the database driver representation of the UUID.
func (u UUID) Value() (driver.Value, error) {
	return u.String(), nil
}
