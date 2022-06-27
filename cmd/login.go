package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

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
	fmt.Println("username: ", username)
	fmt.Println("password: ", password)
}

func getLoginInfo() (username string, password string) {
	if checkIfUserExist() {
		fmt.Println("Existing user found: ", viper.GetString("username"))
		return viper.GetString("username"), viper.GetString("password")
	} else {
		fmt.Println("No user found. Please enter username and password.")
		// ask for username and password
		fmt.Println("Enter username: ")
		fmt.Scanln(&username)
		fmt.Println("Enter password: ")
		// hide password
		pwd, err := term.ReadPassword(0)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(password))
		fmt.Println("logining...")
		password = string(pwd)
		return
	}
}

func checkIfUserExist() bool {
	if viper.GetString("username") == "" {
		fmt.Println("username is required")
		return false
	}
	if viper.GetString("password") == "" {
		fmt.Println("password is required")
		return false
	}
	return true
}
