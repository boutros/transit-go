package transit

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
)

type Decoder struct {
	source *bufio.Reader
}

func NewDecoder(source io.Reader) *Decoder {
	reader := bufio.NewReader(source)
	return &Decoder{source: reader}
}

func (d *Decoder) Decode(v interface{}) error {
	var decodedJson []interface{}

	jsonDecoder := json.NewDecoder(d.source)

	if err := jsonDecoder.Decode(&decodedJson); err != nil {
		return err
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &json.InvalidUnmarshalError{reflect.TypeOf(v)}
	}

	if decodedJson[0] == "~#'" {
		if decodedJson[1] == nil {
			t := reflect.TypeOf(v)
			rv.Elem().Set(reflect.Zero(t))
			return nil
		}

		rv.Elem().Set(reflect.ValueOf(decodedJson[1]))
		return nil
	}

	return fmt.Errorf("not decodable")
}
