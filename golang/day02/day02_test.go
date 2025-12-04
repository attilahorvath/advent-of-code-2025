package main

import "testing"

var ranges string = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

func TestSumInvalidIdsTwiceOnly(t *testing.T) {
	sum := sumInvalidIds(ranges, true)

	if sum != 1227775554 {
		t.Errorf("sumInvalidIds(ranges, true) = %d, want 1227775554", sum)
	}
}

func TestSumInvalidIdsAllSequences(t *testing.T) {
	sum := sumInvalidIds(ranges, false)

	if sum != 4174379265 {
		t.Errorf("sumInvalidIds(ranges, false) = %d, want 4174379265", sum)
	}
}
