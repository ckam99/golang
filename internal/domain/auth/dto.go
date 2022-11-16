package auth

import (
	"errors"
	"regexp"
)

type LoginDTO struct {
	Email    string `json:"email" validate:"empty=false & format=email"`
	Password string `json:"password" validate:"empty=false"`
}

// RegisterDTO: Register DTO
type RegisterDTO struct {
	Email    string `json:"email" validate:"empty=false & format=email"`
	FullName string `json:"full_name" validate:"empty=false & format=alpha & gte=2"`
	/*
	 * Password rules:
	 * at least 6 letters
	 * at least 1 number
	 * at least 1 upper case
	 * at least 1 special character
	 */
	Password string `json:"password"`
	/*
	* Phone Validation rules
	* format: (`xx`)(`xxxxxxxx`)
	*  `xx` should be one of `01`, `05`, `07`, `21`, `25`, `27`
	*  `xxxxxxxx`  should be 8 digits
	 */
	Phone string `json:"phone"`
}

func (r RegisterDTO) Validate() error {

	// ! Password validation
	//msg := "password should be strong: password should contain at least number,upper case and special character"
	if len(r.Password) < 6 {
		return errors.New("password should contain at least 6 letters")
	}
	if !regexp.MustCompile("[a-z]").MatchString(r.Password) {
		return errors.New("password should contain at least 1 lower case character")
	}
	if !regexp.MustCompile("[A-Z]").MatchString(r.Password) {
		return errors.New("password should contain at least 1 upper case character")
	}
	if !regexp.MustCompile("[0-9]").MatchString(r.Password) {
		return errors.New("password should contain at least 1 number [0-9]")
	}
	if !regexp.MustCompile("[//#$(&}?!{;@)*%]").MatchString(r.Password) {
		return errors.New("password should contain at least 1 special character")
	}

	// !Phone number validation
	if r.Phone != "" {
		patern := `(^01|^05|^07|^21|^25|^27)([0-9]{8})$`
		//`(\+|00)225(01|05|07|21|25|27)([0-9]{8})`
		if !regexp.MustCompile(patern).
			MatchString(r.Phone) {
			return errors.New("invalid phone format")
		}
	}
	return nil
}

type TokenDTO struct {
	ID           int64  `json:"id"`
	Email        string `json:"email"`
	AccessToken  string
	RefreshToken string
}
