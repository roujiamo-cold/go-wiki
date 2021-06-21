package testingTechniques

import (
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	const s, sep, want = "chicken", "ken", 4
	got := strings.Index(s, sep)
	if got != want {
		t.Errorf("Index(%q,%q) = %v; want %v", s, sep, got, want)
	}
}

func TestIndexForTableDriven(t *testing.T) {
	var tests = []struct {
		s   string
		sep string
		out int
	}{
		{"", "", 0},
		{"", "a", -1},
		{"fo", "foo", -1},
		{"foo", "foo", 0},
		{"oofofoofooo", "f", 2},
		// etc
	}
	for _, test := range tests {
		actual := strings.Index(test.s, test.sep)
		if actual != test.out {
			t.Errorf("Index(%q,%q) = %v; want %v", test.s, test.sep, actual, test.out)
		}
	}
}

func TestAbs(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_abs_1",
			args: args{
				i: -1,
			},
			want: 1,
		},
		{
			name: "test_abs_1",
			args: args{
				i: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.args.i); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}
