package ports

import (
	"testing"
)

type tcase struct {
	input       []string
	want        []int
	description string
	err         bool
}

type tcase2 struct {
	input       []int
	want        []string
	description string
	err         bool
}

func TestMakePortList(t *testing.T) {

	all, _ := makeRange(1, 65535)

	tcases := []tcase{
		{
			input:       []string{"80"},
			want:        []int{80},
			description: "Testing single port",
		},
		{
			input:       []string{"-1"},
			err:         true,
			description: "Testing single port",
		},
		{
			input:       []string{"80-85"},
			want:        []int{80, 81, 82, 83, 84, 85},
			description: "Testing range of ports",
		},
		{
			input:       []string{"80-23"},
			err:         true,
			description: "Testing range of ports",
		},
		{
			input:       []string{"-1-23"},
			err:         true,
			description: "Testing range of ports",
		},
		{
			input:       []string{"60000-80000"},
			err:         true,
			description: "Testing range of ports",
		},
		{
			input:       []string{"80", "81", "82", "83", "84", "85"},
			want:        []int{80, 81, 82, 83, 84, 85},
			description: "Testing list of ports",
		},
		{
			input:       []string{"80", "81", "82", "800000", "84", "85"},
			err:         true,
			description: "Testing list of ports with out of range port",
		},
		{
			input:       []string{"-1", "80", "81", "82", "84", "85"},
			err:         true,
			description: "Testing list of ports with out of range port",
		},
		{
			input:       []string{"alksdjasl"},
			err:         true,
			description: "Testing list of ports",
		},
		{
			input:       []string{"all"},
			want:        all,
			description: "Testing list of ports",
		},
	}

	for _, tc := range tcases {

		got, err := makePortList(tc.input)

		if err != nil && tc.err != true {
			t.Errorf("Should not have thrown error but did, %v is a valid input", tc.input)
		}
		if err == nil && tc.err == true {
			t.Errorf("Should have thrown error but did not, %v is not a valid input", tc.input)
		}
		for i, w := range tc.want {
			if got[i] != w || err != nil {
				t.Errorf("Incorrect output, got: %d, want: %d.", got, tc.want)
			}
		}
	}

}

func TestStringPortList(t *testing.T) {

	tcases := []tcase2{
		{
			input:       []int{80},
			want:        []string{"80"},
			description: "Testing single port",
		},
		{
			input:       []int{30, 81},
			want:        []string{"30", "81"},
			description: "Testing range of ports",
		},
	}

	for _, tc := range tcases {

		got := stringPortList(tc.input)

		for i, w := range tc.want {
			if got[i] != w {
				t.Errorf("Incorrect output, got: %v, want: %v.", got, tc.want)
			}
		}
	}

}
