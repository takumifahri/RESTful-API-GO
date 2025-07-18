package utils

import (
    "errors"
    "fmt"
    "strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidatorStruct(s interface{}) error {
	err := validate.Struct(s)
	if err != nil {
		 var errorMessages []string
        
        for _, err := range err.(validator.ValidationErrors) {
            switch err.Tag() {
            case "required":
                errorMessages = append(errorMessages, fmt.Sprintf("%s is required", err.Field()))
            case "min":
                errorMessages = append(errorMessages, fmt.Sprintf("%s must be at least %s characters/value", err.Field(), err.Param()))
            case "max":
                errorMessages = append(errorMessages, fmt.Sprintf("%s must be at most %s characters/value", err.Field(), err.Param()))
            case "oneof":
                errorMessages = append(errorMessages, fmt.Sprintf("%s must be one of: %s", err.Field(), err.Param()))
            default:
                errorMessages = append(errorMessages, fmt.Sprintf("%s is not valid", err.Field()))
            }
        }
        
        return errors.New(strings.Join(errorMessages, ", "))
    }
	return nil
}