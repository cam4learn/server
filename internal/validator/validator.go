package validator

import (
	"errors"
	"server/internal/registration"
	"unicode"
)

func Validate(object interface{}) (bool, error) {
	switch concreteType := object.(type) {
	case registration.AuthData:
		return validateInputLoginData(concreteType.Login, concreteType.Password), nil
	case registration.LectorData:
		return validateLectorAddForm(concreteType)
	default:
		return false, nil
	}
}

func validateInputLoginData(login, password string) bool {
	return validateLogin(login) && validatePassword(password)
}

func validateLogin(login string) bool {
	if len(login) < 4 {
		return false
	}
	return true
}
func validatePassword(password string) bool {
	upperCaseLetter := false
	numericChars := false
	for _, c := range password {
		if unicode.IsUpper(c) {
			upperCaseLetter = true
		}
		if unicode.IsDigit(c) {
			numericChars = true
		}
	}
	return upperCaseLetter && numericChars
}

func validateLectorAddForm(form registration.LectorData) (bool, error) {
	if len(form.Name) == 0 {
		return false, errors.New("BAD_NAME")
	}
	if len(form.Surname) == 0 {
		return false, errors.New("BAD_SURNAME")
	}
	if !validatePassword(form.Password) {
		return false, errors.New("BAD_PASSWORD")
	}
	if !validateLogin(form.Login) {
		return false, errors.New("BAD_LOGIN")
	}
	return true, nil
}
