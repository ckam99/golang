package utils

import (
	"errors"
)

var ErrNoEntity = errors.New("entity not found")
var ErrUniqueField = errors.New("unique violation")
var ErrInvalidCredentials = errors.New("invalid credentials")
