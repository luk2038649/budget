package config

import (
	"encoding/json"
	"errors"
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
const configHeaders = "CURRENT\tNAME\tPATH"

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
	_, err := fmt.Fprintln(tw, "CURRENT\tNAME\tPATH")
	for _, i := range c.Items {
		_, err := fmt.Fprintln(tw, i.String())
		if err != nil {
			return fmt.Errorf("config print: %w", err)
		}
	}

	err = tw.Flush()
	if err != nil {
		return fmt.Errorf("config print flush: %w", err)
	}

	return nil
}

func (i Item) String() string {
	if i.Current {
		return fmt.Sprintf("*\t%s\t%s", i.Name, i.DataPath)
	}

	return fmt.Sprintf("\t%s\t%s", i.Name, i.DataPath)
}

func dataPathFromName(name string) (string, error) {
	path, err := getDataPath()
	if err != nil {
		return "", fmt.Errorf("path from name: %w", err)
	}
	path = filepath.Join(path, name+".json")

	return path, nil
}

// Use sets a given Config Item as current
func Use(name string) error {
	err := initDirs()
	if err != nil {
		return fmt.Errorf("config Create: %w", err)
	}
	c, err := loadConfig()
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}
	ok := false
	for idx, i := range c.Items {
		if i.Name == name {
			i.Current = true
			ok = true
		} else {
			i.Current = false
		}
		c.Items[idx] = i
	}
	if !ok {
		return errors.New("no config found")
	}

	err = c.save()
	if err != nil {
		return fmt.Errorf("use: %w", err)
	}

	return nil
}

// save the given config to file.
func (c Config) save() error {
	configFilePath, err := getConfigPath()
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}
	f, err := os.Create(configFilePath)
	if err != nil {
		return fmt.Errorf("create open file: %w", err)
	}

	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	cBytes, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("save: %w", err)
	}

	_, err = f.Write(cBytes)
	if err != nil {
		return fmt.Errorf("create write: %w", err)
	}

	return nil
}

// Create creates a configuration Item and add it to the config file.
func Create(name string) error {
	err := initDirs()
	if err != nil {
		return fmt.Errorf("config Create: %w", err)
	}

	itemDataPath, err := dataPathFromName(name)
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}

	c, err := loadConfig()
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}

	i := Item{
		DataPath: itemDataPath,
		Name:     name,
		Current:  false,
	}

	c.Items = append(c.Items, i)
	err = c.save()
	if len(c.Items) == 1 {
		err = Use(name)
		if err != nil {
			return fmt.Errorf("create: %w", err)
		}
	}
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}
	fmt.Println("Config created.")

	return nil
}

func Show() error {
	err := initDirs()
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
	b, err := load(f)
	if err != nil {
		return config, fmt.Errorf("load config: %w", err)
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, &config)
		if err != nil {
			return config, fmt.Errorf("load config: %w", err)
		}
	}

	return config, nil
}

// load file and return json byte slice.
func load(file string) ([]byte, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("load read file %s: %w", file, err)
	}

	return b, nil
}

func initDirs() error {
	err := initConfigDir()
	if err != nil {
		return fmt.Errorf("initdirs: %w", err)
	}
	err = initDataDir()
	if err != nil {
		return fmt.Errorf("initdirs: %w", err)
	}

	return nil
}

func initDataDir() error {
	path, err := getDataPath()
	if err != nil {
		return fmt.Errorf("InitData dir: %w", err)
	}
	exists, err := file.Exists(path)
	if err != nil {
		return fmt.Errorf("initData dir check: %w", err)
	}
	if !exists {
		err := os.MkdirAll(path, 0755) //nolint:gomnd
		if err != nil {
			return fmt.Errorf("InitDataDir create dir: %w", err)
		}
	}

	return nil
}

// initConfigDir checks for existing config directory and creates if needed.
func initConfigDir() error {
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
		return fmt.Errorf("init dir config: %w", err)
	}
	exists, err = file.Exists(configPath)
	if err != nil {
		return fmt.Errorf("initConfig config file check: %w", err)
	}
	if !exists {
		f, err := os.Create(configPath)
		if err != nil {
			return fmt.Errorf("intidir config create: %w", err)
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
