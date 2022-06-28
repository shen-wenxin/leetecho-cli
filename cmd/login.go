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
	Run:   loginMain,
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

func loginMain(cmd *cobra.Command, args []string) {
	Login()
}

func Login() {
	username, password := getLoginInfo()

	var initClientErr error

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Color("fgHiBlue")
	s.Prefix = "Logging in "

	s.Start()
	// Endpoint is hard coded to CN region for the temporarily
	LeetcodeClient, initClientErr = leetcode_client.Build(username, password, helper.CN)

	if initClientErr != nil {
		s.Stop()
		if leetechoError, ok := initClientErr.(*helper.ErrorResp); ok {
			if leetechoError.Code == 400 {
				color.Red("Login failed. Username or password is incorrect.")
				WipeConfig()
			} else {
				color.Red("An Error occurs when initializing client: " + initClientErr.Error())
				WipeConfig()
			}
			os.Exit(1)
		} else {
			color.Red("An Error occurs when initializing client: " + initClientErr.Error())
			WipeConfig()
			os.Exit(1)
		}
	}

	s.Stop()

	viper.Set("username", username)
	viper.Set("password", password)
	viper.Set("endpoint", helper.CN)
	viper.WriteConfig()

	color.Green(fmt.Sprintf("Login successful. Current user: %s; Current endpoint: %s. ", username, helper.CN))
}

func getLoginInfo() (username string, password string) {
	if checkIfUserExist() {
		info := color.New(color.FgWhite, color.BgHiMagenta).SprintFunc()
		fmt.Println(info("Existing user found: ", viper.GetString("username")))
		username = viper.GetString("username")
		password = viper.GetString("password")
		return
	} else {
		info := color.New(color.FgWhite, color.BgHiYellow).SprintFunc()
		fmt.Println(info("No user found in your configuration file. Please enter your username and password. "))
		// ask for username and password
		asked := color.New(color.Bold).SprintFunc()
		fmt.Println(asked("Enter your LeetCode(CN) username: "))
		fmt.Scanln(&username)
		fmt.Println(asked("Enter your password: "))
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
