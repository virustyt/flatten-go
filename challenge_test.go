package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

var mapTests = []struct {
	name  string
	input map[string]any
	want  map[string]any
}{
	{
		"should flatten a map",
		map[string]any{"a": map[string]any{"b": map[string]any{"c": 1, "f": 2}, "g": 3}, "d": 2},
		map[string]any{"a.b.c": 1, "a.b.f": 2, "a.g": 3, "d": 2},
	},
}

func TestFlattenMap(t *testing.T) {
	for _, test := range mapTests {
		t.Run(test.name, func(t *testing.T) {
			f := Flattened{}

			output := f.FlattenMap(test.input)

			if !reflect.DeepEqual(output, test.want) {
				t.Errorf("NotEqual %v, %v", output, test.want)
			}
		})
	}
}

var jsonTests = []struct {
	name  string
	input json.RawMessage
	want  json.RawMessage
}{
	{
		"should flatten a JSON object",
		[]byte(`{"a":{"b":{"c":1,"f":2},"g":3},"d":2}`),
		[]byte(`{"a.b.c":1,"a.b.f":2,"a.g":3,"d":2}`),
	},
}

func TestFlattenJSON(t *testing.T) {
	for _, test := range jsonTests {
		t.Run(test.name, func(t *testing.T) {
			f := Flattened{}

			output, err := f.FlattenJSON(test.input)
			if err != nil {
				panic(err)
			}

			if !reflect.DeepEqual(output, test.want) {
				t.Errorf("NotEqual %v, %v", string(output), string(test.want))
			}
		})
	}
}
