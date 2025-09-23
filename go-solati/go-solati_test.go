package main

import "testing"

func TestTestAdd(t *testing.T) {
	a, b := 10, 20
	exp := a + b
	res := testAdd(a, b)

	if res != exp {
		t.Errorf("Expected testAdd(%d, %d) = %d, but got %d", a, b, exp, res)
	}
}

func TestTestSub(t *testing.T) {
	a, b := 20, 10
	exp := a - b
	res := testSub(a, b)

	if res != exp {
		t.Errorf("Expected testSub(%d, %d) = %d, but got %d", a, b, exp, res)
	}
}

func BenchmarkTestAdd(b *testing.B) {
	x, y := 10, 20
	for b.Loop() {
		testAdd(x, y)
	}
}

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name  string // Description of this test case
		arr   []int
		value int
		want  int
	}{
		{"test", []int{1, 2, 3, 10, 12, 33, -1}, 10, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinarySearch(tt.arr, tt.value)
			if got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
