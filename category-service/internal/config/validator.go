package config

import (
	"regexp"

	"github.com/CROWNIX/go-utils/validatorx"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// RegisterCustomValidations mendaftarkan semua validasi custom project
func LoadCustomValidations() (err error) {
	v := validatorx.Validate
	t := validatorx.TranslatorID

	err = RegisterPasswordValidation(v, t)

	if err != nil {
		return err
	}

	return nil
}

func RegisterPasswordValidation(v *validator.Validate, t ut.Translator) error {
	v.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()
		if len(password) < 8 {
			return false
		}
		hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
		hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
		hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
		hasSpecial := regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]\{\};':"\\|,.<>\/?]`).MatchString(password)
		return hasUpper && hasLower && hasNumber && hasSpecial
	})

	return v.RegisterTranslation("password", t,
		func(ut ut.Translator) error {
			return ut.Add("password",
				"Password harus minimal 8 karakter, mengandung huruf besar, huruf kecil, angka, dan karakter spesial.",
				true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("password", fe.Field())
			return t
		},
	)
}
