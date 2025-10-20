package util

import (
	"fmt"
	"os"
)

func ReadInput(day int8) ([]byte, error) {
	return os.ReadFile(fmt.Sprintf("src/inputs/%d.txt", day))
}

func ReadExampleInput(day int8) ([]byte, error) {
	return os.ReadFile(fmt.Sprintf("src/inputs/%d-example.txt", day))
}
