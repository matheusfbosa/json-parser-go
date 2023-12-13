package main

import (
	"errors"
	"os"
)

var errInvalidJSON = errors.New("invalid JSON")

func parseJSON(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errInvalidJSON
	}

	return data, nil
}
