package auth

type LoginDTO struct {
	Email    string `json:"email" validate:"empty=false & format=email"`
	Password string `json:"password" validate:"empty=false & format=alnum"`
}

type RegisterDTO struct {
	Email    string `json:"email" validate:"empty=false & format=email"`
	FullName string `json:"full_name" validate:"empty=false & gte=2"`
	/*
	 * Password rules:
	 * at least 7 letters
	 * at least 1 number
	 * at least 1 upper case
	 * at least 1 special character
	 */
	Password string `json:"password" validate:"empty=false & format=alpha & gte=2"`
}

// func (r RegisterDTO) Validate() error {
// 	if len(r.Password) < 6 {
// 		return errors.New("password should contain at least 6 letters")
// 	}
// 	//#$(&}?!{;@)*%
// 	//regx := regexp.MustCompile(`^[a-zA-Z]+\[[0-9]+\]$`)
// 	regx := regexp.MustCompile(`^(?:.*[a-z])(?:.*[A-Z])(?:.*\d)[a-zA-Z\d]{8,}$`)

// 	if regx.MatchString(r.Password) {
// 		return errors.New("password should be strong: password should contain at least number,upper case and special character ")
// 	}
// 	return nil
// }

type TokenDTO struct {
	ID           int64  `json:"id"`
	Email        string `json:"email"`
	AccessToken  string
	RefreshToken string
}
