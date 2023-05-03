package validator

import (
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/domain/dto"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type ValidatorImpl struct {
	V     *validator.Validate
	Trans ut.Translator
}

func (v *ValidatorImpl) Validate(in interface{}) []*dto.BadReqErrResponse {
	err := v.V.Struct(in)

	var errs []*dto.BadReqErrResponse
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			element := dto.BadReqErrResponse{
				Message:     e.Translate(v.Trans),
				FailedField: e.Field(),
				Value:       e.Value(),
			}

			errs = append(errs, &element)
		}
	}
	return errs
}

func (v *ValidatorImpl) GetValidator() *validator.Validate {
	return v.V
}
