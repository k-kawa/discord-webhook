/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/k-kawa/discord-webhook/discord"
	"github.com/k-kawa/discord-webhook/we"
	"github.com/spf13/cobra"
)

const WebhookEnvName = "DISCORD_WEBHOOK_URL"

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send message using a webhook URL",
	Long: `Send message using a webhook URL

You can give a Discord Webhook URL via environment variable named DISCORD_WEBHOOK_URL.
The variable can be derived from We Config file at the path WE_CONFIG.

You can specify the content of the message in the following ways
- via CONTENT environment variable
- via file whose name is set CONTENT_FILE environment variable
- or standard input of this command.
`,
	Run: func(cmd *cobra.Command, args []string) {
		webhookURL := os.Getenv(WebhookEnvName)

		if webhookURL == "" {
			weConfigPath := os.Getenv("WE_CONFIG")
			if weConfigPath != "" {
				weConfig, err := we.OpenURLs(weConfigPath)
				if err != nil {
					log.Fatalf("Failed to open we_config: %s", err.Error())
				}
				envvars, err := weConfig.EnvVars()
				if err != nil {
					log.Fatalf("Failed to get EnvVars: %s", err.Error())
				}
				for _, envvar := range envvars {
					if envvar.Name == WebhookEnvName {
						webhookURL = envvar.Value
						break
					}
				}
			}
		}
		if webhookURL == "" {
			log.Fatalf("Faield to find webhook URL")
		}

		content := os.Getenv("CONTENT")
		if content != "" {
			log.Printf("Content was loaded from CONFIG environment variable")
		}

		if content == "" {
			contentFile := os.Getenv("CONTENT_FILE")
			if contentFile != "" {
				bcontent, err := ioutil.ReadFile(contentFile)
				if err != nil {
					log.Fatalf("Failed to read content from %s: %s", contentFile, err.Error())
				}

				content = string(bcontent)
				log.Printf("Content was loaded from file %s", contentFile)
			}
		}

		if content == "" {
			bcontent, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				log.Fatalf("Failed to read from stdin: %s", err.Error())
			}
			content = string(bcontent)
			log.Printf("Content was loaded from stdin")
		}

		if err := discord.Post(
			discord.WebhookURL(webhookURL),
			&discord.WebhooksPostRequest{
				Content: content,
			},
		); err != nil {
			log.Fatalf("Failed to send: %s", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
