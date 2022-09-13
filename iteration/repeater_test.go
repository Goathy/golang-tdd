package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {

	tt := []struct {
		name  string
		input struct {
			char  string
			times uint
		}
		want string
	}{
		{
			"Want \"a\" repeated 5 times",
			struct {
				char  string
				times uint
			}{"a", 5},
			"aaaaa",
		},
		{
			"Want \"y\" repeated 8 times",
			struct {
				char  string
				times uint
			}{
				"y",
				8,
			},
			"yyyyyyyy",
		},
		{
			"Want \"p\" repeated 0 times",
			struct {
				char  string
				times uint
			}{
				"p",
				0,
			},
			"",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := Repeat(tc.input.char, tc.input.times)

			if got != tc.want {
				t.Errorf("want %q but got %q", tc.want, got)
			}
		})
	}

}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func BenchmarkBuildInRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Repeat("a", 5)
	}

}
