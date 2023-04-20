package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const timeStatic = `{
  "block": {
    "attributes": {
      "day": {
        "computed": true,
        "description": "Number day of timestamp.",
        "description_kind": "plain",
        "type": "number"
      },
      "hour": {
        "computed": true,
        "description": "Number hour of timestamp.",
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description": "RFC3339 format of the offset timestamp, e.g. ` + "`" + `2020-02-12T06:36:13Z` + "`" + `.",
        "description_kind": "plain",
        "type": "string"
      },
      "minute": {
        "computed": true,
        "description": "Number minute of timestamp.",
        "description_kind": "plain",
        "type": "number"
      },
      "month": {
        "computed": true,
        "description": "Number month of timestamp.",
        "description_kind": "plain",
        "type": "number"
      },
      "rfc3339": {
        "computed": true,
        "description": "Base timestamp in [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.8) format (see [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) e.g., ` + "`" + `YYYY-MM-DDTHH:MM:SSZ` + "`" + `). Defaults to the current time.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "second": {
        "computed": true,
        "description": "Number second of timestamp.",
        "description_kind": "plain",
        "type": "number"
      },
      "triggers": {
        "description": "Arbitrary map of values that, when changed, will trigger a new base timestamp value to be saved. See [the main provider documentation](../index.md) for more information.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "unix": {
        "computed": true,
        "description": "Number of seconds since epoch time, e.g. ` + "`" + `1581489373` + "`" + `.",
        "description_kind": "plain",
        "type": "number"
      },
      "year": {
        "computed": true,
        "description": "Number year of timestamp.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Manages a static time resource, which keeps a locally sourced UTC timestamp stored in the Terraform state. This prevents perpetual differences caused by using the [` + "`" + `timestamp()` + "`" + ` function](https://www.terraform.io/docs/configuration/functions/timestamp.html).",
    "description_kind": "plain"
  },
  "version": 0
}`

func TimeStaticSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(timeStatic), &result)
	return &result
}
