package quiz

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseProblems(t *testing.T) {
	in := [][]string{
		{"5+5", "10"},
		{"1+1", "2"},
		{"8+3", "11"},
	}
	// expected output
	expected := []Problem{
		{
			Question: "5+5",
			Answer:   "10",
		},
		{
			Question: "1+1",
			Answer:   "2",
		},
		{
			Question: "8+3",
			Answer:   "11",
		},
	}

	got := ParseProblems(in)

	if diff := cmp.Diff(expected, got); diff != "" {
		t.Fatalf("The difference between expected and got is %s", diff)
	}
}
