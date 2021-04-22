package main

import "testing"

func Test_multipleof(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want string
	}{
		{
			name: "multiple of 3 should return fizz",
			arg:  3,
			want: "fizz",
		},
		{
			name: "multiple of 5 should return buzz",
			arg:  5,
			want: "buzz",
		},
		{
			name: "multiple of both 3 and 5 should return fizzbuzz",
			arg:  15,
			want: "fizzbuzz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := multipleof(tt.arg)
			if got != tt.want {
				t.Errorf("looks like we've got an error! We got: '%s', but expected to get: '%s'", got, tt.want)
			}
		})
	}
}
