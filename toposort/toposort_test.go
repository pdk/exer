package toposort

import (
	"reflect"
	"testing"
)

func TestSortDigraph(t *testing.T) {
	tests := []struct {
		name     string
		digraph  map[string][]string
		expected []string
		hasError bool
	}{
		{
			name: "simple acyclic graph",
			digraph: map[string][]string{
				"a": {"b"},
				"b": {"c"},
				"c": {},
			},
			expected: []string{"a", "b", "c"},
			hasError: false,
		},
		{
			name: "graph with cycle",
			digraph: map[string][]string{
				"a": {"b"},
				"b": {"c"},
				"c": {"a"},
			},
			expected: nil,
			hasError: true,
		},
		{
			name: "disconnected graph",
			digraph: map[string][]string{
				"a": {"b"},
				"c": {"d", "e"},
				"d": {},
				"b": {"c"},
			},
			expected: []string{"a", "b", "c", "d", "e"},
			hasError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SortDigraph(tt.digraph)
			if (err != nil) != tt.hasError {
				t.Errorf("expected error: %v, got: %v", tt.hasError, err)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected: %v, got: %v", tt.expected, result)
			}
		})
	}
}
