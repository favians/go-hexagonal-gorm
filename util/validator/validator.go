package validator

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var lock = &sync.Mutex{}
var validate *validator.Validate

//GetValidator Initiatilize validator in singleton way
func GetValidator() *validator.Validate {
	lock.Lock()
	defer lock.Unlock()

	if validate == nil {
		validate = validator.New()
	}

	return validate
}
