package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type (
	Config struct {
		GistConfig   GistConfig   `toml:"gist"`
		MediumConfig MediumConfig `toml:"medium"`
	}

	GistConfig struct {
		Token string
	}

	MediumConfig struct {
		Token string
	}
)

func (cnf *Config) LoadConfig(filename string) error {
	_, err := os.Stat(filename)
	if err == nil {
		_, err := toml.DecodeFile(filename, cnf)
		if err != nil {
			return err
		}
		return nil
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	var ghToken, midToken string
	fmt.Print("GitHub ACCESS TOKEN> ")
	fmt.Scan(&ghToken)
	fmt.Print("Medium ACCESS TOKEN> ")
	fmt.Scan(&midToken)
	cnf.GistConfig.Token = ghToken
	cnf.MediumConfig.Token = midToken

	return toml.NewEncoder(f).Encode(cnf)
}
