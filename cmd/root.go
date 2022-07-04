/*
Copyright © 2022 CallanBi <mcbihv@126.com>
*/
package cmd

import (
	"fmt"
	"os"

	color "github.com/fatih/color"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "leetecho-cli",
	Short: "A cli tool for Leetecho",
	Long: `A cli tool for leetecho.

Leetecho(GUI) is an elegant, experience-friendly product that
automatically generates LeetCode solutions and notes and
publishes them to personal repositories on code hosting platforms.
Leetecho-cli simply implements basic functionality in Leetecho,
such as publishing problems and notes according to
the configured template, login and logout, and
checking updates.
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		color.Red("An Error occurs when executing command: %v", err)
		os.Exit(1)
	}
}

func init() {
	// OnInitialize sets the passed functions to be run when each command's
	// Execute method is called.
	cobra.OnInitialize(initConfig, checkUpdate)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.leetecho-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		color.Red("An Error occurs when getting home directory: %v", err)
		os.Exit(1)
	}

	sep := string(os.PathSeparator)

	// Search config in home directory. See viper docs -> https://github.com/spf13/viper
	viper.AddConfigPath(home + sep + ".leetecho-cli")
	viper.SetConfigName(".leetecho-cli")

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		// If a config file is found, read it in.
		if IsLogin() {
			info := color.New(color.FgWhite, color.BgHiBlue).SprintFunc()
			if username := viper.GetString("username"); username != "" {
				if endpoint := viper.GetString("endpoint"); endpoint != "" {
					fmt.Printf("Logged in as %s at %s\n\n", info(username), info(endpoint))
				} else {
					fmt.Printf("Logged in as %s\n\n", info("%s", username))
				}
			}
		} else {
			color.Red("No user configuration found. Please login first.")
		}
	} else {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// If no config file is found, create the folder and file.
			if err := os.MkdirAll(home+sep+".leetecho-cli", 0777); err != nil {
				color.Red("An Error occurs when creating config directory: %v", err)
				os.Exit(1)
			}
			if err := viper.WriteConfigAs(home + sep + ".leetecho-cli" + sep + ".leetecho-cli.yaml"); err != nil {
				color.Red("An Error occurs when creating config file: %v", err)
				os.Exit(1)
			}
			fmt.Println("No config file found, created one at", home+sep+".leetecho-cli"+sep+".leetecho-cli.yaml")
		} else {
			color.Red("An Error occurs when reading config file: %v", err)
			os.Exit(1)
		}
	}
}

func checkUpdate() {

}
