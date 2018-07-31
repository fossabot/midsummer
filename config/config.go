package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	homedir "github.com/mitchellh/go-homedir"
)

const (
	cnfDir  string = "~/.config/midsummer"
	cnfFile string = "config.toml"
)

type (
	Config struct {
		GistConfig   GistConfig   `toml:"gist"`
		MediumConfig MediumConfig `toml:"medium"`
	}

	GistConfig struct {
		Token string `toml:"token"`
	}

	MediumConfig struct {
		Token string `toml:"token"`
	}
)

func (cnf *Config) LoadConfig() error {
	dir, err := homedir.Expand(cnfDir)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("cannot create directory: %v", err)
	}
	filename := filepath.Join(dir, cnfFile)

	_, err = os.Stat(filename)
	if err == nil {
		_, err := toml.DecodeFile(filename, cnf)
		if err != nil {
			return err
		}
		return nil
	}

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("cannot create config file: %v", err)
	}

	var ghToken string
	fmt.Print("GitHub ACCESS TOKEN> ")
	fmt.Scan(&ghToken)
	if ghToken == "" {
		return errors.New("GitHub access token is required")
	}
	var midToken string
	fmt.Print("Medium ACCESS TOKEN> ")
	fmt.Scan(&midToken)
	if midToken == "" {
		return errors.New("Medium access token is required")
	}
	cnf.GistConfig.Token = ghToken
	cnf.MediumConfig.Token = midToken

	return toml.NewEncoder(f).Encode(cnf)
}
