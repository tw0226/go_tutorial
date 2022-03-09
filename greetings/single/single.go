package single

import (
	"errors"
	"fmt"
	"math/rand"
)

func Single(name string) (string, error) {
	if name == "" {
		return name, errors.New("Empty Name")
	}
	message := fmt.Sprintf(string(rand.Intn(100)), name)
	return message, nil
}
