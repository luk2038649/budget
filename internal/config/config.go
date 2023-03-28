package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"text/tabwriter"

	"budget/internal/file"
)

const dataPath = "data"
const configPath = "config"
const basePath = ".budget"
const tabWidth = 8

type Config struct {
	Items []Item
}
type Item struct {
	DataPath string
	Name     string
	Current  bool
}

func (c Config) Print() error {
	tw := tabwriter.NewWriter(os.Stdout, 0, tabWidth, 1, '\t', tabwriter.AlignRight)
	for _, i := range c.Items {
		_, err := fmt.Fprintln(tw, i.String())
		if err != nil {
			return fmt.Errorf("config print: %w", err)
		}
	}

	err := tw.Flush()
	if err != nil {
		return fmt.Errorf("config print flush: %w", err)
	}

	return nil
}

func (i Item) String() string {
	return fmt.Sprintf("Name: %s\tPath: %s", i.Name, i.DataPath)
}

func dataPathFromName(name string) (string, error) {
	path, err := getDataPath()
	if err != nil {
		return "", fmt.Errorf("path from name: %w", err)
	}
	path = filepath.Join(path, name+".json")

	return path, nil
}

// Create creates a configuration Item with a budget at the given name in the data directory.
func Create(name string) error {
	err := initDir()
	if err != nil {
		return fmt.Errorf("config Create: %w", err)
	}
	path, err := dataPathFromName(name)
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	i := Item{
		DataPath: path,
		Name:     name,
		Current:  false,
	}
	iBytes, err := json.Marshal(i)
	if err != nil {
		return fmt.Errorf("config json marshal: %w", err)
	}

	_, err = f.Write(iBytes)
	if err != nil {
		fmt.Println()

		return fmt.Errorf("config create write: %w", err)
	}

	return nil
}

func Show() error {
	err := initDir()
	if err != nil {
		return fmt.Errorf("show: %w", err)
	}
	conf, err := loadConfig()
	if err != nil {
		return fmt.Errorf("show: %w", err)
	}
	err = conf.Print()
	if err != nil {
		return fmt.Errorf("show: %w", err)
	}

	return nil
}

func loadConfig() (Config, error) {
	var config = Config{}
	f, err := getConfigPath()
	if err != nil {
		return config, fmt.Errorf("load config: %w", err)
	}
	err = load(f, config)
	if err != nil {
		return config, fmt.Errorf("load config: %w", err)
	}

	return config, nil
}

// load file and marshal json to struct.
func load(file string, v interface{}) error {
	b, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("load read file %s: %w", file, err)
	}
	if len(b) > 0 { // cant unmarshall empty files.
		err = json.Unmarshal(b, &v)
		if err != nil {
			return fmt.Errorf("load unmarshall file %s: %w", file, err)
		}
	}

	return nil
}

// initConfigDir checks for existing config directory and creates if needed.
func initDir() error {
	path, err := getConfigDirPath()
	if err != nil {
		return fmt.Errorf("InitDir: %w", err)
	}
	exists, err := file.Exists(path)
	if err != nil {
		return fmt.Errorf("initConfig dir check: %w", err)
	}
	if !exists {
		err := os.MkdirAll(path, 0755) //nolint:gomnd
		if err != nil {
			return fmt.Errorf("InitConfig create dir: %w", err)
		}
	}
	configPath, err := getConfigPath()
	if err != nil {
		return fmt.Errorf("init dir create config file: %w", err)
	}
	exists, err = file.Exists(configPath)
	if err != nil {
		return fmt.Errorf("initConfig config file check: %w", err)
	}
	if !exists {
		f, err := os.Create(configPath)
		if err != nil {
			return fmt.Errorf("create: %w", err)
		}
		defer func(f *os.File) {
			_ = f.Close()
		}(f)
	}

	return nil
}

// getDataPath returns the path to the .budget/data directory for data storage.
func getDataPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("get data path: %w", err)
	}
	path := filepath.Join(home, basePath, dataPath)

	return path, nil
}

func getConfigDirPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("get config dir path: %w", err)
	}
	path := filepath.Join(home, basePath, configPath)

	return path, nil
}
func getConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("get config path: %w", err)
	}
	path := filepath.Join(home, basePath, configPath, "config.json")

	return path, nil
}
