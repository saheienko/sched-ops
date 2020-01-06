package k8s

import (
	"testing"
)

func TestMapToCSV(t *testing.T) {
	tcs := []struct {
		in       map[string]string
		expected string
	}{
		{},
		{
			in:       map[string]string{"k": "v", "k1": "v1"},
			expected: "k=v,k1=v1",
		},
	}

	for i, tc := range tcs {
		res := mapToCSV(tc.in)
		if res != tc.expected {
			t.Fatalf("TC%d: expected=%q got=%q", i+1, tc.expected, res)
		}
	}
}
