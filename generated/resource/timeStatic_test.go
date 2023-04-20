package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-time-schema/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestTimeStaticSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.TimeStaticSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
