package util

import (
	"encoding/gob"
	"os"
)

func readFile(in string) ([]int16, error) {
	var bytecode []int16
	file, err := os.Open(in)

	// verify if file exists
	if err != nil {
		return nil, err
	}

	defer file.Close()
	dec := gob.NewDecoder(file)
	err = dec.Decode(&bytecode)
	if err != nil {
		return nil, err
	}

	return bytecode, nil
}

func LoadProgram(obj string) ([]int16, error) {
	return readFile(obj)
}
