package utils

import (
	"fmt"
	"math/rand"
)

func RandRanged(min, max int) int {
	return rand.Intn(max-min) + min
}

func CreateIP() string {
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
}
