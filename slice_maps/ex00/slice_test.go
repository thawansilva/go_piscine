package slice

import (
	"testing"
)

func TestMustReturnAnSliceTrimmed(t *testing.T) {
	slice := []int {1, 2, 3, 4, 5, 6}
	got := TrimEdges(slice, 1)
	want := []int { 2, 3, 4, 5}
	if len(got) != len(want) {
		t.Errorf("expected = %d; got %d", want, got)
	}
}

func TestOutOfBoundMustReturnEmptySlice(t *testing.T) {
	slice := []int {1, 2, 3, 4, 5, 6}
	got := TrimEdges(slice, 10)
	var want []int = nil
	if len(got) != len(want) {
		t.Errorf("expected = %d; got %d", want, got)
	}
}

func TestOnLimitBoundShouldReturn(t *testing.T) {
	slice := []int {1, 2, 3, 4, 5, 6}
	got := TrimEdges(slice, 3)
	var want []int = nil
	if len(got) != len(want) {
		t.Errorf("expected = %d; got %d", want, got)
	}
}
