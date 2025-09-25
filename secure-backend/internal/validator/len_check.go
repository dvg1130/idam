package validator

import (
	"errors"
)

// length check
func AuthLenCheck(username string, password string) error {
	if len(username) < 8 || len(password) < 8 {
		return errors.New("username and password must be at least 8 characters")
	}
	return nil
}
