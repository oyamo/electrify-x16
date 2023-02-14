package util

import (
	"bufio"
	"os"
)

// readSource
func readSource(path string) (*bufio.Scanner, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

}
