package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {

	tt := []struct {
		name  string
		input []int
		want  int
	}{
		{
			"Want 4 for 2 + 2",
			[]int{2, 2},
			4,
		},
		{
			"Want 24 for 3 + 4 + 5 + 12",
			[]int{3, 4, 5, 12},
			24,
		},
		{
			"Want 33 for 0 + 1 + 1 + 2 + 3 + 5 + 8 + 13",
			[]int{0, 1, 1, 2, 3, 5, 8, 13},
			33,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := Add(tc.input...)

			if got != tc.want {
				t.Errorf("want '%d' but got '%d'", tc.want, got)
			}
		})
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
