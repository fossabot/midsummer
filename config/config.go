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

	ghToken, err := scanToken("GitHub")
	if err != nil {
		return err
	}
	midToken, err := scanToken("Medium")
	if err != nil {
		return err
	}
	cnf.GistConfig.Token = ghToken
	cnf.MediumConfig.Token = midToken

	return toml.NewEncoder(f).Encode(cnf)
}

func scanToken(appname string) (string, error) {
	fmt.Print(appname + " Access Token> ")
	var token string
	fmt.Scan(&token)
	if token == "" {
		return "", errors.New(appname + " access token is required")
	}
	return token, nil
}
