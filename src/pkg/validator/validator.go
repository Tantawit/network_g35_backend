package validator

import (
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/domain/dto"
	vt "github.com/2110336-2565-2/cu-freelance-chat/src/internal/validator"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"reflect"
	"regexp"
	"strings"
)

type Validator interface {
	Validate(interface{}) []*dto.BadReqErrResponse
}

func NewValidator() (Validator, error) {
	translator := en.New()
	uni := ut.New(translator, translator)

	trans, found := uni.GetTranslator("en")
	if !found {
		return nil, errors.New("translator not found")
	}

	v := validator.New()

	if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		return nil, err
	}

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	_ = v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	_ = v.RegisterTranslation("password", trans, func(ut ut.Translator) error {
		return ut.Add("password", "{0} is not strong enough (must be at lease 8 characters)", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("password", fe.Field())
		return t
	})

	_ = v.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		return len(fl.Field().String()) >= 8
	})

	_ = v.RegisterTranslation("uuid_optional", trans, func(ut ut.Translator) error {
		return ut.Add("uuid_optional", "{0} is not uuid", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("uuid_optional", fe.Field())
		return t
	})

	_ = v.RegisterValidation("uuid_optional", func(fl validator.FieldLevel) bool {
		content := fl.Field().String()
		if len(content) > 0 {
			_, err := uuid.Parse(content)
			return err == nil
		}
		return true
	})

	_ = v.RegisterValidation("student_id", func(fl validator.FieldLevel) bool {
		matched, err := regexp.MatchString(`^6[0-5](3|4)[0-9]{5}(20|21|22|23|24|25|26|27|28|29|30|31|32|33|34|35|36|37|38|39|40|51|53|55|56|58|99|01|02)$`, fl.Field().String())
		return matched && err == nil
	})

	_ = v.RegisterTranslation("student_id", trans, func(ut ut.Translator) error {
		return ut.Add("student_id", "{0} is invalid (eg. 6400000021)", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("student_id", fe.Field())
		return t
	})

	_ = v.RegisterValidation("phone_number", func(fl validator.FieldLevel) bool {
		matched, err := regexp.MatchString(`^(\+\d{1,2}\s)?\(?\d{3}\)?[\s.-]?\d{3}[\s.-]?\d{4}$`, fl.Field().String())
		return matched && err == nil
	})

	_ = v.RegisterTranslation("phone_number", trans, func(ut ut.Translator) error {
		return ut.Add("phone_number", "{0} is invalid", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("phone_number", fe.Field())
		return t
	})

	return &vt.ValidatorImpl{
		V:     v,
		Trans: trans,
	}, nil
}
