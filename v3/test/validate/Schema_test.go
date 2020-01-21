package test_validate

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-commons-go/v3/validate"
	"github.com/stretchr/testify/assert"
)

func TestEmptySchema(t *testing.T) {
	schema := validate.NewSchema()
	results := schema.Validate(nil)
	assert.Equal(t, 0, len(results))
}

func TestSchemaRequired(t *testing.T) {
	schema := validate.NewSchema().MakeRequired()
	results := schema.Validate(nil)
	assert.Equal(t, 1, len(results))
}
