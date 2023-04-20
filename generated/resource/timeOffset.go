package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const timeOffset = `{
  "block": {
    "attributes": {
      "base_rfc3339": {
        "computed": true,
        "description": "Base timestamp in [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.8) format (see [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) e.g., ` + "`" + `YYYY-MM-DDTHH:MM:SSZ` + "`" + `). Defaults to the current time.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "day": {
        "computed": true,
        "description": "Number day of offset timestamp.",
        "description_kind": "plain",
        "type": "number"
      },
      "hour": {
        "computed": true,
        "description": "Number hour of offset timestamp.",
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
        "description": "Number minute of offset timestamp.",
        "description_kind": "plain",
        "type": "number"
      },
      "month": {
        "computed": true,
        "description": "Number month of offset timestamp.",
        "description_kind": "plain",
        "type": "number"
      },
      "offset_days": {
        "description": "Number of days to offset the base timestamp. At least one of the 'offset_' arguments must be configured.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "offset_hours": {
        "description": " Number of hours to offset the base timestamp. At least one of the 'offset_' arguments must be configured.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "offset_minutes": {
        "description": "Number of minutes to offset the base timestamp. At least one of the 'offset_' arguments must be configured.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "offset_months": {
        "description": "Number of months to offset the base timestamp. At least one of the 'offset_' arguments must be configured.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "offset_seconds": {
        "description": "Number of seconds to offset the base timestamp. At least one of the 'offset_' arguments must be configured.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "offset_years": {
        "description": "Number of years to offset the base timestamp. At least one of the 'offset_' arguments must be configured.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "rfc3339": {
        "computed": true,
        "description": "RFC3339 format of the offset timestamp, e.g. ` + "`" + `2020-02-12T06:36:13Z` + "`" + `.",
        "description_kind": "plain",
        "type": "string"
      },
      "second": {
        "computed": true,
        "description": "Number second of offset timestamp.",
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
        "description": "Number year of offset timestamp.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Manages an offset time resource, which keeps an UTC timestamp stored in the Terraform state that is offset from a locally sourced base timestamp. This prevents perpetual differences caused by using the [` + "`" + `timestamp()` + "`" + ` function](https://www.terraform.io/docs/configuration/functions/timestamp.html).",
    "description_kind": "plain"
  },
  "version": 0
}`

func TimeOffsetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(timeOffset), &result)
	return &result
}
