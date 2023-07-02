package grade

import (
	"errors"
)

// Score Will Qualified Here
func PassOrNot(grade string) (int, error) {
	if grade == "" {
		return 0, errors.New("Grade Undefined!")
	}
	if grade == "A" {
		return 100, nil
	} else if grade == "B" {
		return 80, nil
	} else if grade == "C" {
		return 60, nil
	}
	return 0, nil
}
