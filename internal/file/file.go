package file

import (
	"fmt"
	"os"
)

// Exists returns whether the given file or directory exists.
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, fmt.Errorf("exists: %w", err)
}

// Load file and return json byte slice.
func Load(file string) ([]byte, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("load read file %s: %w", file, err)
	}

	return b, nil
}
