package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Config defines a user configuration file
type Config struct {
	Workers  []Worker  `json:workers`
	Mappings []Mapping `json:mappings`
}

// Worker defines a remote worker for building
type Worker struct {
	ID   string `json:"id"`
	Host string `json:"host"`
	User string `json:"user"`
	Port int    `json:"port"`
}

// Mapping maps a local directory to a remote directory.
type Mapping struct {
	// Worker to build on
	Worker string `json:"worker"`
	// Local Directory
	Local string `json:"local"`
	// Remote directory
	Remote string `json:"remote"`
}

func Load() (*Config, error) {
	file, err := ioutil.ReadFile(filepath.Join(os.Getenv("HOME"), ".rbd", "config.json"))

	if err != nil {
		return nil, err
	}

	config := &Config{}

	err = json.Unmarshal(file, config)

	return config, err
}

var ErrMapNotFound = errors.New("Current directory does not match any mappings")

func (c *Config) GetMap() (*Mapping, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	for _, mapping := range c.Mappings {
		absdir, err := filepath.Abs(mapping.Local)
		if err != nil {
			return nil, err
		}
		if absdir == cwd {
			return &mapping, nil
		}
	}

	return nil, ErrMapNotFound
}

var ErrWorkerNotFound = errors.New("Specified worker not found")

func (c *Config) GetWorker(ID string) (*Worker, error) {
	for _, worker := range c.Workers {
		if worker.ID == ID {
			return &worker, nil
		}
	}

	return nil, ErrWorkerNotFound
}
