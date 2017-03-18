package abstraction

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// JSON is a json abstraction
type JSON map[string]interface{}

// ToStringMap returns a string map with the contents of the json
func (j JSON) ToStringMap() map[string]string {
	m := make(map[string]string)
	for k, v := range j {
		switch v.(type) {
		case string:
			m[k] = v.(string)
		case int:
			m[k] = fmt.Sprintf("%d", v.(int))
		}
	}
	return m
}

// Value get value of JSON
func (j JSON) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan scan value into JSON
func (j *JSON) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion .([]byte) failed")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*j, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("type assertion .(map[string]interface{}) failed")
	}

	return nil
}

// NewJSONFromMap creates a new JSON given a string->string map
func NewJSONFromMap(m map[string]string) JSON {
	j := make(JSON)
	for k, v := range m {
		i, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			j[k] = v
			continue
		}
		j[k] = int(i)
	}
	return j
}
