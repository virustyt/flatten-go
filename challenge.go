package main

import (
	"errors"
	"log"
)

type Flattened struct {
	res map[string]any
}

func (f *Flattened) FlattenMap(m map[string]any) map[string]any {
	f.res = map[string]any{}
	for k, v := range m {
		err := f.flattenKey(k, v)
		if err != nil {
			log.Fatal(err)
		}
	}
	return f.res
}

func (f *Flattened) flattenKey(k string, m any) error {
	s, ok := m.(int)
	if ok {
		f.res[k] = s
		return nil
	}
	mm, ok := m.(map[string]any)
	if !ok {
		return errors.New("not a map")
	}
	for kk, vv := range mm {
		f.flattenKey(k+"."+kk, vv)
	}
	return nil
}

// Write your code here
func main() {

}
