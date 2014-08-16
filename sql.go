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

// NullUUID represents a UUID that may be null.
// NullUUID implements the Scanner interface so it can be used as a scan destination.
type NullUUID struct {
	// Valid is true if UUID is not NULL
	Valid bool
	UUID  UUID
}

// Scan scans a uuid or null from the given instance and stores i.
// If the supplied value is nil, Valid will be set to false and the
// UUID will be zeroed.
func (nu *NullUUID) Scan(val interface{}) error {
	if val == nil {
		nu.UUID, nu.Valid = [16]byte{}, false
	}

	nu.Valid = true
	return nu.UUID.Scan(val)
}

// Value gives the database driver representation of the UUID or NULL.
func (nu NullUUID) Value() (driver.Value, error) {
	if !nu.Valid {
		return nil, nil
	}

	return nu.UUID.String(), nil
}
