package validate

import (
	refl "reflect"
	"strconv"

	"github.com/pip-services3-go/pip-services3-commons-go/convert"
	"github.com/pip-services3-go/pip-services3-commons-go/reflect"
)

/*
Schema to validate arrays.

Example:
schema := NewArraySchema(TypeCode.String);

schema.Validate(["A", "B", "C"]);    // Result: no errors
schema.Validate([1, 2, 3]);          // Result: element type mismatch
schema.Validate("A");                // Result: type mismatch
*/
type ArraySchema struct {
	Schema
	valueType interface{}
}

// Creates a new instance of validation schema and sets its values.
// see
// TypeCode
// Parameters:
// 			 - valueType interface{}
// 			 a type of array elements. Null means that elements may have any type.

// Returns *ArraySchema
func NewArraySchema(valueType interface{}) *ArraySchema {
	c := &ArraySchema{
		valueType: valueType,
	}
	c.Schema = *InheritSchema(c)
	return c
}

// Gets the type of array elements. Null means that elements may have any type.
// Returns interface{}
// the type of array elements.
func (c *ArraySchema) ValueType() interface{} {
	return c.valueType
}

// Sets the type of array elements. Null means that elements may have any type.
// Parameters
// value interface{}
// a type of array elements.
func (c *ArraySchema) SetValueType(value interface{}) {
	c.valueType = value
}

// Validates a given value against the schema and configured validation rules.
// Parameters:
// 			- path string
// 			a dot notation path to the value.
//			- value interface{}
// 			a value to be validated.
// Return []*ValidationResult
// a list with validation results to add new results.
func (c *ArraySchema) PerformValidation(path string, value interface{}) []*ValidationResult {
	name := path
	if name == "" {
		name = "value"
	}
	value = reflect.ObjectReader.GetValue(value)

	results := c.Schema.PerformValidation(path, value)
	if results == nil {
		results = []*ValidationResult{}
	}

	if value == nil {
		return results
	}

	val := refl.ValueOf(value)
	if val.Kind() == refl.Ptr {
		val = val.Elem()
	}

	if val.Kind() == refl.Slice || val.Kind() == refl.Array {
		for index := 0; index < val.Len(); index++ {
			elementPath := strconv.Itoa(index)
			if path != "" {
				elementPath = path + "." + elementPath
			}
			elemResults := c.PerformTypeValidation(elementPath, c.valueType, val.Index(index).Interface())
			if elemResults != nil {
				results = append(results, elemResults...)
			}
		}
	} else {
		results = append(results,
			NewValidationResult(
				path,
				Error,
				"VALUE_ISNOT_ARRAY",
				name+" type must to be List or Array",
				convert.Array,
				convert.TypeConverter.ToTypeCode(value),
			),
		)
	}

	return results
}
