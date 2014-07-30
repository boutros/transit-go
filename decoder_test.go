package transit

import (
	"fmt"
	"os"
	"testing"
)

// "keywords"
// "list_empty"
// "list_mixed"
// "list_nested"
// "list_simple"
// "map_10_items"
// "map_10_nested"
// "map_1935_nested"
// "map_1936_nested"
// "map_1937_nested"
// "map_mixed"
// "map_nested"
// "map_numeric_keys"
// "map_simple"
// "map_string_keys"
// "map_unrecognized_vals"
// "map_vector_keys"
// "maps_four_char_string_keys"
// "maps_four_char_sym_keys"
// "maps_three_char_string_keys"
// "maps_three_char_sym_keys"
// "maps_two_char_string_keys"
// "maps_two_char_sym_keys"
// "maps_unrecognized_keys"
// "nil"
// "one"
// "one_date"
// "one_keyword"
// "one_string"
// "one_symbol"
// "one_uri"
// "one_uuid"
// "set_empty"
// "set_mixed"
// "set_nested"
// "set_simple"
// "small_ints"
// "small_strings"
// "strings_hash"
// "strings_hat"
// "strings_tilde"
// "symbols"
// "true"
// "uris"
// "uuids"
// "vector_1935_keywords_repeated_twice"
// "vector_1936_keywords_repeated_twice"
// "vector_1937_keywords_repeated_twice"
// "vector_empty"
// "vector_mixed"
// "vector_nested"
// "vector_simple"
// "vector_unrecognized_vals"
// "zero"
// {"dates_interesting"},
// {"doubles_interesting"},
// {"doubles_small"},
// {"ints"},
// {"ints_interesting"},
// {"ints_interesting_neg"},

func TestDecodingBooleanValues(t *testing.T) {
	booltestcases := []struct {
		file     string
		expected interface{}
	}{
		{"false", false},
		{"true", true},
	}

	for _, tc := range booltestcases {
		var actual bool

		f, err := os.Open("_examples/" + tc.file + ".json")
		if err != nil {
			t.Fatal(err)
		}

		decoder := NewDecoder(f)
		err = decoder.Decode(&actual)
		if err != nil {
			t.Fatalf("Could not decode value for file %s, error was %s", tc.file, err.Error())
		}

		if actual != tc.expected {
			t.Fatalf("Expected %#v to equal %#v", actual, tc.expected)
		}
	}
}

func TestDecodingTopLevelNumber(t *testing.T) {
	booltestcases := []struct {
		file     string
		expected interface{}
	}{
		{"one", 1.0},
	}

	for _, tc := range booltestcases {
		var actual float64

		f, err := os.Open("_examples/" + tc.file + ".json")
		if err != nil {
			t.Fatal(err)
		}

		decoder := NewDecoder(f)
		err = decoder.Decode(&actual)
		if err != nil {
			t.Fatalf("Could not decode value for file %s, error was %s", tc.file, err.Error())
		}

		if actual != tc.expected {
			t.Fatalf("Expected %#v to equal %#v", actual, tc.expected)
		}
	}
}

func TestDecodingTopLevelNull(t *testing.T) {
	booltestcases := []struct {
		file     string
		expected interface{}
	}{
		{"nil", nil},
	}

	for _, tc := range booltestcases {
		actual := &Decoder{}
		fmt.Printf("%#v\n", actual)

		f, err := os.Open("_examples/" + tc.file + ".json")
		if err != nil {
			t.Fatal(err)
		}

		decoder := NewDecoder(f)
		err = decoder.Decode(&actual)
		if err != nil {
			t.Fatalf("Could not decode value for file %s, error was %s", tc.file, err.Error())
		}

		if actual != tc.expected {
			t.Fatalf("Expected %#v to equal %#v", actual, tc.expected)
		}
	}
}
