package main

import (
	"bytes"
	"encoding/json"
	"errors"
)

// Nullable implements a field that can be null and/or optional, meant for JSON serialization / deserialization.
//
// Internal impl detail:
// - map[true]T means a value was provided
// - map[false]T means an explicit null was provided
// - nil or zero map means the field was not provided
type Nullable[T any] map[bool]T

func (t Nullable[T]) Get() (T, error) {
	var empty T
	if t.IsNull() {
		return empty, errors.New("value is null")
	}
	if !t.IsSpecified() {
		return empty, errors.New("value is not specified")
	}
	return t[true], nil
}

func (t *Nullable[T]) Set(value T) {
	*t = map[bool]T{true: value}
}

func (t Nullable[T]) IsNull() bool {
	_, foundNull := t[false]
	return foundNull
}

func (t *Nullable[T]) SetNull() {
	var empty T
	*t = map[bool]T{false: empty}
}

func (t Nullable[T]) IsSpecified() bool {
	return len(t) != 0
}

func (t *Nullable[T]) SetUnspecified() {
	*t = map[bool]T{}
}

func (t Nullable[T]) MarshalJSON() ([]byte, error) {
	if t.IsNull() {
		return []byte("null"), nil
	}
	return json.Marshal(t[true])
}

func (t *Nullable[T]) UnmarshalJSON(data []byte) error {
	// - case not specified: UnmarshalJSON is never called
	// - case of an explicit null: check for that data explicitly
	if bytes.Equal(data, []byte("null")) {
		t.SetNull()
		return nil
	}
	// - case of parsing an actual value
	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	t.Set(v)
	return nil
}

func main() {

}
