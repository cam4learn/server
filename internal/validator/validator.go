package validator

import (
	"errors"
	"server/internal/registration"
)

func Validate(object interface{}) (bool, error) {
	switch concreteType := object.(type) {
	case registration.AuthData:
		return validateInputLoginData(concreteType.Login, concreteType.Password), nil
	case registration.LectorData:
		return validateLectorAddForm(concreteType)
	case registration.DeviceData:
		return validateDeviceAddForm(concreteType)
	case registration.LectorDataEdit:
		return validateLectorAddForm(registration.LectorData(concreteType))
	case registration.DeviceDataEdit:
		return validateDeviceEditForm(concreteType)
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
	//upperCaseLetter := false
	//numericChars := false
	//for _, c := range password {
	//	if unicode.IsUpper(c) {
	//		upperCaseLetter = true
	//	}
	//	if unicode.IsDigit(c) {
	//		numericChars = true
	//	}
	//}
	//return upperCaseLetter && numericChars
	return len(password) > 4
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

func validateDeviceAddForm(form registration.DeviceData) (bool, error) {
	if len(form.MACAdress) == 0 {
		return false, errors.New("BAD_MAC")
	}
	if len(form.Room) == 0 {
		return false, errors.New("BAD_ROOM")
	}
	return true, nil
}

func validateDeviceEditForm(form registration.DeviceDataEdit) (bool, error) {
	if len(form.MACAdress) == 0 {
		return false, errors.New("BAD_MAC")
	}
	if len(form.Room) == 0 {
		return false, errors.New("BAD_ROOM")
	}
	return true, nil
}
