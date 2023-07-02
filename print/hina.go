package print

import (
	"errors"
	"fmt"
)

func Hina(name string) (string, error) {
	// Validation
	if name == "" {
		return "", errors.New("Name Is Empty~")
	}
	message := fmt.Sprintf("Welcome Bro %v", name)
	return message, nil
}
