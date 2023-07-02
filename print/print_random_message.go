package print

import (
	"errors"
	"math/rand"
	"fmt"
)

func PrintRandomMessage(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty Name!")
	}
	customMessage := fmt.Sprintf(getRandomMessage(),name)
	return customMessage, nil
}

func getRandomMessage() string {
	randomMessages := []string{
		"Say Hi To %v",
		"Welcome To Our Board %v",
		"Change The World By %v",
	}
	return randomMessages[rand.Intn(len(randomMessages))]
}
