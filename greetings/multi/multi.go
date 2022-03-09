package multi

import (
	"errors"
	"fmt"
	"math/rand"
)

func Multi(name string) (string, error) {
	if name == "" {
		return name, errors.New("Empty Name")
	}
	message := fmt.Sprintf(string(rand.Intn(100)), name)
	return message, nil
}
