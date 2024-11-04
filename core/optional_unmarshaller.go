package core

import (
	"encoding/json"
)

// DO NOT REVERT - This is a manual change that is required to unmarshall
// input request json into optional type structs
func (o *Optional[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.Null = true
		return nil
	}

	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	o.Value = v
	o.Null = false
	return nil
}
