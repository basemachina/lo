package lo

import (
	"errors"
	"slices"
	"testing"
)

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

func TestReduce(t *testing.T) {
	tests := map[string]struct {
		collection  []int
		accumulator func(agg int, item int) int
		initial     int
		want        int
	}{
		"sum": {
			collection:  []int{1, 2, 3, 4, 5},
			accumulator: func(agg int, item int) int { return agg + item },
			initial:     0,
			want:        15,
		},
		"product": {
			collection:  []int{1, 2, 3, 4},
			accumulator: func(agg int, item int) int { return agg * item },
			initial:     1,
			want:        24,
		},
		"empty": {
			collection:  []int{},
			accumulator: func(agg int, item int) int { return agg + item },
			initial:     10,
			want:        10,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := Reduce(tt.collection, tt.accumulator, tt.initial); got != tt.want {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduceWithIndex(t *testing.T) {
	collection := []int{1, 2, 3}
	accumulator := func(agg int, item int, index int) int { return agg + item + index }
	initial := 0

	result := ReduceWithIndex(collection, accumulator, initial)
	expected := 9 // (0+1+0)=1, (1+2+1)=4, (4+3+2)=9

	if result != expected {
		t.Errorf("ReduceWithIndex() = %v, want %v", result, expected)
	}
}

func TestFilterWithError(t *testing.T) {
	tests := map[string]struct {
		collection []int
		predicate  func(int) (bool, error)
		want       []int
		wantErr    bool
	}{
		"filter even numbers": {
			collection: []int{1, 2, 3, 4, 5, 6},
			predicate:  func(item int) (bool, error) { return item%2 == 0, nil },
			want:       []int{2, 4, 6},
			wantErr:    false,
		},
		"filter with error on specific value": {
			collection: []int{1, 2, 3, 4, 5},
			predicate: func(item int) (bool, error) {
				if item == 3 {
					return false, errors.New("error on value 3")
				}
				return item%2 == 0, nil
			},
			want:    nil,
			wantErr: true,
		},
		// other cases are tested in TestFilterWithIndexError
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := FilterWithError(tt.collection, tt.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("FilterWithError() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("FilterWithError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterWithIndexError(t *testing.T) {
	tests := map[string]struct {
		collection []int
		predicate  func(int, int) (bool, error)
		want       []int
		wantErr    bool
	}{
		"filter by even index": {
			collection: []int{10, 20, 30, 40, 50},
			predicate:  func(item int, index int) (bool, error) { return index%2 == 0, nil },
			want:       []int{10, 30, 50},
			wantErr:    false,
		},
		"filter with error at specific index": {
			collection: []int{1, 2, 3, 4, 5},
			predicate: func(item int, index int) (bool, error) {
				if index == 2 {
					return false, errors.New("error at index 2")
				}
				return item%2 == 0, nil
			},
			want:    nil,
			wantErr: true,
		},
		"empty collection": {
			collection: []int{},
			predicate:  func(item int, index int) (bool, error) { return index%2 == 0, nil },
			want:       []int{},
			wantErr:    false,
		},
		"filter by value and index": {
			collection: []int{1, 2, 3, 4, 5, 6},
			predicate: func(item int, index int) (bool, error) {
				return item%2 == 0 && index < 4, nil
			},
			want:    []int{2, 4},
			wantErr: false,
		},
		"all elements filtered out": {
			collection: []int{1, 3, 5},
			predicate:  func(item int, index int) (bool, error) { return index > 10, nil },
			want:       []int{},
			wantErr:    false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := FilterWithIndexError(tt.collection, tt.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("FilterWithIndexError() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("FilterWithIndexError() = %v, want %v", got, tt.want)
			}
		})
	}
}
