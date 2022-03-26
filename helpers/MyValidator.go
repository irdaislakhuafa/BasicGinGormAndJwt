package helpers

import (
	"log"

	"github.com/go-playground/validator/v10"
)

var myValidator *validator.Validate

func EnableValidator(status bool) {
	if status {
		myValidator = validator.New()
	} else {
		myValidator = nil
	}
}

func ValidateStruct(structType interface{}) ([]string, error) {

	myValidator := validator.New()
	err := myValidator.Struct(structType)

	fieldsError := make([]string, 0)
	if err != nil {
		for _, value := range err.(validator.ValidationErrors) {
			fieldsError = append(fieldsError, value.Error())
			log.Println("=>", value)
		}
		return fieldsError, err
	}
	return nil, nil

}
func GetValidator() *validator.Validate {
	return myValidator
}
