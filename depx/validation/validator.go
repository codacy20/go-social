// validator/validator.go
package validator

import validation "github.com/go-ozzo/ozzo-validation/v4"

// ValidateModel validates any model using the provided field rules.
// The model parameter is of type interface{} so it can accept any struct.
func ValidateModel(model interface{}, rules ...*validation.FieldRules) error {
	return validation.ValidateStruct(model, rules...)
}
