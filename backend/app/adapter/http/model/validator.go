package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

const (
	MinVarcharLength = 2
	MaxVarcharLength = 255
)

func (req *RequestRegisterMessage) ValidateParam() error {
	var fieldRules []*validation.FieldRules
	fieldRules = append(fieldRules, validation.Field(&req.Message, validation.Required, validation.Length(MinVarcharLength, MaxVarcharLength)))
	return validation.ValidateStruct(req, fieldRules...)
}
