package utils

import (
	"testing"
)

func TestAdd(t *testing.T) {
	expected := 7
	actual := Add(1, 2, 3)

	if actual != expected {
		t.Errorf("Addition was incorrect! Expected %d, got %d", expected, actual)
	}
}
