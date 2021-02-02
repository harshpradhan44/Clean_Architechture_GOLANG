package errorConstant

import (
	"fmt"
)

func New(text string) error {
	return &ErrorString{text}
}

type ErrorString struct {
	s string
}

func (e *ErrorString) Error() string {
	return fmt.Sprintf("id %v not present",e.s)
}

const (
	BrandIdNotFound = "brand id not found"
	IdNotFound = "id not found"
	InvalidId= "invalid id"
	InvalidParameterId = "invalid parameter id"
	IdAlreadyExist = "id already exist"
)