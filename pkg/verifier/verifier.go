package verifier

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/guckppap/gukppap-backend/pkg/apperrors"
)

var _validator *validator.Validate

func init() {
	_validator = validator.New()
	_validator.RegisterValidation("custom_shortcut", func(fl validator.FieldLevel) bool {
		shortcut := fl.Field().String()
		// 7자 이어야 유효하다.
		return len(shortcut) == 7
	})
}

func Run(model interface{}) error {
	err := _validator.Struct(model)
	if err != nil {
		var messages []map[string]interface{}
		for _, err := range err.(validator.ValidationErrors) {
			messages = append(messages, map[string]interface{}{
				"field":   err.Field(),
				"message": "this field is " + err.Tag(),
			})
		}

		jsonMessage, _ := json.Marshal(messages)
		return apperrors.New(apperrors.InternalServerError, string(jsonMessage))
	}

	return nil
}
