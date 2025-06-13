package lo

import "testing"

func TestHasDuplicates(t *testing.T) {
	tests := map[string]struct {
		collection []int
		want       bool
	}{
		"has duplicates": {
			collection: []int{1, 1, 2},
			want:       true,
		},
		"no duplicates": {
			collection: []int{1, 2, 3},
			want:       false,
		},
		"empty": {
			collection: []int{},
			want:       false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := HasDuplicates(tt.collection); got != tt.want {
				t.Errorf("HasDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasDuplicatesBy(t *testing.T) {
	type S struct {
		A int
		B int
	}
	tests := map[string]struct {
		collection []S
		iteratee   func(S) int
		want       bool
	}{
		"has duplicates by A": {
			collection: []S{
				{A: 0, B: 1},
				{A: 0, B: 2},
			},
			iteratee: func(item S) int { return item.A },
			want:     true,
		},
		"has duplicates by B": {
			collection: []S{
				{A: 1, B: 0},
				{A: 2, B: 0},
			},
			iteratee: func(item S) int { return item.B },
			want:     true,
		},
		"no duplicates by A": {
			collection: []S{
				{A: 1, B: 0},
				{A: 2, B: 0},
			},
			iteratee: func(item S) int { return item.A },
			want:     false,
		},
		"no duplicates by B": {
			collection: []S{
				{A: 0, B: 1},
				{A: 0, B: 2},
			},
			iteratee: func(item S) int { return item.B },
			want:     false,
		},
		"empty": {
			collection: []S{},
			iteratee:   func(item S) int { return item.A },
			want:       false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := HasDuplicatesBy(tt.collection, tt.iteratee); got != tt.want {
				t.Errorf("HasDuplicatesBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
