package generated

import (
	tfjson "github.com/hashicorp/terraform-json"
	
	"github.com/lonegunmanb/terraform-time-schema/generated/resource"
)

var Resources map[string]*tfjson.Schema
var DataSources map[string]*tfjson.Schema

func init() {
	resources := make(map[string]*tfjson.Schema, 0)
	dataSources := make(map[string]*tfjson.Schema, 0)  
	resources["time_offset"] = resource.TimeOffsetSchema()  
	resources["time_rotating"] = resource.TimeRotatingSchema()  
	resources["time_sleep"] = resource.TimeSleepSchema()  
	resources["time_static"] = resource.TimeStaticSchema()  
	Resources = resources
	DataSources = dataSources
}