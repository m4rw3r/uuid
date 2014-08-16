package uuid

import (
	"database/sql/driver"
)

// Scan scans a uuid from the given interface instance and stores it
func (u *UUID) Scan(val interface{}) error {
	if s, ok := val.(string); ok {
		err := u.SetString(s)
		if err != nil {
			return err
		}
	}

	return ErrInvalid
}

// Value gives the database driver representation of the UUID
func (u UUID) Value() (driver.Value, error) {
	return u.String(), nil
}
