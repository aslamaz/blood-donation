package constant

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrDuplicateEmail     = errors.New("email alredy exists")
	ErrDuplicateMobile    = errors.New("mobile alredy exists")
	ErrInvalidBloodGroup  = errors.New("Invalid blood group")
)
