package generator

import (
	"math/rand"
)

// MaxInt - максимальное значение
var MaxInt int = 5

// GenerateInt - генерация
func GenerateInt() int {
	return rand.Intn(MaxInt)
}
