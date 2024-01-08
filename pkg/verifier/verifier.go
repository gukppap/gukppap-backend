package verifier

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/guckppap/gukppap-backend/pkg/apperrors"
)

var _validator *validator.Validate

func init() {
	_validator = validator.New()
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

		marshaledJSON, _ := json.Marshal(messages)
		return apperrors.New(apperrors.InternalServerError, string(marshaledJSON))
	}

	return nil
}
