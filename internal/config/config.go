package config

import (
	"budget/internal/budget"
	"budget/internal/file"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// As of now a config is just a file.
// the file has a `budget`
// we can save the file as json.

const dataPath = ".budget"

// Create creates a configuration file with a budget at the given name in the data directory
func Create(name string) error {
	err := initDir()
	if err != nil {
		return fmt.Errorf("Config Create: %w", err)
	}
	path, err := getDataPath()
	path = filepath.Join(path, name+".json")
	if err != nil {
		return fmt.Errorf("Create: %w", err)
	}
	f, err := os.Create(path)
	defer f.Close()

	b := budget.New(name)
	bb, err := json.Marshal(b)
	if err != nil {
		return fmt.Errorf("Config json marshal: %w", err)
	}
	
	_, err = f.Write(bb)
	if err != nil {
		fmt.Println()
		return fmt.Errorf("config create write: %w", err)
	}

	return nil
}

// initConfigDir checks for existing config directory and creates if needed.
func initDir() error {
	path, err := getDataPath()
	if err != nil {
		return fmt.Errorf("InitConfigDir: %w", err)
	}
	exists, err := file.Exists(path)
	if err != nil {
		return fmt.Errorf("initConfig file check: %w", err)
	}
	if !exists {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return fmt.Errorf("InitConfig create dir: %w", err)
		}
	}
	return nil
}

// getConfigs returns a list of all json files found in data dir
func list() ([]string, error) {
	path, err := getDataPath()
	if err != nil {
		return nil, fmt.Errorf("GetConfigs: %w", err)
	}
	files, err := filepath.Glob(filepath.Join(path, "*.json"))
	if err != nil {
		return nil, fmt.Errorf("GetConfigs glob: %w", err)
	}
	return files, nil
}

// getDataPath returns the path to the .budget/data directory for data storage
func getDataPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("GetDataPath: %w", err)
	}
	path := filepath.Join(home, dataPath)
	return path, nil
}
