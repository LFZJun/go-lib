package ioc

import (
	"io"

	"github.com/pelletier/go-toml"
)

var load = false

func tomlLoad() {
	if !load {
		RegisterPlugin(BeforeInit, tomlPlugin)
		load = true
	}
}

var tomlPlugin = new(iocToml)

type iocToml toml.TomlTree

func (i *iocToml) Value(path string) interface{} {
	v := (*toml.TomlTree)(i).Get(path)
	if v == nil {
		panic(ErrorMissing.Panic(path))
	}
	return v
}

func (i *iocToml) Prefix() string {
	return "#"
}

func (i *iocToml) Priority() int {
	return 0
}

func TomlLoad(content string) error {
	tomlLoad()
	tree, err := toml.Load(content)
	if err != nil {
		return err
	}
	*tomlPlugin = (iocToml)(*tree)
	return nil
}

func TomlLoadFile(path string) error {
	tomlLoad()
	tree, err := toml.LoadFile(path)
	if err != nil {
		return err
	}
	*tomlPlugin = (iocToml)(*tree)
	return nil
}

func TomlLoadReader(reader io.Reader) error {
	tomlLoad()
	tree, err := toml.LoadReader(reader)
	if err != nil {
		return err
	}
	*tomlPlugin = (iocToml)(*tree)
	return nil
}
