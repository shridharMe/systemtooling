/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var name string
var greeting string
var preview bool
var prompt bool
var debug bool = false

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "motd",
	Short: "A utility to customize the Message of the Day (MOTD)",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// If no arguments passed, show usage
		if !prompt && (name == "" || greeting == "") {
			cmd.Usage()
			os.Exit(1)
		}

		// Optionally print flags and exit if DEBUG is set
		if debug {
			fmt.Println("Name:", name)
			fmt.Println("Greeting:", greeting)
			fmt.Println("Prompt:", prompt)
			os.Exit(0)
		}
		// Conditionally read from stdin
		if prompt {
			name, greeting = renderPrompt()
		}
		// Generate message
		m := buildMessage(name, greeting)

		// Either preview message or write to file
		if preview {
			fmt.Println(m)
		} else {
			f, err := os.OpenFile("/etc/motd", os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println("Error: Unable to open to /etc/motd")
				os.Exit(1)
			}
			defer f.Close()
			_, err = f.Write([]byte(m))
			if err != nil {
				fmt.Println("Error: Failed to write to /etc/motd")
				os.Exit(1)
			}

		}

	},
}

// remainder of file omitted
func buildMessage(name, greeting string) string {
	return fmt.Sprintf("%s, %s", greeting, name)
}

func renderPrompt() (name, greeting string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Your Greeting: ")
	greeting, _ = reader.ReadString('\n')
	greeting = strings.TrimSpace(greeting)

	fmt.Print("Your Name: ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	return
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&name, "name", "n", "", "Name to use in message")
	rootCmd.Flags().StringVarP(&greeting, "greeting", "g", "", "Greeting to use in message")
	rootCmd.Flags().BoolVarP(&preview, "preview", "v", false, "Preview message instead of writing to /etc/motd")
	rootCmd.Flags().BoolVarP(&prompt, "prompt", "p", false, "Prompt for name and greeting")

	if os.Getenv("DEBUG") != "" {
		debug = true
	}
}
