package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const timeSleep = `{
  "block": {
    "attributes": {
      "create_duration": {
        "description": "[Time duration](https://golang.org/pkg/time/#ParseDuration) to delay resource creation. For example, ` + "`" + `30s` + "`" + ` for 30 seconds or ` + "`" + `5m` + "`" + ` for 5 minutes. Updating this value by itself will not trigger a delay.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destroy_duration": {
        "description": "[Time duration](https://golang.org/pkg/time/#ParseDuration) to delay resource destroy. For example, ` + "`" + `30s` + "`" + ` for 30 seconds or ` + "`" + `5m` + "`" + ` for 5 minutes. Updating this value by itself will not trigger a delay. This value or any updates to it must be successfully applied into the Terraform state before destroying this resource to take effect.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "RFC3339 format of the offset timestamp, e.g. ` + "`" + `2020-02-12T06:36:13Z` + "`" + `.",
        "description_kind": "plain",
        "type": "string"
      },
      "triggers": {
        "description": "(Optional) Arbitrary map of values that, when changed, will run any creation or destroy delays again. See [the main provider documentation](../index.md) for more information.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Manages a resource that delays creation and/or destruction, typically for further resources. This prevents cross-platform compatibility and destroy-time issues with using the [` + "`" + `local-exec` + "`" + ` provisioner](https://www.terraform.io/docs/provisioners/local-exec.html).",
    "description_kind": "plain"
  },
  "version": 0
}`

func TimeSleepSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(timeSleep), &result)
	return &result
}
