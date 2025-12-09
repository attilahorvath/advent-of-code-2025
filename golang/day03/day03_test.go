package main

import "testing"

func TestOutputWith2Batteries1(t *testing.T) {
	sum := output([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1}, 2)

	if sum != 98 {
		t.Errorf("output(987654321111111, 2) = %d, want 98", sum)
	}
}

func TestOutputWith2Batteries2(t *testing.T) {
	sum := output([]int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9}, 2)

	if sum != 89 {
		t.Errorf("output(811111111111119, 2) = %d, want 89", sum)
	}
}
func TestOutputWith2Batteries3(t *testing.T) {
	sum := output([]int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8}, 2)

	if sum != 78 {
		t.Errorf("output(234234234234278, 2) = %d, want 78", sum)
	}
}

func TestOutputWith2Batteries4(t *testing.T) {
	sum := output([]int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1}, 2)

	if sum != 92 {
		t.Errorf("output(818181911112111, 2) = %d, want 92", sum)
	}
}

func TestOutputWith12Batteries1(t *testing.T) {
	sum := output([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1}, 12)

	if sum != 987654321111 {
		t.Errorf("output(987654321111111, 12) = %d, want 987654321111", sum)
	}
}

func TestOutputWith12Batteries2(t *testing.T) {
	sum := output([]int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9}, 12)

	if sum != 811111111119 {
		t.Errorf("output(811111111111119, 12) = %d, want 811111111119", sum)
	}
}
func TestOutputWith12Batteries3(t *testing.T) {
	sum := output([]int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8}, 12)

	if sum != 434234234278 {
		t.Errorf("output(234234234234278, 12) = %d, want 434234234278", sum)
	}
}

func TestOutputWith12Batteries4(t *testing.T) {
	sum := output([]int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1}, 12)

	if sum != 888911112111 {
		t.Errorf("output(818181911112111, 12) = %d, want 888911112111", sum)
	}
}
