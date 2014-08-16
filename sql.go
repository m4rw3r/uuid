package uuid

import (
	"database/sql/driver"
	"fmt"
	"reflect"
)

type ErrInvalidType struct {
	Type reflect.Type
}

func (e *ErrInvalidType) Error() string {
	return fmt.Sprintf("uuid Scan(): invalid type '%s', expected string.", e.Type.String())
}

// Scan scans a uuid from the given interface instance and stores it
func (u *UUID) Scan(val interface{}) error {
	if s, ok := val.(string); ok {
		err := u.SetString(s)
		if err != nil {
			return err
		}
	}

	return &ErrInvalidType{reflect.TypeOf(val)}
}

// Value gives the database driver representation of the UUID
func (u UUID) Value() (driver.Value, error) {
	return u.String(), nil
}
