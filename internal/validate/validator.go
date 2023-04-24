package validate

import "github.com/go-playground/validator/v10"

var v *validator.Validate

func Setup() {
	v = validator.New()
}

func Struct(p interface{}) error {
	return v.Struct(p)
}
