package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	goconfluence "github.com/virtomize/confluence-go-api"
)

const pageBaseURL = "https://%s.atlassian.net/wiki/spaces/%s/pages/%s"

type options struct {
	title    string
	spaceKey string
	domain   string
	parent   string
	body     string
	debug    bool
}

var (
	o = &options{}
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
		goconfluence.DebugFlag = o.debug
		apiToken := os.Getenv("CONFLUENCE_API_TOKEN")
		userEmail := os.Getenv("CONFLUENCE_USER_EMAIL")
		if apiToken == "" || userEmail == "" {
			return fmt.Errorf("please set environment variables")
		}

		// reference: https://developer.atlassian.com/cloud/confluence/rest/v1/api-group-content/#api-wiki-rest-api-content-post
		api, err := goconfluence.NewAPI(fmt.Sprintf("https://%s.atlassian.net/wiki/rest/api", o.domain), userEmail, apiToken)
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
				Key: o.spaceKey,
			},
			Status: "current",
		}
		if o.parent != "" {
			data.Ancestors = []goconfluence.Ancestor{{ID: o.parent}}
		}

		c, err := api.CreateContent(data)
		if err != nil {
			return err
		}

		fmt.Println("page generation succeeded!!\n", fmt.Sprintf("title: %s\n", c.Title), fmt.Sprintf(pageBaseURL, o.domain, c.Space.Key, c.ID))
		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&o.title, "title", "t", "", "page title")
	rootCmd.PersistentFlags().StringVarP(&o.spaceKey, "spaceKey", "s", "", "space key to generate page")
	rootCmd.PersistentFlags().StringVarP(&o.domain, "domain", "d", "", "confluence domain: e.g. https://<domain>.atlassian.net/wiki/home")
	rootCmd.MarkPersistentFlagRequired("title")
	rootCmd.MarkPersistentFlagRequired("spaceKey")
	rootCmd.MarkPersistentFlagRequired("domain")

	rootCmd.PersistentFlags().StringVarP(&o.parent, "parent", "p", "", "parent page ID")
	rootCmd.PersistentFlags().StringVarP(&o.body, "body", "b", "", "page body")
	rootCmd.PersistentFlags().BoolVar(&o.debug, "debug", false, "debug flag")
}
