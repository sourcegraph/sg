package main

import (
	"io"
	"os"

	"github.com/cockroachdb/errors"
	"github.com/sourcegraph/sourcegraph/dev/sg/internal/run"
	"gopkg.in/yaml.v2"
)

func ParseConfigFile(name string) (*Config, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot open file %q", name)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.Wrap(err, "reading configuration file")
	}

	return ParseConfig(data)
}

func ParseConfig(data []byte) (*Config, error) {
	var conf Config
	if err := yaml.Unmarshal(data, &conf); err != nil {
		return nil, err
	}

	for name, cmd := range conf.Commands {
		cmd.Name = name
		conf.Commands[name] = cmd
	}

	for name, cmd := range conf.Commandsets {
		cmd.Name = name
		conf.Commandsets[name] = cmd
	}

	for name, cmd := range conf.Tests {
		cmd.Name = name
		conf.Tests[name] = cmd
	}

	for name, check := range conf.Checks {
		check.Name = name
		conf.Checks[name] = check
	}

	return &conf, nil
}

type Commandset struct {
	Name     string   `yaml:"-"`
	Commands []string `yaml:"commands"`
	Checks   []string `yaml:"checks"`
}

// UnmarshalYAML implements the Unmarshaler interface.
func (c *Commandset) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// To be backwards compatible we first try to unmarshal as a simple list.
	var list []string
	if err := unmarshal(&list); err == nil {
		c.Commands = list
		return nil
	}

	// If it's not a list we try to unmarshal it as a Commandset. In order to
	// not recurse infinitely (calling UnmarshalYAML over and over) we create a
	// temporary type alias.
	type rawCommandset Commandset
	if err := unmarshal((*rawCommandset)(c)); err != nil {
		return err
	}

	return nil
}

type Config struct {
	Env         map[string]string      `yaml:"env"`
	Commands    map[string]run.Command `yaml:"commands"`
	Commandsets map[string]*Commandset `yaml:"commandsets"`
	Tests       map[string]run.Command `yaml:"tests"`
	Checks      map[string]run.Check   `yaml:"checks"`
}

// Merges merges the top-level entries of two Config objects, with the receiver
// being modified.
func (c *Config) Merge(other *Config) {
	for k, v := range other.Env {
		c.Env[k] = v
	}

	for k, v := range other.Commands {
		if original, ok := c.Commands[k]; ok {
			c.Commands[k] = original.Merge(v)
		} else {
			c.Commands[k] = v
		}
	}

	for k, v := range other.Commandsets {
		c.Commandsets[k] = v
	}

	for k, v := range other.Tests {
		if original, ok := c.Tests[k]; ok {
			c.Tests[k] = original.Merge(v)
		} else {
			c.Tests[k] = v
		}
	}
}
