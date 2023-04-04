package reflexao

import (
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {

	cases := []struct {
		Name             string
		Input            interface{}
		ExpectedRequests []string
	}{
		{
			"Struct with a string field",
			struct {
				Name string
			}{"Kim"},
			[]string{"Kim"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Kim", "Brazil"},
			[]string{"Kim", "Brazil"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var result []string
			percorre(test.Input, func(input string) {
				result = append(result, input)
			})

			if !reflect.DeepEqual(result, test.ExpectedRequests) {
				t.Errorf("result '%v', expect '%v'", result, test.ExpectedRequests)
			}
		})
	}
}
