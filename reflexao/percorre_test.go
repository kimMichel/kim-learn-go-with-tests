package reflexao

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

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
		{
			"Struct without string field",
			struct {
				Name string
				Age  int
			}{"Kim", 24},
			[]string{"Kim"},
		},
		{
			"Nested Fields",
			Person{
				"Kim",
				Profile{24, "São Paulo"},
			},
			[]string{"Kim", "São Paulo"},
		},
		{
			"Pointer to things",
			&Person{
				"Kim",
				Profile{24, "São Paulo"},
			},
			[]string{"Kim", "São Paulo"},
		},
		{
			"Slices",
			[]Profile{
				{24, "São Paulo"},
				{25, "Seul"},
			},
			[]string{"São Paulo", "Seul"},
		},
		{
			"Arrays",
			[2]Profile{
				{24, "São Paulo"},
				{25, "Seul"},
			},
			[]string{"São Paulo", "Seul"},
		},
		{
			"Maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
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
