package ioc

import (
	"io"

	"github.com/pelletier/go-toml"
)

func init() {
	RegisterPlugin(BeforeInit, tomlPlugin)
}

var tomlPlugin = new(iocToml)

type iocToml toml.TomlTree

func (i *iocToml) Value(path string) interface{} {
	v := (*toml.TomlTree)(i).Get(path)
	if v == nil {
		panic("找不到值: " + path)
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
	tree, err := toml.Load(content)
	if err != nil {
		return err
	}
	*tomlPlugin = (iocToml)(*tree)
	return nil
}

func TomlLoadFile(path string) error {
	tree, err := toml.LoadFile(path)
	if err != nil {
		return err
	}
	*tomlPlugin = (iocToml)(*tree)
	return nil
}

func TomlLoadReader(reader io.Reader) error {
	tree, err := toml.LoadReader(reader)
	if err != nil {
		return err
	}
	*tomlPlugin = (iocToml)(*tree)
	return nil
}
