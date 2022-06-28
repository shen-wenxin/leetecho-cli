package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/CallanBi/leetecho-cli/leetcode_client"
	"github.com/CallanBi/leetecho-cli/leetcode_client/helper"
	"github.com/briandowns/spinner"
	color "github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

var LeetcodeClient *leetcode_client.LeetCodeClient

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Enter username and password to login",
	Long:  `Enter username and password to login to LeetCode(only CN region supported).`,
	Run:   login,
}

func init() {
	RootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// byeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// byeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func login(cmd *cobra.Command, args []string) {
	username, password := getLoginInfo()

	var initClientErr error

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Prefix = "Logging in "

	s.Start()
	LeetcodeClient, initClientErr = leetcode_client.Build(username, password, helper.CN)

	if initClientErr != nil {
		s.Stop()
		if leetechoError, ok := initClientErr.(*helper.ErrorResp); ok {
			if leetechoError.Code == 400 {
				color.Red("Login failed. Username or password is incorrect.")
			} else {
				color.Red("An Error occurs when initializing client: " + initClientErr.Error())
				os.Exit(1)
			}
			os.Exit(1)
		} else {
			color.Red("An Error occurs when initializing client: " + initClientErr.Error())
			os.Exit(1)
		}
	}

	s.Stop()

	color.Green("Login successful. Current user: " + username)

}

func getLoginInfo() (username string, password string) {
	if checkIfUserExist() {
		fmt.Println("Existing user found: ", viper.GetString("username"))
		username = viper.GetString("username")
		password = viper.GetString("password")
		return
	} else {
		info := color.New(color.FgWhite, color.BgHiRed).SprintFunc()
		fmt.Println(info("No user found in your configuration file. Please enter your username and password: "))
		// ask for username and password
		fmt.Println("Enter username: ")
		fmt.Scanln(&username)
		fmt.Println("Enter password: ")
		// hide password
		pwd, err := term.ReadPassword(0)
		if err != nil {
			color.Red("An Error occurs when reading password: " + err.Error())
			os.Exit(1)
		}
		fmt.Println(string(password))

		password = string(pwd)
		return
	}
}

func checkIfUserExist() bool {
	if viper.GetString("username") == "" {
		return false
	}
	if viper.GetString("password") == "" {
		return false
	}
	return true
}
