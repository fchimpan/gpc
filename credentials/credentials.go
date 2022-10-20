package credentials

import (
	"gopkg.in/ini.v1"
)

type Credentials struct {
	ConfluenceAPIToken string `ini:"confluence_api_token"`
	ConfluenceAEmail   string `ini:"confluence_email"`
}

func GetCredentials(credFilePath, sectionName string) (*Credentials, error) {
	c, err := ini.Load(credFilePath)
	if err != nil {
		return nil, err
	}
	cred := &Credentials{
		ConfluenceAPIToken: c.Section(sectionName).Key("confluence_api_token").String(),
		ConfluenceAEmail:   c.Section(sectionName).Key("confluence_email").String(),
	}
	return cred, nil
}
