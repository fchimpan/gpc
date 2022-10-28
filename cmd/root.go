package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fchimpan/gpc/config"
	"github.com/fchimpan/gpc/credentials"
	"github.com/spf13/cobra"
	goconfluence "github.com/virtomize/confluence-go-api"
)

const pageBaseURL = "https://%s.atlassian.net/wiki/spaces/%s/pages/%s"

type options struct {
	title       string
	body        string
	debug       bool
	credentials string
	config      string
}

var (
	o       = &options{}
	homeDir string
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v: try the debug flag\n", err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "gpc",
	Short: "gpc is cli tool to generate pages in confluence",
	Long:  "gpc is cli tool to create any page you want anywhere in confluence",
	RunE: func(cmd *cobra.Command, args []string) error {
		credFilePath := filepath.Join(homeDir, ".gpc", "credentials")
		cred, err := credentials.GetCredentials(credFilePath, o.credentials)
		if err != nil || cred.ConfluenceAPIToken == "" || cred.ConfluenceAEmail == "" {
			return fmt.Errorf("credentials file is not correct: %v", err)
		}

		configFilePath := filepath.Join(homeDir, ".gpc", "config")
		cfg, err := config.GetConfig(configFilePath, o.config)
		if err != nil || cfg.Domain == "" || cfg.SpaceKey == "" {
			return fmt.Errorf("config file is not correct: %v", err)
		}

		goconfluence.DebugFlag = o.debug
		if goconfluence.DebugFlag {
			log.Printf("input: %+v\n", o)
			log.Printf("credentials: %+v\n", cred)
			log.Printf("config: %+v\n", cfg)
		}

		// reference: https://developer.atlassian.com/cloud/confluence/rest/v1/api-group-content/#api-wiki-rest-api-content-post
		api, err := goconfluence.NewAPI(fmt.Sprintf("https://%s.atlassian.net/wiki/rest/api", cfg.Domain), cred.ConfluenceAEmail, cred.ConfluenceAPIToken)
		if err != nil {
			return err
		}
		data := &goconfluence.Content{
			Type:  "page",
			Title: o.title,
			Body: goconfluence.Body{
				Storage: goconfluence.Storage{
					Value:          o.body,
					Representation: "storage",
				},
			},
			Version: &goconfluence.Version{
				Number: 1,
			},
			Space: goconfluence.Space{
				Key: cfg.SpaceKey,
			},
			Status: "current",
		}
		if cfg.Parent != "" {
			data.Ancestors = []goconfluence.Ancestor{{ID: cfg.Parent}}
		}

		c, err := api.CreateContent(data)
		if err != nil {
			return err
		}

		fmt.Println("page generation succeeded!!\n", fmt.Sprintf("title: %s\n", c.Title), fmt.Sprintf(pageBaseURL, cfg.Domain, c.Space.Key, c.ID))
		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&o.title, "title", "t", "", "page title")
	rootCmd.PersistentFlags().StringVarP(&o.config, "profile", "p", "", "config profile name")

	if err := rootCmd.MarkPersistentFlagRequired("title"); err != nil {
		log.Fatal(err)
	}
	if err := rootCmd.MarkPersistentFlagRequired("profile"); err != nil {
		log.Fatal(err)
	}

	rootCmd.PersistentFlags().StringVarP(&o.body, "body", "b", "", "page body")
	rootCmd.PersistentFlags().BoolVar(&o.debug, "debug", false, "debug flag")
	rootCmd.PersistentFlags().StringVarP(&o.credentials, "credentials", "c", "default", "credentials section name")

	homeDir, _ = os.UserHomeDir()
}
