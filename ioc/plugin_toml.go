package ioc

import (
	"io"

	"github.com/pelletier/go-toml"
)

var (
	loadTomlPlugin = false
	tomlPlugin     = new(iToml)
)

func tomlLoad() {
	if !loadTomlPlugin {
		RegisterPlugin(BeforeInit, tomlPlugin)
		loadTomlPlugin = true
	}
}

type iToml toml.Tree

func (i *iToml) Lookup(path string, ice Ice) interface{} {
	v := (*toml.Tree)(i).Get(path)
	if v == nil {
		panic(ErrorMissing.Panic("#." + path))
	}
	return v
}

func (i *iToml) Prefix() string {
	return "#"
}

func (i *iToml) Priority() int {
	return 0
}

func TomlLoad(content string) error {
	tomlLoad()
	tree, err := toml.Load(content)
	if err != nil {
		return err
	}
	*tomlPlugin = (iToml)(*tree)
	return nil
}

func TomlLoadFile(path string) error {
	tomlLoad()
	tree, err := toml.LoadFile(path)
	if err != nil {
		return err
	}
	*tomlPlugin = (iToml)(*tree)
	return nil
}

func TomlLoadReader(reader io.Reader) error {
	tomlLoad()
	tree, err := toml.LoadReader(reader)
	if err != nil {
		return err
	}
	*tomlPlugin = (iToml)(*tree)
	return nil
}
