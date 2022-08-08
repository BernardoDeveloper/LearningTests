package main

import (
	"testing"
)

func TestReturnYear(t *testing.T) {
	var yearTest string = "2004"
	actualInt := YearBorn(yearTest)
	var expectedInteger int = 18

	if actualInt != expectedInteger {
		t.Error("Expected Integer ", expectedInteger, " is not same as actual integer ", actualInt)
	}
}
