package config

import (
	"gopkg.in/ini.v1"
)

type Config struct {
	SpaceKey string `ini:"space_key"`
	Domain   string `ini:"domain"`
	Parent   string `ini:"parent"`
}

func GetConfig(configFilePath, sectionName string) (*Config, error) {
	c, err := ini.Load(configFilePath)
	if err != nil {
		return nil, err
	}
	cfg := &Config{
		SpaceKey: c.Section(sectionName).Key("space_key").String(),
		Domain:   c.Section(sectionName).Key("domain").String(),
		Parent:   c.Section(sectionName).Key("parent").String(),
	}
	return cfg, nil
}

func GetAllConfig(configFilePath string) ([]string, error) {
	c, err := ini.Load(configFilePath)
	if err != nil {
		return nil, err
	}
	return c.SectionStrings()[1:], nil
}
