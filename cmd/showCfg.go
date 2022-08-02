package cmd

import (
	color "github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var showCfgCmd = &cobra.Command{
	Use:   "showcfg",
	Short: "Show current user and its configuration.",
	Long:  `Show current user and its configuration.`,
	Run:   showCfgMain,
}

func init() {
	RootCmd.AddCommand(showCfgCmd)
}

func showCfgMain(cmd *cobra.Command, args []string) {
	info := color.New(color.Bold, color.FgHiYellow).SprintFunc()

	color.Green("Current user: %s", info(viper.GetString("useraccount")))
	color.Green("Current endpoint: %s", info(viper.GetString("endpoint")))
	reponame, repousername, branch, email := viper.GetString("reponame"), viper.GetString("repousername"), viper.GetString("branch"), viper.GetString("email")
	if reponame != "" {
		color.Green("Current repository: %s", info(reponame))
	}
	if repousername != "" {
		color.Green("Current repository owner: %s", info(repousername))
	}
	if branch != "" {
		color.Green("Current branch: %s", info(branch))
	}
	if email != "" {
		color.Green("Current email: %s", info(email))
	}
}

func IsLogin() bool {
	info := color.New(color.Bold, color.FgHiYellow).SprintFunc()
	color.Green("Current settings: %s", info(viper.AllSettings()))
	return len(viper.AllSettings()) != 0
}
