package core

import (
	"encoding/json"
	"fmt"
)

// Optional is a wrapper used to distinguish zero values from
// null or omitted fields.
//
// To instantiate an Optional, use the `Optional()` and `Null()`
// helpers exported from the root package.
type Optional[T any] struct {
	Value T
	Null  bool
}

func (o *Optional[T]) String() string {
	if o == nil {
		return ""
	}
	if s, ok := any(o.Value).(fmt.Stringer); ok {
		return s.String()
	}
	return fmt.Sprintf("%#v", o.Value)
}

func (o *Optional[T]) MarshalJSON() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	if o.Null {
		return []byte("null"), nil
	}
	return json.Marshal(&o.Value)
}

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
